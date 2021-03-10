// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"strings"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/compiler"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
)

const (
	stakingSCName            = "AsDBChainStaking"
	simulatedGasLimit uint64 = 500000000
	GasPriceConfig           = 1000000000
	simulatedBalance         = simulatedGasLimit * GasPriceConfig
)

func (w *wizard) configStakingSC(genesis *core.Genesis, validators []common.Address) error {
	var (
		scPath            string
		stakingSCParmams  []interface{}
		expectedSCAddress *common.Address
		byteCodeString    string
		abiSC             *abi.ABI
	)

	fmt.Println()
	fmt.Println("Do you want to use precompile Staking Contract file? (default = yes)")
	if usePrecompiledSc := w.readDefaultYesNo(true); usePrecompiledSc {
		fmt.Println("Specify your ByteCode file path (default = ./consensus/istanbul/staking_contract/bin/AsDBChainStaking.bin)")
		for {
			if tempValue, err := readFile(w.readDefaultString("./consensus/istanbul/staking_contract/bin/AsDBChainStaking.bin")); err != nil {
				log.Error("Failed to read ByteCode file", "error", err)
			} else {
				byteCodeString = tempValue
				break
			}
		}

		fmt.Println("Specify your ABI path (default = ./consensus/istanbul/staking_contract/bin/AsDBChainStaking.abi)")
		for {
			if tempValue, err := readFile(w.readDefaultString("./consensus/istanbul/staking_contract/bin/AsDBChainStaking.abi")); err != nil {
				log.Error("Failed to read ABI file", "error", err)
			} else {
				if pasedABI, err := abi.JSON(strings.NewReader(tempValue)); err != nil {
					log.Error("Failed to decode input ABI file", "error", err)
				} else {
					abiSC = &pasedABI
				}
			}
		}
	} else {
		fmt.Println()
		fmt.Println("Specify your staking smart contract path (default = ./consensus/istanbul/staking_contract/contracts/AsDBChainStaking.sol)")
		for {
			if scPath = w.readDefaultString("./consensus/istanbul/staking_contract/contracts/AsDBChainStaking.sol"); len(scPath) > 0 {
				break
			}
		}

		//Compile SC file to get Bytecode, ABI
		if compiledByteCode, compiledABI, err := compileSCFile(scPath); err != nil {
			return err
		} else {
			byteCodeString = compiledByteCode
			abiSC = compiledABI
		}
	}

	stakingSCParmams = w.readStakingSCParams(genesis, validators)
	fmt.Println()
	fmt.Println("What is the address of staking smart contract? (avoid special address from 0x0000000000000000000000000000000000000001 to 0x00000000000000000000000000000000000000018)")
	for {
		if expectedSCAddress = w.readAddress(); expectedSCAddress != nil {
			if _, ok := vm.PrecompiledContractsBerlin[*expectedSCAddress]; !ok {
				break
			}
			if _, ok := vm.PrecompiledContractsBLS[*expectedSCAddress]; !ok {
				break
			}
		}
	}

	genesisAccount, err := createGenesisAccountWithStakingSC(genesis, abiSC, byteCodeString, validators, stakingSCParmams)
	if err != nil {
		return err
	}
	genesis.Config.Istanbul.StakingSCAddress = expectedSCAddress
	genesis.Alloc[*expectedSCAddress] = genesisAccount
	return nil
}

func createGenesisAccountWithStakingSC(genesis *core.Genesis, abiSC *abi.ABI, byteCodeSC string, validators []common.Address,
	stakingSCParams []interface{}) (core.GenesisAccount, error) {
	contractBackend, smlSCAddress, err := deployStakingSCToSimulatedBE(genesis, abiSC, byteCodeSC, stakingSCParams)
	if err != nil {
		return core.GenesisAccount{}, err
	}
	codeOfSC, storageOfSC := getStakingSCData(contractBackend, smlSCAddress)
	minValidatorStake, ok := stakingSCParams[len(stakingSCParams)-3].(*big.Int)
	if !ok {
		return core.GenesisAccount{}, errors.New("Failed to convert interface to *big.Int")
	}

	return core.GenesisAccount{
		Balance: new(big.Int).Mul(big.NewInt(int64(len(validators))), minValidatorStake),
		Code:    codeOfSC,
		Storage: storageOfSC,
	}, nil

}

func compileSCFile(scPath string) (string, *abi.ABI, error) {
	contracts, err := compiler.CompileSolidity("solc", scPath)
	if err != nil {
		return "", nil, errors.Errorf("Failed to compile Solidity contract: %v", err)
	}
	byteCodeSC, abiSC, err := getByteCodeAndABIOfSC(fmt.Sprintf("%s:%s", scPath, stakingSCName), contracts)
	if err != nil {
		return "", nil, errors.Errorf("Failed to get Bytecode, ABI from contract: %v", err)
	}
	if len(byteCodeSC) == 0 || abiSC == nil {
		return "", nil, errors.Errorf("Not found any EvrynetStaking contract when compile SC. Error: %+v", err)
	}
	return byteCodeSC, abiSC, err
}

