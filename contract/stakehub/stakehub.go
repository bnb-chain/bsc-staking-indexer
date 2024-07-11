// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stakehub

import (
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// StakeHubCommission is an auto generated low-level Go binding around an user-defined struct.
type StakeHubCommission struct {
	Rate          uint64
	MaxRate       uint64
	MaxChangeRate uint64
}

// StakeHubDescription is an auto generated low-level Go binding around an user-defined struct.
type StakeHubDescription struct {
	Moniker  string
	Identity string
	Website  string
	Details  string
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"BC_FUSION_CHANNELID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"BREATHE_BLOCK_INTERVAL\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEAD_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"LOCK_AMOUNT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REDELEGATE_FEE_RATE_BASE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKING_CHANNELID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addToBlackList\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"blackList\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claim\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"requestNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimBatch\",\"inputs\":[{\"name\":\"operatorAddresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"requestNumbers\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"consensusExpiration\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"consensusToOperator\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createValidator\",\"inputs\":[{\"name\":\"consensusAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"voteAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blsProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"commission\",\"type\":\"tuple\",\"internalType\":\"structStakeHub.Commission\",\"components\":[{\"name\":\"rate\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxRate\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxChangeRate\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"description\",\"type\":\"tuple\",\"internalType\":\"structStakeHub.Description\",\"components\":[{\"name\":\"moniker\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"identity\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"website\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"details\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"delegate\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegateVotePower\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"distributeReward\",\"inputs\":[{\"name\":\"consensusAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"doubleSignSlash\",\"inputs\":[{\"name\":\"consensusAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"downtimeJailTime\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"downtimeSlash\",\"inputs\":[{\"name\":\"consensusAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"downtimeSlashAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"editCommissionRate\",\"inputs\":[{\"name\":\"commissionRate\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"editConsensusAddress\",\"inputs\":[{\"name\":\"newConsensusAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"editDescription\",\"inputs\":[{\"name\":\"description\",\"type\":\"tuple\",\"internalType\":\"structStakeHub.Description\",\"components\":[{\"name\":\"moniker\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"identity\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"website\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"details\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"editVoteAddress\",\"inputs\":[{\"name\":\"newVoteAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blsProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"felonyJailTime\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"felonySlashAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorBasicInfo\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"createdTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"jailed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"jailUntil\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorCommission\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structStakeHub.Commission\",\"components\":[{\"name\":\"rate\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxRate\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxChangeRate\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorConsensusAddress\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"consensusAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorCreditContract\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"creditContract\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorDescription\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structStakeHub.Description\",\"components\":[{\"name\":\"moniker\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"identity\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"website\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"details\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorElectionInfo\",\"inputs\":[{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"consensusAddrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"votingPowers\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"voteAddrs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"totalLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorRewardRecord\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorTotalPooledBNBRecord\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorVoteAddress\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"voteAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidators\",\"inputs\":[{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"operatorAddrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"creditAddrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"totalLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"handleAckPackage\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"msgBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"handleFailAckPackage\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"msgBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"handleSynPackage\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"msgBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isPaused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maliciousVoteSlash\",\"inputs\":[{\"name\":\"voteAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"maxElectedValidators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxFelonyBetweenBreatheBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minDelegationBNBChange\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minSelfDelegationBNB\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"numOfJailed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"redelegate\",\"inputs\":[{\"name\":\"srcValidator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dstValidator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateVotePower\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"redelegateFeeRate\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeFromBlackList\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resume\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"syncGovToken\",\"inputs\":[{\"name\":\"operatorAddresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferGasLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unbondPeriod\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"undelegate\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unjail\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateParam\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"voteExpiration\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"voteToOperator\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"BlackListed\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Claimed\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"delegator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"bnbAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CommissionRateEdited\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newCommissionRate\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConsensusAddressEdited\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newConsensusAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Delegated\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"delegator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"bnbAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DescriptionEdited\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MigrateFailed\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"delegator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"bnbAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"respCode\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumStakeHub.StakeMigrationRespCode\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MigrateSuccess\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"delegator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"bnbAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ParamChange\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProtectorChanged\",\"inputs\":[{\"name\":\"oldProtector\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newProtector\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Redelegated\",\"inputs\":[{\"name\":\"srcValidator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"dstValidator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"delegator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldShares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newShares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"bnbAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Resumed\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardDistributeFailed\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"failReason\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardDistributed\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"reward\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakeCreditInitialized\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"creditContract\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnBlackListed\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Undelegated\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"delegator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"bnbAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnexpectedPackage\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"msgBytes\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorCreated\",\"inputs\":[{\"name\":\"consensusAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"creditContract\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"voteAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorEmptyJailed\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorJailed\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorSlashed\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"jailUntil\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"slashAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"slashType\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumStakeHub.SlashType\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorUnjailed\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VoteAddressEdited\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newVoteAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadySlashed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ConsensusAddressExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DelegationAmountTooSmall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DuplicateConsensusAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DuplicateMoniker\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DuplicateVoteAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InBlackList\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCommission\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConsensusAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMoniker\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRequest\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSynPackage\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidValue\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidVoteAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"JailTimeNotExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoMoreFelonyAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCoinbase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyProtector\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlySelfDelegation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlySystemContract\",\"inputs\":[{\"name\":\"systemContract\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyZeroGasPrice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SameValidator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SelfDelegationNotEnough\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnknownParam\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"UpdateTooFrequently\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ValidatorExisted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ValidatorNotExisted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ValidatorNotJailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"VoteAddressExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroShares\",\"inputs\":[]}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// BCFUSIONCHANNELID is a free data retrieval call binding the contract method 0xf1fad104.
//
// Solidity: function BC_FUSION_CHANNELID() view returns(uint8)
func (_Contract *ContractCaller) BCFUSIONCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "BC_FUSION_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BCFUSIONCHANNELID is a free data retrieval call binding the contract method 0xf1fad104.
//
// Solidity: function BC_FUSION_CHANNELID() view returns(uint8)
func (_Contract *ContractSession) BCFUSIONCHANNELID() (uint8, error) {
	return _Contract.Contract.BCFUSIONCHANNELID(&_Contract.CallOpts)
}

// BCFUSIONCHANNELID is a free data retrieval call binding the contract method 0xf1fad104.
//
// Solidity: function BC_FUSION_CHANNELID() view returns(uint8)
func (_Contract *ContractCallerSession) BCFUSIONCHANNELID() (uint8, error) {
	return _Contract.Contract.BCFUSIONCHANNELID(&_Contract.CallOpts)
}

// BREATHEBLOCKINTERVAL is a free data retrieval call binding the contract method 0x1fa8882b.
//
// Solidity: function BREATHE_BLOCK_INTERVAL() view returns(uint256)
func (_Contract *ContractCaller) BREATHEBLOCKINTERVAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "BREATHE_BLOCK_INTERVAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BREATHEBLOCKINTERVAL is a free data retrieval call binding the contract method 0x1fa8882b.
//
// Solidity: function BREATHE_BLOCK_INTERVAL() view returns(uint256)
func (_Contract *ContractSession) BREATHEBLOCKINTERVAL() (*big.Int, error) {
	return _Contract.Contract.BREATHEBLOCKINTERVAL(&_Contract.CallOpts)
}

// BREATHEBLOCKINTERVAL is a free data retrieval call binding the contract method 0x1fa8882b.
//
// Solidity: function BREATHE_BLOCK_INTERVAL() view returns(uint256)
func (_Contract *ContractCallerSession) BREATHEBLOCKINTERVAL() (*big.Int, error) {
	return _Contract.Contract.BREATHEBLOCKINTERVAL(&_Contract.CallOpts)
}

// DEADADDRESS is a free data retrieval call binding the contract method 0x4e6fd6c4.
//
// Solidity: function DEAD_ADDRESS() view returns(address)
func (_Contract *ContractCaller) DEADADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "DEAD_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DEADADDRESS is a free data retrieval call binding the contract method 0x4e6fd6c4.
//
// Solidity: function DEAD_ADDRESS() view returns(address)
func (_Contract *ContractSession) DEADADDRESS() (common.Address, error) {
	return _Contract.Contract.DEADADDRESS(&_Contract.CallOpts)
}

// DEADADDRESS is a free data retrieval call binding the contract method 0x4e6fd6c4.
//
// Solidity: function DEAD_ADDRESS() view returns(address)
func (_Contract *ContractCallerSession) DEADADDRESS() (common.Address, error) {
	return _Contract.Contract.DEADADDRESS(&_Contract.CallOpts)
}

// LOCKAMOUNT is a free data retrieval call binding the contract method 0x8a4d3fa8.
//
// Solidity: function LOCK_AMOUNT() view returns(uint256)
func (_Contract *ContractCaller) LOCKAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "LOCK_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LOCKAMOUNT is a free data retrieval call binding the contract method 0x8a4d3fa8.
//
// Solidity: function LOCK_AMOUNT() view returns(uint256)
func (_Contract *ContractSession) LOCKAMOUNT() (*big.Int, error) {
	return _Contract.Contract.LOCKAMOUNT(&_Contract.CallOpts)
}

// LOCKAMOUNT is a free data retrieval call binding the contract method 0x8a4d3fa8.
//
// Solidity: function LOCK_AMOUNT() view returns(uint256)
func (_Contract *ContractCallerSession) LOCKAMOUNT() (*big.Int, error) {
	return _Contract.Contract.LOCKAMOUNT(&_Contract.CallOpts)
}

// REDELEGATEFEERATEBASE is a free data retrieval call binding the contract method 0xd115a206.
//
// Solidity: function REDELEGATE_FEE_RATE_BASE() view returns(uint256)
func (_Contract *ContractCaller) REDELEGATEFEERATEBASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "REDELEGATE_FEE_RATE_BASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REDELEGATEFEERATEBASE is a free data retrieval call binding the contract method 0xd115a206.
//
// Solidity: function REDELEGATE_FEE_RATE_BASE() view returns(uint256)
func (_Contract *ContractSession) REDELEGATEFEERATEBASE() (*big.Int, error) {
	return _Contract.Contract.REDELEGATEFEERATEBASE(&_Contract.CallOpts)
}

// REDELEGATEFEERATEBASE is a free data retrieval call binding the contract method 0xd115a206.
//
// Solidity: function REDELEGATE_FEE_RATE_BASE() view returns(uint256)
func (_Contract *ContractCallerSession) REDELEGATEFEERATEBASE() (*big.Int, error) {
	return _Contract.Contract.REDELEGATEFEERATEBASE(&_Contract.CallOpts)
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() view returns(uint8)
func (_Contract *ContractCaller) STAKINGCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "STAKING_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() view returns(uint8)
func (_Contract *ContractSession) STAKINGCHANNELID() (uint8, error) {
	return _Contract.Contract.STAKINGCHANNELID(&_Contract.CallOpts)
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() view returns(uint8)
func (_Contract *ContractCallerSession) STAKINGCHANNELID() (uint8, error) {
	return _Contract.Contract.STAKINGCHANNELID(&_Contract.CallOpts)
}

// BlackList is a free data retrieval call binding the contract method 0x4838d165.
//
// Solidity: function blackList(address ) view returns(bool)
func (_Contract *ContractCaller) BlackList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "blackList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BlackList is a free data retrieval call binding the contract method 0x4838d165.
//
// Solidity: function blackList(address ) view returns(bool)
func (_Contract *ContractSession) BlackList(arg0 common.Address) (bool, error) {
	return _Contract.Contract.BlackList(&_Contract.CallOpts, arg0)
}

// BlackList is a free data retrieval call binding the contract method 0x4838d165.
//
// Solidity: function blackList(address ) view returns(bool)
func (_Contract *ContractCallerSession) BlackList(arg0 common.Address) (bool, error) {
	return _Contract.Contract.BlackList(&_Contract.CallOpts, arg0)
}

// ConsensusExpiration is a free data retrieval call binding the contract method 0x663706d3.
//
// Solidity: function consensusExpiration(address ) view returns(uint256)
func (_Contract *ContractCaller) ConsensusExpiration(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "consensusExpiration", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConsensusExpiration is a free data retrieval call binding the contract method 0x663706d3.
//
// Solidity: function consensusExpiration(address ) view returns(uint256)
func (_Contract *ContractSession) ConsensusExpiration(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.ConsensusExpiration(&_Contract.CallOpts, arg0)
}

// ConsensusExpiration is a free data retrieval call binding the contract method 0x663706d3.
//
// Solidity: function consensusExpiration(address ) view returns(uint256)
func (_Contract *ContractCallerSession) ConsensusExpiration(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.ConsensusExpiration(&_Contract.CallOpts, arg0)
}

// ConsensusToOperator is a free data retrieval call binding the contract method 0x86d54506.
//
// Solidity: function consensusToOperator(address ) view returns(address)
func (_Contract *ContractCaller) ConsensusToOperator(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "consensusToOperator", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConsensusToOperator is a free data retrieval call binding the contract method 0x86d54506.
//
// Solidity: function consensusToOperator(address ) view returns(address)
func (_Contract *ContractSession) ConsensusToOperator(arg0 common.Address) (common.Address, error) {
	return _Contract.Contract.ConsensusToOperator(&_Contract.CallOpts, arg0)
}

// ConsensusToOperator is a free data retrieval call binding the contract method 0x86d54506.
//
// Solidity: function consensusToOperator(address ) view returns(address)
func (_Contract *ContractCallerSession) ConsensusToOperator(arg0 common.Address) (common.Address, error) {
	return _Contract.Contract.ConsensusToOperator(&_Contract.CallOpts, arg0)
}

// DowntimeJailTime is a free data retrieval call binding the contract method 0x76e7d6d6.
//
// Solidity: function downtimeJailTime() view returns(uint256)
func (_Contract *ContractCaller) DowntimeJailTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "downtimeJailTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DowntimeJailTime is a free data retrieval call binding the contract method 0x76e7d6d6.
//
// Solidity: function downtimeJailTime() view returns(uint256)
func (_Contract *ContractSession) DowntimeJailTime() (*big.Int, error) {
	return _Contract.Contract.DowntimeJailTime(&_Contract.CallOpts)
}

// DowntimeJailTime is a free data retrieval call binding the contract method 0x76e7d6d6.
//
// Solidity: function downtimeJailTime() view returns(uint256)
func (_Contract *ContractCallerSession) DowntimeJailTime() (*big.Int, error) {
	return _Contract.Contract.DowntimeJailTime(&_Contract.CallOpts)
}

// DowntimeSlashAmount is a free data retrieval call binding the contract method 0xd8ca511f.
//
// Solidity: function downtimeSlashAmount() view returns(uint256)
func (_Contract *ContractCaller) DowntimeSlashAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "downtimeSlashAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DowntimeSlashAmount is a free data retrieval call binding the contract method 0xd8ca511f.
//
// Solidity: function downtimeSlashAmount() view returns(uint256)
func (_Contract *ContractSession) DowntimeSlashAmount() (*big.Int, error) {
	return _Contract.Contract.DowntimeSlashAmount(&_Contract.CallOpts)
}

// DowntimeSlashAmount is a free data retrieval call binding the contract method 0xd8ca511f.
//
// Solidity: function downtimeSlashAmount() view returns(uint256)
func (_Contract *ContractCallerSession) DowntimeSlashAmount() (*big.Int, error) {
	return _Contract.Contract.DowntimeSlashAmount(&_Contract.CallOpts)
}

// FelonyJailTime is a free data retrieval call binding the contract method 0xf1f74d84.
//
// Solidity: function felonyJailTime() view returns(uint256)
func (_Contract *ContractCaller) FelonyJailTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "felonyJailTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FelonyJailTime is a free data retrieval call binding the contract method 0xf1f74d84.
//
// Solidity: function felonyJailTime() view returns(uint256)
func (_Contract *ContractSession) FelonyJailTime() (*big.Int, error) {
	return _Contract.Contract.FelonyJailTime(&_Contract.CallOpts)
}

// FelonyJailTime is a free data retrieval call binding the contract method 0xf1f74d84.
//
// Solidity: function felonyJailTime() view returns(uint256)
func (_Contract *ContractCallerSession) FelonyJailTime() (*big.Int, error) {
	return _Contract.Contract.FelonyJailTime(&_Contract.CallOpts)
}

// FelonySlashAmount is a free data retrieval call binding the contract method 0xbdceadf3.
//
// Solidity: function felonySlashAmount() view returns(uint256)
func (_Contract *ContractCaller) FelonySlashAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "felonySlashAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FelonySlashAmount is a free data retrieval call binding the contract method 0xbdceadf3.
//
// Solidity: function felonySlashAmount() view returns(uint256)
func (_Contract *ContractSession) FelonySlashAmount() (*big.Int, error) {
	return _Contract.Contract.FelonySlashAmount(&_Contract.CallOpts)
}

// FelonySlashAmount is a free data retrieval call binding the contract method 0xbdceadf3.
//
// Solidity: function felonySlashAmount() view returns(uint256)
func (_Contract *ContractCallerSession) FelonySlashAmount() (*big.Int, error) {
	return _Contract.Contract.FelonySlashAmount(&_Contract.CallOpts)
}

// GetValidatorBasicInfo is a free data retrieval call binding the contract method 0xcbb04d9d.
//
// Solidity: function getValidatorBasicInfo(address operatorAddress) view returns(uint256 createdTime, bool jailed, uint256 jailUntil)
func (_Contract *ContractCaller) GetValidatorBasicInfo(opts *bind.CallOpts, operatorAddress common.Address) (struct {
	CreatedTime *big.Int
	Jailed      bool
	JailUntil   *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getValidatorBasicInfo", operatorAddress)

	outstruct := new(struct {
		CreatedTime *big.Int
		Jailed      bool
		JailUntil   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CreatedTime = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Jailed = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.JailUntil = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetValidatorBasicInfo is a free data retrieval call binding the contract method 0xcbb04d9d.
//
// Solidity: function getValidatorBasicInfo(address operatorAddress) view returns(uint256 createdTime, bool jailed, uint256 jailUntil)
func (_Contract *ContractSession) GetValidatorBasicInfo(operatorAddress common.Address) (struct {
	CreatedTime *big.Int
	Jailed      bool
	JailUntil   *big.Int
}, error) {
	return _Contract.Contract.GetValidatorBasicInfo(&_Contract.CallOpts, operatorAddress)
}

// GetValidatorBasicInfo is a free data retrieval call binding the contract method 0xcbb04d9d.
//
// Solidity: function getValidatorBasicInfo(address operatorAddress) view returns(uint256 createdTime, bool jailed, uint256 jailUntil)
func (_Contract *ContractCallerSession) GetValidatorBasicInfo(operatorAddress common.Address) (struct {
	CreatedTime *big.Int
	Jailed      bool
	JailUntil   *big.Int
}, error) {
	return _Contract.Contract.GetValidatorBasicInfo(&_Contract.CallOpts, operatorAddress)
}

// GetValidatorCommission is a free data retrieval call binding the contract method 0x6ec01b27.
//
// Solidity: function getValidatorCommission(address operatorAddress) view returns((uint64,uint64,uint64))
func (_Contract *ContractCaller) GetValidatorCommission(opts *bind.CallOpts, operatorAddress common.Address) (StakeHubCommission, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getValidatorCommission", operatorAddress)

	if err != nil {
		return *new(StakeHubCommission), err
	}

	out0 := *abi.ConvertType(out[0], new(StakeHubCommission)).(*StakeHubCommission)

	return out0, err

}

// GetValidatorCommission is a free data retrieval call binding the contract method 0x6ec01b27.
//
// Solidity: function getValidatorCommission(address operatorAddress) view returns((uint64,uint64,uint64))
func (_Contract *ContractSession) GetValidatorCommission(operatorAddress common.Address) (StakeHubCommission, error) {
	return _Contract.Contract.GetValidatorCommission(&_Contract.CallOpts, operatorAddress)
}

// GetValidatorCommission is a free data retrieval call binding the contract method 0x6ec01b27.
//
// Solidity: function getValidatorCommission(address operatorAddress) view returns((uint64,uint64,uint64))
func (_Contract *ContractCallerSession) GetValidatorCommission(operatorAddress common.Address) (StakeHubCommission, error) {
	return _Contract.Contract.GetValidatorCommission(&_Contract.CallOpts, operatorAddress)
}

// GetValidatorConsensusAddress is a free data retrieval call binding the contract method 0x059ddd22.
//
// Solidity: function getValidatorConsensusAddress(address operatorAddress) view returns(address consensusAddress)
func (_Contract *ContractCaller) GetValidatorConsensusAddress(opts *bind.CallOpts, operatorAddress common.Address) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getValidatorConsensusAddress", operatorAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetValidatorConsensusAddress is a free data retrieval call binding the contract method 0x059ddd22.
//
// Solidity: function getValidatorConsensusAddress(address operatorAddress) view returns(address consensusAddress)
func (_Contract *ContractSession) GetValidatorConsensusAddress(operatorAddress common.Address) (common.Address, error) {
	return _Contract.Contract.GetValidatorConsensusAddress(&_Contract.CallOpts, operatorAddress)
}

// GetValidatorConsensusAddress is a free data retrieval call binding the contract method 0x059ddd22.
//
// Solidity: function getValidatorConsensusAddress(address operatorAddress) view returns(address consensusAddress)
func (_Contract *ContractCallerSession) GetValidatorConsensusAddress(operatorAddress common.Address) (common.Address, error) {
	return _Contract.Contract.GetValidatorConsensusAddress(&_Contract.CallOpts, operatorAddress)
}

// GetValidatorCreditContract is a free data retrieval call binding the contract method 0xdbda7fb3.
//
// Solidity: function getValidatorCreditContract(address operatorAddress) view returns(address creditContract)
func (_Contract *ContractCaller) GetValidatorCreditContract(opts *bind.CallOpts, operatorAddress common.Address) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getValidatorCreditContract", operatorAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetValidatorCreditContract is a free data retrieval call binding the contract method 0xdbda7fb3.
//
// Solidity: function getValidatorCreditContract(address operatorAddress) view returns(address creditContract)
func (_Contract *ContractSession) GetValidatorCreditContract(operatorAddress common.Address) (common.Address, error) {
	return _Contract.Contract.GetValidatorCreditContract(&_Contract.CallOpts, operatorAddress)
}

// GetValidatorCreditContract is a free data retrieval call binding the contract method 0xdbda7fb3.
//
// Solidity: function getValidatorCreditContract(address operatorAddress) view returns(address creditContract)
func (_Contract *ContractCallerSession) GetValidatorCreditContract(operatorAddress common.Address) (common.Address, error) {
	return _Contract.Contract.GetValidatorCreditContract(&_Contract.CallOpts, operatorAddress)
}

// GetValidatorDescription is a free data retrieval call binding the contract method 0xa43569b3.
//
// Solidity: function getValidatorDescription(address operatorAddress) view returns((string,string,string,string))
func (_Contract *ContractCaller) GetValidatorDescription(opts *bind.CallOpts, operatorAddress common.Address) (StakeHubDescription, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getValidatorDescription", operatorAddress)

	if err != nil {
		return *new(StakeHubDescription), err
	}

	out0 := *abi.ConvertType(out[0], new(StakeHubDescription)).(*StakeHubDescription)

	return out0, err

}

// GetValidatorDescription is a free data retrieval call binding the contract method 0xa43569b3.
//
// Solidity: function getValidatorDescription(address operatorAddress) view returns((string,string,string,string))
func (_Contract *ContractSession) GetValidatorDescription(operatorAddress common.Address) (StakeHubDescription, error) {
	return _Contract.Contract.GetValidatorDescription(&_Contract.CallOpts, operatorAddress)
}

// GetValidatorDescription is a free data retrieval call binding the contract method 0xa43569b3.
//
// Solidity: function getValidatorDescription(address operatorAddress) view returns((string,string,string,string))
func (_Contract *ContractCallerSession) GetValidatorDescription(operatorAddress common.Address) (StakeHubDescription, error) {
	return _Contract.Contract.GetValidatorDescription(&_Contract.CallOpts, operatorAddress)
}

// GetValidatorElectionInfo is a free data retrieval call binding the contract method 0x63a036b5.
//
// Solidity: function getValidatorElectionInfo(uint256 offset, uint256 limit) view returns(address[] consensusAddrs, uint256[] votingPowers, bytes[] voteAddrs, uint256 totalLength)
func (_Contract *ContractCaller) GetValidatorElectionInfo(opts *bind.CallOpts, offset *big.Int, limit *big.Int) (struct {
	ConsensusAddrs []common.Address
	VotingPowers   []*big.Int
	VoteAddrs      [][]byte
	TotalLength    *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getValidatorElectionInfo", offset, limit)

	outstruct := new(struct {
		ConsensusAddrs []common.Address
		VotingPowers   []*big.Int
		VoteAddrs      [][]byte
		TotalLength    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConsensusAddrs = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.VotingPowers = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.VoteAddrs = *abi.ConvertType(out[2], new([][]byte)).(*[][]byte)
	outstruct.TotalLength = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetValidatorElectionInfo is a free data retrieval call binding the contract method 0x63a036b5.
//
// Solidity: function getValidatorElectionInfo(uint256 offset, uint256 limit) view returns(address[] consensusAddrs, uint256[] votingPowers, bytes[] voteAddrs, uint256 totalLength)
func (_Contract *ContractSession) GetValidatorElectionInfo(offset *big.Int, limit *big.Int) (struct {
	ConsensusAddrs []common.Address
	VotingPowers   []*big.Int
	VoteAddrs      [][]byte
	TotalLength    *big.Int
}, error) {
	return _Contract.Contract.GetValidatorElectionInfo(&_Contract.CallOpts, offset, limit)
}

// GetValidatorElectionInfo is a free data retrieval call binding the contract method 0x63a036b5.
//
// Solidity: function getValidatorElectionInfo(uint256 offset, uint256 limit) view returns(address[] consensusAddrs, uint256[] votingPowers, bytes[] voteAddrs, uint256 totalLength)
func (_Contract *ContractCallerSession) GetValidatorElectionInfo(offset *big.Int, limit *big.Int) (struct {
	ConsensusAddrs []common.Address
	VotingPowers   []*big.Int
	VoteAddrs      [][]byte
	TotalLength    *big.Int
}, error) {
	return _Contract.Contract.GetValidatorElectionInfo(&_Contract.CallOpts, offset, limit)
}

// GetValidatorRewardRecord is a free data retrieval call binding the contract method 0xf80a3402.
//
// Solidity: function getValidatorRewardRecord(address operatorAddress, uint256 index) view returns(uint256)
func (_Contract *ContractCaller) GetValidatorRewardRecord(opts *bind.CallOpts, operatorAddress common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getValidatorRewardRecord", operatorAddress, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidatorRewardRecord is a free data retrieval call binding the contract method 0xf80a3402.
//
// Solidity: function getValidatorRewardRecord(address operatorAddress, uint256 index) view returns(uint256)
func (_Contract *ContractSession) GetValidatorRewardRecord(operatorAddress common.Address, index *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetValidatorRewardRecord(&_Contract.CallOpts, operatorAddress, index)
}

// GetValidatorRewardRecord is a free data retrieval call binding the contract method 0xf80a3402.
//
// Solidity: function getValidatorRewardRecord(address operatorAddress, uint256 index) view returns(uint256)
func (_Contract *ContractCallerSession) GetValidatorRewardRecord(operatorAddress common.Address, index *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetValidatorRewardRecord(&_Contract.CallOpts, operatorAddress, index)
}

// GetValidatorTotalPooledBNBRecord is a free data retrieval call binding the contract method 0x8cd22b22.
//
// Solidity: function getValidatorTotalPooledBNBRecord(address operatorAddress, uint256 index) view returns(uint256)
func (_Contract *ContractCaller) GetValidatorTotalPooledBNBRecord(opts *bind.CallOpts, operatorAddress common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getValidatorTotalPooledBNBRecord", operatorAddress, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidatorTotalPooledBNBRecord is a free data retrieval call binding the contract method 0x8cd22b22.
//
// Solidity: function getValidatorTotalPooledBNBRecord(address operatorAddress, uint256 index) view returns(uint256)
func (_Contract *ContractSession) GetValidatorTotalPooledBNBRecord(operatorAddress common.Address, index *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetValidatorTotalPooledBNBRecord(&_Contract.CallOpts, operatorAddress, index)
}

// GetValidatorTotalPooledBNBRecord is a free data retrieval call binding the contract method 0x8cd22b22.
//
// Solidity: function getValidatorTotalPooledBNBRecord(address operatorAddress, uint256 index) view returns(uint256)
func (_Contract *ContractCallerSession) GetValidatorTotalPooledBNBRecord(operatorAddress common.Address, index *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetValidatorTotalPooledBNBRecord(&_Contract.CallOpts, operatorAddress, index)
}

// GetValidatorVoteAddress is a free data retrieval call binding the contract method 0x6f8e2fa4.
//
// Solidity: function getValidatorVoteAddress(address operatorAddress) view returns(bytes voteAddress)
func (_Contract *ContractCaller) GetValidatorVoteAddress(opts *bind.CallOpts, operatorAddress common.Address) ([]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getValidatorVoteAddress", operatorAddress)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetValidatorVoteAddress is a free data retrieval call binding the contract method 0x6f8e2fa4.
//
// Solidity: function getValidatorVoteAddress(address operatorAddress) view returns(bytes voteAddress)
func (_Contract *ContractSession) GetValidatorVoteAddress(operatorAddress common.Address) ([]byte, error) {
	return _Contract.Contract.GetValidatorVoteAddress(&_Contract.CallOpts, operatorAddress)
}

// GetValidatorVoteAddress is a free data retrieval call binding the contract method 0x6f8e2fa4.
//
// Solidity: function getValidatorVoteAddress(address operatorAddress) view returns(bytes voteAddress)
func (_Contract *ContractCallerSession) GetValidatorVoteAddress(operatorAddress common.Address) ([]byte, error) {
	return _Contract.Contract.GetValidatorVoteAddress(&_Contract.CallOpts, operatorAddress)
}

// GetValidators is a free data retrieval call binding the contract method 0xbff02e20.
//
// Solidity: function getValidators(uint256 offset, uint256 limit) view returns(address[] operatorAddrs, address[] creditAddrs, uint256 totalLength)
func (_Contract *ContractCaller) GetValidators(opts *bind.CallOpts, offset *big.Int, limit *big.Int) (struct {
	OperatorAddrs []common.Address
	CreditAddrs   []common.Address
	TotalLength   *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getValidators", offset, limit)

	outstruct := new(struct {
		OperatorAddrs []common.Address
		CreditAddrs   []common.Address
		TotalLength   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OperatorAddrs = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.CreditAddrs = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.TotalLength = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetValidators is a free data retrieval call binding the contract method 0xbff02e20.
//
// Solidity: function getValidators(uint256 offset, uint256 limit) view returns(address[] operatorAddrs, address[] creditAddrs, uint256 totalLength)
func (_Contract *ContractSession) GetValidators(offset *big.Int, limit *big.Int) (struct {
	OperatorAddrs []common.Address
	CreditAddrs   []common.Address
	TotalLength   *big.Int
}, error) {
	return _Contract.Contract.GetValidators(&_Contract.CallOpts, offset, limit)
}

// GetValidators is a free data retrieval call binding the contract method 0xbff02e20.
//
// Solidity: function getValidators(uint256 offset, uint256 limit) view returns(address[] operatorAddrs, address[] creditAddrs, uint256 totalLength)
func (_Contract *ContractCallerSession) GetValidators(offset *big.Int, limit *big.Int) (struct {
	OperatorAddrs []common.Address
	CreditAddrs   []common.Address
	TotalLength   *big.Int
}, error) {
	return _Contract.Contract.GetValidators(&_Contract.CallOpts, offset, limit)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Contract *ContractCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Contract *ContractSession) IsPaused() (bool, error) {
	return _Contract.Contract.IsPaused(&_Contract.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Contract *ContractCallerSession) IsPaused() (bool, error) {
	return _Contract.Contract.IsPaused(&_Contract.CallOpts)
}

// MaxElectedValidators is a free data retrieval call binding the contract method 0xc473318f.
//
// Solidity: function maxElectedValidators() view returns(uint256)
func (_Contract *ContractCaller) MaxElectedValidators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "maxElectedValidators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxElectedValidators is a free data retrieval call binding the contract method 0xc473318f.
//
// Solidity: function maxElectedValidators() view returns(uint256)
func (_Contract *ContractSession) MaxElectedValidators() (*big.Int, error) {
	return _Contract.Contract.MaxElectedValidators(&_Contract.CallOpts)
}

// MaxElectedValidators is a free data retrieval call binding the contract method 0xc473318f.
//
// Solidity: function maxElectedValidators() view returns(uint256)
func (_Contract *ContractCallerSession) MaxElectedValidators() (*big.Int, error) {
	return _Contract.Contract.MaxElectedValidators(&_Contract.CallOpts)
}

// MaxFelonyBetweenBreatheBlock is a free data retrieval call binding the contract method 0xff69ab61.
//
// Solidity: function maxFelonyBetweenBreatheBlock() view returns(uint256)
func (_Contract *ContractCaller) MaxFelonyBetweenBreatheBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "maxFelonyBetweenBreatheBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxFelonyBetweenBreatheBlock is a free data retrieval call binding the contract method 0xff69ab61.
//
// Solidity: function maxFelonyBetweenBreatheBlock() view returns(uint256)
func (_Contract *ContractSession) MaxFelonyBetweenBreatheBlock() (*big.Int, error) {
	return _Contract.Contract.MaxFelonyBetweenBreatheBlock(&_Contract.CallOpts)
}

// MaxFelonyBetweenBreatheBlock is a free data retrieval call binding the contract method 0xff69ab61.
//
// Solidity: function maxFelonyBetweenBreatheBlock() view returns(uint256)
func (_Contract *ContractCallerSession) MaxFelonyBetweenBreatheBlock() (*big.Int, error) {
	return _Contract.Contract.MaxFelonyBetweenBreatheBlock(&_Contract.CallOpts)
}

// MinDelegationBNBChange is a free data retrieval call binding the contract method 0x38409988.
//
// Solidity: function minDelegationBNBChange() view returns(uint256)
func (_Contract *ContractCaller) MinDelegationBNBChange(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "minDelegationBNBChange")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDelegationBNBChange is a free data retrieval call binding the contract method 0x38409988.
//
// Solidity: function minDelegationBNBChange() view returns(uint256)
func (_Contract *ContractSession) MinDelegationBNBChange() (*big.Int, error) {
	return _Contract.Contract.MinDelegationBNBChange(&_Contract.CallOpts)
}

// MinDelegationBNBChange is a free data retrieval call binding the contract method 0x38409988.
//
// Solidity: function minDelegationBNBChange() view returns(uint256)
func (_Contract *ContractCallerSession) MinDelegationBNBChange() (*big.Int, error) {
	return _Contract.Contract.MinDelegationBNBChange(&_Contract.CallOpts)
}

// MinSelfDelegationBNB is a free data retrieval call binding the contract method 0x0661806e.
//
// Solidity: function minSelfDelegationBNB() view returns(uint256)
func (_Contract *ContractCaller) MinSelfDelegationBNB(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "minSelfDelegationBNB")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinSelfDelegationBNB is a free data retrieval call binding the contract method 0x0661806e.
//
// Solidity: function minSelfDelegationBNB() view returns(uint256)
func (_Contract *ContractSession) MinSelfDelegationBNB() (*big.Int, error) {
	return _Contract.Contract.MinSelfDelegationBNB(&_Contract.CallOpts)
}

// MinSelfDelegationBNB is a free data retrieval call binding the contract method 0x0661806e.
//
// Solidity: function minSelfDelegationBNB() view returns(uint256)
func (_Contract *ContractCallerSession) MinSelfDelegationBNB() (*big.Int, error) {
	return _Contract.Contract.MinSelfDelegationBNB(&_Contract.CallOpts)
}

// NumOfJailed is a free data retrieval call binding the contract method 0xdaacdb66.
//
// Solidity: function numOfJailed() view returns(uint256)
func (_Contract *ContractCaller) NumOfJailed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "numOfJailed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumOfJailed is a free data retrieval call binding the contract method 0xdaacdb66.
//
// Solidity: function numOfJailed() view returns(uint256)
func (_Contract *ContractSession) NumOfJailed() (*big.Int, error) {
	return _Contract.Contract.NumOfJailed(&_Contract.CallOpts)
}

// NumOfJailed is a free data retrieval call binding the contract method 0xdaacdb66.
//
// Solidity: function numOfJailed() view returns(uint256)
func (_Contract *ContractCallerSession) NumOfJailed() (*big.Int, error) {
	return _Contract.Contract.NumOfJailed(&_Contract.CallOpts)
}

// RedelegateFeeRate is a free data retrieval call binding the contract method 0xe992aaf5.
//
// Solidity: function redelegateFeeRate() view returns(uint256)
func (_Contract *ContractCaller) RedelegateFeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "redelegateFeeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RedelegateFeeRate is a free data retrieval call binding the contract method 0xe992aaf5.
//
// Solidity: function redelegateFeeRate() view returns(uint256)
func (_Contract *ContractSession) RedelegateFeeRate() (*big.Int, error) {
	return _Contract.Contract.RedelegateFeeRate(&_Contract.CallOpts)
}

// RedelegateFeeRate is a free data retrieval call binding the contract method 0xe992aaf5.
//
// Solidity: function redelegateFeeRate() view returns(uint256)
func (_Contract *ContractCallerSession) RedelegateFeeRate() (*big.Int, error) {
	return _Contract.Contract.RedelegateFeeRate(&_Contract.CallOpts)
}

// TransferGasLimit is a free data retrieval call binding the contract method 0xe8f67c3b.
//
// Solidity: function transferGasLimit() view returns(uint256)
func (_Contract *ContractCaller) TransferGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "transferGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransferGasLimit is a free data retrieval call binding the contract method 0xe8f67c3b.
//
// Solidity: function transferGasLimit() view returns(uint256)
func (_Contract *ContractSession) TransferGasLimit() (*big.Int, error) {
	return _Contract.Contract.TransferGasLimit(&_Contract.CallOpts)
}

// TransferGasLimit is a free data retrieval call binding the contract method 0xe8f67c3b.
//
// Solidity: function transferGasLimit() view returns(uint256)
func (_Contract *ContractCallerSession) TransferGasLimit() (*big.Int, error) {
	return _Contract.Contract.TransferGasLimit(&_Contract.CallOpts)
}

// UnbondPeriod is a free data retrieval call binding the contract method 0xfc0c5ff1.
//
// Solidity: function unbondPeriod() view returns(uint256)
func (_Contract *ContractCaller) UnbondPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "unbondPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnbondPeriod is a free data retrieval call binding the contract method 0xfc0c5ff1.
//
// Solidity: function unbondPeriod() view returns(uint256)
func (_Contract *ContractSession) UnbondPeriod() (*big.Int, error) {
	return _Contract.Contract.UnbondPeriod(&_Contract.CallOpts)
}

// UnbondPeriod is a free data retrieval call binding the contract method 0xfc0c5ff1.
//
// Solidity: function unbondPeriod() view returns(uint256)
func (_Contract *ContractCallerSession) UnbondPeriod() (*big.Int, error) {
	return _Contract.Contract.UnbondPeriod(&_Contract.CallOpts)
}

// VoteExpiration is a free data retrieval call binding the contract method 0xefdbf0e1.
//
// Solidity: function voteExpiration(bytes ) view returns(uint256)
func (_Contract *ContractCaller) VoteExpiration(opts *bind.CallOpts, arg0 []byte) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "voteExpiration", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VoteExpiration is a free data retrieval call binding the contract method 0xefdbf0e1.
//
// Solidity: function voteExpiration(bytes ) view returns(uint256)
func (_Contract *ContractSession) VoteExpiration(arg0 []byte) (*big.Int, error) {
	return _Contract.Contract.VoteExpiration(&_Contract.CallOpts, arg0)
}

// VoteExpiration is a free data retrieval call binding the contract method 0xefdbf0e1.
//
// Solidity: function voteExpiration(bytes ) view returns(uint256)
func (_Contract *ContractCallerSession) VoteExpiration(arg0 []byte) (*big.Int, error) {
	return _Contract.Contract.VoteExpiration(&_Contract.CallOpts, arg0)
}

// VoteToOperator is a free data retrieval call binding the contract method 0x17b4f353.
//
// Solidity: function voteToOperator(bytes ) view returns(address)
func (_Contract *ContractCaller) VoteToOperator(opts *bind.CallOpts, arg0 []byte) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "voteToOperator", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VoteToOperator is a free data retrieval call binding the contract method 0x17b4f353.
//
// Solidity: function voteToOperator(bytes ) view returns(address)
func (_Contract *ContractSession) VoteToOperator(arg0 []byte) (common.Address, error) {
	return _Contract.Contract.VoteToOperator(&_Contract.CallOpts, arg0)
}

// VoteToOperator is a free data retrieval call binding the contract method 0x17b4f353.
//
// Solidity: function voteToOperator(bytes ) view returns(address)
func (_Contract *ContractCallerSession) VoteToOperator(arg0 []byte) (common.Address, error) {
	return _Contract.Contract.VoteToOperator(&_Contract.CallOpts, arg0)
}

// AddToBlackList is a paid mutator transaction binding the contract method 0x417c73a7.
//
// Solidity: function addToBlackList(address account) returns()
func (_Contract *ContractTransactor) AddToBlackList(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "addToBlackList", account)
}

// AddToBlackList is a paid mutator transaction binding the contract method 0x417c73a7.
//
// Solidity: function addToBlackList(address account) returns()
func (_Contract *ContractSession) AddToBlackList(account common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AddToBlackList(&_Contract.TransactOpts, account)
}

// AddToBlackList is a paid mutator transaction binding the contract method 0x417c73a7.
//
// Solidity: function addToBlackList(address account) returns()
func (_Contract *ContractTransactorSession) AddToBlackList(account common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AddToBlackList(&_Contract.TransactOpts, account)
}

// Claim is a paid mutator transaction binding the contract method 0xaad3ec96.
//
// Solidity: function claim(address operatorAddress, uint256 requestNumber) returns()
func (_Contract *ContractTransactor) Claim(opts *bind.TransactOpts, operatorAddress common.Address, requestNumber *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "claim", operatorAddress, requestNumber)
}

// Claim is a paid mutator transaction binding the contract method 0xaad3ec96.
//
// Solidity: function claim(address operatorAddress, uint256 requestNumber) returns()
func (_Contract *ContractSession) Claim(operatorAddress common.Address, requestNumber *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Claim(&_Contract.TransactOpts, operatorAddress, requestNumber)
}

// Claim is a paid mutator transaction binding the contract method 0xaad3ec96.
//
// Solidity: function claim(address operatorAddress, uint256 requestNumber) returns()
func (_Contract *ContractTransactorSession) Claim(operatorAddress common.Address, requestNumber *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Claim(&_Contract.TransactOpts, operatorAddress, requestNumber)
}

// ClaimBatch is a paid mutator transaction binding the contract method 0xd7c2dfc8.
//
// Solidity: function claimBatch(address[] operatorAddresses, uint256[] requestNumbers) returns()
func (_Contract *ContractTransactor) ClaimBatch(opts *bind.TransactOpts, operatorAddresses []common.Address, requestNumbers []*big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "claimBatch", operatorAddresses, requestNumbers)
}

// ClaimBatch is a paid mutator transaction binding the contract method 0xd7c2dfc8.
//
// Solidity: function claimBatch(address[] operatorAddresses, uint256[] requestNumbers) returns()
func (_Contract *ContractSession) ClaimBatch(operatorAddresses []common.Address, requestNumbers []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ClaimBatch(&_Contract.TransactOpts, operatorAddresses, requestNumbers)
}

// ClaimBatch is a paid mutator transaction binding the contract method 0xd7c2dfc8.
//
// Solidity: function claimBatch(address[] operatorAddresses, uint256[] requestNumbers) returns()
func (_Contract *ContractTransactorSession) ClaimBatch(operatorAddresses []common.Address, requestNumbers []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ClaimBatch(&_Contract.TransactOpts, operatorAddresses, requestNumbers)
}

// CreateValidator is a paid mutator transaction binding the contract method 0x64028fbd.
//
// Solidity: function createValidator(address consensusAddress, bytes voteAddress, bytes blsProof, (uint64,uint64,uint64) commission, (string,string,string,string) description) payable returns()
func (_Contract *ContractTransactor) CreateValidator(opts *bind.TransactOpts, consensusAddress common.Address, voteAddress []byte, blsProof []byte, commission StakeHubCommission, description StakeHubDescription) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "createValidator", consensusAddress, voteAddress, blsProof, commission, description)
}

// CreateValidator is a paid mutator transaction binding the contract method 0x64028fbd.
//
// Solidity: function createValidator(address consensusAddress, bytes voteAddress, bytes blsProof, (uint64,uint64,uint64) commission, (string,string,string,string) description) payable returns()
func (_Contract *ContractSession) CreateValidator(consensusAddress common.Address, voteAddress []byte, blsProof []byte, commission StakeHubCommission, description StakeHubDescription) (*types.Transaction, error) {
	return _Contract.Contract.CreateValidator(&_Contract.TransactOpts, consensusAddress, voteAddress, blsProof, commission, description)
}

// CreateValidator is a paid mutator transaction binding the contract method 0x64028fbd.
//
// Solidity: function createValidator(address consensusAddress, bytes voteAddress, bytes blsProof, (uint64,uint64,uint64) commission, (string,string,string,string) description) payable returns()
func (_Contract *ContractTransactorSession) CreateValidator(consensusAddress common.Address, voteAddress []byte, blsProof []byte, commission StakeHubCommission, description StakeHubDescription) (*types.Transaction, error) {
	return _Contract.Contract.CreateValidator(&_Contract.TransactOpts, consensusAddress, voteAddress, blsProof, commission, description)
}

// Delegate is a paid mutator transaction binding the contract method 0x982ef0a7.
//
// Solidity: function delegate(address operatorAddress, bool delegateVotePower) payable returns()
func (_Contract *ContractTransactor) Delegate(opts *bind.TransactOpts, operatorAddress common.Address, delegateVotePower bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "delegate", operatorAddress, delegateVotePower)
}

// Delegate is a paid mutator transaction binding the contract method 0x982ef0a7.
//
// Solidity: function delegate(address operatorAddress, bool delegateVotePower) payable returns()
func (_Contract *ContractSession) Delegate(operatorAddress common.Address, delegateVotePower bool) (*types.Transaction, error) {
	return _Contract.Contract.Delegate(&_Contract.TransactOpts, operatorAddress, delegateVotePower)
}

// Delegate is a paid mutator transaction binding the contract method 0x982ef0a7.
//
// Solidity: function delegate(address operatorAddress, bool delegateVotePower) payable returns()
func (_Contract *ContractTransactorSession) Delegate(operatorAddress common.Address, delegateVotePower bool) (*types.Transaction, error) {
	return _Contract.Contract.Delegate(&_Contract.TransactOpts, operatorAddress, delegateVotePower)
}

// DistributeReward is a paid mutator transaction binding the contract method 0x092193ab.
//
// Solidity: function distributeReward(address consensusAddress) payable returns()
func (_Contract *ContractTransactor) DistributeReward(opts *bind.TransactOpts, consensusAddress common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "distributeReward", consensusAddress)
}

// DistributeReward is a paid mutator transaction binding the contract method 0x092193ab.
//
// Solidity: function distributeReward(address consensusAddress) payable returns()
func (_Contract *ContractSession) DistributeReward(consensusAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.DistributeReward(&_Contract.TransactOpts, consensusAddress)
}

// DistributeReward is a paid mutator transaction binding the contract method 0x092193ab.
//
// Solidity: function distributeReward(address consensusAddress) payable returns()
func (_Contract *ContractTransactorSession) DistributeReward(consensusAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.DistributeReward(&_Contract.TransactOpts, consensusAddress)
}

// DoubleSignSlash is a paid mutator transaction binding the contract method 0xc38fbec8.
//
// Solidity: function doubleSignSlash(address consensusAddress) returns()
func (_Contract *ContractTransactor) DoubleSignSlash(opts *bind.TransactOpts, consensusAddress common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "doubleSignSlash", consensusAddress)
}

// DoubleSignSlash is a paid mutator transaction binding the contract method 0xc38fbec8.
//
// Solidity: function doubleSignSlash(address consensusAddress) returns()
func (_Contract *ContractSession) DoubleSignSlash(consensusAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.DoubleSignSlash(&_Contract.TransactOpts, consensusAddress)
}

// DoubleSignSlash is a paid mutator transaction binding the contract method 0xc38fbec8.
//
// Solidity: function doubleSignSlash(address consensusAddress) returns()
func (_Contract *ContractTransactorSession) DoubleSignSlash(consensusAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.DoubleSignSlash(&_Contract.TransactOpts, consensusAddress)
}

// DowntimeSlash is a paid mutator transaction binding the contract method 0x75cc7d89.
//
// Solidity: function downtimeSlash(address consensusAddress) returns()
func (_Contract *ContractTransactor) DowntimeSlash(opts *bind.TransactOpts, consensusAddress common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "downtimeSlash", consensusAddress)
}

// DowntimeSlash is a paid mutator transaction binding the contract method 0x75cc7d89.
//
// Solidity: function downtimeSlash(address consensusAddress) returns()
func (_Contract *ContractSession) DowntimeSlash(consensusAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.DowntimeSlash(&_Contract.TransactOpts, consensusAddress)
}

// DowntimeSlash is a paid mutator transaction binding the contract method 0x75cc7d89.
//
// Solidity: function downtimeSlash(address consensusAddress) returns()
func (_Contract *ContractTransactorSession) DowntimeSlash(consensusAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.DowntimeSlash(&_Contract.TransactOpts, consensusAddress)
}

// EditCommissionRate is a paid mutator transaction binding the contract method 0x5e7cc1c9.
//
// Solidity: function editCommissionRate(uint64 commissionRate) returns()
func (_Contract *ContractTransactor) EditCommissionRate(opts *bind.TransactOpts, commissionRate uint64) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "editCommissionRate", commissionRate)
}

// EditCommissionRate is a paid mutator transaction binding the contract method 0x5e7cc1c9.
//
// Solidity: function editCommissionRate(uint64 commissionRate) returns()
func (_Contract *ContractSession) EditCommissionRate(commissionRate uint64) (*types.Transaction, error) {
	return _Contract.Contract.EditCommissionRate(&_Contract.TransactOpts, commissionRate)
}

// EditCommissionRate is a paid mutator transaction binding the contract method 0x5e7cc1c9.
//
// Solidity: function editCommissionRate(uint64 commissionRate) returns()
func (_Contract *ContractTransactorSession) EditCommissionRate(commissionRate uint64) (*types.Transaction, error) {
	return _Contract.Contract.EditCommissionRate(&_Contract.TransactOpts, commissionRate)
}

// EditConsensusAddress is a paid mutator transaction binding the contract method 0x45211bfd.
//
// Solidity: function editConsensusAddress(address newConsensusAddress) returns()
func (_Contract *ContractTransactor) EditConsensusAddress(opts *bind.TransactOpts, newConsensusAddress common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "editConsensusAddress", newConsensusAddress)
}

// EditConsensusAddress is a paid mutator transaction binding the contract method 0x45211bfd.
//
// Solidity: function editConsensusAddress(address newConsensusAddress) returns()
func (_Contract *ContractSession) EditConsensusAddress(newConsensusAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.EditConsensusAddress(&_Contract.TransactOpts, newConsensusAddress)
}

// EditConsensusAddress is a paid mutator transaction binding the contract method 0x45211bfd.
//
// Solidity: function editConsensusAddress(address newConsensusAddress) returns()
func (_Contract *ContractTransactorSession) EditConsensusAddress(newConsensusAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.EditConsensusAddress(&_Contract.TransactOpts, newConsensusAddress)
}

// EditDescription is a paid mutator transaction binding the contract method 0xd6ca429d.
//
// Solidity: function editDescription((string,string,string,string) description) returns()
func (_Contract *ContractTransactor) EditDescription(opts *bind.TransactOpts, description StakeHubDescription) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "editDescription", description)
}

// EditDescription is a paid mutator transaction binding the contract method 0xd6ca429d.
//
// Solidity: function editDescription((string,string,string,string) description) returns()
func (_Contract *ContractSession) EditDescription(description StakeHubDescription) (*types.Transaction, error) {
	return _Contract.Contract.EditDescription(&_Contract.TransactOpts, description)
}

// EditDescription is a paid mutator transaction binding the contract method 0xd6ca429d.
//
// Solidity: function editDescription((string,string,string,string) description) returns()
func (_Contract *ContractTransactorSession) EditDescription(description StakeHubDescription) (*types.Transaction, error) {
	return _Contract.Contract.EditDescription(&_Contract.TransactOpts, description)
}

// EditVoteAddress is a paid mutator transaction binding the contract method 0xfb50b31f.
//
// Solidity: function editVoteAddress(bytes newVoteAddress, bytes blsProof) returns()
func (_Contract *ContractTransactor) EditVoteAddress(opts *bind.TransactOpts, newVoteAddress []byte, blsProof []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "editVoteAddress", newVoteAddress, blsProof)
}

// EditVoteAddress is a paid mutator transaction binding the contract method 0xfb50b31f.
//
// Solidity: function editVoteAddress(bytes newVoteAddress, bytes blsProof) returns()
func (_Contract *ContractSession) EditVoteAddress(newVoteAddress []byte, blsProof []byte) (*types.Transaction, error) {
	return _Contract.Contract.EditVoteAddress(&_Contract.TransactOpts, newVoteAddress, blsProof)
}

// EditVoteAddress is a paid mutator transaction binding the contract method 0xfb50b31f.
//
// Solidity: function editVoteAddress(bytes newVoteAddress, bytes blsProof) returns()
func (_Contract *ContractTransactorSession) EditVoteAddress(newVoteAddress []byte, blsProof []byte) (*types.Transaction, error) {
	return _Contract.Contract.EditVoteAddress(&_Contract.TransactOpts, newVoteAddress, blsProof)
}

// HandleAckPackage is a paid mutator transaction binding the contract method 0x831d65d1.
//
// Solidity: function handleAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_Contract *ContractTransactor) HandleAckPackage(opts *bind.TransactOpts, channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "handleAckPackage", channelId, msgBytes)
}

// HandleAckPackage is a paid mutator transaction binding the contract method 0x831d65d1.
//
// Solidity: function handleAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_Contract *ContractSession) HandleAckPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Contract.Contract.HandleAckPackage(&_Contract.TransactOpts, channelId, msgBytes)
}

// HandleAckPackage is a paid mutator transaction binding the contract method 0x831d65d1.
//
// Solidity: function handleAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_Contract *ContractTransactorSession) HandleAckPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Contract.Contract.HandleAckPackage(&_Contract.TransactOpts, channelId, msgBytes)
}

// HandleFailAckPackage is a paid mutator transaction binding the contract method 0xc8509d81.
//
// Solidity: function handleFailAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_Contract *ContractTransactor) HandleFailAckPackage(opts *bind.TransactOpts, channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "handleFailAckPackage", channelId, msgBytes)
}

// HandleFailAckPackage is a paid mutator transaction binding the contract method 0xc8509d81.
//
// Solidity: function handleFailAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_Contract *ContractSession) HandleFailAckPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Contract.Contract.HandleFailAckPackage(&_Contract.TransactOpts, channelId, msgBytes)
}

// HandleFailAckPackage is a paid mutator transaction binding the contract method 0xc8509d81.
//
// Solidity: function handleFailAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_Contract *ContractTransactorSession) HandleFailAckPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Contract.Contract.HandleFailAckPackage(&_Contract.TransactOpts, channelId, msgBytes)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 , bytes msgBytes) returns(bytes)
func (_Contract *ContractTransactor) HandleSynPackage(opts *bind.TransactOpts, arg0 uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "handleSynPackage", arg0, msgBytes)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 , bytes msgBytes) returns(bytes)
func (_Contract *ContractSession) HandleSynPackage(arg0 uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Contract.Contract.HandleSynPackage(&_Contract.TransactOpts, arg0, msgBytes)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 , bytes msgBytes) returns(bytes)
func (_Contract *ContractTransactorSession) HandleSynPackage(arg0 uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Contract.Contract.HandleSynPackage(&_Contract.TransactOpts, arg0, msgBytes)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Contract *ContractTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Contract *ContractSession) Initialize() (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Contract *ContractTransactorSession) Initialize() (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts)
}

