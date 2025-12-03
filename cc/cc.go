package cc

import (
	"bytes"
	"context"
	"io/ioutil" // nolint:staticcheck
	"net/http"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-co-op/gocron"
	jsoniter "github.com/json-iterator/go"
	"github.com/node-real/go-pkg/log"
	"github.com/shopspring/decimal"

	"github.com/bnb-chain/bsc-staking-indexer/model"
	"github.com/bnb-chain/bsc-staking-indexer/store"
)

const url = "https://ccalert.bk.nodereal.cc/sendalerts"

var weiPerEth = decimal.NewFromInt(1e18)

// OperatorRewardConfig holds configuration for a single operator whose rewards are to be calculated.
type OperatorRewardConfig struct {
	OperatorAddress       common.Address // The address of the validator operator.
	SelfDelegationAddress common.Address // A specific delegator address whose rewards on this OperatorAddress should be included (e.g., an address controlled by the operator). Can be zero if not applicable.
	OperatorName          string         // A human-readable name for the operator (e.g., "NodeReal", "OperatorX").
}

// Config now holds a list of operators to process and the CC group ID.
type Config struct {
	Operators []OperatorRewardConfig // List of operator configurations.
	CCGroupID string
}

type CC struct {
	cfg       Config
	store     store.Store
	scheduler *gocron.Scheduler
}

func New(cfg Config, store store.Store) *CC {
	cc := &CC{
		cfg:       cfg,
		store:     store,
		scheduler: gocron.NewScheduler(time.UTC),
	}

	if _, err := cc.scheduler.Every(1).Month(1).At("10:00").Do(func() {
		cc.ComputeAndSend()
	}); err != nil {
		log.Debugw("error while setting up scheduler", "err", err)
	}

	return cc
}

func (c *CC) Start() {
	c.scheduler.StartAsync()
	// for test
	c.ComputeAndSend()
}

func (c *CC) ComputeAndSend() {
	var (
		ctx   = context.Background()
		year  = time.Now().Year()
		month = time.Now().Month()
		from  time.Time // the beginning of last month
		to    time.Time // the beginning of this month
	)

	if month == 1 {
		from = time.Date(year-1, 12, 1, 0, 0, 0, 0, time.UTC)
	} else {
		from = time.Date(year, month-1, 1, 0, 0, 0, 0, time.UTC)
	}
	to = time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	fromUnix := from.Unix()
	toUnix := to.Unix()

	log.Infow("[cc] Starting reward computation for period", "from", from, "to", to, "operators_count", len(c.cfg.Operators))

	for _, opCfg := range c.cfg.Operators {
		log.Infow("[cc] Processing operator", "name", opCfg.OperatorName, "address", opCfg.OperatorAddress.Hex())

		// 1. Calculate commission reward for the current operator
		opCommissionWei := decimal.Zero
		commissions, err := c.store.QueryBreathBlockRewardEvent(ctx, opCfg.OperatorAddress.Hex(), fromUnix, toUnix)
		if err != nil {
			log.Errorw("[cc] Failed to query commission events", "operator_name", opCfg.OperatorName, "operator_address", opCfg.OperatorAddress.Hex(), "err", err)
			continue // Skip to the next operator
		}
		for _, commission := range commissions {
			opCommissionWei = opCommissionWei.Add(commission.Commission)
		}
		// This is the commission amount to be reported and used in subtraction
		opCommissionEthForReport := opCommissionWei.Div(weiPerEth)

		// 2. Calculate total stake increase (operator's own + specific self-delegator's)
		totalStakeIncreaseWei := decimal.Zero
		totalNetDelegationWei := decimal.Zero

		// 2a. Operator's own self-delegation reward
		lastMonthOperatorStake, err := c.store.QueryDelegator(ctx, opCfg.OperatorAddress.Hex(),
			opCfg.OperatorAddress.Hex(), fromUnix)
		if err != nil {
			log.Errorw("[cc] Failed to query operator's own last month stake", "operator_name", opCfg.OperatorName, "delegator_address", opCfg.OperatorAddress.Hex(), "validator_address", opCfg.OperatorAddress.Hex(), "err", err)
			continue // Skip to the next operator
		}
		thisMonthOperatorStake, err := c.store.QueryDelegator(ctx, opCfg.OperatorAddress.Hex(),
			opCfg.OperatorAddress.Hex(), toUnix)
		if err != nil {
			log.Errorw("[cc] Failed to query operator's own this month stake", "operator_name", opCfg.OperatorName, "delegator_address", opCfg.OperatorAddress.Hex(), "validator_address", opCfg.OperatorAddress.Hex(), "err", err)
			continue // Skip to the next operator
		}
		operatorOwnStakeIncrease := thisMonthOperatorStake.Amount.Sub(lastMonthOperatorStake.Amount)
		totalStakeIncreaseWei = totalStakeIncreaseWei.Add(operatorOwnStakeIncrease)
		operatorNetDelegationWei, err := c.netDelegationChange(ctx, opCfg.OperatorAddress, opCfg.OperatorAddress, fromUnix, toUnix)
		if err != nil {
			log.Errorw("[cc] Failed to compute operator net delegation change", "operator_name", opCfg.OperatorName, "delegator_address", opCfg.OperatorAddress.Hex(), "err", err)
			continue
		}
		totalNetDelegationWei = totalNetDelegationWei.Add(operatorNetDelegationWei)

		// 2b. Reward from the specific "self-delegation" address, if configured
		if opCfg.SelfDelegationAddress != (common.Address{}) { // Check if a specific self-delegation address is provided
			lastMonthSpecialDelegatorStake, err := c.store.QueryDelegator(ctx, opCfg.SelfDelegationAddress.Hex(),
				opCfg.OperatorAddress.Hex(), fromUnix)
			if err != nil {
				// Log error but continue, as this part might be optional or fail independently
				log.Errorw("[cc] Failed to query special delegator's last month stake", "operator_name", opCfg.OperatorName, "delegator_address", opCfg.SelfDelegationAddress.Hex(), "validator_address", opCfg.OperatorAddress.Hex(), "err", err)
			} else {
				thisMonthSpecialDelegatorStake, err := c.store.QueryDelegator(ctx, opCfg.SelfDelegationAddress.Hex(),
					opCfg.OperatorAddress.Hex(), toUnix)
				if err != nil {
					log.Errorw("[cc] Failed to query special delegator's this month stake", "operator_name", opCfg.OperatorName, "delegator_address", opCfg.SelfDelegationAddress.Hex(), "validator_address", opCfg.OperatorAddress.Hex(), "err", err)
				} else {
					specialNetDelegationWei, err := c.netDelegationChange(ctx, opCfg.SelfDelegationAddress, opCfg.OperatorAddress, fromUnix, toUnix)
					if err != nil {
						log.Errorw("[cc] Failed to compute special delegator net delegation change", "operator_name", opCfg.OperatorName, "delegator_address", opCfg.SelfDelegationAddress.Hex(), "err", err)
						// Skip adding special delegator's stake increase if we can't get net delegation
					} else {
						specialDelegatorStakeIncrease := thisMonthSpecialDelegatorStake.Amount.Sub(lastMonthSpecialDelegatorStake.Amount)
						totalStakeIncreaseWei = totalStakeIncreaseWei.Add(specialDelegatorStakeIncrease)
						totalNetDelegationWei = totalNetDelegationWei.Add(specialNetDelegationWei)
					}
				}
			}
		}

		totalStakeIncreaseNetWei := totalStakeIncreaseWei.Sub(totalNetDelegationWei)
		totalStakeIncreaseEth := totalStakeIncreaseNetWei.Div(weiPerEth)

		// This calculation mirrors the original logic for "Self-delegation reward"
		finalSelfDelegationRewardForReport := totalStakeIncreaseEth.Sub(opCommissionEthForReport)

		msg := opCfg.OperatorName + " staking reward during " + strconv.Itoa(year) + "-" + from.Month().String() + ". \n" +
			opCfg.OperatorName + " commission reward: " + opCommissionEthForReport.String() + "\n" +
			opCfg.OperatorName + " Self-delegation reward: " + finalSelfDelegationRewardForReport.String()

		log.Infow("[cc] Sending CC message for operator", "operator_name", opCfg.OperatorName, "commission_reward", opCommissionEthForReport.String(), "self_delegation_reward", finalSelfDelegationRewardForReport.String())
		SendCCMessage(c.cfg.CCGroupID, msg)
	}
	log.Info("[cc] Finished reward computation cycle.")
}