func getByteCodeAndABIOfSC(contractName string, contracts map[string]*compiler.Contract) (string, *abi.ABI, error) {
	var byteCodeSC string
	ct := contracts[contractName]
	if ct == nil {
		return "", nil, errors.Errorf("Not found any contract by key %s", contractName)
	}
	if byteCodeSC = ct.Code; len(byteCodeSC) == 0 {
		return "", nil, errors.New("Failed to get code of contract")
	}
	abiBytes, err := json.Marshal(ct.Info.AbiDefinition)
	if err != nil {
		return "", nil, errors.Errorf("Failed to parse ABI from compiler output: %v", err)
	}

	parsedABI, err := abi.JSON(strings.NewReader(string(abiBytes)))
	if err != nil {
		return "", nil, errors.Errorf("Failed to parse bytes to ABI: %v", err)
	}
	return byteCodeSC, &parsedABI, nil
}

func deployStakingSCToSimulatedBE(genesis *core.Genesis, parsedABI *abi.ABI, byteCode string, stakingSCParams []interface{}) (*backends.SimulatedBackend, common.Address, error) {
	pKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, common.Address{}, err
	}
	addr := crypto.PubkeyToAddress(pKey.PublicKey)
	contractBackend := backends.NewSimulatedBackend(core.GenesisAlloc{addr: {Balance: big.NewInt(int64(simulatedBalance))}}, simulatedGasLimit)

	transactOpts, err := bind.NewKeyedTransactorWithChainID(pKey, params.AllEthashProtocolChanges.ChainID)
	if err != nil {
		return nil, common.Address{}, err
	}
	smlSCAddress, _, _, err := bind.DeployContract(transactOpts, *parsedABI, common.FromHex(byteCode), contractBackend, stakingSCParams...)
	if err != nil {
		utils.Fatalf("Failed to deploy contract: %v", err)
	}
	contractBackend.Commit()

	return contractBackend, smlSCAddress, nil
}

func getStakingSCData(contractBackend *backends.SimulatedBackend, smlSCAddress common.Address) ([]byte, map[common.Hash]common.Hash) {
	codeOfStakingSC, err := contractBackend.CodeAt(context.Background(), smlSCAddress, nil)
	if err != nil || len(codeOfStakingSC) == 0 {
		utils.Fatalf("Failed to get code contract: %v", err)
	}

	// Read data of contract in statedb & put to Storage of genesis account
	storage := make(map[common.Hash]common.Hash)

	if err := contractBackend.ForEachStorageAt(context.Background(), smlSCAddress, nil, getDataForStorage(storage)); err != nil {
		utils.Fatalf("Failed to to read all keys, values in the storage: %v", err)
	}
	return codeOfStakingSC, storage

}

// readStakingSCParams returns the params to deploy staking smart-contract and writes epoch to genesis config
func (w *wizard) readStakingSCParams(genesis *core.Genesis, validators []common.Address) []interface{} {
	fmt.Println()
	fmt.Println("Input params to init staking SC:")
	fmt.Printf("- What is the address of candidates owner? (expected %d address)\n", len(validators))
	var _candidatesOwners []common.Address
	for {
		if address := w.readAddress(); address != nil {
			_candidatesOwners = append(_candidatesOwners, *address)
			continue
		}
		if len(_candidatesOwners) == len(validators) {
			break
		}
	}
	fmt.Println("- What is the admin address of staking SC?")
	_admin := w.readMandatoryAddress()
	fmt.Println("- How many blocks for epoch period? (default = 1024)")
	_epochPeriod := w.readDefaultBigInt(big.NewInt(1024))
	_startBlock := big.NewInt(0)
	fmt.Println("- What is the max size of validators? (max number of candidates to be selected as validators for producing blocks)")
	_maxValidatorSize := w.readMandatoryBigInt()
	fmt.Println("- What is the min stake of validator? (minimum (his own) stake of each candidate to become a validator (use to slash if validator is doing malicious things))")
	_minValidatorStake := w.readMandatoryBigInt()
	fmt.Println("- What is the min cap of vote? (minimum amount of EVR tokens to vote for a candidate)")
	_minVoteCap := w.readMandatoryBigInt()
	genesis.Config.Istanbul.Epoch = _epochPeriod.Uint64()
	return []interface{}{validators, _candidatesOwners, _epochPeriod, _startBlock, _maxValidatorSize, _minValidatorStake, _minVoteCap, *_admin}
}

func getDataForStorage(storage map[common.Hash]common.Hash) func(key common.Hash, val common.Hash) bool {
	return func(key, val common.Hash) bool {
		storage[key] = val
		return true
	}
}

func readFile(path string) (string, error) {
	res, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(res), nil
}