// MaliciousVoteSlash is a paid mutator transaction binding the contract method 0x0e9fbf51.
//
// Solidity: function maliciousVoteSlash(bytes voteAddress) returns()
func (_Contract *ContractTransactor) MaliciousVoteSlash(opts *bind.TransactOpts, voteAddress []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "maliciousVoteSlash", voteAddress)
}

// MaliciousVoteSlash is a paid mutator transaction binding the contract method 0x0e9fbf51.
//
// Solidity: function maliciousVoteSlash(bytes voteAddress) returns()
func (_Contract *ContractSession) MaliciousVoteSlash(voteAddress []byte) (*types.Transaction, error) {
	return _Contract.Contract.MaliciousVoteSlash(&_Contract.TransactOpts, voteAddress)
}

// MaliciousVoteSlash is a paid mutator transaction binding the contract method 0x0e9fbf51.
//
// Solidity: function maliciousVoteSlash(bytes voteAddress) returns()
func (_Contract *ContractTransactorSession) MaliciousVoteSlash(voteAddress []byte) (*types.Transaction, error) {
	return _Contract.Contract.MaliciousVoteSlash(&_Contract.TransactOpts, voteAddress)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Contract *ContractTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Contract *ContractSession) Pause() (*types.Transaction, error) {
	return _Contract.Contract.Pause(&_Contract.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Contract *ContractTransactorSession) Pause() (*types.Transaction, error) {
	return _Contract.Contract.Pause(&_Contract.TransactOpts)
}

// Redelegate is a paid mutator transaction binding the contract method 0x59491871.
//
// Solidity: function redelegate(address srcValidator, address dstValidator, uint256 shares, bool delegateVotePower) returns()
func (_Contract *ContractTransactor) Redelegate(opts *bind.TransactOpts, srcValidator common.Address, dstValidator common.Address, shares *big.Int, delegateVotePower bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "redelegate", srcValidator, dstValidator, shares, delegateVotePower)
}

