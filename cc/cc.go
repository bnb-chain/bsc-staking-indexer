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

	"github.com/bnb-chain/bsc-staking-indexer/store"
)

const url = "https://ccalert.bk.nodereal.cc/sendalerts"

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

	var (
		commissionReward     = decimal.Zero
		selfDelegationReward = decimal.Zero
	)

	// commission reward
	commissions, err := c.store.QueryBreathBlockRewardEvent(ctx, c.cfg.NodeRealOperator.Hex(), from.Unix(), to.Unix())
	if err != nil {
		log.Errorw("[cc] query commission error", "err", err)
		return
	}
	for _, commission := range commissions {
		commissionReward = commissionReward.Add(commission.Commission)
	}
	commissionReward = commissionReward.Div(decimal.NewFromInt(1e18))

	// self delegation reward of operator
	lastMonth, err := c.store.QueryDelegator(ctx, c.cfg.NodeRealOperator.Hex(),
		c.cfg.NodeRealOperator.Hex(), from.Unix())
	if err != nil {
		log.Errorw("[cc] query self delegation error", "err", err)
		return
	}

	thisMonth, err := c.store.QueryDelegator(ctx, c.cfg.NodeRealOperator.Hex(),
		c.cfg.NodeRealOperator.Hex(), to.Unix())
	if err != nil {
		log.Errorw("[cc] query self delegation error", "err", err)
		return
	}

	selfDelegationReward = thisMonth.Amount.Sub(lastMonth.Amount)

	// self delegation reward of delegator
	lastMonth, err = c.store.QueryDelegator(ctx, c.cfg.NodeRealSelfDelegation.Hex(),
		c.cfg.NodeRealOperator.Hex(), from.Unix())
	if err != nil {
		log.Errorw("[cc] query self delegation error", "err", err)
		return
	}

	thisMonth, err = c.store.QueryDelegator(ctx, c.cfg.NodeRealSelfDelegation.Hex(),
		c.cfg.NodeRealOperator.Hex(), to.Unix())
	if err != nil {
		log.Errorw("[cc] query self delegation error", "err", err)
		return
	}

	selfDelegationReward = selfDelegationReward.Add(thisMonth.Amount.Sub(lastMonth.Amount))
	selfDelegationReward = selfDelegationReward.Div(decimal.NewFromInt(1e18))
	selfDelegationReward = selfDelegationReward.Sub(commissionReward)

	msg := "Staking reward during " + strconv.Itoa(year) + "-" + from.Month().String() + ". \n" +
		"NodeReal commission reward: " + commissionReward.String() + "\n" +
		"NodeReal Self-delegation reward: " + selfDelegationReward.String()

	SendCCMessage(c.cfg.CCGroupID, msg)
}

type Config struct {
	NodeRealOperator       common.Address
	NodeRealSelfDelegation common.Address

	CCGroupID string
}

type Payload struct {
	GroupID string `json:"groupID"`
	MSG     string `json:"msg"`
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
	response, _ := ioutil.ReadAll(resp.Body)

	log.Infof("[cc] send message success, message=%s, res=%s", msg, string(response))
}
