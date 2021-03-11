package staking

import (
	"context"
	"errors"
	"math/big"
	"sort"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/istanbul/staking_contracts"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/params"
)

type evmStakingCaller struct {
	blockNumber  *big.Int
	header       *types.Header
	stateDB      *state.StateDB
	chainContext core.ChainContext
	chainConfig  *params.ChainConfig
	vmConfig     vm.Config
}

// Deprecated: Using NewStateDbStakingCaller instead of
// NewBECaller returns staking caller which reads data from staking smart-contract by execute a call from evm
func NewEVMStakingCaller(stateDB *state.StateDB, chainContext core.ChainContext, header *types.Header,
	chainConfig *params.ChainConfig, vmConfig vm.Config) StakingCaller {
	return &evmStakingCaller{
		stateDB:      stateDB,
		chainContext: chainContext,
		blockNumber:  header.Number,
		header:       header,
		chainConfig:  chainConfig,
		vmConfig:     vmConfig,
	}
}

func (e evmStakingCaller) GetValidators(stakingContractAddr common.Address) ([]common.Address, error) {
	var (
		candidatesAddr []common.Address
		stakes         = make(map[common.Address]*big.Int)
	)
	sc, err := staking_contracts.NewAsDBChainStakingCaller(stakingContractAddr, e)
	if err != nil {
		return nil, err
	}
	data, err := sc.GetListCandidates(nil)
	if err != nil {
		return nil, err
	}
	if len(data.Candidates) != len(data.Stakes) {
		return nil, ErrLengthOfCandidatesAndStakesMisMatch
	}

	minValidatorStake := data.MinValidatorCap
	for i, candidate := range data.Candidates {
		owner, err := sc.GetCandidateOwner(nil, candidate)
		if err != nil {
			return nil, err
		}
		stake, err := sc.GetVoterStake(nil, candidate, owner)
		if err != nil {
			return nil, err
		}
		if stake.Cmp(minValidatorStake) < 0 {
			continue
		}
		candidatesAddr = append(candidatesAddr, data.Candidates[i])
		stakes[data.Candidates[i]] = stake
	}

	if len(candidatesAddr) == 0 || len(stakes) == 0 {
		return nil, ErrEmptyValidatorSet
	}

	if len(candidatesAddr) < int(data.ValidatorSize.Int64()) {
		return candidatesAddr, nil
	}

	sort.Slice(candidatesAddr, func(i, j int) bool {
		if stakes[candidatesAddr[i]].Cmp(stakes[candidatesAddr[j]]) == 0 {
			return strings.Compare(candidatesAddr[i].String(), candidatesAddr[j].String()) > 0
		}
		return stakes[candidatesAddr[i]].Cmp(stakes[candidatesAddr[j]]) > 0
	})
	return candidatesAddr[:int(data.ValidatorSize.Int64())], err

}

func (e evmStakingCaller) GetValidatorsData(stakingContractAddr common.Address, candidates []common.Address) (map[common.Address]CandidateData, error) {
	sc, err := staking_contracts.NewAsDBChainStakingCaller(stakingContractAddr, e)
	if err != nil {
		return nil, err
	}
	allValidatorsData := make(map[common.Address]CandidateData)
	for _, candidate := range candidates {
		candidateData, err := sc.GetCandidateData(nil, candidate)
		if err != nil {
			return nil, err
		}
		voters, err := sc.GetVoters(nil, candidate)
		if err != nil {
			return nil, err
		}
		voterStakes, err := sc.GetVoterStakes(nil, candidate, voters)
		if err != nil {
			return nil, err
		}
		if len(voterStakes) != len(voters) {
			return nil, ErrLengthOfVotesAndStakesMisMatch
		}
		voteStakes := make(map[common.Address]*big.Int)
		for i := range voterStakes {
			voteStakes[voters[i]] = voterStakes[i]
		}
		allValidatorsData[candidate] = CandidateData{
			VoterStakes: voteStakes,
			Owner:       candidateData.Owner,
			TotalStake:  candidateData.TotalStake,
		}
	}
	return allValidatorsData, nil
}

func (e evmStakingCaller) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	return e.stateDB.GetCode(contract), nil
}

func (e evmStakingCaller) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	clonedStateDB := e.stateDB.Copy()
	if blockNumber != nil && blockNumber.Cmp(e.blockNumber) != 0 {
		return nil, errors.New("blockNumber is not supported")
	}

	if call.GasPrice == nil {
		call.GasPrice = big.NewInt(1)
	}
	if call.Gas == 0 {
		call.Gas = maxGasGetValSet
	}
	if call.Value == nil {
		call.Value = new(big.Int)
	}
	from := clonedStateDB.GetOrNewStateObject(call.From)
	from.SetBalance(math.MaxBig256)
	//core.NewEVMContext()
	msg := callMsg{call}
	txContext := core.NewEVMTxContext(msg)
	evmContext := core.NewEVMBlockContext(e.header, e.chainContext, nil)
	// Create a new environment which holds all relevant information
	// about the transaction and calling mechanisms.
	vmEnv := vm.NewEVM(evmContext, txContext, e.stateDB, e.chainConfig, vm.Config{})
	defer vmEnv.Cancel()
	gasPool := new(core.GasPool).AddGas(math.MaxUint64)
	res, err := core.NewStateTransition(vmEnv, msg, gasPool).TransitionDb()
	return res.ReturnData, err
}

type callMsg struct {
	ethereum.CallMsg
}

func (m callMsg) From() common.Address         { return m.CallMsg.From }
func (m callMsg) Nonce() uint64                { return 0 }
func (m callMsg) CheckNonce() bool             { return false }
func (m callMsg) To() *common.Address          { return m.CallMsg.To }
func (m callMsg) GasPrice() *big.Int           { return m.CallMsg.GasPrice }
func (m callMsg) Gas() uint64                  { return m.CallMsg.Gas }
func (m callMsg) Value() *big.Int              { return m.CallMsg.Value }
func (m callMsg) Data() []byte                 { return m.CallMsg.Data }
func (m callMsg) AccessList() types.AccessList { return m.CallMsg.AccessList }

type chainContextWrapper struct {
	engine      consensus.Engine
	getHeaderFn func(common.Hash, uint64) *types.Header
}

func (w *chainContextWrapper) Engine() consensus.Engine {
	return w.engine
}

func (w *chainContextWrapper) GetHeader(hash common.Hash, height uint64) *types.Header {
	return w.getHeaderFn(hash, height)
}

func NewChainContextWrapper(engine consensus.Engine, getHeaderFn func(common.Hash, uint64) *types.Header) core.ChainContext {
	return &chainContextWrapper{
		engine:      engine,
		getHeaderFn: getHeaderFn,
	}
}