// Redelegate is a paid mutator transaction binding the contract method 0x59491871.
//
// Solidity: function redelegate(address srcValidator, address dstValidator, uint256 shares, bool delegateVotePower) returns()
func (_Contract *ContractSession) Redelegate(srcValidator common.Address, dstValidator common.Address, shares *big.Int, delegateVotePower bool) (*types.Transaction, error) {
	return _Contract.Contract.Redelegate(&_Contract.TransactOpts, srcValidator, dstValidator, shares, delegateVotePower)
}

// Redelegate is a paid mutator transaction binding the contract method 0x59491871.
//
// Solidity: function redelegate(address srcValidator, address dstValidator, uint256 shares, bool delegateVotePower) returns()
func (_Contract *ContractTransactorSession) Redelegate(srcValidator common.Address, dstValidator common.Address, shares *big.Int, delegateVotePower bool) (*types.Transaction, error) {
	return _Contract.Contract.Redelegate(&_Contract.TransactOpts, srcValidator, dstValidator, shares, delegateVotePower)
}

// RemoveFromBlackList is a paid mutator transaction binding the contract method 0x4a49ac4c.
//
// Solidity: function removeFromBlackList(address account) returns()
func (_Contract *ContractTransactor) RemoveFromBlackList(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "removeFromBlackList", account)
}