/*
type Config struct {
	NodeRealOperator       common.Address
	NodeRealSelfDelegation common.Address

	CCGroupID string
}
*/

type Payload struct {
	GroupID string `json:"groupID"`
	MSG     string `json:"msg"`
}

func (c *CC) netDelegationChange(ctx context.Context, delegator common.Address, operator common.Address, fromTimestamp, toTimestamp int64) (decimal.Decimal, error) {
	if delegator == (common.Address{}) {
		return decimal.Zero, nil
	}

	txs, err := c.store.QueryDelegateTxs(ctx, delegator.Hex(), fromTimestamp, toTimestamp)
	if err != nil {
		return decimal.Zero, err
	}

	net := decimal.Zero
	operatorHex := operator.Hex()

	for _, tx := range txs {
		switch tx.Action {
		case model.Delegate:
			if tx.SrcOperator == operatorHex {
				net = net.Add(tx.Amount)
			}
		case model.UnDelegate:
			if tx.SrcOperator == operatorHex {
				net = net.Sub(tx.Amount)
			}
		case model.ReDelegate:
			if tx.SrcOperator == operatorHex {
				net = net.Sub(tx.Amount)
			}
			if tx.DstOperator == operatorHex {
				net = net.Add(tx.Amount)
			}
		}
	}

	return net, nil
}

func SendCCMessage(groupID string, msg string) {
	data := Payload{
		GroupID: groupID,
		MSG:     msg,
	}

	payloadBytes, err := jsoniter.Marshal(data)
	if err != nil {
		log.Errorf("[cc] data format is not supported, err=%s", err.Error())
		return
	}

	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Errorf("[cc] send message error, err=%s", err.Error())
		return
	}

	req.Header.Set("app", "ccalert")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Errorf("[cc] send message error, err=%s", err.Error())
		return
	}

	defer resp.Body.Close()
	response, _ := ioutil.ReadAll(resp.Body) // Error from ReadAll is ignored in original, kept same.

	log.Infof("[cc] send message success, message=%s, res=%s", msg, string(response))
}
