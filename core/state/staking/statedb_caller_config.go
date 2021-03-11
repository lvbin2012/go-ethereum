package staking

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// Layout represents the Offset and Slot order of a state variable
type Layout struct {
	Offset uint64
	Slot   uint64
}

// CandidateDataStructIndex represents the struct of CandidateData index information
type CandidateDataStructIndex struct {
	IsCandidateLayout Layout
	TotalStakeLayout  Layout
	OwnerLayout       Layout
	VoterStakeLayout  Layout
}

// IndexConfigs represents the configuration index of state variables.
type IndexConfigs struct {
	WithdrawsStateLayout    Layout //1
	CandidateVotersLayout   Layout //2
	CandidateDataLayout     Layout //3
	CandidatesLayout        Layout //4
	StartBlockLayout        Layout //5
	EpochPeriodLayout       Layout //6
	MaxValidatorSizeLayout  Layout //7
	MinValidatorStakeLayout Layout //8
	MinVoterCapLayout       Layout //9
	AdminLayout             Layout //10

	CandidateDataStruct CandidateDataStructIndex
}

// DefaultConfig represents he default configuration.
var DefaultConfig = &IndexConfigs{
	WithdrawsStateLayout:    NewLayout(1, 0),
	CandidateVotersLayout:   NewLayout(2, 0),
	CandidateDataLayout:     NewLayout(3, 0),
	CandidatesLayout:        NewLayout(4, 0),
	StartBlockLayout:        NewLayout(5, 0),
	EpochPeriodLayout:       NewLayout(6, 0),
	MaxValidatorSizeLayout:  NewLayout(7, 0),
	MinValidatorStakeLayout: NewLayout(8, 0),
	MinVoterCapLayout:       NewLayout(9, 0),
	AdminLayout:             NewLayout(10, 0),

	CandidateDataStruct: CandidateDataStructIndex{
		IsCandidateLayout: NewLayout(0, 0),
		TotalStakeLayout:  NewLayout(1, 0),
		OwnerLayout:       NewLayout(2, 0),
		VoterStakeLayout:  NewLayout(3, 0),
	},
}

func (layout *Layout) slotHash() common.Hash {
	return common.BigToHash(new(big.Int).SetUint64(layout.Slot))
}

// NewLayout returns a instance of Layout
func NewLayout(slot uint64, offset uint64) Layout {
	return Layout{Offset: offset, Slot: slot}
}