// RemoveFromBlackList is a paid mutator transaction binding the contract method 0x4a49ac4c.
//
// Solidity: function removeFromBlackList(address account) returns()
func (_Contract *ContractSession) RemoveFromBlackList(account common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RemoveFromBlackList(&_Contract.TransactOpts, account)
}

// RemoveFromBlackList is a paid mutator transaction binding the contract method 0x4a49ac4c.
//
// Solidity: function removeFromBlackList(address account) returns()
func (_Contract *ContractTransactorSession) RemoveFromBlackList(account common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RemoveFromBlackList(&_Contract.TransactOpts, account)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Contract *ContractTransactor) Resume(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "resume")
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Contract *ContractSession) Resume() (*types.Transaction, error) {
	return _Contract.Contract.Resume(&_Contract.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Contract *ContractTransactorSession) Resume() (*types.Transaction, error) {
	return _Contract.Contract.Resume(&_Contract.TransactOpts)
}

// SyncGovToken is a paid mutator transaction binding the contract method 0xbaa7199e.
//
// Solidity: function syncGovToken(address[] operatorAddresses, address account) returns()
func (_Contract *ContractTransactor) SyncGovToken(opts *bind.TransactOpts, operatorAddresses []common.Address, account common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "syncGovToken", operatorAddresses, account)
}

