// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package staking_contracts

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AsDBChainStakingABI is the input ABI used to generate the binding from.
const AsDBChainStakingABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_candidates\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_candidateOwners\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_epochPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxValidatorSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minValidatorStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minVoteCap\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_candidate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"}],\"name\":\"Resigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Unvoted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Voted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_destAddress\",\"type\":\"address\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"candidates\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"getCandidateData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_isActiveCandidate\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_totalStake\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"getCandidateOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"getCandidateStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getListCandidates\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_candidates\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"stakes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValidatorCap\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_candidate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getVoterStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_candidate\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"voters\",\"type\":\"address[]\"}],\"name\":\"getVoterStakes\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"stakes\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"getVoters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"voters\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"getWithdrawCap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWithdrawEpochs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"epochs\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWithdrawEpochsAndCaps\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"epochs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"caps\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"isCandidate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxValidatorSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minValidatorStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minVoterCap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_candidate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"resign\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"transferAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"unvote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMaxValidatorSize\",\"type\":\"uint256\"}],\"name\":\"updateMaxValidatorSize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newCap\",\"type\":\"uint256\"}],\"name\":\"updateMinValidateStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newCap\",\"type\":\"uint256\"}],\"name\":\"updateMinVoteCap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"destAddress\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"destAddress\",\"type\":\"address\"}],\"name\":\"withdrawWithIndex\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AsDBChainStakingFuncSigs maps the 4-byte function signature to its string representation.
var AsDBChainStakingFuncSigs = map[string]string{
	"f851a440": "admin()",
	"3477ee2e": "candidates(uint256)",
	"b5b7a184": "epochPeriod()",
	"2a466ac7": "getCandidateData(address)",
	"b642facd": "getCandidateOwner(address)",
	"484da961": "getCandidateStake(address)",
	"b97dd9e2": "getCurrentEpoch()",
	"690ff8a1": "getListCandidates()",
	"158a65f6": "getVoterStake(address,address)",
	"e2db89b5": "getVoterStakes(address,address[])",
	"2d15cc04": "getVoters(address)",
	"15febd68": "getWithdrawCap(uint256)",
	"0e0516aa": "getWithdrawEpochs()",
	"d5816bfa": "getWithdrawEpochsAndCaps()",
	"d51b9e93": "isCandidate(address)",
	"2de7dd5f": "maxValidatorSize()",
	"017ddd35": "minValidatorStake()",
	"f8ac9dd5": "minVoterCap()",
	"aa677354": "register(address,address)",
	"ae6e43f5": "resign(address)",
	"48cd4cb1": "startBlock()",
	"75829def": "transferAdmin(address)",
	"02aa9be2": "unvote(address,uint256)",
	"0619624f": "updateMaxValidatorSize(uint256)",
	"b2c76f10": "updateMinValidateStake(uint256)",
	"3a1d8c5a": "updateMinVoteCap(uint256)",
	"6dd7d8ea": "vote(address)",
	"00f714ce": "withdraw(uint256,address)",
	"c2df409b": "withdrawWithIndex(uint256,uint256,address)",
}

// AsDBChainStakingBin is the compiled bytecode used for deploying new contracts.
var AsDBChainStakingBin = "0x60806040523480156200001157600080fd5b50604051620025743803806200257483398181016040526101008110156200003857600080fd5b81019080805160405193929190846401000000008211156200005957600080fd5b9083019060208201858111156200006f57600080fd5b82518660208202830111640100000000821117156200008d57600080fd5b82525081516020918201928201910280838360005b83811015620000bc578181015183820152602001620000a2565b5050505090500160405260200180516040519392919084640100000000821115620000e657600080fd5b908301906020820185811115620000fc57600080fd5b82518660208202830111640100000000821117156200011a57600080fd5b82525081516020918201928201910280838360005b83811015620001495781810151838201526020016200012f565b505050509190910160409081526020830151908301516060840151608085015160a086015160c0909601516001600055939750919550939092509085620001d7576040805162461bcd60e51b815260206004820152601660248201527f65706f6368206d75737420626520706f73697469766500000000000000000000604482015290519081900360640190fd5b865188511462000221576040805162461bcd60e51b815260206004820152601060248201526f0d8cadccee8d040dcdee840dac2e8c6d60831b604482015290519081900360640190fd5b875184101562000278576040805162461bcd60e51b815260206004820152601960248201527f696e76616c6964205f6d617856616c696461746f7253697a6500000000000000604482015290519081900360640190fd5b60068690556007849055600883905560098290558751620002a19060049060208b01906200042a565b5060005b8851811015620003f957600060048281548110620002bf57fe5b60009182526020808320909101546001600160a01b031680835260039091526040909120805460ff191660011781558a51919250908a90849081106200030157fe5b60200260200101518160020160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550858160010181905550858160030160008c86815181106200034e57fe5b60200260200101516001600160a01b03166001600160a01b031681526020019081526020016000208190555060026000836001600160a01b03166001600160a01b031681526020019081526020016000208a8481518110620003ac57fe5b60209081029190910181015182546001808201855560009485529290932090920180546001600160a01b0319166001600160a01b0390931692909217909155929092019150620002a59050565b50600a80546001600160a01b0319166001600160a01b039290921691909117905550505060055550620004ab915050565b82805482825590600052602060002090810192821562000482579160200282015b828111156200048257825182546001600160a01b0319166001600160a01b039091161782556020909201916001909101906200044b565b506200049092915062000494565b5090565b5b8082111562000490576000815560010162000495565b6120b980620004bb6000396000f3fe6080604052600436106101c15760003560e01c8063690ff8a1116100f7578063b642facd11610095578063d5816bfa11610064578063d5816bfa14610747578063e2db89b5146107f5578063f851a440146108b5578063f8ac9dd5146108ca576101c1565b8063b642facd1461068d578063b97dd9e2146106c0578063c2df409b146106d5578063d51b9e9314610714576101c1565b8063aa677354116100d1578063aa677354146105e0578063ae6e43f51461061b578063b2c76f101461064e578063b5b7a18414610678576101c1565b8063690ff8a1146104c45780636dd7d8ea1461058757806375829def146105ad576101c1565b80632a466ac7116101645780633477ee2e1161013e5780633477ee2e1461040c5780633a1d8c5a14610452578063484da9611461047c57806348cd4cb1146104af576101c1565b80632a466ac7146103695780632d15cc04146103c45780632de7dd5f146103f7576101c1565b80630619624f116101a05780630619624f146102755780630e0516aa1461029f578063158a65f61461030457806315febd681461033f576101c1565b8062f714ce146101c6578063017ddd351461021357806302aa9be21461023a575b600080fd5b3480156101d257600080fd5b506101ff600480360360408110156101e957600080fd5b50803590602001356001600160a01b03166108df565b604080519115158252519081900360200190f35b34801561021f57600080fd5b50610228610a9b565b60408051918252519081900360200190f35b34801561024657600080fd5b506102736004803603604081101561025d57600080fd5b506001600160a01b038135169060200135610aa1565b005b34801561028157600080fd5b506102736004803603602081101561029857600080fd5b5035610d93565b3480156102ab57600080fd5b506102b4610de4565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156102f05781810151838201526020016102d8565b505050509050019250505060405180910390f35b34801561031057600080fd5b506102286004803603604081101561032757600080fd5b506001600160a01b0381358116916020013516610e48565b34801561034b57600080fd5b506102286004803603602081101561036257600080fd5b5035610e74565b34801561037557600080fd5b5061039c6004803603602081101561038c57600080fd5b50356001600160a01b0316610e91565b6040805193151584526001600160a01b03909216602084015282820152519081900360600190f35b3480156103d057600080fd5b506102b4600480360360208110156103e757600080fd5b50356001600160a01b0316610ec5565b34801561040357600080fd5b50610228610f3b565b34801561041857600080fd5b506104366004803603602081101561042f57600080fd5b5035610f41565b604080516001600160a01b039092168252519081900360200190f35b34801561045e57600080fd5b506102736004803603602081101561047557600080fd5b5035610f6b565b34801561048857600080fd5b506102286004803603602081101561049f57600080fd5b50356001600160a01b0316610fbc565b3480156104bb57600080fd5b50610228610fda565b3480156104d057600080fd5b506104d9610fe0565b604051808060200180602001868152602001858152602001848152602001838103835288818151815260200191508051906020019060200280838360005b8381101561052f578181015183820152602001610517565b50505050905001838103825287818151815260200191508051906020019060200280838360005b8381101561056e578181015183820152602001610556565b5050505090500197505050505050505060405180910390f35b6102736004803603602081101561059d57600080fd5b50356001600160a01b031661110d565b3480156105b957600080fd5b50610273600480360360208110156105d057600080fd5b50356001600160a01b031661130a565b3480156105ec57600080fd5b506102736004803603604081101561060357600080fd5b506001600160a01b03813581169160200135166113c0565b34801561062757600080fd5b506102736004803603602081101561063e57600080fd5b50356001600160a01b0316611653565b34801561065a57600080fd5b506102736004803603602081101561067157600080fd5b503561195a565b34801561068457600080fd5b506102286119ab565b34801561069957600080fd5b50610436600480360360208110156106b057600080fd5b50356001600160a01b03166119b1565b3480156106cc57600080fd5b506102286119d2565b3480156106e157600080fd5b506101ff600480360360608110156106f857600080fd5b50803590602081013590604001356001600160a01b03166119fa565b34801561072057600080fd5b506101ff6004803603602081101561073757600080fd5b50356001600160a01b0316611d20565b34801561075357600080fd5b5061075c611d3e565b604051808060200180602001838103835285818151815260200191508051906020019060200280838360005b838110156107a0578181015183820152602001610788565b50505050905001838103825284818151815260200191508051906020019060200280838360005b838110156107df5781810151838201526020016107c7565b5050505090500194505050505060405180910390f35b34801561080157600080fd5b506102b46004803603604081101561081857600080fd5b6001600160a01b03823516919081019060408101602082013564010000000081111561084357600080fd5b82018360208201111561085557600080fd5b8035906020019184602083028401116401000000008311171561087757600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550611e50945050505050565b3480156108c157600080fd5b50610436611f28565b3480156108d657600080fd5b50610228611f37565b600060026000541415610939576040805162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b600260009081556109486119d2565b9050838110156109895760405162461bcd60e51b81526004018080602001828103825260218152602001806120636021913960400191505060405180910390fd5b336000818152600160209081526040808320888452909152902054806109ea576040805162461bcd60e51b81526020600482015260116024820152700776974686472617720636170206973203607c1b604482015290519081900360640190fd5b6001600160a01b0380831660009081526001602090815260408083208a8452909152808220829055519187169183156108fc0291849190818181858888f19350505050158015610a3e573d6000803e3d6000fd5b50604080516001600160a01b0380851682526020820184905287168183015290517f56c54ba9bd38d8fd62012e42c7ee564519b09763c426d331b3661b537ead19b29181900360600190a160019350505050600160005592915050565b60085481565b60026000541415610af9576040805162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b600260005580610b50576040805162461bcd60e51b815260206004820152601960248201527f616d6f756e742073686f756c6420626520706f73697469766500000000000000604482015290519081900360640190fd5b6000610b5a6119d2565b6001600160a01b038416600090815260036020818152604080842033808652930190915282205492935091610b8f9085611f3d565b6001600160a01b0386811660009081526003602052604090206002015491925083811691161415610c1657600854811015610c11576040805162461bcd60e51b815260206004820152601e60248201527f6e6577207374616b6573203c206d696e56616c696461746f725374616b650000604482015290519081900360640190fd5b610c6b565b801580610c2557506009548110155b610c6b576040805162461bcd60e51b81526020600482015260126024820152711a5b9d985b1a59081d5b9d9bdd1948185b5d60721b604482015290519081900360640190fd5b6001600160a01b03808616600081815260036020818152604080842095881684528583018252832086905592909152905260010154610caa9085611f3d565b6001600160a01b038616600090815260036020526040812060010191909155610cd4846002611f9a565b6001600160a01b0384166000908152600160209081526040808320848452909152902054909150610d059086611f9a565b6001600160a01b03808516600081815260016020818152604080842088855280835281852097909755828252958201805492830181558352918290200185905583519182529189169181019190915280820187905290517f7958395da8e26969cc7c09ee58e9507a2601574c3bd232617e2d6354224ff8369181900360600190a15050600160005550505050565b600a546001600160a01b03163314610ddf576040805162461bcd60e51b815260206004820152600a60248201526941444d494e204f4e4c5960b01b604482015290519081900360640190fd5b600755565b33600090815260016020818152604092839020909101805483518184028101840190945280845260609392830182828015610e3e57602002820191906000526020600020905b815481526020019060010190808311610e2a575b5050505050905090565b6001600160a01b0391821660009081526003602081815260408084209490951683529201909152205490565b336000908152600160209081526040808320938352929052205490565b6001600160a01b0390811660009081526003602052604090208054600282015460019092015460ff90911693919092169190565b6001600160a01b038116600090815260026020908152604091829020805483518184028101840190945280845260609392830182828015610f2f57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610f11575b50505050509050919050565b60075481565b60048181548110610f5157600080fd5b6000918252602090912001546001600160a01b0316905081565b600a546001600160a01b03163314610fb7576040805162461bcd60e51b815260206004820152600a60248201526941444d494e204f4e4c5960b01b604482015290519081900360640190fd5b600955565b6001600160a01b031660009081526003602052604090206001015490565b60055481565b6060806000806000610ff06119d2565b925060075491506008549050600480548060200260200160405190810160405280929190818152602001828054801561105257602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611034575b50505050509450845167ffffffffffffffff8111801561107157600080fd5b5060405190808252806020026020018201604052801561109b578160200160208202803683370190505b50935060005b855181101561110557600360008783815181106110ba57fe5b60200260200101516001600160a01b03166001600160a01b03168152602001908152602001600020600101548582815181106110f257fe5b60209081029190910101526001016110a1565b509091929394565b600954341015611156576040805162461bcd60e51b815260206004820152600f60248201526e1b1bddc81d9bdd1948185b5bdd5b9d608a1b604482015290519081900360640190fd5b6001600160a01b038116600090815260036020526040902054819060ff1615156001146111c2576040805162461bcd60e51b81526020600482015260156024820152746f6e6c79206163746976652063616e64696461746560581b604482015290519081900360640190fd5b6001600160a01b038216600090815260036020818152604080842033808652930190915290912054349190611230576001600160a01b0384811660009081526002602090815260408220805460018101825590835291200180546001600160a01b0319169183169190911790555b6001600160a01b038085166000908152600360208181526040808420948616845293909101905220546112639083611f9a565b6001600160a01b03808616600081815260036020818152604080842095881684528583018252832095909555919052909152600101546112a39083611f9a565b6001600160a01b0380861660008181526003602090815260409182902060010194909455805192851683529282015280820184905290517f174ba19ba3c8bb5c679c87e51db79fff7c3f04bb84c1fd55b7dacb470b674aa69181900360600190a150505050565b600a546001600160a01b03163314611356576040805162461bcd60e51b815260206004820152600a60248201526941444d494e204f4e4c5960b01b604482015290519081900360640190fd5b6001600160a01b03811661139e576040805162461bcd60e51b815260206004820152600a602482015269041444d494e20697320360b41b604482015290519081900360640190fd5b600a80546001600160a01b0319166001600160a01b0392909216919091179055565b600a546001600160a01b0316331461140c576040805162461bcd60e51b815260206004820152600a60248201526941444d494e204f4e4c5960b01b604482015290519081900360640190fd5b6001600160a01b038216600090815260036020526040902054829060ff161561147c576040805162461bcd60e51b815260206004820152601960248201527f6f6e6c79206e6f74206163746976652063616e64696461746500000000000000604482015290519081900360640190fd5b6001600160a01b0383166114d7576040805162461bcd60e51b815260206004820152601760248201527f5f63616e64696461746520616464726573732069732030000000000000000000604482015290519081900360640190fd5b6001600160a01b038216611528576040805162461bcd60e51b815260206004820152601360248201527205f6f776e65722061646472657373206973203606c1b604482015290519081900360640190fd5b600454608011611575576040805162461bcd60e51b8152602060048201526013602482015272746f6f206d616e792063616e6469646174657360681b604482015290519081900360640190fd5b6001600160a01b03808416600081815260036020908152604080832060018082015460028084018054998c166001600160a01b03199a8b168117909155845460ff191684178555600480548086019091557f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b0180548b168a1790558888529086528487208054938401815587529585902090910180549097168517909655815194855291840192909252815190927f0a31ee9d46a828884b81003c8498156ea6aa15b9b54bdd0ef0b533d9eba57e5592908290030190a15050505050565b6001600160a01b038116600090815260036020526040902054819060ff1615156001146116bf576040805162461bcd60e51b81526020600482015260156024820152746f6e6c79206163746976652063616e64696461746560581b604482015290519081900360640190fd5b6001600160a01b03808316600090815260036020526040902060020154839116331461171e576040805162461bcd60e51b81526020600482015260096024820152683737ba1037bbb732b960b91b604482015290519081900360640190fd5b3360006117296119d2565b905060005b60045481101561183057856001600160a01b03166004828154811061174f57fe5b6000918252602090912001546001600160a01b031614156118285760048054600019810190811061177c57fe5b600091825260209091200154600480546001600160a01b0390921691839081106117a257fe5b600091825260209091200180546001600160a01b0319166001600160a01b03929092169190911790556004805460001981019081106117dd57fe5b600091825260209091200180546001600160a01b0319169055600480548061180157fe5b600082815260209020810160001990810180546001600160a01b0319169055019055611830565b60010161172e565b506001600160a01b038086166000818152600360208181526040808420805460ff1916815595881684528583018252832080549084905593909252905260019091015461187d9082611f3d565b6001600160a01b0387166000908152600360205260408120600101919091556118a7836002611f9a565b6001600160a01b03851660009081526001602090815260408083208484529091529020549091506118d89083611f9a565b6001600160a01b0380861660009081526001602081815260408084208785528083528185209690965582825294820180549283018155835291829020018490558251918a168252810185905281517f886e0db046874dde595498040d176412e81183750ceb33fc46f0450362bbc241929181900390910190a150505050505050565b600a546001600160a01b031633146119a6576040805162461bcd60e51b815260206004820152600a60248201526941444d494e204f4e4c5960b01b604482015290519081900360640190fd5b600855565b60065481565b6001600160a01b039081166000908152600360205260409020600201541690565b60006119f56006546119ef60055443611f3d90919063ffffffff16565b90611ffb565b905090565b600060026000541415611a54576040805162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b60026000908155611a636119d2565b905084811015611aa45760405162461bcd60e51b81526004018080602001828103825260218152602001806120636021913960400191505060405180910390fd5b33600081815260016020908152604080832089845290915290205480611b05576040805162461bcd60e51b81526020600482015260116024820152700776974686472617720636170206973203607c1b604482015290519081900360640190fd5b8660016000846001600160a01b03166001600160a01b031681526020019081526020016000206001018781548110611b3957fe5b906000526020600020015414611b8a576040805162461bcd60e51b81526020600482015260116024820152700dcdee840c6dee4e4cac6e840d2dcc8caf607b1b604482015290519081900360640190fd5b6001600160a01b03821660008181526001602081815260408084208c85528083529084208490559390925290819052018054906000198201828110611bcb57fe5b906000526020600020015460016000856001600160a01b03166001600160a01b031681526020019081526020016000206001018881548110611c0957fe5b60009182526020808320909101929092556001600160a01b0385168152600191829052604090200180546000198301908110611c4157fe5b600091825260208083209091018290556001600160a01b0385168252600190819052604090912001805480611c7257fe5b60019003818190600052602060002001600090559055856001600160a01b03166108fc839081150290604051600060405180830381858888f19350505050158015611cc1573d6000803e3d6000fd5b50604080516001600160a01b0380861682526020820185905288168183015290517f56c54ba9bd38d8fd62012e42c7ee564519b09763c426d331b3661b537ead19b29181900360600190a1600194505050505060016000559392505050565b6001600160a01b031660009081526003602052604090205460ff1690565b336000908152600160208181526040928390209091018054835181840281018401909452808452606093849390929190830182828015611d9d57602002820191906000526020600020905b815481526020019060010190808311611d89575b50505050509150815167ffffffffffffffff81118015611dbc57600080fd5b50604051908082528060200260200182016040528015611de6578160200160208202803683370190505b50905060005b8251811015611e4b573360009081526001602052604081208451909190859084908110611e1557fe5b6020026020010151815260200190815260200160002054828281518110611e3857fe5b6020908102919091010152600101611dec565b509091565b6060815167ffffffffffffffff81118015611e6a57600080fd5b50604051908082528060200260200182016040528015611e94578160200160208202803683370190505b50905060005b8251811015611f215760036000856001600160a01b03166001600160a01b031681526020019081526020016000206003016000848381518110611ed957fe5b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002054828281518110611f0e57fe5b6020908102919091010152600101611e9a565b5092915050565b600a546001600160a01b031681565b60095481565b600082821115611f94576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b600082820183811015611ff4576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b6000808211612051576040805162461bcd60e51b815260206004820152601a60248201527f536166654d6174683a206469766973696f6e206279207a65726f000000000000604482015290519081900360640190fd5b81838161205a57fe5b04939250505056fe63616e206e6f7420776974686472617720666f72206675747572652065706f6368a2646970667358221220af1de754dd70a867efe8d67a5ca37fbf3b47b2d8258ac6ee163343572b13fc3364736f6c63430007060033"

