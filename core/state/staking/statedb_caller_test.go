package staking

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/istanbul/staking_contracts"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
)

const (
	gasLimit = 10000000
)

func getRandomAddress(t *testing.T) common.Address {
	privateKey, err := crypto.GenerateKey()
	NoError(t, err)
	return crypto.PubkeyToAddress(privateKey.PublicKey)
}

func NoError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Get error: %v", err)
	}
}

func getCandidateDataExp(validates []common.Address, owners []common.Address, stake *big.Int) (map[common.Address]CandidateData, error) {
	if len(validates) != len(owners) {
		return nil, errors.New("length of validates is not equal the length owner")
	}
	res := make(map[common.Address]CandidateData)
	for i := 0; i < len(validates); i++ {
		res[validates[i]] = CandidateData{
			Owner:       owners[i],
			TotalStake:  stake,
			VoterStakes: map[common.Address]*big.Int{owners[i]: stake},
		}
	}
	return res, nil
}

func TestCheckIndex(t *testing.T) {
	var (
		candidatesExp = []common.Address{
			getRandomAddress(t),
			getRandomAddress(t),
			getRandomAddress(t),
		}
		ownersExp = []common.Address{
			getRandomAddress(t),
			getRandomAddress(t),
			getRandomAddress(t),
		}

		epochExp             = big.NewInt(300000)
		startBlockExp        = big.NewInt(1)
		maxValidatorSizeExp  = big.NewInt(3)
		minValidatorStakeExp = big.NewInt(10000000000000)
		minVoterCapExp       = big.NewInt(2)
		adminAddrExp         = getRandomAddress(t)
	)
	validatorsDataExp, err := getCandidateDataExp(candidatesExp, ownersExp, minValidatorStakeExp)
	NoError(t, err)
	privateKey, err := crypto.GenerateKey()
	NoError(t, err)
	addr := crypto.PubkeyToAddress(privateKey.PublicKey)

	be := backends.NewSimulatedBackend(
		core.GenesisAlloc{
			addr: core.GenesisAccount{
				Balance: big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil),
			},
		}, gasLimit)

	transOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, params.AllEthashProtocolChanges.ChainID)
	NoError(t, err)
	transOpts.Nonce = big.NewInt(0)

	scAddress, tx, _, err := staking_contracts.DeployAsDBChainStaking(transOpts, be, candidatesExp, ownersExp, epochExp, startBlockExp, maxValidatorSizeExp, minValidatorStakeExp, minVoterCapExp, adminAddrExp)
	NoError(t, err)

	be.Commit()

	receipt, err := be.TransactionReceipt(context.Background(), tx.Hash())
	NoError(t, err)
	if receipt.Status != uint64(1) {
		t.Error("Deploy contract failed")
	}

	stateDB, err := be.Blockchain().State()
	NoError(t, err)

	caller := &stateDBStakingCaller{
		stateDB: stateDB,
		config:  DefaultConfig,
	}

	// Check startBlock
	startBlock := caller.GetStartBlock(scAddress)
	if startBlock.Cmp(startBlockExp) != 0 {
		t.Errorf("Get startBlock error: have %v, want  %v", startBlock.String(), startBlockExp.String())
	}

	// Check Epoch
	epoch := caller.GetEpochPeriod(scAddress)
	if epoch.Cmp(epochExp) != 0 {
		t.Errorf("Get Epoch error: have %v, want  %v", epoch.String(), epochExp.String())
	}

	// Check maxValidatorSize
	maxValidatorSize := caller.GetMaxValidatorSize(scAddress)
	if maxValidatorSize.Cmp(maxValidatorSizeExp) != 0 {
		t.Errorf("Get maxValidatorSize error: have %v, want  %v", maxValidatorSize.String(), maxValidatorSizeExp.String())
	}

	// Check minValidatorStake
	minValidatorStake := caller.GetMinValidatorStake(scAddress)
	if minValidatorStake.Cmp(minValidatorStakeExp) != 0 {
		t.Errorf("Get minValidatorStake error: have %v, want  %v", minValidatorStake.String(), minValidatorStakeExp.String())
	}

	// Check minVoteCap
	minVoterCap := caller.GetMinVoterCap(scAddress)
	if minVoterCap.Cmp(minVoterCapExp) != 0 {
		t.Errorf("Get minVoteCap error: have %v, want  %v", minVoterCap.String(), minVoterCapExp.String())
	}

	// check admin address
	adminAddr := caller.GetAdmin(scAddress)
	if !reflect.DeepEqual(adminAddr, adminAddrExp) {
		t.Errorf("Get admin address error: have %v, want  %v", adminAddr.String(), adminAddrExp.String())
	}

	// check  function GetValidators
	candidates, err := caller.GetValidators(scAddress)
	NoError(t, err)

	if len(candidates) != len(candidatesExp) {
		t.Errorf("Validators length is not equal, error: have %v, want %v", len(candidates), len(candidatesExp))
	}

	sort.Slice(candidates, func(i, j int) bool {
		return strings.Compare(candidates[i].String(), candidates[j].String()) < 0
	})
	sort.Slice(candidatesExp, func(i, j int) bool {
		return strings.Compare(candidatesExp[i].String(), candidatesExp[j].String()) < 0
	})

	if !reflect.DeepEqual(candidates, candidatesExp) {
		t.Errorf("validator check failed, error: have %v, want %v", candidates, candidatesExp)
	}

	// check  function GetValidatorsData
	validatorsData, err := caller.GetValidatorsData(scAddress, candidates)
	NoError(t, err)
	if len(validatorsData) != len(validatorsDataExp) {
		t.Errorf("candidateData length is not equal, error: have %v, want %v", len(validatorsData), len(validatorsDataExp))
	}

	for addr, validatorData := range validatorsData {
		if validatorDataExp, ok := validatorsDataExp[addr]; ok {
			if validatorData.TotalStake.Cmp(validatorDataExp.TotalStake) != 0 {
				t.Errorf("Check validatorData.TotalStake at address %v, error: have %v, want %v", addr,
					validatorData.TotalStake, validatorDataExp.TotalStake)
			}
			if !reflect.DeepEqual(validatorData.Owner, validatorDataExp.Owner) {
				t.Errorf("Check validatorData.Owner at address %v, error: have %v, want %v", addr,
					validatorData.Owner, validatorDataExp.Owner)

			}
			voteStakeExpStr := fmt.Sprintf("%v", validatorDataExp.VoterStakes)
			voteStakeStr := fmt.Sprintf("%v", validatorData.VoterStakes)
			if voteStakeExpStr != voteStakeStr {
				t.Errorf("Check validatorData.VoterStakes at address %v, error: have %v, want %v", addr,
					voteStakeStr, voteStakeExpStr)
			}
		}
	}
}