// SyncGovToken is a paid mutator transaction binding the contract method 0xbaa7199e.
//
// Solidity: function syncGovToken(address[] operatorAddresses, address account) returns()
func (_Contract *ContractSession) SyncGovToken(operatorAddresses []common.Address, account common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SyncGovToken(&_Contract.TransactOpts, operatorAddresses, account)
}

// SyncGovToken is a paid mutator transaction binding the contract method 0xbaa7199e.
//
// Solidity: function syncGovToken(address[] operatorAddresses, address account) returns()
func (_Contract *ContractTransactorSession) SyncGovToken(operatorAddresses []common.Address, account common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SyncGovToken(&_Contract.TransactOpts, operatorAddresses, account)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address operatorAddress, uint256 shares) returns()
func (_Contract *ContractTransactor) Undelegate(opts *bind.TransactOpts, operatorAddress common.Address, shares *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "undelegate", operatorAddress, shares)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address operatorAddress, uint256 shares) returns()
func (_Contract *ContractSession) Undelegate(operatorAddress common.Address, shares *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Undelegate(&_Contract.TransactOpts, operatorAddress, shares)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address operatorAddress, uint256 shares) returns()
func (_Contract *ContractTransactorSession) Undelegate(operatorAddress common.Address, shares *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Undelegate(&_Contract.TransactOpts, operatorAddress, shares)
}

// Unjail is a paid mutator transaction binding the contract method 0x449ecfe6.
//
// Solidity: function unjail(address operatorAddress) returns()
func (_Contract *ContractTransactor) Unjail(opts *bind.TransactOpts, operatorAddress common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "unjail", operatorAddress)
}

// Unjail is a paid mutator transaction binding the contract method 0x449ecfe6.
//
// Solidity: function unjail(address operatorAddress) returns()
func (_Contract *ContractSession) Unjail(operatorAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Unjail(&_Contract.TransactOpts, operatorAddress)
}

