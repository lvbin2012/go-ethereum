package backend

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/state/staking"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

var (
	validatorRewardPercentage int64 = 50
	voterRewardPercentage     int64 = 50
)

// AccumulateRewards credits the coinbase of the given block with the proposing
// reward.
func (sb *backend) accumulateRewards(chainReader consensus.FullChainReader, state *state.StateDB, header *types.Header) error {
	var (
		currentNumber = header.Number.Uint64()
		epoch         = chainReader.Config().Istanbul.Epoch
		start         = time.Now()
	)

	if currentNumber == 0 {
		return istanbul.ErrFinalizeZeroBlock
	}

	if currentNumber%epoch != 0 {
		return nil
	}
	validatorsRewards, err := calculateTotalValidatorsRewards(chainReader, epoch, header)
	if err != nil {
		return err
	}
	transitionHeader := chainReader.GetHeaderByNumber(currentNumber - epoch)
	snap, err := sb.snapshot(chainReader, currentNumber, header.ParentHash, nil)
	if err != nil {
		return err
	}
	stateDB, err := chainReader.StateAt(transitionHeader.Root)
	if err != nil {
		return err
	}
	stakingCaller := sb.getStakingCaller(chainReader, stateDB, header)
	validatorsData, err := stakingCaller.GetValidatorsData(*sb.config.StakingSCAddress, snap.validators())
	if err != nil {
		return err
	}

	finalReward := calculateReward(validatorsData, validatorsRewards)
	for addr, value := range finalReward {
		state.AddBalance(addr, value)
	}
	log.Debug("accumulateRewards", "number", currentNumber, "elapsed", common.PrettyDuration(time.Since(start)))
	return nil
}

// calculateTotalValidatorsRewards gets reward from chainReader and current header (from finalize)
// reward includes block rewards and tx fee from block number currentBlock - epoch +1
func calculateTotalValidatorsRewards(chainReader consensus.ChainReader, epoch uint64,
	header *types.Header) (map[common.Address]*big.Int, error) {
	current := header.Number.Uint64()
	validatorsRewards := make(map[common.Address]*big.Int)
	for i := current - epoch + 1; i <= current; i++ {
		var currentHeader *types.Header
		if i != current {
			currentHeader = chainReader.GetHeaderByNumber(i)
		} else {
			currentHeader = header
		}
		istanbulExtra, err := types.ExtractIstanbulExtra(currentHeader)
		if err != nil {
			return nil, err
		}

		if value, ok := validatorsRewards[currentHeader.Coinbase]; ok {
			validatorsRewards[currentHeader.Coinbase] = new(big.Int).Add(value, istanbulExtra.Reward)
		} else {
			validatorsRewards[currentHeader.Coinbase] = istanbulExtra.Reward
		}
	}
	return validatorsRewards, nil
}

// calculateReward divides rewards into 50% to owner and 50% among voters
// rewards for voters is proportional to voters'stake
func calculateReward(validatorsData map[common.Address]staking.CandidateData, validatorsReward map[common.Address]*big.Int) map[common.Address]*big.Int {
	finalReward := make(map[common.Address]*big.Int)
	addReward := func(addr common.Address, value *big.Int) {
		if current, ok := finalReward[addr]; ok {
			finalReward[addr] = new(big.Int).Add(current, value)
		} else {
			finalReward[addr] = new(big.Int).Set(value)
		}
	}
	for addr, validatorData := range validatorsData {
		totalReward, ok := validatorsReward[addr]
		if !ok {
			continue
		}
		// remainingReward to ensure the total reward for the voters and owner is equals to the wei validator earns
		remainingReward := new(big.Int).Set(totalReward)
		totalVoterReward := new(big.Int).Mul(totalReward, big.NewInt(voterRewardPercentage))
		totalVoterReward = new(big.Int).Div(totalVoterReward, big.NewInt(100))
		for voter, voterStake := range validatorData.VoterStakes {
			voterReward := new(big.Int).Mul(totalVoterReward, voterStake)
			voterReward = new(big.Int).Div(voterReward, validatorData.TotalStake)
			addReward(voter, voterReward)
			remainingReward.Sub(remainingReward, voterReward)
		}
		validatorReward := new(big.Int).Mul(totalReward, big.NewInt(validatorRewardPercentage))
		validatorReward = new(big.Int).Div(validatorReward, big.NewInt(100))
		addReward(validatorData.Owner, remainingReward)
	}
	return finalReward
}
