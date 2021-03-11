package staking

import (
	"math/big"
	"sort"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	// default element size for address, uint array
	defaultElementSize = 1
)

// stateDBStakingCaller creates a wrapper with statedb to implements ContractCaller
type stateDBStakingCaller struct {
	stateDB *state.StateDB
	config  *IndexConfigs
}

func NewStateDBStakingCaller(state *state.StateDB, cfg *IndexConfigs) StakingCaller {
	return &stateDBStakingCaller{
		stateDB: state,
		config:  cfg,
	}
}

func (s *stateDBStakingCaller) GetValidators(stakingContractAddr common.Address) ([]common.Address, error) {
	if codes := s.stateDB.GetCode(stakingContractAddr); len(codes) == 0 {
		return nil, bind.ErrNoCode
	}

	candidates, err := s.GetCandidates(stakingContractAddr)
	if err != nil {
		return nil, err
	}

	var (
		validators []common.Address
		stakes     = make(map[common.Address]*big.Int)
	)

	minValStake := s.GetMinValidatorStake(stakingContractAddr)
	for _, candidate := range candidates {
		stake := s.GetCandidateStake(stakingContractAddr, candidate)
		if stake.Cmp(minValStake) < 0 {
			continue
		}
		validators = append(validators, candidate)
		stakes[candidate] = stake
	}

	maxValSize := int(s.GetMaxValidatorSize(stakingContractAddr).Uint64())
	if len(validators) <= maxValSize {
		return validators, nil
	}

	sort.Slice(validators, func(i, j int) bool {
		if stakes[validators[i]].Cmp(stakes[validators[j]]) == 0 {
			return strings.Compare(validators[i].String(), validators[j].String()) > 0
		}
		return stakes[validators[i]].Cmp(stakes[validators[j]]) > 0
	})
	return candidates[:maxValSize], nil
}

func (s *stateDBStakingCaller) GetValidatorsData(stakingContractAddr common.Address, candidates []common.Address) (map[common.Address]CandidateData, error) {
	allValidatorsData := make(map[common.Address]CandidateData)
	for _, candidate := range candidates {
		candidateData := s.GetCandidateData(stakingContractAddr, candidate)
		allValidatorsData[candidate] = candidateData
	}
	return allValidatorsData, nil
}

func (s *stateDBStakingCaller) GetCandidates(stakingContractAddr common.Address) ([]common.Address, error) {
	slotHash := s.config.CandidatesLayout.slotHash()
	candidatesNum := s.getBigInt(stakingContractAddr, slotHash).Uint64()
	if candidatesNum == 0 {
		return nil, ErrEmptyValidatorSet
	}
	candidates := make([]common.Address, candidatesNum)

	for i := uint64(0); i < candidatesNum; i++ {
		ret := s.stateDB.GetState(stakingContractAddr, getElementArrayLoc(slotHash, i, defaultElementSize))
		candidates[i] = common.BytesToAddress(ret.Bytes())
	}
	return candidates, nil
}

func (s *stateDBStakingCaller) GetCandidateData(stakingContractAddr common.Address, candidate common.Address) CandidateData {
	loc := getMappingElementLoc(s.config.CandidateDataLayout.slotHash(), candidate.Hash())
	totalStakeLoc := addOffsetToLoc(loc, new(big.Int).SetUint64(s.config.CandidateDataStruct.TotalStakeLayout.Slot))
	totalStake := s.stateDB.GetState(stakingContractAddr, totalStakeLoc).Big()

	ownerLoc := addOffsetToLoc(loc, new(big.Int).SetUint64(s.config.CandidateDataStruct.OwnerLayout.Slot))
	owner := common.BytesToAddress(s.stateDB.GetState(stakingContractAddr, ownerLoc).Bytes())

	voterStakesLoc := addOffsetToLoc(loc, new(big.Int).SetUint64(s.config.CandidateDataStruct.VoterStakeLayout.Slot))
	voterStakes := make(map[common.Address]*big.Int)
	for _, voter := range s.GetVoters(stakingContractAddr, candidate) {
		voterStakeLoc := getMappingElementLoc(voterStakesLoc, voter.Hash())
		voterStakes[voter] = s.getBigInt(stakingContractAddr, voterStakeLoc)
	}
	return CandidateData{
		Owner:       owner,
		TotalStake:  totalStake,
		VoterStakes: voterStakes,
	}
}