// Unjail is a paid mutator transaction binding the contract method 0x449ecfe6.
//
// Solidity: function unjail(address operatorAddress) returns()
func (_Contract *ContractTransactorSession) Unjail(operatorAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Unjail(&_Contract.TransactOpts, operatorAddress)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Contract *ContractTransactor) UpdateParam(opts *bind.TransactOpts, key string, value []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateParam", key, value)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Contract *ContractSession) UpdateParam(key string, value []byte) (*types.Transaction, error) {
	return _Contract.Contract.UpdateParam(&_Contract.TransactOpts, key, value)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Contract *ContractTransactorSession) UpdateParam(key string, value []byte) (*types.Transaction, error) {
	return _Contract.Contract.UpdateParam(&_Contract.TransactOpts, key, value)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contract *ContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contract *ContractSession) Receive() (*types.Transaction, error) {
	return _Contract.Contract.Receive(&_Contract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contract *ContractTransactorSession) Receive() (*types.Transaction, error) {
	return _Contract.Contract.Receive(&_Contract.TransactOpts)
}

// ContractBlackListedIterator is returned from FilterBlackListed and is used to iterate over the raw logs and unpacked data for BlackListed events raised by the Contract contract.
type ContractBlackListedIterator struct {
	Event *ContractBlackListed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractBlackListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractBlackListed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractBlackListed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractBlackListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractBlackListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractBlackListed represents a BlackListed event raised by the Contract contract.
type ContractBlackListed struct {
	Target common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBlackListed is a free log retrieval operation binding the contract event 0x7fd26be6fc92aff63f1f4409b2b2ddeb272a888031d7f55ec830485ec6194186.
//
// Solidity: event BlackListed(address indexed target)
func (_Contract *ContractFilterer) FilterBlackListed(opts *bind.FilterOpts, target []common.Address) (*ContractBlackListedIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "BlackListed", targetRule)
	if err != nil {
		return nil, err
	}
	return &ContractBlackListedIterator{contract: _Contract.contract, event: "BlackListed", logs: logs, sub: sub}, nil
}

// WatchBlackListed is a free log subscription operation binding the contract event 0x7fd26be6fc92aff63f1f4409b2b2ddeb272a888031d7f55ec830485ec6194186.
//
// Solidity: event BlackListed(address indexed target)
func (_Contract *ContractFilterer) WatchBlackListed(opts *bind.WatchOpts, sink chan<- *ContractBlackListed, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "BlackListed", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractBlackListed)
				if err := _Contract.contract.UnpackLog(event, "BlackListed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBlackListed is a log parse operation binding the contract event 0x7fd26be6fc92aff63f1f4409b2b2ddeb272a888031d7f55ec830485ec6194186.
//
// Solidity: event BlackListed(address indexed target)
func (_Contract *ContractFilterer) ParseBlackListed(log types.Log) (*ContractBlackListed, error) {
	event := new(ContractBlackListed)
	if err := _Contract.contract.UnpackLog(event, "BlackListed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the Contract contract.
type ContractClaimedIterator struct {
	Event *ContractClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractClaimed represents a Claimed event raised by the Contract contract.
type ContractClaimed struct {
	OperatorAddress common.Address
	Delegator       common.Address
	BnbAmount       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0xf7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd3992683.
//
// Solidity: event Claimed(address indexed operatorAddress, address indexed delegator, uint256 bnbAmount)
func (_Contract *ContractFilterer) FilterClaimed(opts *bind.FilterOpts, operatorAddress []common.Address, delegator []common.Address) (*ContractClaimedIterator, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Claimed", operatorAddressRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractClaimedIterator{contract: _Contract.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0xf7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd3992683.
//
// Solidity: event Claimed(address indexed operatorAddress, address indexed delegator, uint256 bnbAmount)
func (_Contract *ContractFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *ContractClaimed, operatorAddress []common.Address, delegator []common.Address) (event.Subscription, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Claimed", operatorAddressRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractClaimed)
				if err := _Contract.contract.UnpackLog(event, "Claimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaimed is a log parse operation binding the contract event 0xf7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd3992683.
//
// Solidity: event Claimed(address indexed operatorAddress, address indexed delegator, uint256 bnbAmount)
func (_Contract *ContractFilterer) ParseClaimed(log types.Log) (*ContractClaimed, error) {
	event := new(ContractClaimed)
	if err := _Contract.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractCommissionRateEditedIterator is returned from FilterCommissionRateEdited and is used to iterate over the raw logs and unpacked data for CommissionRateEdited events raised by the Contract contract.
type ContractCommissionRateEditedIterator struct {
	Event *ContractCommissionRateEdited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractCommissionRateEditedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractCommissionRateEdited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractCommissionRateEdited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractCommissionRateEditedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractCommissionRateEditedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractCommissionRateEdited represents a CommissionRateEdited event raised by the Contract contract.
type ContractCommissionRateEdited struct {
	OperatorAddress   common.Address
	NewCommissionRate uint64
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCommissionRateEdited is a free log retrieval operation binding the contract event 0x78cdd96edf59e09cfd4d26ef6ef6c92d166effe6a40970c54821206d541932cb.
//
// Solidity: event CommissionRateEdited(address indexed operatorAddress, uint64 newCommissionRate)
func (_Contract *ContractFilterer) FilterCommissionRateEdited(opts *bind.FilterOpts, operatorAddress []common.Address) (*ContractCommissionRateEditedIterator, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "CommissionRateEdited", operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractCommissionRateEditedIterator{contract: _Contract.contract, event: "CommissionRateEdited", logs: logs, sub: sub}, nil
}

// WatchCommissionRateEdited is a free log subscription operation binding the contract event 0x78cdd96edf59e09cfd4d26ef6ef6c92d166effe6a40970c54821206d541932cb.
//
// Solidity: event CommissionRateEdited(address indexed operatorAddress, uint64 newCommissionRate)
func (_Contract *ContractFilterer) WatchCommissionRateEdited(opts *bind.WatchOpts, sink chan<- *ContractCommissionRateEdited, operatorAddress []common.Address) (event.Subscription, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "CommissionRateEdited", operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractCommissionRateEdited)
				if err := _Contract.contract.UnpackLog(event, "CommissionRateEdited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCommissionRateEdited is a log parse operation binding the contract event 0x78cdd96edf59e09cfd4d26ef6ef6c92d166effe6a40970c54821206d541932cb.
//
// Solidity: event CommissionRateEdited(address indexed operatorAddress, uint64 newCommissionRate)
func (_Contract *ContractFilterer) ParseCommissionRateEdited(log types.Log) (*ContractCommissionRateEdited, error) {
	event := new(ContractCommissionRateEdited)
	if err := _Contract.contract.UnpackLog(event, "CommissionRateEdited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractConsensusAddressEditedIterator is returned from FilterConsensusAddressEdited and is used to iterate over the raw logs and unpacked data for ConsensusAddressEdited events raised by the Contract contract.
type ContractConsensusAddressEditedIterator struct {
	Event *ContractConsensusAddressEdited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractConsensusAddressEditedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractConsensusAddressEdited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractConsensusAddressEdited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractConsensusAddressEditedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractConsensusAddressEditedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractConsensusAddressEdited represents a ConsensusAddressEdited event raised by the Contract contract.
type ContractConsensusAddressEdited struct {
	OperatorAddress     common.Address
	NewConsensusAddress common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterConsensusAddressEdited is a free log retrieval operation binding the contract event 0x6e4e747ca35203f16401c69805c7dd52fff67ef60b0ebc5c7fe16890530f2235.
//
// Solidity: event ConsensusAddressEdited(address indexed operatorAddress, address indexed newConsensusAddress)
func (_Contract *ContractFilterer) FilterConsensusAddressEdited(opts *bind.FilterOpts, operatorAddress []common.Address, newConsensusAddress []common.Address) (*ContractConsensusAddressEditedIterator, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}
	var newConsensusAddressRule []interface{}
	for _, newConsensusAddressItem := range newConsensusAddress {
		newConsensusAddressRule = append(newConsensusAddressRule, newConsensusAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ConsensusAddressEdited", operatorAddressRule, newConsensusAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractConsensusAddressEditedIterator{contract: _Contract.contract, event: "ConsensusAddressEdited", logs: logs, sub: sub}, nil
}

// WatchConsensusAddressEdited is a free log subscription operation binding the contract event 0x6e4e747ca35203f16401c69805c7dd52fff67ef60b0ebc5c7fe16890530f2235.
//
// Solidity: event ConsensusAddressEdited(address indexed operatorAddress, address indexed newConsensusAddress)
func (_Contract *ContractFilterer) WatchConsensusAddressEdited(opts *bind.WatchOpts, sink chan<- *ContractConsensusAddressEdited, operatorAddress []common.Address, newConsensusAddress []common.Address) (event.Subscription, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}
	var newConsensusAddressRule []interface{}
	for _, newConsensusAddressItem := range newConsensusAddress {
		newConsensusAddressRule = append(newConsensusAddressRule, newConsensusAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ConsensusAddressEdited", operatorAddressRule, newConsensusAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractConsensusAddressEdited)
				if err := _Contract.contract.UnpackLog(event, "ConsensusAddressEdited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConsensusAddressEdited is a log parse operation binding the contract event 0x6e4e747ca35203f16401c69805c7dd52fff67ef60b0ebc5c7fe16890530f2235.
//
// Solidity: event ConsensusAddressEdited(address indexed operatorAddress, address indexed newConsensusAddress)
func (_Contract *ContractFilterer) ParseConsensusAddressEdited(log types.Log) (*ContractConsensusAddressEdited, error) {
	event := new(ContractConsensusAddressEdited)
	if err := _Contract.contract.UnpackLog(event, "ConsensusAddressEdited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegatedIterator is returned from FilterDelegated and is used to iterate over the raw logs and unpacked data for Delegated events raised by the Contract contract.
type ContractDelegatedIterator struct {
	Event *ContractDelegated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractDelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractDelegated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractDelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegated represents a Delegated event raised by the Contract contract.
type ContractDelegated struct {
	OperatorAddress common.Address
	Delegator       common.Address
	Shares          *big.Int
	BnbAmount       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDelegated is a free log retrieval operation binding the contract event 0x24d7bda8602b916d64417f0dbfe2e2e88ec9b1157bd9f596dfdb91ba26624e04.
//
// Solidity: event Delegated(address indexed operatorAddress, address indexed delegator, uint256 shares, uint256 bnbAmount)
func (_Contract *ContractFilterer) FilterDelegated(opts *bind.FilterOpts, operatorAddress []common.Address, delegator []common.Address) (*ContractDelegatedIterator, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Delegated", operatorAddressRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractDelegatedIterator{contract: _Contract.contract, event: "Delegated", logs: logs, sub: sub}, nil
}

// WatchDelegated is a free log subscription operation binding the contract event 0x24d7bda8602b916d64417f0dbfe2e2e88ec9b1157bd9f596dfdb91ba26624e04.
//
// Solidity: event Delegated(address indexed operatorAddress, address indexed delegator, uint256 shares, uint256 bnbAmount)
func (_Contract *ContractFilterer) WatchDelegated(opts *bind.WatchOpts, sink chan<- *ContractDelegated, operatorAddress []common.Address, delegator []common.Address) (event.Subscription, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Delegated", operatorAddressRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegated)
				if err := _Contract.contract.UnpackLog(event, "Delegated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelegated is a log parse operation binding the contract event 0x24d7bda8602b916d64417f0dbfe2e2e88ec9b1157bd9f596dfdb91ba26624e04.
//
// Solidity: event Delegated(address indexed operatorAddress, address indexed delegator, uint256 shares, uint256 bnbAmount)
func (_Contract *ContractFilterer) ParseDelegated(log types.Log) (*ContractDelegated, error) {
	event := new(ContractDelegated)
	if err := _Contract.contract.UnpackLog(event, "Delegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDescriptionEditedIterator is returned from FilterDescriptionEdited and is used to iterate over the raw logs and unpacked data for DescriptionEdited events raised by the Contract contract.
type ContractDescriptionEditedIterator struct {
	Event *ContractDescriptionEdited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractDescriptionEditedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDescriptionEdited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractDescriptionEdited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractDescriptionEditedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDescriptionEditedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDescriptionEdited represents a DescriptionEdited event raised by the Contract contract.
type ContractDescriptionEdited struct {
	OperatorAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDescriptionEdited is a free log retrieval operation binding the contract event 0x85d6366b336ade7f106987ec7a8eac1e8799e508aeab045a39d2f63e0dc969d9.
//
// Solidity: event DescriptionEdited(address indexed operatorAddress)
func (_Contract *ContractFilterer) FilterDescriptionEdited(opts *bind.FilterOpts, operatorAddress []common.Address) (*ContractDescriptionEditedIterator, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "DescriptionEdited", operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractDescriptionEditedIterator{contract: _Contract.contract, event: "DescriptionEdited", logs: logs, sub: sub}, nil
}

// WatchDescriptionEdited is a free log subscription operation binding the contract event 0x85d6366b336ade7f106987ec7a8eac1e8799e508aeab045a39d2f63e0dc969d9.
//
// Solidity: event DescriptionEdited(address indexed operatorAddress)
func (_Contract *ContractFilterer) WatchDescriptionEdited(opts *bind.WatchOpts, sink chan<- *ContractDescriptionEdited, operatorAddress []common.Address) (event.Subscription, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "DescriptionEdited", operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDescriptionEdited)
				if err := _Contract.contract.UnpackLog(event, "DescriptionEdited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDescriptionEdited is a log parse operation binding the contract event 0x85d6366b336ade7f106987ec7a8eac1e8799e508aeab045a39d2f63e0dc969d9.
//
// Solidity: event DescriptionEdited(address indexed operatorAddress)
func (_Contract *ContractFilterer) ParseDescriptionEdited(log types.Log) (*ContractDescriptionEdited, error) {
	event := new(ContractDescriptionEdited)
	if err := _Contract.contract.UnpackLog(event, "DescriptionEdited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Contract contract.
type ContractInitializedIterator struct {
	Event *ContractInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractInitialized represents a Initialized event raised by the Contract contract.
type ContractInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Contract *ContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractInitializedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractInitializedIterator{contract: _Contract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Contract *ContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractInitialized) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractInitialized)
				if err := _Contract.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Contract *ContractFilterer) ParseInitialized(log types.Log) (*ContractInitialized, error) {
	event := new(ContractInitialized)
	if err := _Contract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractMigrateFailedIterator is returned from FilterMigrateFailed and is used to iterate over the raw logs and unpacked data for MigrateFailed events raised by the Contract contract.
type ContractMigrateFailedIterator struct {
	Event *ContractMigrateFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractMigrateFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMigrateFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractMigrateFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractMigrateFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMigrateFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMigrateFailed represents a MigrateFailed event raised by the Contract contract.
type ContractMigrateFailed struct {
	OperatorAddress common.Address
	Delegator       common.Address
	BnbAmount       *big.Int
	RespCode        uint8
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMigrateFailed is a free log retrieval operation binding the contract event 0xa9084c89a291b43bc984e045671d394974730a159b9a826b577bb148ab504c3a.
//
// Solidity: event MigrateFailed(address indexed operatorAddress, address indexed delegator, uint256 bnbAmount, uint8 respCode)
func (_Contract *ContractFilterer) FilterMigrateFailed(opts *bind.FilterOpts, operatorAddress []common.Address, delegator []common.Address) (*ContractMigrateFailedIterator, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "MigrateFailed", operatorAddressRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractMigrateFailedIterator{contract: _Contract.contract, event: "MigrateFailed", logs: logs, sub: sub}, nil
}

// WatchMigrateFailed is a free log subscription operation binding the contract event 0xa9084c89a291b43bc984e045671d394974730a159b9a826b577bb148ab504c3a.
//
// Solidity: event MigrateFailed(address indexed operatorAddress, address indexed delegator, uint256 bnbAmount, uint8 respCode)
func (_Contract *ContractFilterer) WatchMigrateFailed(opts *bind.WatchOpts, sink chan<- *ContractMigrateFailed, operatorAddress []common.Address, delegator []common.Address) (event.Subscription, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "MigrateFailed", operatorAddressRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMigrateFailed)
				if err := _Contract.contract.UnpackLog(event, "MigrateFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMigrateFailed is a log parse operation binding the contract event 0xa9084c89a291b43bc984e045671d394974730a159b9a826b577bb148ab504c3a.
//
// Solidity: event MigrateFailed(address indexed operatorAddress, address indexed delegator, uint256 bnbAmount, uint8 respCode)
func (_Contract *ContractFilterer) ParseMigrateFailed(log types.Log) (*ContractMigrateFailed, error) {
	event := new(ContractMigrateFailed)
	if err := _Contract.contract.UnpackLog(event, "MigrateFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractMigrateSuccessIterator is returned from FilterMigrateSuccess and is used to iterate over the raw logs and unpacked data for MigrateSuccess events raised by the Contract contract.
type ContractMigrateSuccessIterator struct {
	Event *ContractMigrateSuccess // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractMigrateSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMigrateSuccess)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractMigrateSuccess)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractMigrateSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMigrateSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMigrateSuccess represents a MigrateSuccess event raised by the Contract contract.
type ContractMigrateSuccess struct {
	OperatorAddress common.Address
	Delegator       common.Address
	Shares          *big.Int
	BnbAmount       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMigrateSuccess is a free log retrieval operation binding the contract event 0x607b17598da6bdca05650a2fc08bd2bc8e38c3236806a0fa8e0daabc1d6cb1d8.
//
// Solidity: event MigrateSuccess(address indexed operatorAddress, address indexed delegator, uint256 shares, uint256 bnbAmount)
func (_Contract *ContractFilterer) FilterMigrateSuccess(opts *bind.FilterOpts, operatorAddress []common.Address, delegator []common.Address) (*ContractMigrateSuccessIterator, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "MigrateSuccess", operatorAddressRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractMigrateSuccessIterator{contract: _Contract.contract, event: "MigrateSuccess", logs: logs, sub: sub}, nil
}

// WatchMigrateSuccess is a free log subscription operation binding the contract event 0x607b17598da6bdca05650a2fc08bd2bc8e38c3236806a0fa8e0daabc1d6cb1d8.
//
// Solidity: event MigrateSuccess(address indexed operatorAddress, address indexed delegator, uint256 shares, uint256 bnbAmount)
func (_Contract *ContractFilterer) WatchMigrateSuccess(opts *bind.WatchOpts, sink chan<- *ContractMigrateSuccess, operatorAddress []common.Address, delegator []common.Address) (event.Subscription, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "MigrateSuccess", operatorAddressRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMigrateSuccess)
				if err := _Contract.contract.UnpackLog(event, "MigrateSuccess", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMigrateSuccess is a log parse operation binding the contract event 0x607b17598da6bdca05650a2fc08bd2bc8e38c3236806a0fa8e0daabc1d6cb1d8.
//
// Solidity: event MigrateSuccess(address indexed operatorAddress, address indexed delegator, uint256 shares, uint256 bnbAmount)
func (_Contract *ContractFilterer) ParseMigrateSuccess(log types.Log) (*ContractMigrateSuccess, error) {
	event := new(ContractMigrateSuccess)
	if err := _Contract.contract.UnpackLog(event, "MigrateSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractParamChangeIterator is returned from FilterParamChange and is used to iterate over the raw logs and unpacked data for ParamChange events raised by the Contract contract.
type ContractParamChangeIterator struct {
	Event *ContractParamChange // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractParamChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractParamChange)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractParamChange)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractParamChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractParamChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractParamChange represents a ParamChange event raised by the Contract contract.
type ContractParamChange struct {
	Key   string
	Value []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterParamChange is a free log retrieval operation binding the contract event 0xf1ce9b2cbf50eeb05769a29e2543fd350cab46894a7dd9978a12d534bb20e633.
//
// Solidity: event ParamChange(string key, bytes value)
func (_Contract *ContractFilterer) FilterParamChange(opts *bind.FilterOpts) (*ContractParamChangeIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ParamChange")
	if err != nil {
		return nil, err
	}
	return &ContractParamChangeIterator{contract: _Contract.contract, event: "ParamChange", logs: logs, sub: sub}, nil
}

// WatchParamChange is a free log subscription operation binding the contract event 0xf1ce9b2cbf50eeb05769a29e2543fd350cab46894a7dd9978a12d534bb20e633.
//
// Solidity: event ParamChange(string key, bytes value)
func (_Contract *ContractFilterer) WatchParamChange(opts *bind.WatchOpts, sink chan<- *ContractParamChange) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ParamChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractParamChange)
				if err := _Contract.contract.UnpackLog(event, "ParamChange", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseParamChange is a log parse operation binding the contract event 0xf1ce9b2cbf50eeb05769a29e2543fd350cab46894a7dd9978a12d534bb20e633.
//
// Solidity: event ParamChange(string key, bytes value)
func (_Contract *ContractFilterer) ParseParamChange(log types.Log) (*ContractParamChange, error) {
	event := new(ContractParamChange)
	if err := _Contract.contract.UnpackLog(event, "ParamChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Contract contract.
type ContractPausedIterator struct {
	Event *ContractPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPaused represents a Paused event raised by the Contract contract.
type ContractPaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_Contract *ContractFilterer) FilterPaused(opts *bind.FilterOpts) (*ContractPausedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ContractPausedIterator{contract: _Contract.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_Contract *ContractFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractPaused) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPaused)
				if err := _Contract.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_Contract *ContractFilterer) ParsePaused(log types.Log) (*ContractPaused, error) {
	event := new(ContractPaused)
	if err := _Contract.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractProtectorChangedIterator is returned from FilterProtectorChanged and is used to iterate over the raw logs and unpacked data for ProtectorChanged events raised by the Contract contract.
type ContractProtectorChangedIterator struct {
	Event *ContractProtectorChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractProtectorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractProtectorChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractProtectorChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractProtectorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractProtectorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractProtectorChanged represents a ProtectorChanged event raised by the Contract contract.
type ContractProtectorChanged struct {
	OldProtector common.Address
	NewProtector common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProtectorChanged is a free log retrieval operation binding the contract event 0x44fc1b38a4abaa91ebd1b628a5b259a698f86238c8217d68f516e87769c60c0b.
//
// Solidity: event ProtectorChanged(address indexed oldProtector, address indexed newProtector)
func (_Contract *ContractFilterer) FilterProtectorChanged(opts *bind.FilterOpts, oldProtector []common.Address, newProtector []common.Address) (*ContractProtectorChangedIterator, error) {

	var oldProtectorRule []interface{}
	for _, oldProtectorItem := range oldProtector {
		oldProtectorRule = append(oldProtectorRule, oldProtectorItem)
	}
	var newProtectorRule []interface{}
	for _, newProtectorItem := range newProtector {
		newProtectorRule = append(newProtectorRule, newProtectorItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ProtectorChanged", oldProtectorRule, newProtectorRule)
	if err != nil {
		return nil, err
	}
	return &ContractProtectorChangedIterator{contract: _Contract.contract, event: "ProtectorChanged", logs: logs, sub: sub}, nil
}

// WatchProtectorChanged is a free log subscription operation binding the contract event 0x44fc1b38a4abaa91ebd1b628a5b259a698f86238c8217d68f516e87769c60c0b.
//
// Solidity: event ProtectorChanged(address indexed oldProtector, address indexed newProtector)
func (_Contract *ContractFilterer) WatchProtectorChanged(opts *bind.WatchOpts, sink chan<- *ContractProtectorChanged, oldProtector []common.Address, newProtector []common.Address) (event.Subscription, error) {

	var oldProtectorRule []interface{}
	for _, oldProtectorItem := range oldProtector {
		oldProtectorRule = append(oldProtectorRule, oldProtectorItem)
	}
	var newProtectorRule []interface{}
	for _, newProtectorItem := range newProtector {
		newProtectorRule = append(newProtectorRule, newProtectorItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ProtectorChanged", oldProtectorRule, newProtectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractProtectorChanged)
				if err := _Contract.contract.UnpackLog(event, "ProtectorChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProtectorChanged is a log parse operation binding the contract event 0x44fc1b38a4abaa91ebd1b628a5b259a698f86238c8217d68f516e87769c60c0b.
//
// Solidity: event ProtectorChanged(address indexed oldProtector, address indexed newProtector)
func (_Contract *ContractFilterer) ParseProtectorChanged(log types.Log) (*ContractProtectorChanged, error) {
	event := new(ContractProtectorChanged)
	if err := _Contract.contract.UnpackLog(event, "ProtectorChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRedelegatedIterator is returned from FilterRedelegated and is used to iterate over the raw logs and unpacked data for Redelegated events raised by the Contract contract.
type ContractRedelegatedIterator struct {
	Event *ContractRedelegated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractRedelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRedelegated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractRedelegated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractRedelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRedelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRedelegated represents a Redelegated event raised by the Contract contract.
type ContractRedelegated struct {
	SrcValidator common.Address
	DstValidator common.Address
	Delegator    common.Address
	OldShares    *big.Int
	NewShares    *big.Int
	BnbAmount    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRedelegated is a free log retrieval operation binding the contract event 0xfdac6e81913996d95abcc289e90f2d8bd235487ce6fe6f821e7d21002a1915b4.
//
// Solidity: event Redelegated(address indexed srcValidator, address indexed dstValidator, address indexed delegator, uint256 oldShares, uint256 newShares, uint256 bnbAmount)
func (_Contract *ContractFilterer) FilterRedelegated(opts *bind.FilterOpts, srcValidator []common.Address, dstValidator []common.Address, delegator []common.Address) (*ContractRedelegatedIterator, error) {

	var srcValidatorRule []interface{}
	for _, srcValidatorItem := range srcValidator {
		srcValidatorRule = append(srcValidatorRule, srcValidatorItem)
	}
	var dstValidatorRule []interface{}
	for _, dstValidatorItem := range dstValidator {
		dstValidatorRule = append(dstValidatorRule, dstValidatorItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Redelegated", srcValidatorRule, dstValidatorRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractRedelegatedIterator{contract: _Contract.contract, event: "Redelegated", logs: logs, sub: sub}, nil
}

// WatchRedelegated is a free log subscription operation binding the contract event 0xfdac6e81913996d95abcc289e90f2d8bd235487ce6fe6f821e7d21002a1915b4.
//
// Solidity: event Redelegated(address indexed srcValidator, address indexed dstValidator, address indexed delegator, uint256 oldShares, uint256 newShares, uint256 bnbAmount)
func (_Contract *ContractFilterer) WatchRedelegated(opts *bind.WatchOpts, sink chan<- *ContractRedelegated, srcValidator []common.Address, dstValidator []common.Address, delegator []common.Address) (event.Subscription, error) {

	var srcValidatorRule []interface{}
	for _, srcValidatorItem := range srcValidator {
		srcValidatorRule = append(srcValidatorRule, srcValidatorItem)
	}
	var dstValidatorRule []interface{}
	for _, dstValidatorItem := range dstValidator {
		dstValidatorRule = append(dstValidatorRule, dstValidatorItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Redelegated", srcValidatorRule, dstValidatorRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRedelegated)
				if err := _Contract.contract.UnpackLog(event, "Redelegated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRedelegated is a log parse operation binding the contract event 0xfdac6e81913996d95abcc289e90f2d8bd235487ce6fe6f821e7d21002a1915b4.
//
// Solidity: event Redelegated(address indexed srcValidator, address indexed dstValidator, address indexed delegator, uint256 oldShares, uint256 newShares, uint256 bnbAmount)
func (_Contract *ContractFilterer) ParseRedelegated(log types.Log) (*ContractRedelegated, error) {
	event := new(ContractRedelegated)
	if err := _Contract.contract.UnpackLog(event, "Redelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractResumedIterator is returned from FilterResumed and is used to iterate over the raw logs and unpacked data for Resumed events raised by the Contract contract.
type ContractResumedIterator struct {
	Event *ContractResumed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractResumedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractResumed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractResumed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractResumedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractResumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractResumed represents a Resumed event raised by the Contract contract.
type ContractResumed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterResumed is a free log retrieval operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_Contract *ContractFilterer) FilterResumed(opts *bind.FilterOpts) (*ContractResumedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Resumed")
	if err != nil {
		return nil, err
	}
	return &ContractResumedIterator{contract: _Contract.contract, event: "Resumed", logs: logs, sub: sub}, nil
}

// WatchResumed is a free log subscription operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_Contract *ContractFilterer) WatchResumed(opts *bind.WatchOpts, sink chan<- *ContractResumed) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Resumed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractResumed)
				if err := _Contract.contract.UnpackLog(event, "Resumed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseResumed is a log parse operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_Contract *ContractFilterer) ParseResumed(log types.Log) (*ContractResumed, error) {
	event := new(ContractResumed)
	if err := _Contract.contract.UnpackLog(event, "Resumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRewardDistributeFailedIterator is returned from FilterRewardDistributeFailed and is used to iterate over the raw logs and unpacked data for RewardDistributeFailed events raised by the Contract contract.
type ContractRewardDistributeFailedIterator struct {
	Event *ContractRewardDistributeFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractRewardDistributeFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRewardDistributeFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractRewardDistributeFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractRewardDistributeFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRewardDistributeFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRewardDistributeFailed represents a RewardDistributeFailed event raised by the Contract contract.
type ContractRewardDistributeFailed struct {
	OperatorAddress common.Address
	FailReason      []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRewardDistributeFailed is a free log retrieval operation binding the contract event 0xfc8bff675087dd2da069cc3fb517b9ed001e19750c0865241a5542dba1ba170d.
//
// Solidity: event RewardDistributeFailed(address indexed operatorAddress, bytes failReason)
func (_Contract *ContractFilterer) FilterRewardDistributeFailed(opts *bind.FilterOpts, operatorAddress []common.Address) (*ContractRewardDistributeFailedIterator, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RewardDistributeFailed", operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractRewardDistributeFailedIterator{contract: _Contract.contract, event: "RewardDistributeFailed", logs: logs, sub: sub}, nil
}

// WatchRewardDistributeFailed is a free log subscription operation binding the contract event 0xfc8bff675087dd2da069cc3fb517b9ed001e19750c0865241a5542dba1ba170d.
//
// Solidity: event RewardDistributeFailed(address indexed operatorAddress, bytes failReason)
func (_Contract *ContractFilterer) WatchRewardDistributeFailed(opts *bind.WatchOpts, sink chan<- *ContractRewardDistributeFailed, operatorAddress []common.Address) (event.Subscription, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RewardDistributeFailed", operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRewardDistributeFailed)
				if err := _Contract.contract.UnpackLog(event, "RewardDistributeFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardDistributeFailed is a log parse operation binding the contract event 0xfc8bff675087dd2da069cc3fb517b9ed001e19750c0865241a5542dba1ba170d.
//
// Solidity: event RewardDistributeFailed(address indexed operatorAddress, bytes failReason)
func (_Contract *ContractFilterer) ParseRewardDistributeFailed(log types.Log) (*ContractRewardDistributeFailed, error) {
	event := new(ContractRewardDistributeFailed)
	if err := _Contract.contract.UnpackLog(event, "RewardDistributeFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRewardDistributedIterator is returned from FilterRewardDistributed and is used to iterate over the raw logs and unpacked data for RewardDistributed events raised by the Contract contract.
type ContractRewardDistributedIterator struct {
	Event *ContractRewardDistributed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractRewardDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRewardDistributed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractRewardDistributed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractRewardDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRewardDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRewardDistributed represents a RewardDistributed event raised by the Contract contract.
type ContractRewardDistributed struct {
	OperatorAddress common.Address
	Reward          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRewardDistributed is a free log retrieval operation binding the contract event 0xe34918ff1c7084970068b53fd71ad6d8b04e9f15d3886cbf006443e6cdc52ea6.
//
// Solidity: event RewardDistributed(address indexed operatorAddress, uint256 reward)
func (_Contract *ContractFilterer) FilterRewardDistributed(opts *bind.FilterOpts, operatorAddress []common.Address) (*ContractRewardDistributedIterator, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RewardDistributed", operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractRewardDistributedIterator{contract: _Contract.contract, event: "RewardDistributed", logs: logs, sub: sub}, nil
}

// WatchRewardDistributed is a free log subscription operation binding the contract event 0xe34918ff1c7084970068b53fd71ad6d8b04e9f15d3886cbf006443e6cdc52ea6.
//
// Solidity: event RewardDistributed(address indexed operatorAddress, uint256 reward)
func (_Contract *ContractFilterer) WatchRewardDistributed(opts *bind.WatchOpts, sink chan<- *ContractRewardDistributed, operatorAddress []common.Address) (event.Subscription, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RewardDistributed", operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRewardDistributed)
				if err := _Contract.contract.UnpackLog(event, "RewardDistributed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardDistributed is a log parse operation binding the contract event 0xe34918ff1c7084970068b53fd71ad6d8b04e9f15d3886cbf006443e6cdc52ea6.
//
// Solidity: event RewardDistributed(address indexed operatorAddress, uint256 reward)
func (_Contract *ContractFilterer) ParseRewardDistributed(log types.Log) (*ContractRewardDistributed, error) {
	event := new(ContractRewardDistributed)
	if err := _Contract.contract.UnpackLog(event, "RewardDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractStakeCreditInitializedIterator is returned from FilterStakeCreditInitialized and is used to iterate over the raw logs and unpacked data for StakeCreditInitialized events raised by the Contract contract.
type ContractStakeCreditInitializedIterator struct {
	Event *ContractStakeCreditInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractStakeCreditInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractStakeCreditInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractStakeCreditInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractStakeCreditInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractStakeCreditInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractStakeCreditInitialized represents a StakeCreditInitialized event raised by the Contract contract.
type ContractStakeCreditInitialized struct {
	OperatorAddress common.Address
	CreditContract  common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStakeCreditInitialized is a free log retrieval operation binding the contract event 0xd481492e4e93bb36b4c12a5af93f03be3bf04b454dfbc35dd2663fa26f44d5b0.
//
// Solidity: event StakeCreditInitialized(address indexed operatorAddress, address indexed creditContract)
func (_Contract *ContractFilterer) FilterStakeCreditInitialized(opts *bind.FilterOpts, operatorAddress []common.Address, creditContract []common.Address) (*ContractStakeCreditInitializedIterator, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}
	var creditContractRule []interface{}
	for _, creditContractItem := range creditContract {
		creditContractRule = append(creditContractRule, creditContractItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "StakeCreditInitialized", operatorAddressRule, creditContractRule)
	if err != nil {
		return nil, err
	}
	return &ContractStakeCreditInitializedIterator{contract: _Contract.contract, event: "StakeCreditInitialized", logs: logs, sub: sub}, nil
}

// WatchStakeCreditInitialized is a free log subscription operation binding the contract event 0xd481492e4e93bb36b4c12a5af93f03be3bf04b454dfbc35dd2663fa26f44d5b0.
//
// Solidity: event StakeCreditInitialized(address indexed operatorAddress, address indexed creditContract)
func (_Contract *ContractFilterer) WatchStakeCreditInitialized(opts *bind.WatchOpts, sink chan<- *ContractStakeCreditInitialized, operatorAddress []common.Address, creditContract []common.Address) (event.Subscription, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}
	var creditContractRule []interface{}
	for _, creditContractItem := range creditContract {
		creditContractRule = append(creditContractRule, creditContractItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "StakeCreditInitialized", operatorAddressRule, creditContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractStakeCreditInitialized)
				if err := _Contract.contract.UnpackLog(event, "StakeCreditInitialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStakeCreditInitialized is a log parse operation binding the contract event 0xd481492e4e93bb36b4c12a5af93f03be3bf04b454dfbc35dd2663fa26f44d5b0.
//
// Solidity: event StakeCreditInitialized(address indexed operatorAddress, address indexed creditContract)
func (_Contract *ContractFilterer) ParseStakeCreditInitialized(log types.Log) (*ContractStakeCreditInitialized, error) {
	event := new(ContractStakeCreditInitialized)
	if err := _Contract.contract.UnpackLog(event, "StakeCreditInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractUnBlackListedIterator is returned from FilterUnBlackListed and is used to iterate over the raw logs and unpacked data for UnBlackListed events raised by the Contract contract.
type ContractUnBlackListedIterator struct {
	Event *ContractUnBlackListed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractUnBlackListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUnBlackListed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractUnBlackListed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractUnBlackListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUnBlackListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUnBlackListed represents a UnBlackListed event raised by the Contract contract.
type ContractUnBlackListed struct {
	Target common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnBlackListed is a free log retrieval operation binding the contract event 0xe0db3499b7fdc3da4cddff5f45d694549c19835e7f719fb5606d3ad1a5de4011.
//
// Solidity: event UnBlackListed(address indexed target)
func (_Contract *ContractFilterer) FilterUnBlackListed(opts *bind.FilterOpts, target []common.Address) (*ContractUnBlackListedIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "UnBlackListed", targetRule)
	if err != nil {
		return nil, err
	}
	return &ContractUnBlackListedIterator{contract: _Contract.contract, event: "UnBlackListed", logs: logs, sub: sub}, nil
}

// WatchUnBlackListed is a free log subscription operation binding the contract event 0xe0db3499b7fdc3da4cddff5f45d694549c19835e7f719fb5606d3ad1a5de4011.
//
// Solidity: event UnBlackListed(address indexed target)
func (_Contract *ContractFilterer) WatchUnBlackListed(opts *bind.WatchOpts, sink chan<- *ContractUnBlackListed, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "UnBlackListed", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUnBlackListed)
				if err := _Contract.contract.UnpackLog(event, "UnBlackListed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnBlackListed is a log parse operation binding the contract event 0xe0db3499b7fdc3da4cddff5f45d694549c19835e7f719fb5606d3ad1a5de4011.
//
// Solidity: event UnBlackListed(address indexed target)
func (_Contract *ContractFilterer) ParseUnBlackListed(log types.Log) (*ContractUnBlackListed, error) {
	event := new(ContractUnBlackListed)
	if err := _Contract.contract.UnpackLog(event, "UnBlackListed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractUndelegatedIterator is returned from FilterUndelegated and is used to iterate over the raw logs and unpacked data for Undelegated events raised by the Contract contract.
type ContractUndelegatedIterator struct {
	Event *ContractUndelegated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractUndelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUndelegated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractUndelegated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractUndelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUndelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUndelegated represents a Undelegated event raised by the Contract contract.
type ContractUndelegated struct {
	OperatorAddress common.Address
	Delegator       common.Address
	Shares          *big.Int
	BnbAmount       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUndelegated is a free log retrieval operation binding the contract event 0x3aace7340547de7b9156593a7652dc07ee900cea3fd8f82cb6c9d38b40829802.
//
// Solidity: event Undelegated(address indexed operatorAddress, address indexed delegator, uint256 shares, uint256 bnbAmount)
func (_Contract *ContractFilterer) FilterUndelegated(opts *bind.FilterOpts, operatorAddress []common.Address, delegator []common.Address) (*ContractUndelegatedIterator, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Undelegated", operatorAddressRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractUndelegatedIterator{contract: _Contract.contract, event: "Undelegated", logs: logs, sub: sub}, nil
}

// WatchUndelegated is a free log subscription operation binding the contract event 0x3aace7340547de7b9156593a7652dc07ee900cea3fd8f82cb6c9d38b40829802.
//
// Solidity: event Undelegated(address indexed operatorAddress, address indexed delegator, uint256 shares, uint256 bnbAmount)
func (_Contract *ContractFilterer) WatchUndelegated(opts *bind.WatchOpts, sink chan<- *ContractUndelegated, operatorAddress []common.Address, delegator []common.Address) (event.Subscription, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Undelegated", operatorAddressRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUndelegated)
				if err := _Contract.contract.UnpackLog(event, "Undelegated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUndelegated is a log parse operation binding the contract event 0x3aace7340547de7b9156593a7652dc07ee900cea3fd8f82cb6c9d38b40829802.
//
// Solidity: event Undelegated(address indexed operatorAddress, address indexed delegator, uint256 shares, uint256 bnbAmount)
func (_Contract *ContractFilterer) ParseUndelegated(log types.Log) (*ContractUndelegated, error) {
	event := new(ContractUndelegated)
	if err := _Contract.contract.UnpackLog(event, "Undelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractUnexpectedPackageIterator is returned from FilterUnexpectedPackage and is used to iterate over the raw logs and unpacked data for UnexpectedPackage events raised by the Contract contract.
type ContractUnexpectedPackageIterator struct {
	Event *ContractUnexpectedPackage // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractUnexpectedPackageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUnexpectedPackage)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractUnexpectedPackage)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractUnexpectedPackageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUnexpectedPackageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUnexpectedPackage represents a UnexpectedPackage event raised by the Contract contract.
type ContractUnexpectedPackage struct {
	ChannelId uint8
	MsgBytes  []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnexpectedPackage is a free log retrieval operation binding the contract event 0xaa5ba621c8b3d7d05bb9e51a7506108251d4d5dbe542ca66fc7bb52aacb02b65.
//
// Solidity: event UnexpectedPackage(uint8 channelId, bytes msgBytes)
func (_Contract *ContractFilterer) FilterUnexpectedPackage(opts *bind.FilterOpts) (*ContractUnexpectedPackageIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "UnexpectedPackage")
	if err != nil {
		return nil, err
	}
	return &ContractUnexpectedPackageIterator{contract: _Contract.contract, event: "UnexpectedPackage", logs: logs, sub: sub}, nil
}

// WatchUnexpectedPackage is a free log subscription operation binding the contract event 0xaa5ba621c8b3d7d05bb9e51a7506108251d4d5dbe542ca66fc7bb52aacb02b65.
//
// Solidity: event UnexpectedPackage(uint8 channelId, bytes msgBytes)
func (_Contract *ContractFilterer) WatchUnexpectedPackage(opts *bind.WatchOpts, sink chan<- *ContractUnexpectedPackage) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "UnexpectedPackage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUnexpectedPackage)
				if err := _Contract.contract.UnpackLog(event, "UnexpectedPackage", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnexpectedPackage is a log parse operation binding the contract event 0xaa5ba621c8b3d7d05bb9e51a7506108251d4d5dbe542ca66fc7bb52aacb02b65.
//
// Solidity: event UnexpectedPackage(uint8 channelId, bytes msgBytes)
func (_Contract *ContractFilterer) ParseUnexpectedPackage(log types.Log) (*ContractUnexpectedPackage, error) {
	event := new(ContractUnexpectedPackage)
	if err := _Contract.contract.UnpackLog(event, "UnexpectedPackage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractValidatorCreatedIterator is returned from FilterValidatorCreated and is used to iterate over the raw logs and unpacked data for ValidatorCreated events raised by the Contract contract.
type ContractValidatorCreatedIterator struct {
	Event *ContractValidatorCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractValidatorCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractValidatorCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractValidatorCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractValidatorCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractValidatorCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractValidatorCreated represents a ValidatorCreated event raised by the Contract contract.
type ContractValidatorCreated struct {
	ConsensusAddress common.Address
	OperatorAddress  common.Address
	CreditContract   common.Address
	VoteAddress      []byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterValidatorCreated is a free log retrieval operation binding the contract event 0xaecd9fb95e79c75a3a1de93362c6be5fe6ab65770d8614be583884161cd8228d.
//
// Solidity: event ValidatorCreated(address indexed consensusAddress, address indexed operatorAddress, address indexed creditContract, bytes voteAddress)
func (_Contract *ContractFilterer) FilterValidatorCreated(opts *bind.FilterOpts, consensusAddress []common.Address, operatorAddress []common.Address, creditContract []common.Address) (*ContractValidatorCreatedIterator, error) {

	var consensusAddressRule []interface{}
	for _, consensusAddressItem := range consensusAddress {
		consensusAddressRule = append(consensusAddressRule, consensusAddressItem)
	}
	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}
	var creditContractRule []interface{}
	for _, creditContractItem := range creditContract {
		creditContractRule = append(creditContractRule, creditContractItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ValidatorCreated", consensusAddressRule, operatorAddressRule, creditContractRule)
	if err != nil {
		return nil, err
	}
	return &ContractValidatorCreatedIterator{contract: _Contract.contract, event: "ValidatorCreated", logs: logs, sub: sub}, nil
}

// WatchValidatorCreated is a free log subscription operation binding the contract event 0xaecd9fb95e79c75a3a1de93362c6be5fe6ab65770d8614be583884161cd8228d.
//
// Solidity: event ValidatorCreated(address indexed consensusAddress, address indexed operatorAddress, address indexed creditContract, bytes voteAddress)
func (_Contract *ContractFilterer) WatchValidatorCreated(opts *bind.WatchOpts, sink chan<- *ContractValidatorCreated, consensusAddress []common.Address, operatorAddress []common.Address, creditContract []common.Address) (event.Subscription, error) {

	var consensusAddressRule []interface{}
	for _, consensusAddressItem := range consensusAddress {
		consensusAddressRule = append(consensusAddressRule, consensusAddressItem)
	}
	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}
	var creditContractRule []interface{}
	for _, creditContractItem := range creditContract {
		creditContractRule = append(creditContractRule, creditContractItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ValidatorCreated", consensusAddressRule, operatorAddressRule, creditContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractValidatorCreated)
				if err := _Contract.contract.UnpackLog(event, "ValidatorCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorCreated is a log parse operation binding the contract event 0xaecd9fb95e79c75a3a1de93362c6be5fe6ab65770d8614be583884161cd8228d.
//
// Solidity: event ValidatorCreated(address indexed consensusAddress, address indexed operatorAddress, address indexed creditContract, bytes voteAddress)
func (_Contract *ContractFilterer) ParseValidatorCreated(log types.Log) (*ContractValidatorCreated, error) {
	event := new(ContractValidatorCreated)
	if err := _Contract.contract.UnpackLog(event, "ValidatorCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractValidatorEmptyJailedIterator is returned from FilterValidatorEmptyJailed and is used to iterate over the raw logs and unpacked data for ValidatorEmptyJailed events raised by the Contract contract.
type ContractValidatorEmptyJailedIterator struct {
	Event *ContractValidatorEmptyJailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractValidatorEmptyJailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractValidatorEmptyJailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractValidatorEmptyJailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractValidatorEmptyJailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractValidatorEmptyJailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractValidatorEmptyJailed represents a ValidatorEmptyJailed event raised by the Contract contract.
type ContractValidatorEmptyJailed struct {
	OperatorAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterValidatorEmptyJailed is a free log retrieval operation binding the contract event 0x2afdc18061ac21cff7d9f11527ab9c8dec6fabd4edf6f894ed634bebd6a20d45.
//
// Solidity: event ValidatorEmptyJailed(address indexed operatorAddress)
func (_Contract *ContractFilterer) FilterValidatorEmptyJailed(opts *bind.FilterOpts, operatorAddress []common.Address) (*ContractValidatorEmptyJailedIterator, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ValidatorEmptyJailed", operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractValidatorEmptyJailedIterator{contract: _Contract.contract, event: "ValidatorEmptyJailed", logs: logs, sub: sub}, nil
}

// WatchValidatorEmptyJailed is a free log subscription operation binding the contract event 0x2afdc18061ac21cff7d9f11527ab9c8dec6fabd4edf6f894ed634bebd6a20d45.
//
// Solidity: event ValidatorEmptyJailed(address indexed operatorAddress)
func (_Contract *ContractFilterer) WatchValidatorEmptyJailed(opts *bind.WatchOpts, sink chan<- *ContractValidatorEmptyJailed, operatorAddress []common.Address) (event.Subscription, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ValidatorEmptyJailed", operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractValidatorEmptyJailed)
				if err := _Contract.contract.UnpackLog(event, "ValidatorEmptyJailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorEmptyJailed is a log parse operation binding the contract event 0x2afdc18061ac21cff7d9f11527ab9c8dec6fabd4edf6f894ed634bebd6a20d45.
//
// Solidity: event ValidatorEmptyJailed(address indexed operatorAddress)
func (_Contract *ContractFilterer) ParseValidatorEmptyJailed(log types.Log) (*ContractValidatorEmptyJailed, error) {
	event := new(ContractValidatorEmptyJailed)
	if err := _Contract.contract.UnpackLog(event, "ValidatorEmptyJailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractValidatorJailedIterator is returned from FilterValidatorJailed and is used to iterate over the raw logs and unpacked data for ValidatorJailed events raised by the Contract contract.
type ContractValidatorJailedIterator struct {
	Event *ContractValidatorJailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractValidatorJailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractValidatorJailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractValidatorJailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractValidatorJailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractValidatorJailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractValidatorJailed represents a ValidatorJailed event raised by the Contract contract.
type ContractValidatorJailed struct {
	OperatorAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterValidatorJailed is a free log retrieval operation binding the contract event 0x4905ac32602da3fb8b4b7b00c285e5fc4c6c2308cc908b4a1e4e9625a29c90a3.
//
// Solidity: event ValidatorJailed(address indexed operatorAddress)
func (_Contract *ContractFilterer) FilterValidatorJailed(opts *bind.FilterOpts, operatorAddress []common.Address) (*ContractValidatorJailedIterator, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ValidatorJailed", operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractValidatorJailedIterator{contract: _Contract.contract, event: "ValidatorJailed", logs: logs, sub: sub}, nil
}

// WatchValidatorJailed is a free log subscription operation binding the contract event 0x4905ac32602da3fb8b4b7b00c285e5fc4c6c2308cc908b4a1e4e9625a29c90a3.
//
// Solidity: event ValidatorJailed(address indexed operatorAddress)
func (_Contract *ContractFilterer) WatchValidatorJailed(opts *bind.WatchOpts, sink chan<- *ContractValidatorJailed, operatorAddress []common.Address) (event.Subscription, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ValidatorJailed", operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractValidatorJailed)
				if err := _Contract.contract.UnpackLog(event, "ValidatorJailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorJailed is a log parse operation binding the contract event 0x4905ac32602da3fb8b4b7b00c285e5fc4c6c2308cc908b4a1e4e9625a29c90a3.
//
// Solidity: event ValidatorJailed(address indexed operatorAddress)
func (_Contract *ContractFilterer) ParseValidatorJailed(log types.Log) (*ContractValidatorJailed, error) {
	event := new(ContractValidatorJailed)
	if err := _Contract.contract.UnpackLog(event, "ValidatorJailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractValidatorSlashedIterator is returned from FilterValidatorSlashed and is used to iterate over the raw logs and unpacked data for ValidatorSlashed events raised by the Contract contract.
type ContractValidatorSlashedIterator struct {
	Event *ContractValidatorSlashed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractValidatorSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractValidatorSlashed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractValidatorSlashed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractValidatorSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractValidatorSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractValidatorSlashed represents a ValidatorSlashed event raised by the Contract contract.
type ContractValidatorSlashed struct {
	OperatorAddress common.Address
	JailUntil       *big.Int
	SlashAmount     *big.Int
	SlashType       uint8
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterValidatorSlashed is a free log retrieval operation binding the contract event 0x6e9a2ee7aee95665e3a774a212eb11441b217e3e4656ab9563793094689aabb2.
//
// Solidity: event ValidatorSlashed(address indexed operatorAddress, uint256 jailUntil, uint256 slashAmount, uint8 slashType)
func (_Contract *ContractFilterer) FilterValidatorSlashed(opts *bind.FilterOpts, operatorAddress []common.Address) (*ContractValidatorSlashedIterator, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ValidatorSlashed", operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractValidatorSlashedIterator{contract: _Contract.contract, event: "ValidatorSlashed", logs: logs, sub: sub}, nil
}

// WatchValidatorSlashed is a free log subscription operation binding the contract event 0x6e9a2ee7aee95665e3a774a212eb11441b217e3e4656ab9563793094689aabb2.
//
// Solidity: event ValidatorSlashed(address indexed operatorAddress, uint256 jailUntil, uint256 slashAmount, uint8 slashType)
func (_Contract *ContractFilterer) WatchValidatorSlashed(opts *bind.WatchOpts, sink chan<- *ContractValidatorSlashed, operatorAddress []common.Address) (event.Subscription, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ValidatorSlashed", operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractValidatorSlashed)
				if err := _Contract.contract.UnpackLog(event, "ValidatorSlashed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorSlashed is a log parse operation binding the contract event 0x6e9a2ee7aee95665e3a774a212eb11441b217e3e4656ab9563793094689aabb2.
//
// Solidity: event ValidatorSlashed(address indexed operatorAddress, uint256 jailUntil, uint256 slashAmount, uint8 slashType)
func (_Contract *ContractFilterer) ParseValidatorSlashed(log types.Log) (*ContractValidatorSlashed, error) {
	event := new(ContractValidatorSlashed)
	if err := _Contract.contract.UnpackLog(event, "ValidatorSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractValidatorUnjailedIterator is returned from FilterValidatorUnjailed and is used to iterate over the raw logs and unpacked data for ValidatorUnjailed events raised by the Contract contract.
type ContractValidatorUnjailedIterator struct {
	Event *ContractValidatorUnjailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractValidatorUnjailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractValidatorUnjailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractValidatorUnjailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractValidatorUnjailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractValidatorUnjailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractValidatorUnjailed represents a ValidatorUnjailed event raised by the Contract contract.
type ContractValidatorUnjailed struct {
	OperatorAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterValidatorUnjailed is a free log retrieval operation binding the contract event 0x9390b453426557da5ebdc31f19a37753ca04addf656d32f35232211bb2af3f19.
//
// Solidity: event ValidatorUnjailed(address indexed operatorAddress)
func (_Contract *ContractFilterer) FilterValidatorUnjailed(opts *bind.FilterOpts, operatorAddress []common.Address) (*ContractValidatorUnjailedIterator, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ValidatorUnjailed", operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractValidatorUnjailedIterator{contract: _Contract.contract, event: "ValidatorUnjailed", logs: logs, sub: sub}, nil
}

// WatchValidatorUnjailed is a free log subscription operation binding the contract event 0x9390b453426557da5ebdc31f19a37753ca04addf656d32f35232211bb2af3f19.
//
// Solidity: event ValidatorUnjailed(address indexed operatorAddress)
func (_Contract *ContractFilterer) WatchValidatorUnjailed(opts *bind.WatchOpts, sink chan<- *ContractValidatorUnjailed, operatorAddress []common.Address) (event.Subscription, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ValidatorUnjailed", operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractValidatorUnjailed)
				if err := _Contract.contract.UnpackLog(event, "ValidatorUnjailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorUnjailed is a log parse operation binding the contract event 0x9390b453426557da5ebdc31f19a37753ca04addf656d32f35232211bb2af3f19.
//
// Solidity: event ValidatorUnjailed(address indexed operatorAddress)
func (_Contract *ContractFilterer) ParseValidatorUnjailed(log types.Log) (*ContractValidatorUnjailed, error) {
	event := new(ContractValidatorUnjailed)
	if err := _Contract.contract.UnpackLog(event, "ValidatorUnjailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractVoteAddressEditedIterator is returned from FilterVoteAddressEdited and is used to iterate over the raw logs and unpacked data for VoteAddressEdited events raised by the Contract contract.
type ContractVoteAddressEditedIterator struct {
	Event *ContractVoteAddressEdited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractVoteAddressEditedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractVoteAddressEdited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractVoteAddressEdited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractVoteAddressEditedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractVoteAddressEditedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractVoteAddressEdited represents a VoteAddressEdited event raised by the Contract contract.
type ContractVoteAddressEdited struct {
	OperatorAddress common.Address
	NewVoteAddress  []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterVoteAddressEdited is a free log retrieval operation binding the contract event 0x783156582145bd0ff7924fae6953ba054cf1233eb60739a200ddb10de068ff0d.
//
// Solidity: event VoteAddressEdited(address indexed operatorAddress, bytes newVoteAddress)
func (_Contract *ContractFilterer) FilterVoteAddressEdited(opts *bind.FilterOpts, operatorAddress []common.Address) (*ContractVoteAddressEditedIterator, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "VoteAddressEdited", operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractVoteAddressEditedIterator{contract: _Contract.contract, event: "VoteAddressEdited", logs: logs, sub: sub}, nil
}

// WatchVoteAddressEdited is a free log subscription operation binding the contract event 0x783156582145bd0ff7924fae6953ba054cf1233eb60739a200ddb10de068ff0d.
//
// Solidity: event VoteAddressEdited(address indexed operatorAddress, bytes newVoteAddress)
func (_Contract *ContractFilterer) WatchVoteAddressEdited(opts *bind.WatchOpts, sink chan<- *ContractVoteAddressEdited, operatorAddress []common.Address) (event.Subscription, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "VoteAddressEdited", operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractVoteAddressEdited)
				if err := _Contract.contract.UnpackLog(event, "VoteAddressEdited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVoteAddressEdited is a log parse operation binding the contract event 0x783156582145bd0ff7924fae6953ba054cf1233eb60739a200ddb10de068ff0d.
//
// Solidity: event VoteAddressEdited(address indexed operatorAddress, bytes newVoteAddress)
func (_Contract *ContractFilterer) ParseVoteAddressEdited(log types.Log) (*ContractVoteAddressEdited, error) {
	event := new(ContractVoteAddressEdited)
	if err := _Contract.contract.UnpackLog(event, "VoteAddressEdited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