// DeployAsDBChainStaking deploys a new Ethereum contract, binding an instance of AsDBChainStaking to it.
func DeployAsDBChainStaking(auth *bind.TransactOpts, backend bind.ContractBackend, _candidates []common.Address, _candidateOwners []common.Address, _epochPeriod *big.Int, _startBlock *big.Int, _maxValidatorSize *big.Int, _minValidatorStake *big.Int, _minVoteCap *big.Int, _admin common.Address) (common.Address, *types.Transaction, *AsDBChainStaking, error) {
	parsed, err := abi.JSON(strings.NewReader(AsDBChainStakingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AsDBChainStakingBin), backend, _candidates, _candidateOwners, _epochPeriod, _startBlock, _maxValidatorSize, _minValidatorStake, _minVoteCap, _admin)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AsDBChainStaking{AsDBChainStakingCaller: AsDBChainStakingCaller{contract: contract}, AsDBChainStakingTransactor: AsDBChainStakingTransactor{contract: contract}, AsDBChainStakingFilterer: AsDBChainStakingFilterer{contract: contract}}, nil
}

// AsDBChainStaking is an auto generated Go binding around an Ethereum contract.
type AsDBChainStaking struct {
	AsDBChainStakingCaller     // Read-only binding to the contract
	AsDBChainStakingTransactor // Write-only binding to the contract
	AsDBChainStakingFilterer   // Log filterer for contract events
}

// AsDBChainStakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type AsDBChainStakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AsDBChainStakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AsDBChainStakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AsDBChainStakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AsDBChainStakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AsDBChainStakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AsDBChainStakingSession struct {
	Contract     *AsDBChainStaking // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AsDBChainStakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AsDBChainStakingCallerSession struct {
	Contract *AsDBChainStakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// AsDBChainStakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AsDBChainStakingTransactorSession struct {
	Contract     *AsDBChainStakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// AsDBChainStakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type AsDBChainStakingRaw struct {
	Contract *AsDBChainStaking // Generic contract binding to access the raw methods on
}

// AsDBChainStakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AsDBChainStakingCallerRaw struct {
	Contract *AsDBChainStakingCaller // Generic read-only contract binding to access the raw methods on
}

// AsDBChainStakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AsDBChainStakingTransactorRaw struct {
	Contract *AsDBChainStakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAsDBChainStaking creates a new instance of AsDBChainStaking, bound to a specific deployed contract.
func NewAsDBChainStaking(address common.Address, backend bind.ContractBackend) (*AsDBChainStaking, error) {
	contract, err := bindAsDBChainStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AsDBChainStaking{AsDBChainStakingCaller: AsDBChainStakingCaller{contract: contract}, AsDBChainStakingTransactor: AsDBChainStakingTransactor{contract: contract}, AsDBChainStakingFilterer: AsDBChainStakingFilterer{contract: contract}}, nil
}

// NewAsDBChainStakingCaller creates a new read-only instance of AsDBChainStaking, bound to a specific deployed contract.
func NewAsDBChainStakingCaller(address common.Address, caller bind.ContractCaller) (*AsDBChainStakingCaller, error) {
	contract, err := bindAsDBChainStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AsDBChainStakingCaller{contract: contract}, nil
}

// NewAsDBChainStakingTransactor creates a new write-only instance of AsDBChainStaking, bound to a specific deployed contract.
func NewAsDBChainStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*AsDBChainStakingTransactor, error) {
	contract, err := bindAsDBChainStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AsDBChainStakingTransactor{contract: contract}, nil
}

// NewAsDBChainStakingFilterer creates a new log filterer instance of AsDBChainStaking, bound to a specific deployed contract.
func NewAsDBChainStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*AsDBChainStakingFilterer, error) {
	contract, err := bindAsDBChainStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AsDBChainStakingFilterer{contract: contract}, nil
}