func (s *stateDBStakingCaller) GetVoters(stakingContractAddr common.Address, candidate common.Address) []common.Address {
	votersArrLoc := getMappingElementLoc(s.config.CandidateVotersLayout.slotHash(), candidate.Hash())
	votersLength := s.getBigInt(stakingContractAddr, votersArrLoc).Uint64()
	voters := make([]common.Address, votersLength)
	for i := uint64(0); i < votersLength; i++ {
		elemLoc := getElementArrayLoc(votersArrLoc, i, defaultElementSize)
		voters[i] = s.getAddress(stakingContractAddr, elemLoc)
	}
	return voters
}

func (s *stateDBStakingCaller) GetCandidateStake(stakingContractAddr common.Address, candidate common.Address) *big.Int {
	loc := getMappingElementLoc(s.config.CandidateDataLayout.slotHash(), candidate.Hash())
	totalStakeLoc := addOffsetToLoc(loc, new(big.Int).SetUint64(s.config.CandidateDataStruct.TotalStakeLayout.Slot))
	return s.stateDB.GetState(stakingContractAddr, totalStakeLoc).Big()
}

func (s *stateDBStakingCaller) GetStartBlock(stakingContractAddr common.Address) *big.Int {
	return s.getBigInt(stakingContractAddr, s.config.StartBlockLayout.slotHash())
}

func (s *stateDBStakingCaller) GetEpochPeriod(stakingContractAddr common.Address) *big.Int {
	return s.getBigInt(stakingContractAddr, s.config.EpochPeriodLayout.slotHash())
}

func (s *stateDBStakingCaller) GetMaxValidatorSize(stakingContractAddr common.Address) *big.Int {
	return s.getBigInt(stakingContractAddr, s.config.MaxValidatorSizeLayout.slotHash())
}

func (s *stateDBStakingCaller) GetMinValidatorStake(stakingContractAddr common.Address) *big.Int {
	return s.getBigInt(stakingContractAddr, s.config.MinValidatorStakeLayout.slotHash())
}

func (s *stateDBStakingCaller) GetMinVoterCap(stakingContractAddr common.Address) *big.Int {
	return s.getBigInt(stakingContractAddr, s.config.MinVoterCapLayout.slotHash())
}

func (s *stateDBStakingCaller) GetAdmin(stakingContractAddr common.Address) common.Address {
	return s.getAddress(stakingContractAddr, s.config.AdminLayout.slotHash())
}

func (s *stateDBStakingCaller) GetCandidateOwner(stakingContractAddr common.Address, candidate common.Address) common.Address {
	loc := getMappingElementLoc(s.config.CandidateDataLayout.slotHash(), candidate.Hash())
	ownerLoc := addOffsetToLoc(loc, new(big.Int).SetUint64(s.config.CandidateDataStruct.OwnerLayout.Slot))
	return s.getAddress(stakingContractAddr, ownerLoc)
}

func (s *stateDBStakingCaller) getAddress(contractAddr common.Address, hash common.Hash) common.Address {
	return common.HexToAddress(s.stateDB.GetState(contractAddr, hash).Hex())
}

func (s *stateDBStakingCaller) getBigInt(contractAddr common.Address, hash common.Hash) *big.Int {
	return s.stateDB.GetState(contractAddr, hash).Big()
}

/**
 * The value to a mapping key k at position p is located at keccak256(k . p) where . is concatenation.
 */
func getElementArrayLoc(slotHash common.Hash, index uint64, elementSize uint64) common.Hash {
	slotKecBig := crypto.Keccak256Hash(slotHash.Bytes()).Big()
	arrBig := slotKecBig.Add(slotKecBig, new(big.Int).SetUint64(index*elementSize))
	return common.BigToHash(arrBig)
}

/**
 * Get the position for a field inside a struct
 */
func getMappingElementLoc(slotHash common.Hash, key common.Hash) common.Hash {
	return common.BytesToHash(crypto.Keccak256(key.Bytes(), slotHash.Bytes()))
}

/**
 * Get the position for a field inside a struct
 */
func addOffsetToLoc(slotHash common.Hash, slot *big.Int) common.Hash {
	rootBig := slotHash.Big()
	arrBig := new(big.Int).Add(rootBig, slot)
	return common.BigToHash(arrBig)
}