// bindAsDBChainStaking binds a generic wrapper to an already deployed contract.
func bindAsDBChainStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AsDBChainStakingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AsDBChainStaking *AsDBChainStakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AsDBChainStaking.Contract.AsDBChainStakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AsDBChainStaking *AsDBChainStakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AsDBChainStaking.Contract.AsDBChainStakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AsDBChainStaking *AsDBChainStakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AsDBChainStaking.Contract.AsDBChainStakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AsDBChainStaking *AsDBChainStakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AsDBChainStaking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AsDBChainStaking *AsDBChainStakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AsDBChainStaking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AsDBChainStaking *AsDBChainStakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AsDBChainStaking.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_AsDBChainStaking *AsDBChainStakingCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AsDBChainStaking.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_AsDBChainStaking *AsDBChainStakingSession) Admin() (common.Address, error) {
	return _AsDBChainStaking.Contract.Admin(&_AsDBChainStaking.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_AsDBChainStaking *AsDBChainStakingCallerSession) Admin() (common.Address, error) {
	return _AsDBChainStaking.Contract.Admin(&_AsDBChainStaking.CallOpts)
}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates(uint256 ) view returns(address)
func (_AsDBChainStaking *AsDBChainStakingCaller) Candidates(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AsDBChainStaking.contract.Call(opts, &out, "candidates", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates(uint256 ) view returns(address)
func (_AsDBChainStaking *AsDBChainStakingSession) Candidates(arg0 *big.Int) (common.Address, error) {
	return _AsDBChainStaking.Contract.Candidates(&_AsDBChainStaking.CallOpts, arg0)
}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates(uint256 ) view returns(address)
func (_AsDBChainStaking *AsDBChainStakingCallerSession) Candidates(arg0 *big.Int) (common.Address, error) {
	return _AsDBChainStaking.Contract.Candidates(&_AsDBChainStaking.CallOpts, arg0)
}

// EpochPeriod is a free data retrieval call binding the contract method 0xb5b7a184.
//
// Solidity: function epochPeriod() view returns(uint256)
func (_AsDBChainStaking *AsDBChainStakingCaller) EpochPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AsDBChainStaking.contract.Call(opts, &out, "epochPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochPeriod is a free data retrieval call binding the contract method 0xb5b7a184.
//
// Solidity: function epochPeriod() view returns(uint256)
func (_AsDBChainStaking *AsDBChainStakingSession) EpochPeriod() (*big.Int, error) {
	return _AsDBChainStaking.Contract.EpochPeriod(&_AsDBChainStaking.CallOpts)
}

// EpochPeriod is a free data retrieval call binding the contract method 0xb5b7a184.
//
// Solidity: function epochPeriod() view returns(uint256)
func (_AsDBChainStaking *AsDBChainStakingCallerSession) EpochPeriod() (*big.Int, error) {
	return _AsDBChainStaking.Contract.EpochPeriod(&_AsDBChainStaking.CallOpts)
}

// GetCandidateData is a free data retrieval call binding the contract method 0x2a466ac7.
//
// Solidity: function getCandidateData(address _candidate) view returns(bool _isActiveCandidate, address _owner, uint256 _totalStake)
func (_AsDBChainStaking *AsDBChainStakingCaller) GetCandidateData(opts *bind.CallOpts, _candidate common.Address) (struct {
	IsActiveCandidate bool
	Owner             common.Address
	TotalStake        *big.Int
}, error) {
	var out []interface{}
	err := _AsDBChainStaking.contract.Call(opts, &out, "getCandidateData", _candidate)

	outstruct := new(struct {
		IsActiveCandidate bool
		Owner             common.Address
		TotalStake        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsActiveCandidate = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Owner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.TotalStake = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetCandidateData is a free data retrieval call binding the contract method 0x2a466ac7.
//
// Solidity: function getCandidateData(address _candidate) view returns(bool _isActiveCandidate, address _owner, uint256 _totalStake)
func (_AsDBChainStaking *AsDBChainStakingSession) GetCandidateData(_candidate common.Address) (struct {
	IsActiveCandidate bool
	Owner             common.Address
	TotalStake        *big.Int
}, error) {
	return _AsDBChainStaking.Contract.GetCandidateData(&_AsDBChainStaking.CallOpts, _candidate)
}

// GetCandidateData is a free data retrieval call binding the contract method 0x2a466ac7.
//
// Solidity: function getCandidateData(address _candidate) view returns(bool _isActiveCandidate, address _owner, uint256 _totalStake)
func (_AsDBChainStaking *AsDBChainStakingCallerSession) GetCandidateData(_candidate common.Address) (struct {
	IsActiveCandidate bool
	Owner             common.Address
	TotalStake        *big.Int
}, error) {
	return _AsDBChainStaking.Contract.GetCandidateData(&_AsDBChainStaking.CallOpts, _candidate)
}

// GetCandidateOwner is a free data retrieval call binding the contract method 0xb642facd.
//
// Solidity: function getCandidateOwner(address _candidate) view returns(address)
func (_AsDBChainStaking *AsDBChainStakingCaller) GetCandidateOwner(opts *bind.CallOpts, _candidate common.Address) (common.Address, error) {
	var out []interface{}
	err := _AsDBChainStaking.contract.Call(opts, &out, "getCandidateOwner", _candidate)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCandidateOwner is a free data retrieval call binding the contract method 0xb642facd.
//
// Solidity: function getCandidateOwner(address _candidate) view returns(address)
func (_AsDBChainStaking *AsDBChainStakingSession) GetCandidateOwner(_candidate common.Address) (common.Address, error) {
	return _AsDBChainStaking.Contract.GetCandidateOwner(&_AsDBChainStaking.CallOpts, _candidate)
}

// GetCandidateOwner is a free data retrieval call binding the contract method 0xb642facd.
//
// Solidity: function getCandidateOwner(address _candidate) view returns(address)
func (_AsDBChainStaking *AsDBChainStakingCallerSession) GetCandidateOwner(_candidate common.Address) (common.Address, error) {
	return _AsDBChainStaking.Contract.GetCandidateOwner(&_AsDBChainStaking.CallOpts, _candidate)
}

// GetCandidateStake is a free data retrieval call binding the contract method 0x484da961.
//
// Solidity: function getCandidateStake(address _candidate) view returns(uint256)
func (_AsDBChainStaking *AsDBChainStakingCaller) GetCandidateStake(opts *bind.CallOpts, _candidate common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AsDBChainStaking.contract.Call(opts, &out, "getCandidateStake", _candidate)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCandidateStake is a free data retrieval call binding the contract method 0x484da961.
//
// Solidity: function getCandidateStake(address _candidate) view returns(uint256)
func (_AsDBChainStaking *AsDBChainStakingSession) GetCandidateStake(_candidate common.Address) (*big.Int, error) {
	return _AsDBChainStaking.Contract.GetCandidateStake(&_AsDBChainStaking.CallOpts, _candidate)
}

// GetCandidateStake is a free data retrieval call binding the contract method 0x484da961.
//
// Solidity: function getCandidateStake(address _candidate) view returns(uint256)
func (_AsDBChainStaking *AsDBChainStakingCallerSession) GetCandidateStake(_candidate common.Address) (*big.Int, error) {
	return _AsDBChainStaking.Contract.GetCandidateStake(&_AsDBChainStaking.CallOpts, _candidate)
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint256)
func (_AsDBChainStaking *AsDBChainStakingCaller) GetCurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AsDBChainStaking.contract.Call(opts, &out, "getCurrentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint256)
func (_AsDBChainStaking *AsDBChainStakingSession) GetCurrentEpoch() (*big.Int, error) {
	return _AsDBChainStaking.Contract.GetCurrentEpoch(&_AsDBChainStaking.CallOpts)
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint256)
func (_AsDBChainStaking *AsDBChainStakingCallerSession) GetCurrentEpoch() (*big.Int, error) {
	return _AsDBChainStaking.Contract.GetCurrentEpoch(&_AsDBChainStaking.CallOpts)
}

// GetListCandidates is a free data retrieval call binding the contract method 0x690ff8a1.
//
// Solidity: function getListCandidates() view returns(address[] _candidates, uint256[] stakes, uint256 epoch, uint256 validatorSize, uint256 minValidatorCap)
func (_AsDBChainStaking *AsDBChainStakingCaller) GetListCandidates(opts *bind.CallOpts) (struct {
	Candidates      []common.Address
	Stakes          []*big.Int
	Epoch           *big.Int
	ValidatorSize   *big.Int
	MinValidatorCap *big.Int
}, error) {
	var out []interface{}
	err := _AsDBChainStaking.contract.Call(opts, &out, "getListCandidates")

	outstruct := new(struct {
		Candidates      []common.Address
		Stakes          []*big.Int
		Epoch           *big.Int
		ValidatorSize   *big.Int
		MinValidatorCap *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Candidates = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Stakes = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.Epoch = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ValidatorSize = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.MinValidatorCap = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetListCandidates is a free data retrieval call binding the contract method 0x690ff8a1.
//
// Solidity: function getListCandidates() view returns(address[] _candidates, uint256[] stakes, uint256 epoch, uint256 validatorSize, uint256 minValidatorCap)
func (_AsDBChainStaking *AsDBChainStakingSession) GetListCandidates() (struct {
	Candidates      []common.Address
	Stakes          []*big.Int
	Epoch           *big.Int
	ValidatorSize   *big.Int
	MinValidatorCap *big.Int
}, error) {
	return _AsDBChainStaking.Contract.GetListCandidates(&_AsDBChainStaking.CallOpts)
}

// GetListCandidates is a free data retrieval call binding the contract method 0x690ff8a1.
//
// Solidity: function getListCandidates() view returns(address[] _candidates, uint256[] stakes, uint256 epoch, uint256 validatorSize, uint256 minValidatorCap)
func (_AsDBChainStaking *AsDBChainStakingCallerSession) GetListCandidates() (struct {
	Candidates      []common.Address
	Stakes          []*big.Int
	Epoch           *big.Int
	ValidatorSize   *big.Int
	MinValidatorCap *big.Int
}, error) {
	return _AsDBChainStaking.Contract.GetListCandidates(&_AsDBChainStaking.CallOpts)
}

// GetVoterStake is a free data retrieval call binding the contract method 0x158a65f6.
//
// Solidity: function getVoterStake(address _candidate, address _voter) view returns(uint256 stake)
func (_AsDBChainStaking *AsDBChainStakingCaller) GetVoterStake(opts *bind.CallOpts, _candidate common.Address, _voter common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AsDBChainStaking.contract.Call(opts, &out, "getVoterStake", _candidate, _voter)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVoterStake is a free data retrieval call binding the contract method 0x158a65f6.
//
// Solidity: function getVoterStake(address _candidate, address _voter) view returns(uint256 stake)
func (_AsDBChainStaking *AsDBChainStakingSession) GetVoterStake(_candidate common.Address, _voter common.Address) (*big.Int, error) {
	return _AsDBChainStaking.Contract.GetVoterStake(&_AsDBChainStaking.CallOpts, _candidate, _voter)
}

// GetVoterStake is a free data retrieval call binding the contract method 0x158a65f6.
//
// Solidity: function getVoterStake(address _candidate, address _voter) view returns(uint256 stake)
func (_AsDBChainStaking *AsDBChainStakingCallerSession) GetVoterStake(_candidate common.Address, _voter common.Address) (*big.Int, error) {
	return _AsDBChainStaking.Contract.GetVoterStake(&_AsDBChainStaking.CallOpts, _candidate, _voter)
}

// GetVoterStakes is a free data retrieval call binding the contract method 0xe2db89b5.
//
// Solidity: function getVoterStakes(address _candidate, address[] voters) view returns(uint256[] stakes)
func (_AsDBChainStaking *AsDBChainStakingCaller) GetVoterStakes(opts *bind.CallOpts, _candidate common.Address, voters []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _AsDBChainStaking.contract.Call(opts, &out, "getVoterStakes", _candidate, voters)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetVoterStakes is a free data retrieval call binding the contract method 0xe2db89b5.
//
// Solidity: function getVoterStakes(address _candidate, address[] voters) view returns(uint256[] stakes)
func (_AsDBChainStaking *AsDBChainStakingSession) GetVoterStakes(_candidate common.Address, voters []common.Address) ([]*big.Int, error) {
	return _AsDBChainStaking.Contract.GetVoterStakes(&_AsDBChainStaking.CallOpts, _candidate, voters)
}

// GetVoterStakes is a free data retrieval call binding the contract method 0xe2db89b5.
//
// Solidity: function getVoterStakes(address _candidate, address[] voters) view returns(uint256[] stakes)
func (_AsDBChainStaking *AsDBChainStakingCallerSession) GetVoterStakes(_candidate common.Address, voters []common.Address) ([]*big.Int, error) {
	return _AsDBChainStaking.Contract.GetVoterStakes(&_AsDBChainStaking.CallOpts, _candidate, voters)
}

// GetVoters is a free data retrieval call binding the contract method 0x2d15cc04.
//
// Solidity: function getVoters(address _candidate) view returns(address[] voters)
func (_AsDBChainStaking *AsDBChainStakingCaller) GetVoters(opts *bind.CallOpts, _candidate common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _AsDBChainStaking.contract.Call(opts, &out, "getVoters", _candidate)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetVoters is a free data retrieval call binding the contract method 0x2d15cc04.
//
// Solidity: function getVoters(address _candidate) view returns(address[] voters)
func (_AsDBChainStaking *AsDBChainStakingSession) GetVoters(_candidate common.Address) ([]common.Address, error) {
	return _AsDBChainStaking.Contract.GetVoters(&_AsDBChainStaking.CallOpts, _candidate)
}

// GetVoters is a free data retrieval call binding the contract method 0x2d15cc04.
//
// Solidity: function getVoters(address _candidate) view returns(address[] voters)
func (_AsDBChainStaking *AsDBChainStakingCallerSession) GetVoters(_candidate common.Address) ([]common.Address, error) {
	return _AsDBChainStaking.Contract.GetVoters(&_AsDBChainStaking.CallOpts, _candidate)
}

// GetWithdrawCap is a free data retrieval call binding the contract method 0x15febd68.
//
// Solidity: function getWithdrawCap(uint256 epoch) view returns(uint256 cap)
func (_AsDBChainStaking *AsDBChainStakingCaller) GetWithdrawCap(opts *bind.CallOpts, epoch *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AsDBChainStaking.contract.Call(opts, &out, "getWithdrawCap", epoch)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWithdrawCap is a free data retrieval call binding the contract method 0x15febd68.
//
// Solidity: function getWithdrawCap(uint256 epoch) view returns(uint256 cap)
func (_AsDBChainStaking *AsDBChainStakingSession) GetWithdrawCap(epoch *big.Int) (*big.Int, error) {
	return _AsDBChainStaking.Contract.GetWithdrawCap(&_AsDBChainStaking.CallOpts, epoch)
}

// GetWithdrawCap is a free data retrieval call binding the contract method 0x15febd68.
//
// Solidity: function getWithdrawCap(uint256 epoch) view returns(uint256 cap)
func (_AsDBChainStaking *AsDBChainStakingCallerSession) GetWithdrawCap(epoch *big.Int) (*big.Int, error) {
	return _AsDBChainStaking.Contract.GetWithdrawCap(&_AsDBChainStaking.CallOpts, epoch)
}

// GetWithdrawEpochs is a free data retrieval call binding the contract method 0x0e0516aa.
//
// Solidity: function getWithdrawEpochs() view returns(uint256[] epochs)
func (_AsDBChainStaking *AsDBChainStakingCaller) GetWithdrawEpochs(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _AsDBChainStaking.contract.Call(opts, &out, "getWithdrawEpochs")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetWithdrawEpochs is a free data retrieval call binding the contract method 0x0e0516aa.
//
// Solidity: function getWithdrawEpochs() view returns(uint256[] epochs)
func (_AsDBChainStaking *AsDBChainStakingSession) GetWithdrawEpochs() ([]*big.Int, error) {
	return _AsDBChainStaking.Contract.GetWithdrawEpochs(&_AsDBChainStaking.CallOpts)
}

// GetWithdrawEpochs is a free data retrieval call binding the contract method 0x0e0516aa.
//
// Solidity: function getWithdrawEpochs() view returns(uint256[] epochs)
func (_AsDBChainStaking *AsDBChainStakingCallerSession) GetWithdrawEpochs() ([]*big.Int, error) {
	return _AsDBChainStaking.Contract.GetWithdrawEpochs(&_AsDBChainStaking.CallOpts)
}

// GetWithdrawEpochsAndCaps is a free data retrieval call binding the contract method 0xd5816bfa.
//
// Solidity: function getWithdrawEpochsAndCaps() view returns(uint256[] epochs, uint256[] caps)
func (_AsDBChainStaking *AsDBChainStakingCaller) GetWithdrawEpochsAndCaps(opts *bind.CallOpts) (struct {
	Epochs []*big.Int
	Caps   []*big.Int
}, error) {
	var out []interface{}
	err := _AsDBChainStaking.contract.Call(opts, &out, "getWithdrawEpochsAndCaps")

	outstruct := new(struct {
		Epochs []*big.Int
		Caps   []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Epochs = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.Caps = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetWithdrawEpochsAndCaps is a free data retrieval call binding the contract method 0xd5816bfa.
//
// Solidity: function getWithdrawEpochsAndCaps() view returns(uint256[] epochs, uint256[] caps)
func (_AsDBChainStaking *AsDBChainStakingSession) GetWithdrawEpochsAndCaps() (struct {
	Epochs []*big.Int
	Caps   []*big.Int
}, error) {
	return _AsDBChainStaking.Contract.GetWithdrawEpochsAndCaps(&_AsDBChainStaking.CallOpts)
}

// GetWithdrawEpochsAndCaps is a free data retrieval call binding the contract method 0xd5816bfa.
//
// Solidity: function getWithdrawEpochsAndCaps() view returns(uint256[] epochs, uint256[] caps)
func (_AsDBChainStaking *AsDBChainStakingCallerSession) GetWithdrawEpochsAndCaps() (struct {
	Epochs []*big.Int
	Caps   []*big.Int
}, error) {
	return _AsDBChainStaking.Contract.GetWithdrawEpochsAndCaps(&_AsDBChainStaking.CallOpts)
}

// IsCandidate is a free data retrieval call binding the contract method 0xd51b9e93.
//
// Solidity: function isCandidate(address _candidate) view returns(bool)
func (_AsDBChainStaking *AsDBChainStakingCaller) IsCandidate(opts *bind.CallOpts, _candidate common.Address) (bool, error) {
	var out []interface{}
	err := _AsDBChainStaking.contract.Call(opts, &out, "isCandidate", _candidate)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCandidate is a free data retrieval call binding the contract method 0xd51b9e93.
//
// Solidity: function isCandidate(address _candidate) view returns(bool)
func (_AsDBChainStaking *AsDBChainStakingSession) IsCandidate(_candidate common.Address) (bool, error) {
	return _AsDBChainStaking.Contract.IsCandidate(&_AsDBChainStaking.CallOpts, _candidate)
}

// IsCandidate is a free data retrieval call binding the contract method 0xd51b9e93.
//
// Solidity: function isCandidate(address _candidate) view returns(bool)
func (_AsDBChainStaking *AsDBChainStakingCallerSession) IsCandidate(_candidate common.Address) (bool, error) {
	return _AsDBChainStaking.Contract.IsCandidate(&_AsDBChainStaking.CallOpts, _candidate)
}

// MaxValidatorSize is a free data retrieval call binding the contract method 0x2de7dd5f.
//
// Solidity: function maxValidatorSize() view returns(uint256)
func (_AsDBChainStaking *AsDBChainStakingCaller) MaxValidatorSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AsDBChainStaking.contract.Call(opts, &out, "maxValidatorSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxValidatorSize is a free data retrieval call binding the contract method 0x2de7dd5f.
//
// Solidity: function maxValidatorSize() view returns(uint256)
func (_AsDBChainStaking *AsDBChainStakingSession) MaxValidatorSize() (*big.Int, error) {
	return _AsDBChainStaking.Contract.MaxValidatorSize(&_AsDBChainStaking.CallOpts)
}

// MaxValidatorSize is a free data retrieval call binding the contract method 0x2de7dd5f.
//
// Solidity: function maxValidatorSize() view returns(uint256)
func (_AsDBChainStaking *AsDBChainStakingCallerSession) MaxValidatorSize() (*big.Int, error) {
	return _AsDBChainStaking.Contract.MaxValidatorSize(&_AsDBChainStaking.CallOpts)
}

// MinValidatorStake is a free data retrieval call binding the contract method 0x017ddd35.
//
// Solidity: function minValidatorStake() view returns(uint256)
func (_AsDBChainStaking *AsDBChainStakingCaller) MinValidatorStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AsDBChainStaking.contract.Call(opts, &out, "minValidatorStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinValidatorStake is a free data retrieval call binding the contract method 0x017ddd35.
//
// Solidity: function minValidatorStake() view returns(uint256)
func (_AsDBChainStaking *AsDBChainStakingSession) MinValidatorStake() (*big.Int, error) {
	return _AsDBChainStaking.Contract.MinValidatorStake(&_AsDBChainStaking.CallOpts)
}

// MinValidatorStake is a free data retrieval call binding the contract method 0x017ddd35.
//
// Solidity: function minValidatorStake() view returns(uint256)
func (_AsDBChainStaking *AsDBChainStakingCallerSession) MinValidatorStake() (*big.Int, error) {
	return _AsDBChainStaking.Contract.MinValidatorStake(&_AsDBChainStaking.CallOpts)
}

// MinVoterCap is a free data retrieval call binding the contract method 0xf8ac9dd5.
//
// Solidity: function minVoterCap() view returns(uint256)
func (_AsDBChainStaking *AsDBChainStakingCaller) MinVoterCap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AsDBChainStaking.contract.Call(opts, &out, "minVoterCap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinVoterCap is a free data retrieval call binding the contract method 0xf8ac9dd5.
//
// Solidity: function minVoterCap() view returns(uint256)
func (_AsDBChainStaking *AsDBChainStakingSession) MinVoterCap() (*big.Int, error) {
	return _AsDBChainStaking.Contract.MinVoterCap(&_AsDBChainStaking.CallOpts)
}

// MinVoterCap is a free data retrieval call binding the contract method 0xf8ac9dd5.
//
// Solidity: function minVoterCap() view returns(uint256)
func (_AsDBChainStaking *AsDBChainStakingCallerSession) MinVoterCap() (*big.Int, error) {
	return _AsDBChainStaking.Contract.MinVoterCap(&_AsDBChainStaking.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_AsDBChainStaking *AsDBChainStakingCaller) StartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AsDBChainStaking.contract.Call(opts, &out, "startBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_AsDBChainStaking *AsDBChainStakingSession) StartBlock() (*big.Int, error) {
	return _AsDBChainStaking.Contract.StartBlock(&_AsDBChainStaking.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_AsDBChainStaking *AsDBChainStakingCallerSession) StartBlock() (*big.Int, error) {
	return _AsDBChainStaking.Contract.StartBlock(&_AsDBChainStaking.CallOpts)
}

// Register is a paid mutator transaction binding the contract method 0xaa677354.
//
// Solidity: function register(address _candidate, address _owner) returns()
func (_AsDBChainStaking *AsDBChainStakingTransactor) Register(opts *bind.TransactOpts, _candidate common.Address, _owner common.Address) (*types.Transaction, error) {
	return _AsDBChainStaking.contract.Transact(opts, "register", _candidate, _owner)
}

// Register is a paid mutator transaction binding the contract method 0xaa677354.
//
// Solidity: function register(address _candidate, address _owner) returns()
func (_AsDBChainStaking *AsDBChainStakingSession) Register(_candidate common.Address, _owner common.Address) (*types.Transaction, error) {
	return _AsDBChainStaking.Contract.Register(&_AsDBChainStaking.TransactOpts, _candidate, _owner)
}

// Register is a paid mutator transaction binding the contract method 0xaa677354.
//
// Solidity: function register(address _candidate, address _owner) returns()
func (_AsDBChainStaking *AsDBChainStakingTransactorSession) Register(_candidate common.Address, _owner common.Address) (*types.Transaction, error) {
	return _AsDBChainStaking.Contract.Register(&_AsDBChainStaking.TransactOpts, _candidate, _owner)
}

// Resign is a paid mutator transaction binding the contract method 0xae6e43f5.
//
// Solidity: function resign(address _candidate) returns()
func (_AsDBChainStaking *AsDBChainStakingTransactor) Resign(opts *bind.TransactOpts, _candidate common.Address) (*types.Transaction, error) {
	return _AsDBChainStaking.contract.Transact(opts, "resign", _candidate)
}

// Resign is a paid mutator transaction binding the contract method 0xae6e43f5.
//
// Solidity: function resign(address _candidate) returns()
func (_AsDBChainStaking *AsDBChainStakingSession) Resign(_candidate common.Address) (*types.Transaction, error) {
	return _AsDBChainStaking.Contract.Resign(&_AsDBChainStaking.TransactOpts, _candidate)
}

// Resign is a paid mutator transaction binding the contract method 0xae6e43f5.
//
// Solidity: function resign(address _candidate) returns()
func (_AsDBChainStaking *AsDBChainStakingTransactorSession) Resign(_candidate common.Address) (*types.Transaction, error) {
	return _AsDBChainStaking.Contract.Resign(&_AsDBChainStaking.TransactOpts, _candidate)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_AsDBChainStaking *AsDBChainStakingTransactor) TransferAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _AsDBChainStaking.contract.Transact(opts, "transferAdmin", newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_AsDBChainStaking *AsDBChainStakingSession) TransferAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _AsDBChainStaking.Contract.TransferAdmin(&_AsDBChainStaking.TransactOpts, newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_AsDBChainStaking *AsDBChainStakingTransactorSession) TransferAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _AsDBChainStaking.Contract.TransferAdmin(&_AsDBChainStaking.TransactOpts, newAdmin)
}

// Unvote is a paid mutator transaction binding the contract method 0x02aa9be2.
//
// Solidity: function unvote(address candidate, uint256 amount) returns()
func (_AsDBChainStaking *AsDBChainStakingTransactor) Unvote(opts *bind.TransactOpts, candidate common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AsDBChainStaking.contract.Transact(opts, "unvote", candidate, amount)
}

// Unvote is a paid mutator transaction binding the contract method 0x02aa9be2.
//
// Solidity: function unvote(address candidate, uint256 amount) returns()
func (_AsDBChainStaking *AsDBChainStakingSession) Unvote(candidate common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AsDBChainStaking.Contract.Unvote(&_AsDBChainStaking.TransactOpts, candidate, amount)
}

// Unvote is a paid mutator transaction binding the contract method 0x02aa9be2.
//
// Solidity: function unvote(address candidate, uint256 amount) returns()
func (_AsDBChainStaking *AsDBChainStakingTransactorSession) Unvote(candidate common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AsDBChainStaking.Contract.Unvote(&_AsDBChainStaking.TransactOpts, candidate, amount)
}

// UpdateMaxValidatorSize is a paid mutator transaction binding the contract method 0x0619624f.
//
// Solidity: function updateMaxValidatorSize(uint256 newMaxValidatorSize) returns()
func (_AsDBChainStaking *AsDBChainStakingTransactor) UpdateMaxValidatorSize(opts *bind.TransactOpts, newMaxValidatorSize *big.Int) (*types.Transaction, error) {
	return _AsDBChainStaking.contract.Transact(opts, "updateMaxValidatorSize", newMaxValidatorSize)
}

// UpdateMaxValidatorSize is a paid mutator transaction binding the contract method 0x0619624f.
//
// Solidity: function updateMaxValidatorSize(uint256 newMaxValidatorSize) returns()
func (_AsDBChainStaking *AsDBChainStakingSession) UpdateMaxValidatorSize(newMaxValidatorSize *big.Int) (*types.Transaction, error) {
	return _AsDBChainStaking.Contract.UpdateMaxValidatorSize(&_AsDBChainStaking.TransactOpts, newMaxValidatorSize)
}

// UpdateMaxValidatorSize is a paid mutator transaction binding the contract method 0x0619624f.
//
// Solidity: function updateMaxValidatorSize(uint256 newMaxValidatorSize) returns()
func (_AsDBChainStaking *AsDBChainStakingTransactorSession) UpdateMaxValidatorSize(newMaxValidatorSize *big.Int) (*types.Transaction, error) {
	return _AsDBChainStaking.Contract.UpdateMaxValidatorSize(&_AsDBChainStaking.TransactOpts, newMaxValidatorSize)
}

// UpdateMinValidateStake is a paid mutator transaction binding the contract method 0xb2c76f10.
//
// Solidity: function updateMinValidateStake(uint256 _newCap) returns()
func (_AsDBChainStaking *AsDBChainStakingTransactor) UpdateMinValidateStake(opts *bind.TransactOpts, _newCap *big.Int) (*types.Transaction, error) {
	return _AsDBChainStaking.contract.Transact(opts, "updateMinValidateStake", _newCap)
}

// UpdateMinValidateStake is a paid mutator transaction binding the contract method 0xb2c76f10.
//
// Solidity: function updateMinValidateStake(uint256 _newCap) returns()
func (_AsDBChainStaking *AsDBChainStakingSession) UpdateMinValidateStake(_newCap *big.Int) (*types.Transaction, error) {
	return _AsDBChainStaking.Contract.UpdateMinValidateStake(&_AsDBChainStaking.TransactOpts, _newCap)
}

// UpdateMinValidateStake is a paid mutator transaction binding the contract method 0xb2c76f10.
//
// Solidity: function updateMinValidateStake(uint256 _newCap) returns()
func (_AsDBChainStaking *AsDBChainStakingTransactorSession) UpdateMinValidateStake(_newCap *big.Int) (*types.Transaction, error) {
	return _AsDBChainStaking.Contract.UpdateMinValidateStake(&_AsDBChainStaking.TransactOpts, _newCap)
}

// UpdateMinVoteCap is a paid mutator transaction binding the contract method 0x3a1d8c5a.
//
// Solidity: function updateMinVoteCap(uint256 _newCap) returns()
func (_AsDBChainStaking *AsDBChainStakingTransactor) UpdateMinVoteCap(opts *bind.TransactOpts, _newCap *big.Int) (*types.Transaction, error) {
	return _AsDBChainStaking.contract.Transact(opts, "updateMinVoteCap", _newCap)
}

// UpdateMinVoteCap is a paid mutator transaction binding the contract method 0x3a1d8c5a.
//
// Solidity: function updateMinVoteCap(uint256 _newCap) returns()
func (_AsDBChainStaking *AsDBChainStakingSession) UpdateMinVoteCap(_newCap *big.Int) (*types.Transaction, error) {
	return _AsDBChainStaking.Contract.UpdateMinVoteCap(&_AsDBChainStaking.TransactOpts, _newCap)
}

// UpdateMinVoteCap is a paid mutator transaction binding the contract method 0x3a1d8c5a.
//
// Solidity: function updateMinVoteCap(uint256 _newCap) returns()
func (_AsDBChainStaking *AsDBChainStakingTransactorSession) UpdateMinVoteCap(_newCap *big.Int) (*types.Transaction, error) {
	return _AsDBChainStaking.Contract.UpdateMinVoteCap(&_AsDBChainStaking.TransactOpts, _newCap)
}

// Vote is a paid mutator transaction binding the contract method 0x6dd7d8ea.
//
// Solidity: function vote(address candidate) payable returns()
func (_AsDBChainStaking *AsDBChainStakingTransactor) Vote(opts *bind.TransactOpts, candidate common.Address) (*types.Transaction, error) {
	return _AsDBChainStaking.contract.Transact(opts, "vote", candidate)
}

// Vote is a paid mutator transaction binding the contract method 0x6dd7d8ea.
//
// Solidity: function vote(address candidate) payable returns()
func (_AsDBChainStaking *AsDBChainStakingSession) Vote(candidate common.Address) (*types.Transaction, error) {
	return _AsDBChainStaking.Contract.Vote(&_AsDBChainStaking.TransactOpts, candidate)
}

// Vote is a paid mutator transaction binding the contract method 0x6dd7d8ea.
//
// Solidity: function vote(address candidate) payable returns()
func (_AsDBChainStaking *AsDBChainStakingTransactorSession) Vote(candidate common.Address) (*types.Transaction, error) {
	return _AsDBChainStaking.Contract.Vote(&_AsDBChainStaking.TransactOpts, candidate)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 epoch, address destAddress) returns(bool)
func (_AsDBChainStaking *AsDBChainStakingTransactor) Withdraw(opts *bind.TransactOpts, epoch *big.Int, destAddress common.Address) (*types.Transaction, error) {
	return _AsDBChainStaking.contract.Transact(opts, "withdraw", epoch, destAddress)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 epoch, address destAddress) returns(bool)
func (_AsDBChainStaking *AsDBChainStakingSession) Withdraw(epoch *big.Int, destAddress common.Address) (*types.Transaction, error) {
	return _AsDBChainStaking.Contract.Withdraw(&_AsDBChainStaking.TransactOpts, epoch, destAddress)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 epoch, address destAddress) returns(bool)
func (_AsDBChainStaking *AsDBChainStakingTransactorSession) Withdraw(epoch *big.Int, destAddress common.Address) (*types.Transaction, error) {
	return _AsDBChainStaking.Contract.Withdraw(&_AsDBChainStaking.TransactOpts, epoch, destAddress)
}

// WithdrawWithIndex is a paid mutator transaction binding the contract method 0xc2df409b.
//
// Solidity: function withdrawWithIndex(uint256 epoch, uint256 index, address destAddress) returns(bool)
func (_AsDBChainStaking *AsDBChainStakingTransactor) WithdrawWithIndex(opts *bind.TransactOpts, epoch *big.Int, index *big.Int, destAddress common.Address) (*types.Transaction, error) {
	return _AsDBChainStaking.contract.Transact(opts, "withdrawWithIndex", epoch, index, destAddress)
}

// WithdrawWithIndex is a paid mutator transaction binding the contract method 0xc2df409b.
//
// Solidity: function withdrawWithIndex(uint256 epoch, uint256 index, address destAddress) returns(bool)
func (_AsDBChainStaking *AsDBChainStakingSession) WithdrawWithIndex(epoch *big.Int, index *big.Int, destAddress common.Address) (*types.Transaction, error) {
	return _AsDBChainStaking.Contract.WithdrawWithIndex(&_AsDBChainStaking.TransactOpts, epoch, index, destAddress)
}

// WithdrawWithIndex is a paid mutator transaction binding the contract method 0xc2df409b.
//
// Solidity: function withdrawWithIndex(uint256 epoch, uint256 index, address destAddress) returns(bool)
func (_AsDBChainStaking *AsDBChainStakingTransactorSession) WithdrawWithIndex(epoch *big.Int, index *big.Int, destAddress common.Address) (*types.Transaction, error) {
	return _AsDBChainStaking.Contract.WithdrawWithIndex(&_AsDBChainStaking.TransactOpts, epoch, index, destAddress)
}

// AsDBChainStakingRegisteredIterator is returned from FilterRegistered and is used to iterate over the raw logs and unpacked data for Registered events raised by the AsDBChainStaking contract.
type AsDBChainStakingRegisteredIterator struct {
	Event *AsDBChainStakingRegistered // Event containing the contract specifics and raw log

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
func (it *AsDBChainStakingRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AsDBChainStakingRegistered)
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
		it.Event = new(AsDBChainStakingRegistered)
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
func (it *AsDBChainStakingRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AsDBChainStakingRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AsDBChainStakingRegistered represents a Registered event raised by the AsDBChainStaking contract.
type AsDBChainStakingRegistered struct {
	Candidate common.Address
	Owner     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRegistered is a free log retrieval operation binding the contract event 0x0a31ee9d46a828884b81003c8498156ea6aa15b9b54bdd0ef0b533d9eba57e55.
//
// Solidity: event Registered(address candidate, address owner)
func (_AsDBChainStaking *AsDBChainStakingFilterer) FilterRegistered(opts *bind.FilterOpts) (*AsDBChainStakingRegisteredIterator, error) {

	logs, sub, err := _AsDBChainStaking.contract.FilterLogs(opts, "Registered")
	if err != nil {
		return nil, err
	}
	return &AsDBChainStakingRegisteredIterator{contract: _AsDBChainStaking.contract, event: "Registered", logs: logs, sub: sub}, nil
}

// WatchRegistered is a free log subscription operation binding the contract event 0x0a31ee9d46a828884b81003c8498156ea6aa15b9b54bdd0ef0b533d9eba57e55.
//
// Solidity: event Registered(address candidate, address owner)
func (_AsDBChainStaking *AsDBChainStakingFilterer) WatchRegistered(opts *bind.WatchOpts, sink chan<- *AsDBChainStakingRegistered) (event.Subscription, error) {

	logs, sub, err := _AsDBChainStaking.contract.WatchLogs(opts, "Registered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AsDBChainStakingRegistered)
				if err := _AsDBChainStaking.contract.UnpackLog(event, "Registered", log); err != nil {
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

// ParseRegistered is a log parse operation binding the contract event 0x0a31ee9d46a828884b81003c8498156ea6aa15b9b54bdd0ef0b533d9eba57e55.
//
// Solidity: event Registered(address candidate, address owner)
func (_AsDBChainStaking *AsDBChainStakingFilterer) ParseRegistered(log types.Log) (*AsDBChainStakingRegistered, error) {
	event := new(AsDBChainStakingRegistered)
	if err := _AsDBChainStaking.contract.UnpackLog(event, "Registered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AsDBChainStakingResignedIterator is returned from FilterResigned and is used to iterate over the raw logs and unpacked data for Resigned events raised by the AsDBChainStaking contract.
type AsDBChainStakingResignedIterator struct {
	Event *AsDBChainStakingResigned // Event containing the contract specifics and raw log

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
func (it *AsDBChainStakingResignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AsDBChainStakingResigned)
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
		it.Event = new(AsDBChainStakingResigned)
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
func (it *AsDBChainStakingResignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AsDBChainStakingResignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AsDBChainStakingResigned represents a Resigned event raised by the AsDBChainStaking contract.
type AsDBChainStakingResigned struct {
	Candidate common.Address
	Epoch     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterResigned is a free log retrieval operation binding the contract event 0x886e0db046874dde595498040d176412e81183750ceb33fc46f0450362bbc241.
//
// Solidity: event Resigned(address _candidate, uint256 _epoch)
func (_AsDBChainStaking *AsDBChainStakingFilterer) FilterResigned(opts *bind.FilterOpts) (*AsDBChainStakingResignedIterator, error) {

	logs, sub, err := _AsDBChainStaking.contract.FilterLogs(opts, "Resigned")
	if err != nil {
		return nil, err
	}
	return &AsDBChainStakingResignedIterator{contract: _AsDBChainStaking.contract, event: "Resigned", logs: logs, sub: sub}, nil
}

// WatchResigned is a free log subscription operation binding the contract event 0x886e0db046874dde595498040d176412e81183750ceb33fc46f0450362bbc241.
//
// Solidity: event Resigned(address _candidate, uint256 _epoch)
func (_AsDBChainStaking *AsDBChainStakingFilterer) WatchResigned(opts *bind.WatchOpts, sink chan<- *AsDBChainStakingResigned) (event.Subscription, error) {

	logs, sub, err := _AsDBChainStaking.contract.WatchLogs(opts, "Resigned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AsDBChainStakingResigned)
				if err := _AsDBChainStaking.contract.UnpackLog(event, "Resigned", log); err != nil {
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

// ParseResigned is a log parse operation binding the contract event 0x886e0db046874dde595498040d176412e81183750ceb33fc46f0450362bbc241.
//
// Solidity: event Resigned(address _candidate, uint256 _epoch)
func (_AsDBChainStaking *AsDBChainStakingFilterer) ParseResigned(log types.Log) (*AsDBChainStakingResigned, error) {
	event := new(AsDBChainStakingResigned)
	if err := _AsDBChainStaking.contract.UnpackLog(event, "Resigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AsDBChainStakingUnvotedIterator is returned from FilterUnvoted and is used to iterate over the raw logs and unpacked data for Unvoted events raised by the AsDBChainStaking contract.
type AsDBChainStakingUnvotedIterator struct {
	Event *AsDBChainStakingUnvoted // Event containing the contract specifics and raw log

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
func (it *AsDBChainStakingUnvotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AsDBChainStakingUnvoted)
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
		it.Event = new(AsDBChainStakingUnvoted)
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
func (it *AsDBChainStakingUnvotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AsDBChainStakingUnvotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AsDBChainStakingUnvoted represents a Unvoted event raised by the AsDBChainStaking contract.
type AsDBChainStakingUnvoted struct {
	Voter     common.Address
	Candidate common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnvoted is a free log retrieval operation binding the contract event 0x7958395da8e26969cc7c09ee58e9507a2601574c3bd232617e2d6354224ff836.
//
// Solidity: event Unvoted(address voter, address candidate, uint256 amount)
func (_AsDBChainStaking *AsDBChainStakingFilterer) FilterUnvoted(opts *bind.FilterOpts) (*AsDBChainStakingUnvotedIterator, error) {

	logs, sub, err := _AsDBChainStaking.contract.FilterLogs(opts, "Unvoted")
	if err != nil {
		return nil, err
	}
	return &AsDBChainStakingUnvotedIterator{contract: _AsDBChainStaking.contract, event: "Unvoted", logs: logs, sub: sub}, nil
}

// WatchUnvoted is a free log subscription operation binding the contract event 0x7958395da8e26969cc7c09ee58e9507a2601574c3bd232617e2d6354224ff836.
//
// Solidity: event Unvoted(address voter, address candidate, uint256 amount)
func (_AsDBChainStaking *AsDBChainStakingFilterer) WatchUnvoted(opts *bind.WatchOpts, sink chan<- *AsDBChainStakingUnvoted) (event.Subscription, error) {

	logs, sub, err := _AsDBChainStaking.contract.WatchLogs(opts, "Unvoted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AsDBChainStakingUnvoted)
				if err := _AsDBChainStaking.contract.UnpackLog(event, "Unvoted", log); err != nil {
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

// ParseUnvoted is a log parse operation binding the contract event 0x7958395da8e26969cc7c09ee58e9507a2601574c3bd232617e2d6354224ff836.
//
// Solidity: event Unvoted(address voter, address candidate, uint256 amount)
func (_AsDBChainStaking *AsDBChainStakingFilterer) ParseUnvoted(log types.Log) (*AsDBChainStakingUnvoted, error) {
	event := new(AsDBChainStakingUnvoted)
	if err := _AsDBChainStaking.contract.UnpackLog(event, "Unvoted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AsDBChainStakingVotedIterator is returned from FilterVoted and is used to iterate over the raw logs and unpacked data for Voted events raised by the AsDBChainStaking contract.
type AsDBChainStakingVotedIterator struct {
	Event *AsDBChainStakingVoted // Event containing the contract specifics and raw log

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
func (it *AsDBChainStakingVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AsDBChainStakingVoted)
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
		it.Event = new(AsDBChainStakingVoted)
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
func (it *AsDBChainStakingVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AsDBChainStakingVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AsDBChainStakingVoted represents a Voted event raised by the AsDBChainStaking contract.
type AsDBChainStakingVoted struct {
	Voter     common.Address
	Candidate common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVoted is a free log retrieval operation binding the contract event 0x174ba19ba3c8bb5c679c87e51db79fff7c3f04bb84c1fd55b7dacb470b674aa6.
//
// Solidity: event Voted(address voter, address candidate, uint256 amount)
func (_AsDBChainStaking *AsDBChainStakingFilterer) FilterVoted(opts *bind.FilterOpts) (*AsDBChainStakingVotedIterator, error) {

	logs, sub, err := _AsDBChainStaking.contract.FilterLogs(opts, "Voted")
	if err != nil {
		return nil, err
	}
	return &AsDBChainStakingVotedIterator{contract: _AsDBChainStaking.contract, event: "Voted", logs: logs, sub: sub}, nil
}

// WatchVoted is a free log subscription operation binding the contract event 0x174ba19ba3c8bb5c679c87e51db79fff7c3f04bb84c1fd55b7dacb470b674aa6.
//
// Solidity: event Voted(address voter, address candidate, uint256 amount)
func (_AsDBChainStaking *AsDBChainStakingFilterer) WatchVoted(opts *bind.WatchOpts, sink chan<- *AsDBChainStakingVoted) (event.Subscription, error) {

	logs, sub, err := _AsDBChainStaking.contract.WatchLogs(opts, "Voted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AsDBChainStakingVoted)
				if err := _AsDBChainStaking.contract.UnpackLog(event, "Voted", log); err != nil {
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

// ParseVoted is a log parse operation binding the contract event 0x174ba19ba3c8bb5c679c87e51db79fff7c3f04bb84c1fd55b7dacb470b674aa6.
//
// Solidity: event Voted(address voter, address candidate, uint256 amount)
func (_AsDBChainStaking *AsDBChainStakingFilterer) ParseVoted(log types.Log) (*AsDBChainStakingVoted, error) {
	event := new(AsDBChainStakingVoted)
	if err := _AsDBChainStaking.contract.UnpackLog(event, "Voted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AsDBChainStakingWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the AsDBChainStaking contract.
type AsDBChainStakingWithdrawIterator struct {
	Event *AsDBChainStakingWithdraw // Event containing the contract specifics and raw log

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
func (it *AsDBChainStakingWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AsDBChainStakingWithdraw)
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
		it.Event = new(AsDBChainStakingWithdraw)
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
func (it *AsDBChainStakingWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AsDBChainStakingWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AsDBChainStakingWithdraw represents a Withdraw event raised by the AsDBChainStaking contract.
type AsDBChainStakingWithdraw struct {
	Staker      common.Address
	Amount      *big.Int
	DestAddress common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x56c54ba9bd38d8fd62012e42c7ee564519b09763c426d331b3661b537ead19b2.
//
// Solidity: event Withdraw(address _staker, uint256 _amount, address _destAddress)
func (_AsDBChainStaking *AsDBChainStakingFilterer) FilterWithdraw(opts *bind.FilterOpts) (*AsDBChainStakingWithdrawIterator, error) {

	logs, sub, err := _AsDBChainStaking.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &AsDBChainStakingWithdrawIterator{contract: _AsDBChainStaking.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x56c54ba9bd38d8fd62012e42c7ee564519b09763c426d331b3661b537ead19b2.
//
// Solidity: event Withdraw(address _staker, uint256 _amount, address _destAddress)
func (_AsDBChainStaking *AsDBChainStakingFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *AsDBChainStakingWithdraw) (event.Subscription, error) {

	logs, sub, err := _AsDBChainStaking.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AsDBChainStakingWithdraw)
				if err := _AsDBChainStaking.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x56c54ba9bd38d8fd62012e42c7ee564519b09763c426d331b3661b537ead19b2.
//
// Solidity: event Withdraw(address _staker, uint256 _amount, address _destAddress)
func (_AsDBChainStaking *AsDBChainStakingFilterer) ParseWithdraw(log types.Log) (*AsDBChainStakingWithdraw, error) {
	event := new(AsDBChainStakingWithdraw)
	if err := _AsDBChainStaking.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAsDBChainStakingABI is the input ABI used to generate the binding from.
const IAsDBChainStakingABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_candidate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"}],\"name\":\"Resigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Unvoted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Voted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_destAddress\",\"type\":\"address\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getListCandidates\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_candidates\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"stakes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValidatorCap\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_candidate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"resign\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"unvote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"destAddress\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"destAddress\",\"type\":\"address\"}],\"name\":\"withdrawWithIndex\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IAsDBChainStakingFuncSigs maps the 4-byte function signature to its string representation.
var IAsDBChainStakingFuncSigs = map[string]string{
	"690ff8a1": "getListCandidates()",
	"aa677354": "register(address,address)",
	"ae6e43f5": "resign(address)",
	"02aa9be2": "unvote(address,uint256)",
	"6dd7d8ea": "vote(address)",
	"00f714ce": "withdraw(uint256,address)",
	"c2df409b": "withdrawWithIndex(uint256,uint256,address)",
}

// IAsDBChainStaking is an auto generated Go binding around an Ethereum contract.
type IAsDBChainStaking struct {
	IAsDBChainStakingCaller     // Read-only binding to the contract
	IAsDBChainStakingTransactor // Write-only binding to the contract
	IAsDBChainStakingFilterer   // Log filterer for contract events
}

// IAsDBChainStakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAsDBChainStakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAsDBChainStakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAsDBChainStakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAsDBChainStakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAsDBChainStakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAsDBChainStakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAsDBChainStakingSession struct {
	Contract     *IAsDBChainStaking // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IAsDBChainStakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAsDBChainStakingCallerSession struct {
	Contract *IAsDBChainStakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IAsDBChainStakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAsDBChainStakingTransactorSession struct {
	Contract     *IAsDBChainStakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IAsDBChainStakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAsDBChainStakingRaw struct {
	Contract *IAsDBChainStaking // Generic contract binding to access the raw methods on
}

// IAsDBChainStakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAsDBChainStakingCallerRaw struct {
	Contract *IAsDBChainStakingCaller // Generic read-only contract binding to access the raw methods on
}

// IAsDBChainStakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAsDBChainStakingTransactorRaw struct {
	Contract *IAsDBChainStakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAsDBChainStaking creates a new instance of IAsDBChainStaking, bound to a specific deployed contract.
func NewIAsDBChainStaking(address common.Address, backend bind.ContractBackend) (*IAsDBChainStaking, error) {
	contract, err := bindIAsDBChainStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAsDBChainStaking{IAsDBChainStakingCaller: IAsDBChainStakingCaller{contract: contract}, IAsDBChainStakingTransactor: IAsDBChainStakingTransactor{contract: contract}, IAsDBChainStakingFilterer: IAsDBChainStakingFilterer{contract: contract}}, nil
}

// NewIAsDBChainStakingCaller creates a new read-only instance of IAsDBChainStaking, bound to a specific deployed contract.
func NewIAsDBChainStakingCaller(address common.Address, caller bind.ContractCaller) (*IAsDBChainStakingCaller, error) {
	contract, err := bindIAsDBChainStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAsDBChainStakingCaller{contract: contract}, nil
}

// NewIAsDBChainStakingTransactor creates a new write-only instance of IAsDBChainStaking, bound to a specific deployed contract.
func NewIAsDBChainStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*IAsDBChainStakingTransactor, error) {
	contract, err := bindIAsDBChainStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAsDBChainStakingTransactor{contract: contract}, nil
}

// NewIAsDBChainStakingFilterer creates a new log filterer instance of IAsDBChainStaking, bound to a specific deployed contract.
func NewIAsDBChainStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*IAsDBChainStakingFilterer, error) {
	contract, err := bindIAsDBChainStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAsDBChainStakingFilterer{contract: contract}, nil
}

// bindIAsDBChainStaking binds a generic wrapper to an already deployed contract.
func bindIAsDBChainStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IAsDBChainStakingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAsDBChainStaking *IAsDBChainStakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAsDBChainStaking.Contract.IAsDBChainStakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAsDBChainStaking *IAsDBChainStakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAsDBChainStaking.Contract.IAsDBChainStakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAsDBChainStaking *IAsDBChainStakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAsDBChainStaking.Contract.IAsDBChainStakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAsDBChainStaking *IAsDBChainStakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAsDBChainStaking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAsDBChainStaking *IAsDBChainStakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAsDBChainStaking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAsDBChainStaking *IAsDBChainStakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAsDBChainStaking.Contract.contract.Transact(opts, method, params...)
}

// GetListCandidates is a free data retrieval call binding the contract method 0x690ff8a1.
//
// Solidity: function getListCandidates() view returns(address[] _candidates, uint256[] stakes, uint256 epoch, uint256 validatorSize, uint256 minValidatorCap)
func (_IAsDBChainStaking *IAsDBChainStakingCaller) GetListCandidates(opts *bind.CallOpts) (struct {
	Candidates      []common.Address
	Stakes          []*big.Int
	Epoch           *big.Int
	ValidatorSize   *big.Int
	MinValidatorCap *big.Int
}, error) {
	var out []interface{}
	err := _IAsDBChainStaking.contract.Call(opts, &out, "getListCandidates")

	outstruct := new(struct {
		Candidates      []common.Address
		Stakes          []*big.Int
		Epoch           *big.Int
		ValidatorSize   *big.Int
		MinValidatorCap *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Candidates = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Stakes = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.Epoch = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ValidatorSize = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.MinValidatorCap = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetListCandidates is a free data retrieval call binding the contract method 0x690ff8a1.
//
// Solidity: function getListCandidates() view returns(address[] _candidates, uint256[] stakes, uint256 epoch, uint256 validatorSize, uint256 minValidatorCap)
func (_IAsDBChainStaking *IAsDBChainStakingSession) GetListCandidates() (struct {
	Candidates      []common.Address
	Stakes          []*big.Int
	Epoch           *big.Int
	ValidatorSize   *big.Int
	MinValidatorCap *big.Int
}, error) {
	return _IAsDBChainStaking.Contract.GetListCandidates(&_IAsDBChainStaking.CallOpts)
}

// GetListCandidates is a free data retrieval call binding the contract method 0x690ff8a1.
//
// Solidity: function getListCandidates() view returns(address[] _candidates, uint256[] stakes, uint256 epoch, uint256 validatorSize, uint256 minValidatorCap)
func (_IAsDBChainStaking *IAsDBChainStakingCallerSession) GetListCandidates() (struct {
	Candidates      []common.Address
	Stakes          []*big.Int
	Epoch           *big.Int
	ValidatorSize   *big.Int
	MinValidatorCap *big.Int
}, error) {
	return _IAsDBChainStaking.Contract.GetListCandidates(&_IAsDBChainStaking.CallOpts)
}

// Register is a paid mutator transaction binding the contract method 0xaa677354.
//
// Solidity: function register(address _candidate, address _owner) returns()
func (_IAsDBChainStaking *IAsDBChainStakingTransactor) Register(opts *bind.TransactOpts, _candidate common.Address, _owner common.Address) (*types.Transaction, error) {
	return _IAsDBChainStaking.contract.Transact(opts, "register", _candidate, _owner)
}

// Register is a paid mutator transaction binding the contract method 0xaa677354.
//
// Solidity: function register(address _candidate, address _owner) returns()
func (_IAsDBChainStaking *IAsDBChainStakingSession) Register(_candidate common.Address, _owner common.Address) (*types.Transaction, error) {
	return _IAsDBChainStaking.Contract.Register(&_IAsDBChainStaking.TransactOpts, _candidate, _owner)
}

// Register is a paid mutator transaction binding the contract method 0xaa677354.
//
// Solidity: function register(address _candidate, address _owner) returns()
func (_IAsDBChainStaking *IAsDBChainStakingTransactorSession) Register(_candidate common.Address, _owner common.Address) (*types.Transaction, error) {
	return _IAsDBChainStaking.Contract.Register(&_IAsDBChainStaking.TransactOpts, _candidate, _owner)
}

// Resign is a paid mutator transaction binding the contract method 0xae6e43f5.
//
// Solidity: function resign(address _candidate) returns()
func (_IAsDBChainStaking *IAsDBChainStakingTransactor) Resign(opts *bind.TransactOpts, _candidate common.Address) (*types.Transaction, error) {
	return _IAsDBChainStaking.contract.Transact(opts, "resign", _candidate)
}

// Resign is a paid mutator transaction binding the contract method 0xae6e43f5.
//
// Solidity: function resign(address _candidate) returns()
func (_IAsDBChainStaking *IAsDBChainStakingSession) Resign(_candidate common.Address) (*types.Transaction, error) {
	return _IAsDBChainStaking.Contract.Resign(&_IAsDBChainStaking.TransactOpts, _candidate)
}

// Resign is a paid mutator transaction binding the contract method 0xae6e43f5.
//
// Solidity: function resign(address _candidate) returns()
func (_IAsDBChainStaking *IAsDBChainStakingTransactorSession) Resign(_candidate common.Address) (*types.Transaction, error) {
	return _IAsDBChainStaking.Contract.Resign(&_IAsDBChainStaking.TransactOpts, _candidate)
}

// Unvote is a paid mutator transaction binding the contract method 0x02aa9be2.
//
// Solidity: function unvote(address candidate, uint256 amount) returns()
func (_IAsDBChainStaking *IAsDBChainStakingTransactor) Unvote(opts *bind.TransactOpts, candidate common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IAsDBChainStaking.contract.Transact(opts, "unvote", candidate, amount)
}

// Unvote is a paid mutator transaction binding the contract method 0x02aa9be2.
//
// Solidity: function unvote(address candidate, uint256 amount) returns()
func (_IAsDBChainStaking *IAsDBChainStakingSession) Unvote(candidate common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IAsDBChainStaking.Contract.Unvote(&_IAsDBChainStaking.TransactOpts, candidate, amount)
}

// Unvote is a paid mutator transaction binding the contract method 0x02aa9be2.
//
// Solidity: function unvote(address candidate, uint256 amount) returns()
func (_IAsDBChainStaking *IAsDBChainStakingTransactorSession) Unvote(candidate common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IAsDBChainStaking.Contract.Unvote(&_IAsDBChainStaking.TransactOpts, candidate, amount)
}

// Vote is a paid mutator transaction binding the contract method 0x6dd7d8ea.
//
// Solidity: function vote(address candidate) payable returns()
func (_IAsDBChainStaking *IAsDBChainStakingTransactor) Vote(opts *bind.TransactOpts, candidate common.Address) (*types.Transaction, error) {
	return _IAsDBChainStaking.contract.Transact(opts, "vote", candidate)
}

// Vote is a paid mutator transaction binding the contract method 0x6dd7d8ea.
//
// Solidity: function vote(address candidate) payable returns()
func (_IAsDBChainStaking *IAsDBChainStakingSession) Vote(candidate common.Address) (*types.Transaction, error) {
	return _IAsDBChainStaking.Contract.Vote(&_IAsDBChainStaking.TransactOpts, candidate)
}

// Vote is a paid mutator transaction binding the contract method 0x6dd7d8ea.
//
// Solidity: function vote(address candidate) payable returns()
func (_IAsDBChainStaking *IAsDBChainStakingTransactorSession) Vote(candidate common.Address) (*types.Transaction, error) {
	return _IAsDBChainStaking.Contract.Vote(&_IAsDBChainStaking.TransactOpts, candidate)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 epoch, address destAddress) returns(bool)
func (_IAsDBChainStaking *IAsDBChainStakingTransactor) Withdraw(opts *bind.TransactOpts, epoch *big.Int, destAddress common.Address) (*types.Transaction, error) {
	return _IAsDBChainStaking.contract.Transact(opts, "withdraw", epoch, destAddress)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 epoch, address destAddress) returns(bool)
func (_IAsDBChainStaking *IAsDBChainStakingSession) Withdraw(epoch *big.Int, destAddress common.Address) (*types.Transaction, error) {
	return _IAsDBChainStaking.Contract.Withdraw(&_IAsDBChainStaking.TransactOpts, epoch, destAddress)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 epoch, address destAddress) returns(bool)
func (_IAsDBChainStaking *IAsDBChainStakingTransactorSession) Withdraw(epoch *big.Int, destAddress common.Address) (*types.Transaction, error) {
	return _IAsDBChainStaking.Contract.Withdraw(&_IAsDBChainStaking.TransactOpts, epoch, destAddress)
}

// WithdrawWithIndex is a paid mutator transaction binding the contract method 0xc2df409b.
//
// Solidity: function withdrawWithIndex(uint256 epoch, uint256 index, address destAddress) returns(bool)
func (_IAsDBChainStaking *IAsDBChainStakingTransactor) WithdrawWithIndex(opts *bind.TransactOpts, epoch *big.Int, index *big.Int, destAddress common.Address) (*types.Transaction, error) {
	return _IAsDBChainStaking.contract.Transact(opts, "withdrawWithIndex", epoch, index, destAddress)
}

// WithdrawWithIndex is a paid mutator transaction binding the contract method 0xc2df409b.
//
// Solidity: function withdrawWithIndex(uint256 epoch, uint256 index, address destAddress) returns(bool)
func (_IAsDBChainStaking *IAsDBChainStakingSession) WithdrawWithIndex(epoch *big.Int, index *big.Int, destAddress common.Address) (*types.Transaction, error) {
	return _IAsDBChainStaking.Contract.WithdrawWithIndex(&_IAsDBChainStaking.TransactOpts, epoch, index, destAddress)
}

// WithdrawWithIndex is a paid mutator transaction binding the contract method 0xc2df409b.
//
// Solidity: function withdrawWithIndex(uint256 epoch, uint256 index, address destAddress) returns(bool)
func (_IAsDBChainStaking *IAsDBChainStakingTransactorSession) WithdrawWithIndex(epoch *big.Int, index *big.Int, destAddress common.Address) (*types.Transaction, error) {
	return _IAsDBChainStaking.Contract.WithdrawWithIndex(&_IAsDBChainStaking.TransactOpts, epoch, index, destAddress)
}

// IAsDBChainStakingRegisteredIterator is returned from FilterRegistered and is used to iterate over the raw logs and unpacked data for Registered events raised by the IAsDBChainStaking contract.
type IAsDBChainStakingRegisteredIterator struct {
	Event *IAsDBChainStakingRegistered // Event containing the contract specifics and raw log

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
func (it *IAsDBChainStakingRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAsDBChainStakingRegistered)
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
		it.Event = new(IAsDBChainStakingRegistered)
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
func (it *IAsDBChainStakingRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAsDBChainStakingRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAsDBChainStakingRegistered represents a Registered event raised by the IAsDBChainStaking contract.
type IAsDBChainStakingRegistered struct {
	Candidate common.Address
	Owner     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRegistered is a free log retrieval operation binding the contract event 0x0a31ee9d46a828884b81003c8498156ea6aa15b9b54bdd0ef0b533d9eba57e55.
//
// Solidity: event Registered(address candidate, address owner)
func (_IAsDBChainStaking *IAsDBChainStakingFilterer) FilterRegistered(opts *bind.FilterOpts) (*IAsDBChainStakingRegisteredIterator, error) {

	logs, sub, err := _IAsDBChainStaking.contract.FilterLogs(opts, "Registered")
	if err != nil {
		return nil, err
	}
	return &IAsDBChainStakingRegisteredIterator{contract: _IAsDBChainStaking.contract, event: "Registered", logs: logs, sub: sub}, nil
}

// WatchRegistered is a free log subscription operation binding the contract event 0x0a31ee9d46a828884b81003c8498156ea6aa15b9b54bdd0ef0b533d9eba57e55.
//
// Solidity: event Registered(address candidate, address owner)
func (_IAsDBChainStaking *IAsDBChainStakingFilterer) WatchRegistered(opts *bind.WatchOpts, sink chan<- *IAsDBChainStakingRegistered) (event.Subscription, error) {

	logs, sub, err := _IAsDBChainStaking.contract.WatchLogs(opts, "Registered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAsDBChainStakingRegistered)
				if err := _IAsDBChainStaking.contract.UnpackLog(event, "Registered", log); err != nil {
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

// ParseRegistered is a log parse operation binding the contract event 0x0a31ee9d46a828884b81003c8498156ea6aa15b9b54bdd0ef0b533d9eba57e55.
//
// Solidity: event Registered(address candidate, address owner)
func (_IAsDBChainStaking *IAsDBChainStakingFilterer) ParseRegistered(log types.Log) (*IAsDBChainStakingRegistered, error) {
	event := new(IAsDBChainStakingRegistered)
	if err := _IAsDBChainStaking.contract.UnpackLog(event, "Registered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAsDBChainStakingResignedIterator is returned from FilterResigned and is used to iterate over the raw logs and unpacked data for Resigned events raised by the IAsDBChainStaking contract.
type IAsDBChainStakingResignedIterator struct {
	Event *IAsDBChainStakingResigned // Event containing the contract specifics and raw log

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
func (it *IAsDBChainStakingResignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAsDBChainStakingResigned)
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
		it.Event = new(IAsDBChainStakingResigned)
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
func (it *IAsDBChainStakingResignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAsDBChainStakingResignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAsDBChainStakingResigned represents a Resigned event raised by the IAsDBChainStaking contract.
type IAsDBChainStakingResigned struct {
	Candidate common.Address
	Epoch     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterResigned is a free log retrieval operation binding the contract event 0x886e0db046874dde595498040d176412e81183750ceb33fc46f0450362bbc241.
//
// Solidity: event Resigned(address _candidate, uint256 _epoch)
func (_IAsDBChainStaking *IAsDBChainStakingFilterer) FilterResigned(opts *bind.FilterOpts) (*IAsDBChainStakingResignedIterator, error) {

	logs, sub, err := _IAsDBChainStaking.contract.FilterLogs(opts, "Resigned")
	if err != nil {
		return nil, err
	}
	return &IAsDBChainStakingResignedIterator{contract: _IAsDBChainStaking.contract, event: "Resigned", logs: logs, sub: sub}, nil
}

// WatchResigned is a free log subscription operation binding the contract event 0x886e0db046874dde595498040d176412e81183750ceb33fc46f0450362bbc241.
//
// Solidity: event Resigned(address _candidate, uint256 _epoch)
func (_IAsDBChainStaking *IAsDBChainStakingFilterer) WatchResigned(opts *bind.WatchOpts, sink chan<- *IAsDBChainStakingResigned) (event.Subscription, error) {

	logs, sub, err := _IAsDBChainStaking.contract.WatchLogs(opts, "Resigned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAsDBChainStakingResigned)
				if err := _IAsDBChainStaking.contract.UnpackLog(event, "Resigned", log); err != nil {
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

// ParseResigned is a log parse operation binding the contract event 0x886e0db046874dde595498040d176412e81183750ceb33fc46f0450362bbc241.
//
// Solidity: event Resigned(address _candidate, uint256 _epoch)
func (_IAsDBChainStaking *IAsDBChainStakingFilterer) ParseResigned(log types.Log) (*IAsDBChainStakingResigned, error) {
	event := new(IAsDBChainStakingResigned)
	if err := _IAsDBChainStaking.contract.UnpackLog(event, "Resigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAsDBChainStakingUnvotedIterator is returned from FilterUnvoted and is used to iterate over the raw logs and unpacked data for Unvoted events raised by the IAsDBChainStaking contract.
type IAsDBChainStakingUnvotedIterator struct {
	Event *IAsDBChainStakingUnvoted // Event containing the contract specifics and raw log

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
func (it *IAsDBChainStakingUnvotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAsDBChainStakingUnvoted)
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
		it.Event = new(IAsDBChainStakingUnvoted)
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
func (it *IAsDBChainStakingUnvotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAsDBChainStakingUnvotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAsDBChainStakingUnvoted represents a Unvoted event raised by the IAsDBChainStaking contract.
type IAsDBChainStakingUnvoted struct {
	Voter     common.Address
	Candidate common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnvoted is a free log retrieval operation binding the contract event 0x7958395da8e26969cc7c09ee58e9507a2601574c3bd232617e2d6354224ff836.
//
// Solidity: event Unvoted(address voter, address candidate, uint256 amount)
func (_IAsDBChainStaking *IAsDBChainStakingFilterer) FilterUnvoted(opts *bind.FilterOpts) (*IAsDBChainStakingUnvotedIterator, error) {

	logs, sub, err := _IAsDBChainStaking.contract.FilterLogs(opts, "Unvoted")
	if err != nil {
		return nil, err
	}
	return &IAsDBChainStakingUnvotedIterator{contract: _IAsDBChainStaking.contract, event: "Unvoted", logs: logs, sub: sub}, nil
}

// WatchUnvoted is a free log subscription operation binding the contract event 0x7958395da8e26969cc7c09ee58e9507a2601574c3bd232617e2d6354224ff836.
//
// Solidity: event Unvoted(address voter, address candidate, uint256 amount)
func (_IAsDBChainStaking *IAsDBChainStakingFilterer) WatchUnvoted(opts *bind.WatchOpts, sink chan<- *IAsDBChainStakingUnvoted) (event.Subscription, error) {

	logs, sub, err := _IAsDBChainStaking.contract.WatchLogs(opts, "Unvoted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAsDBChainStakingUnvoted)
				if err := _IAsDBChainStaking.contract.UnpackLog(event, "Unvoted", log); err != nil {
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

// ParseUnvoted is a log parse operation binding the contract event 0x7958395da8e26969cc7c09ee58e9507a2601574c3bd232617e2d6354224ff836.
//
// Solidity: event Unvoted(address voter, address candidate, uint256 amount)
func (_IAsDBChainStaking *IAsDBChainStakingFilterer) ParseUnvoted(log types.Log) (*IAsDBChainStakingUnvoted, error) {
	event := new(IAsDBChainStakingUnvoted)
	if err := _IAsDBChainStaking.contract.UnpackLog(event, "Unvoted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAsDBChainStakingVotedIterator is returned from FilterVoted and is used to iterate over the raw logs and unpacked data for Voted events raised by the IAsDBChainStaking contract.
type IAsDBChainStakingVotedIterator struct {
	Event *IAsDBChainStakingVoted // Event containing the contract specifics and raw log

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
func (it *IAsDBChainStakingVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAsDBChainStakingVoted)
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
		it.Event = new(IAsDBChainStakingVoted)
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
func (it *IAsDBChainStakingVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAsDBChainStakingVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAsDBChainStakingVoted represents a Voted event raised by the IAsDBChainStaking contract.
type IAsDBChainStakingVoted struct {
	Voter     common.Address
	Candidate common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVoted is a free log retrieval operation binding the contract event 0x174ba19ba3c8bb5c679c87e51db79fff7c3f04bb84c1fd55b7dacb470b674aa6.
//
// Solidity: event Voted(address voter, address candidate, uint256 amount)
func (_IAsDBChainStaking *IAsDBChainStakingFilterer) FilterVoted(opts *bind.FilterOpts) (*IAsDBChainStakingVotedIterator, error) {

	logs, sub, err := _IAsDBChainStaking.contract.FilterLogs(opts, "Voted")
	if err != nil {
		return nil, err
	}
	return &IAsDBChainStakingVotedIterator{contract: _IAsDBChainStaking.contract, event: "Voted", logs: logs, sub: sub}, nil
}

// WatchVoted is a free log subscription operation binding the contract event 0x174ba19ba3c8bb5c679c87e51db79fff7c3f04bb84c1fd55b7dacb470b674aa6.
//
// Solidity: event Voted(address voter, address candidate, uint256 amount)
func (_IAsDBChainStaking *IAsDBChainStakingFilterer) WatchVoted(opts *bind.WatchOpts, sink chan<- *IAsDBChainStakingVoted) (event.Subscription, error) {

	logs, sub, err := _IAsDBChainStaking.contract.WatchLogs(opts, "Voted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAsDBChainStakingVoted)
				if err := _IAsDBChainStaking.contract.UnpackLog(event, "Voted", log); err != nil {
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

// ParseVoted is a log parse operation binding the contract event 0x174ba19ba3c8bb5c679c87e51db79fff7c3f04bb84c1fd55b7dacb470b674aa6.
//
// Solidity: event Voted(address voter, address candidate, uint256 amount)
func (_IAsDBChainStaking *IAsDBChainStakingFilterer) ParseVoted(log types.Log) (*IAsDBChainStakingVoted, error) {
	event := new(IAsDBChainStakingVoted)
	if err := _IAsDBChainStaking.contract.UnpackLog(event, "Voted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAsDBChainStakingWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the IAsDBChainStaking contract.
type IAsDBChainStakingWithdrawIterator struct {
	Event *IAsDBChainStakingWithdraw // Event containing the contract specifics and raw log

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
func (it *IAsDBChainStakingWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAsDBChainStakingWithdraw)
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
		it.Event = new(IAsDBChainStakingWithdraw)
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
func (it *IAsDBChainStakingWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAsDBChainStakingWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAsDBChainStakingWithdraw represents a Withdraw event raised by the IAsDBChainStaking contract.
type IAsDBChainStakingWithdraw struct {
	Staker      common.Address
	Amount      *big.Int
	DestAddress common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x56c54ba9bd38d8fd62012e42c7ee564519b09763c426d331b3661b537ead19b2.
//
// Solidity: event Withdraw(address _staker, uint256 _amount, address _destAddress)
func (_IAsDBChainStaking *IAsDBChainStakingFilterer) FilterWithdraw(opts *bind.FilterOpts) (*IAsDBChainStakingWithdrawIterator, error) {

	logs, sub, err := _IAsDBChainStaking.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &IAsDBChainStakingWithdrawIterator{contract: _IAsDBChainStaking.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x56c54ba9bd38d8fd62012e42c7ee564519b09763c426d331b3661b537ead19b2.
//
// Solidity: event Withdraw(address _staker, uint256 _amount, address _destAddress)
func (_IAsDBChainStaking *IAsDBChainStakingFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *IAsDBChainStakingWithdraw) (event.Subscription, error) {

	logs, sub, err := _IAsDBChainStaking.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAsDBChainStakingWithdraw)
				if err := _IAsDBChainStaking.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x56c54ba9bd38d8fd62012e42c7ee564519b09763c426d331b3661b537ead19b2.
//
// Solidity: event Withdraw(address _staker, uint256 _amount, address _destAddress)
func (_IAsDBChainStaking *IAsDBChainStakingFilterer) ParseWithdraw(log types.Log) (*IAsDBChainStakingWithdraw, error) {
	event := new(IAsDBChainStakingWithdraw)
	if err := _IAsDBChainStaking.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReentrancyGuardABI is the input ABI used to generate the binding from.
const ReentrancyGuardABI = "[]"

// ReentrancyGuard is an auto generated Go binding around an Ethereum contract.
type ReentrancyGuard struct {
	ReentrancyGuardCaller     // Read-only binding to the contract
	ReentrancyGuardTransactor // Write-only binding to the contract
	ReentrancyGuardFilterer   // Log filterer for contract events
}

// ReentrancyGuardCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReentrancyGuardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReentrancyGuardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReentrancyGuardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReentrancyGuardSession struct {
	Contract     *ReentrancyGuard  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReentrancyGuardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReentrancyGuardCallerSession struct {
	Contract *ReentrancyGuardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ReentrancyGuardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReentrancyGuardTransactorSession struct {
	Contract     *ReentrancyGuardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ReentrancyGuardRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReentrancyGuardRaw struct {
	Contract *ReentrancyGuard // Generic contract binding to access the raw methods on
}

// ReentrancyGuardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReentrancyGuardCallerRaw struct {
	Contract *ReentrancyGuardCaller // Generic read-only contract binding to access the raw methods on
}

// ReentrancyGuardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReentrancyGuardTransactorRaw struct {
	Contract *ReentrancyGuardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReentrancyGuard creates a new instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuard(address common.Address, backend bind.ContractBackend) (*ReentrancyGuard, error) {
	contract, err := bindReentrancyGuard(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuard{ReentrancyGuardCaller: ReentrancyGuardCaller{contract: contract}, ReentrancyGuardTransactor: ReentrancyGuardTransactor{contract: contract}, ReentrancyGuardFilterer: ReentrancyGuardFilterer{contract: contract}}, nil
}

// NewReentrancyGuardCaller creates a new read-only instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardCaller(address common.Address, caller bind.ContractCaller) (*ReentrancyGuardCaller, error) {
	contract, err := bindReentrancyGuard(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardCaller{contract: contract}, nil
}

// NewReentrancyGuardTransactor creates a new write-only instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardTransactor(address common.Address, transactor bind.ContractTransactor) (*ReentrancyGuardTransactor, error) {
	contract, err := bindReentrancyGuard(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardTransactor{contract: contract}, nil
}

// NewReentrancyGuardFilterer creates a new log filterer instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardFilterer(address common.Address, filterer bind.ContractFilterer) (*ReentrancyGuardFilterer, error) {
	contract, err := bindReentrancyGuard(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardFilterer{contract: contract}, nil
}

// bindReentrancyGuard binds a generic wrapper to an already deployed contract.
func bindReentrancyGuard(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReentrancyGuardABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuard *ReentrancyGuardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuard.Contract.ReentrancyGuardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuard *ReentrancyGuardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.ReentrancyGuardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuard *ReentrancyGuardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.ReentrancyGuardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuard *ReentrancyGuardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuard.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuard *ReentrancyGuardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuard *ReentrancyGuardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.contract.Transact(opts, method, params...)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212205932f71d5e06919df9473ed0b3828447aa2ff6efe581743be70b9905dc8a4ae664736f6c63430007060033"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}
