// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package consensus

import (
	"math/big"
	"strings"

	kowala "github.com/kowala-tech/kcoin/client"
	"github.com/kowala-tech/kcoin/client/accounts/abi"
	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/core/types"
	"github.com/kowala-tech/kcoin/client/event"
)

// MiningTokenABI is the input ABI used to generate the binding from.
const MiningTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"mintingFinished\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"_name\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"_totalSupply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"_decimals\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_cap\",\"type\":\"uint256\"},{\"name\":\"_decimals\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finishMinting\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"_symbol\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"},{\"name\":\"_custom_fallback\",\"type\":\"string\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_cap\",\"type\":\"uint256\"},{\"name\":\"_decimals\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"MintFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// MiningTokenBin is the compiled bytecode used for deploying new contracts.
const MiningTokenBin = `60806040526000600760146101000a81548160ff0219169083151502179055503480156200002c57600080fd5b5060405162001dc438038062001dc4833981018060405281019080805182019291906020018051820192919060200180519060200190929190805190602001909291905050508133600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600081111515620000c457600080fd5b80600881905550508360039080519060200190620000e492919062000123565b508260049080519060200190620000fd92919062000123565b5080600560006101000a81548160ff021916908360ff16021790555050505050620001d2565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200016657805160ff191683800117855562000197565b8280016001018555821562000197579182015b828111156200019657825182559160200191906001019062000179565b5b509050620001a69190620001aa565b5090565b620001cf91905b80821115620001cb576000816000905550600101620001b1565b5090565b90565b611be280620001e26000396000f3006080604052600436106100fc576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806305d2035b1461010157806306fdde0314610130578063158ef93e146101c057806318160ddd146101ef578063313ce5671461021a578063355274ea1461024b57806340c10f191461027657806341658f3c146102db57806370a08231146103a1578063715018a6146103f85780637d64bcb41461040f578063893d20e81461043e5780638da5cb5b1461049557806395d89b41146104ec578063a9059cbb1461057c578063be45fd62146105e1578063f2fde38b1461068c578063f6368f8a146106cf575b600080fd5b34801561010d57600080fd5b506101166107c0565b604051808215151515815260200191505060405180910390f35b34801561013c57600080fd5b506101456107d3565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561018557808201518184015260208101905061016a565b50505050905090810190601f1680156101b25780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101cc57600080fd5b506101d5610875565b604051808215151515815260200191505060405180910390f35b3480156101fb57600080fd5b50610204610888565b6040518082815260200191505060405180910390f35b34801561022657600080fd5b5061022f610892565b604051808260ff1660ff16815260200191505060405180910390f35b34801561025757600080fd5b506102606108a9565b6040518082815260200191505060405180910390f35b34801561028257600080fd5b506102c1600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506108af565b604051808215151515815260200191505060405180910390f35b3480156102e757600080fd5b5061039f600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929080359060200190929190803560ff169060200190929190505050610960565b005b3480156103ad57600080fd5b506103e2600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610acb565b6040518082815260200191505060405180910390f35b34801561040457600080fd5b5061040d610b14565b005b34801561041b57600080fd5b50610424610c19565b604051808215151515815260200191505060405180910390f35b34801561044a57600080fd5b50610453610ce1565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156104a157600080fd5b506104aa610d0b565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156104f857600080fd5b50610501610d31565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610541578082015181840152602081019050610526565b50505050905090810190601f16801561056e5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561058857600080fd5b506105c7600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610dd3565b604051808215151515815260200191505060405180910390f35b3480156105ed57600080fd5b50610672600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610e0c565b604051808215151515815260200191505060405180910390f35b34801561069857600080fd5b506106cd600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610e43565b005b3480156106db57600080fd5b506107a6600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610eab565b604051808215151515815260200191505060405180910390f35b600760149054906101000a900460ff1681565b606060038054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561086b5780601f106108405761010080835404028352916020019161086b565b820191906000526020600020905b81548152906001019060200180831161084e57829003601f168201915b5050505050905090565b600160009054906101000a900460ff1681565b6000600654905090565b6000600560009054906101000a900460ff16905090565b60085481565b6000600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561090d57600080fd5b600760149054906101000a900460ff1615151561092957600080fd5b6008546109418360065461126890919063ffffffff16565b1115151561094e57600080fd5b6109588383611284565b905092915050565b600160009054906101000a900460ff16151515610a0b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602e8152602001807f436f6e747261637420696e7374616e63652068617320616c726561647920626581526020017f656e20696e697469616c697a656400000000000000000000000000000000000081525060400191505060405180910390fd5b600082111515610a1a57600080fd5b8160088190555033600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508360039080519060200190610a78929190611b11565b508260049080519060200190610a8f929190611b11565b5080600560006101000a81548160ff021916908360ff16021790555060018060006101000a81548160ff02191690831515021790555050505050565b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610b7057600080fd5b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482060405160405180910390a26000600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6000600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610c7757600080fd5b600760149054906101000a900460ff16151515610c9357600080fd5b6001600760146101000a81548160ff0219169083151502179055507fae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa0860405160405180910390a16001905090565b6000600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b606060048054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610dc95780601f10610d9e57610100808354040283529160200191610dc9565b820191906000526020600020905b815481529060010190602001808311610dac57829003601f168201915b5050505050905090565b60006060610de084611480565b15610df757610df0848483611493565b9150610e05565b610e028484836117de565b91505b5092915050565b6000610e1784611480565b15610e2e57610e27848484611493565b9050610e3c565b610e398484846117de565b90505b9392505050565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610e9f57600080fd5b610ea8816119fc565b50565b6000610eb685611480565b156112525783610ec533610acb565b1015610ed057600080fd5b610f2284600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611af890919063ffffffff16565b600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610fb784600260008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461126890919063ffffffff16565b600260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508473ffffffffffffffffffffffffffffffffffffffff166000836040518082805190602001908083835b6020831015156110495780518252602082019150602081019050602083039250611024565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390207c01000000000000000000000000000000000000000000000000000000009004903387876040518563ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001838152602001828051906020019080838360005b8381101561112a57808201518184015260208101905061110f565b50505050905090810190601f1680156111575780820380516001836020036101000a031916815260200191505b50935050505060006040518083038185885af19350505050151561117757fe5b8473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c1686866040518083815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561120e5780820151818401526020810190506111f3565b50505050905090810190601f16801561123b5780820380516001836020036101000a031916815260200191505b50935050505060405180910390a360019050611260565b61125d8585856117de565b90505b949350505050565b6000818301905082811015151561127b57fe5b80905092915050565b6000600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156112e257600080fd5b600760149054906101000a900460ff161515156112fe57600080fd5b6113138260065461126890919063ffffffff16565b60068190555061136b82600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461126890919063ffffffff16565b600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885836040518082815260200191505060405180910390a28273ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c168460405180828152602001806020018281038252600081526020016020019250505060405180910390a36001905092915050565b600080823b905060008111915050919050565b600080836114a033610acb565b10156114ab57600080fd5b6114fd84600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611af890919063ffffffff16565b600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555061159284600260008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461126890919063ffffffff16565b600260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508490508073ffffffffffffffffffffffffffffffffffffffff1663c0ee0b8a3386866040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561169a57808201518184015260208101905061167f565b50505050905090810190601f1680156116c75780820380516001836020036101000a031916815260200191505b50945050505050600060405180830381600087803b1580156116e857600080fd5b505af11580156116fc573d6000803e3d6000fd5b505050508473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c1686866040518083815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561179757808201518184015260208101905061177c565b50505050905090810190601f1680156117c45780820380516001836020036101000a031916815260200191505b50935050505060405180910390a360019150509392505050565b6000826117ea33610acb565b10156117f557600080fd5b61184783600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611af890919063ffffffff16565b600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506118dc83600260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461126890919063ffffffff16565b600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508373ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c1685856040518083815260200180602001828103825283818151815260200191508051906020019080838360005b838110156119b657808201518184015260208101905061199b565b50505050905090810190601f1680156119e35780820380516001836020036101000a031916815260200191505b50935050505060405180910390a3600190509392505050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515611a3857600080fd5b8073ffffffffffffffffffffffffffffffffffffffff16600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a380600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000828211151515611b0657fe5b818303905092915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611b5257805160ff1916838001178555611b80565b82800160010185558215611b80579182015b82811115611b7f578251825591602001919060010190611b64565b5b509050611b8d9190611b91565b5090565b611bb391905b80821115611baf576000816000905550600101611b97565b5090565b905600a165627a7a723058205e5cc50b54a1435f03d0f8490f628e5477e3c71169f264d34c47d0a49253bbfc0029`

// DeployMiningToken deploys a new Kowala contract, binding an instance of MiningToken to it.
func DeployMiningToken(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _symbol string, _cap *big.Int, _decimals uint8) (common.Address, *types.Transaction, *MiningToken, error) {
	parsed, err := abi.JSON(strings.NewReader(MiningTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MiningTokenBin), backend, _name, _symbol, _cap, _decimals)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MiningToken{MiningTokenCaller: MiningTokenCaller{contract: contract}, MiningTokenTransactor: MiningTokenTransactor{contract: contract}, MiningTokenFilterer: MiningTokenFilterer{contract: contract}}, nil
}

// MiningToken is an auto generated Go binding around a Kowala contract.
type MiningToken struct {
	MiningTokenCaller     // Read-only binding to the contract
	MiningTokenTransactor // Write-only binding to the contract
	MiningTokenFilterer   // Log filterer for contract events
}

// MiningTokenCaller is an auto generated read-only Go binding around a Kowala contract.
type MiningTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MiningTokenTransactor is an auto generated write-only Go binding around a Kowala contract.
type MiningTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MiningTokenFilterer is an auto generated log filtering Go binding around a Kowala contract events.
type MiningTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MiningTokenSession is an auto generated Go binding around a Kowala contract,
// with pre-set call and transact options.
type MiningTokenSession struct {
	Contract     *MiningToken      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MiningTokenCallerSession is an auto generated read-only Go binding around a Kowala contract,
// with pre-set call options.
type MiningTokenCallerSession struct {
	Contract *MiningTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MiningTokenTransactorSession is an auto generated write-only Go binding around a Kowala contract,
// with pre-set transact options.
type MiningTokenTransactorSession struct {
	Contract     *MiningTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MiningTokenRaw is an auto generated low-level Go binding around a Kowala contract.
type MiningTokenRaw struct {
	Contract *MiningToken // Generic contract binding to access the raw methods on
}

// MiningTokenCallerRaw is an auto generated low-level read-only Go binding around a Kowala contract.
type MiningTokenCallerRaw struct {
	Contract *MiningTokenCaller // Generic read-only contract binding to access the raw methods on
}

// MiningTokenTransactorRaw is an auto generated low-level write-only Go binding around a Kowala contract.
type MiningTokenTransactorRaw struct {
	Contract *MiningTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMiningToken creates a new instance of MiningToken, bound to a specific deployed contract.
func NewMiningToken(address common.Address, backend bind.ContractBackend) (*MiningToken, error) {
	contract, err := bindMiningToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MiningToken{MiningTokenCaller: MiningTokenCaller{contract: contract}, MiningTokenTransactor: MiningTokenTransactor{contract: contract}, MiningTokenFilterer: MiningTokenFilterer{contract: contract}}, nil
}

// NewMiningTokenCaller creates a new read-only instance of MiningToken, bound to a specific deployed contract.
func NewMiningTokenCaller(address common.Address, caller bind.ContractCaller) (*MiningTokenCaller, error) {
	contract, err := bindMiningToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MiningTokenCaller{contract: contract}, nil
}

// NewMiningTokenTransactor creates a new write-only instance of MiningToken, bound to a specific deployed contract.
func NewMiningTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*MiningTokenTransactor, error) {
	contract, err := bindMiningToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MiningTokenTransactor{contract: contract}, nil
}

// NewMiningTokenFilterer creates a new log filterer instance of MiningToken, bound to a specific deployed contract.
func NewMiningTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*MiningTokenFilterer, error) {
	contract, err := bindMiningToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MiningTokenFilterer{contract: contract}, nil
}

// bindMiningToken binds a generic wrapper to an already deployed contract.
func bindMiningToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MiningTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MiningToken *MiningTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MiningToken.Contract.MiningTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MiningToken *MiningTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MiningToken.Contract.MiningTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MiningToken *MiningTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MiningToken.Contract.MiningTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MiningToken *MiningTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MiningToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MiningToken *MiningTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MiningToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MiningToken *MiningTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MiningToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_MiningToken *MiningTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_MiningToken *MiningTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _MiningToken.Contract.BalanceOf(&_MiningToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_MiningToken *MiningTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _MiningToken.Contract.BalanceOf(&_MiningToken.CallOpts, _owner)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() constant returns(uint256)
func (_MiningToken *MiningTokenCaller) Cap(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningToken.contract.Call(opts, out, "cap")
	return *ret0, err
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() constant returns(uint256)
func (_MiningToken *MiningTokenSession) Cap() (*big.Int, error) {
	return _MiningToken.Contract.Cap(&_MiningToken.CallOpts)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() constant returns(uint256)
func (_MiningToken *MiningTokenCallerSession) Cap() (*big.Int, error) {
	return _MiningToken.Contract.Cap(&_MiningToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(_decimals uint8)
func (_MiningToken *MiningTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _MiningToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(_decimals uint8)
func (_MiningToken *MiningTokenSession) Decimals() (uint8, error) {
	return _MiningToken.Contract.Decimals(&_MiningToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(_decimals uint8)
func (_MiningToken *MiningTokenCallerSession) Decimals() (uint8, error) {
	return _MiningToken.Contract.Decimals(&_MiningToken.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_MiningToken *MiningTokenCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MiningToken.contract.Call(opts, out, "initialized")
	return *ret0, err
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_MiningToken *MiningTokenSession) Initialized() (bool, error) {
	return _MiningToken.Contract.Initialized(&_MiningToken.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_MiningToken *MiningTokenCallerSession) Initialized() (bool, error) {
	return _MiningToken.Contract.Initialized(&_MiningToken.CallOpts)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_MiningToken *MiningTokenCaller) MintingFinished(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MiningToken.contract.Call(opts, out, "mintingFinished")
	return *ret0, err
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_MiningToken *MiningTokenSession) MintingFinished() (bool, error) {
	return _MiningToken.Contract.MintingFinished(&_MiningToken.CallOpts)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_MiningToken *MiningTokenCallerSession) MintingFinished() (bool, error) {
	return _MiningToken.Contract.MintingFinished(&_MiningToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(_name string)
func (_MiningToken *MiningTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MiningToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(_name string)
func (_MiningToken *MiningTokenSession) Name() (string, error) {
	return _MiningToken.Contract.Name(&_MiningToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(_name string)
func (_MiningToken *MiningTokenCallerSession) Name() (string, error) {
	return _MiningToken.Contract.Name(&_MiningToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MiningToken *MiningTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MiningToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MiningToken *MiningTokenSession) Owner() (common.Address, error) {
	return _MiningToken.Contract.Owner(&_MiningToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MiningToken *MiningTokenCallerSession) Owner() (common.Address, error) {
	return _MiningToken.Contract.Owner(&_MiningToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(_symbol string)
func (_MiningToken *MiningTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MiningToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(_symbol string)
func (_MiningToken *MiningTokenSession) Symbol() (string, error) {
	return _MiningToken.Contract.Symbol(&_MiningToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(_symbol string)
func (_MiningToken *MiningTokenCallerSession) Symbol() (string, error) {
	return _MiningToken.Contract.Symbol(&_MiningToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(_totalSupply uint256)
func (_MiningToken *MiningTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(_totalSupply uint256)
func (_MiningToken *MiningTokenSession) TotalSupply() (*big.Int, error) {
	return _MiningToken.Contract.TotalSupply(&_MiningToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(_totalSupply uint256)
func (_MiningToken *MiningTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _MiningToken.Contract.TotalSupply(&_MiningToken.CallOpts)
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_MiningToken *MiningTokenTransactor) FinishMinting(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MiningToken.contract.Transact(opts, "finishMinting")
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_MiningToken *MiningTokenSession) FinishMinting() (*types.Transaction, error) {
	return _MiningToken.Contract.FinishMinting(&_MiningToken.TransactOpts)
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_MiningToken *MiningTokenTransactorSession) FinishMinting() (*types.Transaction, error) {
	return _MiningToken.Contract.FinishMinting(&_MiningToken.TransactOpts)
}

// GetOwner is a paid mutator transaction binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() returns(address)
func (_MiningToken *MiningTokenTransactor) GetOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MiningToken.contract.Transact(opts, "getOwner")
}

// GetOwner is a paid mutator transaction binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() returns(address)
func (_MiningToken *MiningTokenSession) GetOwner() (*types.Transaction, error) {
	return _MiningToken.Contract.GetOwner(&_MiningToken.TransactOpts)
}

// GetOwner is a paid mutator transaction binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() returns(address)
func (_MiningToken *MiningTokenTransactorSession) GetOwner() (*types.Transaction, error) {
	return _MiningToken.Contract.GetOwner(&_MiningToken.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x41658f3c.
//
// Solidity: function initialize(_name string, _symbol string, _cap uint256, _decimals uint8) returns()
func (_MiningToken *MiningTokenTransactor) Initialize(opts *bind.TransactOpts, _name string, _symbol string, _cap *big.Int, _decimals uint8) (*types.Transaction, error) {
	return _MiningToken.contract.Transact(opts, "initialize", _name, _symbol, _cap, _decimals)
}

// Initialize is a paid mutator transaction binding the contract method 0x41658f3c.
//
// Solidity: function initialize(_name string, _symbol string, _cap uint256, _decimals uint8) returns()
func (_MiningToken *MiningTokenSession) Initialize(_name string, _symbol string, _cap *big.Int, _decimals uint8) (*types.Transaction, error) {
	return _MiningToken.Contract.Initialize(&_MiningToken.TransactOpts, _name, _symbol, _cap, _decimals)
}

// Initialize is a paid mutator transaction binding the contract method 0x41658f3c.
//
// Solidity: function initialize(_name string, _symbol string, _cap uint256, _decimals uint8) returns()
func (_MiningToken *MiningTokenTransactorSession) Initialize(_name string, _symbol string, _cap *big.Int, _decimals uint8) (*types.Transaction, error) {
	return _MiningToken.Contract.Initialize(&_MiningToken.TransactOpts, _name, _symbol, _cap, _decimals)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_MiningToken *MiningTokenTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MiningToken.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_MiningToken *MiningTokenSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MiningToken.Contract.Mint(&_MiningToken.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_MiningToken *MiningTokenTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MiningToken.Contract.Mint(&_MiningToken.TransactOpts, _to, _amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MiningToken *MiningTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MiningToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MiningToken *MiningTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _MiningToken.Contract.RenounceOwnership(&_MiningToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MiningToken *MiningTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MiningToken.Contract.RenounceOwnership(&_MiningToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xf6368f8a.
//
// Solidity: function transfer(_to address, _value uint256, _data bytes, _custom_fallback string) returns(success bool)
func (_MiningToken *MiningTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _data []byte, _custom_fallback string) (*types.Transaction, error) {
	return _MiningToken.contract.Transact(opts, "transfer", _to, _value, _data, _custom_fallback)
}

// Transfer is a paid mutator transaction binding the contract method 0xf6368f8a.
//
// Solidity: function transfer(_to address, _value uint256, _data bytes, _custom_fallback string) returns(success bool)
func (_MiningToken *MiningTokenSession) Transfer(_to common.Address, _value *big.Int, _data []byte, _custom_fallback string) (*types.Transaction, error) {
	return _MiningToken.Contract.Transfer(&_MiningToken.TransactOpts, _to, _value, _data, _custom_fallback)
}

// Transfer is a paid mutator transaction binding the contract method 0xf6368f8a.
//
// Solidity: function transfer(_to address, _value uint256, _data bytes, _custom_fallback string) returns(success bool)
func (_MiningToken *MiningTokenTransactorSession) Transfer(_to common.Address, _value *big.Int, _data []byte, _custom_fallback string) (*types.Transaction, error) {
	return _MiningToken.Contract.Transfer(&_MiningToken.TransactOpts, _to, _value, _data, _custom_fallback)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_MiningToken *MiningTokenTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _MiningToken.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_MiningToken *MiningTokenSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _MiningToken.Contract.TransferOwnership(&_MiningToken.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_MiningToken *MiningTokenTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _MiningToken.Contract.TransferOwnership(&_MiningToken.TransactOpts, _newOwner)
}

// MiningTokenMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the MiningToken contract.
type MiningTokenMintIterator struct {
	Event *MiningTokenMint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  kowala.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MiningTokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MiningTokenMint)
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
		it.Event = new(MiningTokenMint)
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
func (it *MiningTokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MiningTokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MiningTokenMint represents a Mint event raised by the MiningToken contract.
type MiningTokenMint struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: e Mint(to indexed address, amount uint256)
func (_MiningToken *MiningTokenFilterer) FilterMint(opts *bind.FilterOpts, to []common.Address) (*MiningTokenMintIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MiningToken.contract.FilterLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return &MiningTokenMintIterator{contract: _MiningToken.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: e Mint(to indexed address, amount uint256)
func (_MiningToken *MiningTokenFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *MiningTokenMint, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MiningToken.contract.WatchLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MiningTokenMint)
				if err := _MiningToken.contract.UnpackLog(event, "Mint", log); err != nil {
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

// MiningTokenMintFinishedIterator is returned from FilterMintFinished and is used to iterate over the raw logs and unpacked data for MintFinished events raised by the MiningToken contract.
type MiningTokenMintFinishedIterator struct {
	Event *MiningTokenMintFinished // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  kowala.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MiningTokenMintFinishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MiningTokenMintFinished)
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
		it.Event = new(MiningTokenMintFinished)
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
func (it *MiningTokenMintFinishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MiningTokenMintFinishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MiningTokenMintFinished represents a MintFinished event raised by the MiningToken contract.
type MiningTokenMintFinished struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMintFinished is a free log retrieval operation binding the contract event 0xae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa08.
//
// Solidity: e MintFinished()
func (_MiningToken *MiningTokenFilterer) FilterMintFinished(opts *bind.FilterOpts) (*MiningTokenMintFinishedIterator, error) {

	logs, sub, err := _MiningToken.contract.FilterLogs(opts, "MintFinished")
	if err != nil {
		return nil, err
	}
	return &MiningTokenMintFinishedIterator{contract: _MiningToken.contract, event: "MintFinished", logs: logs, sub: sub}, nil
}

// WatchMintFinished is a free log subscription operation binding the contract event 0xae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa08.
//
// Solidity: e MintFinished()
func (_MiningToken *MiningTokenFilterer) WatchMintFinished(opts *bind.WatchOpts, sink chan<- *MiningTokenMintFinished) (event.Subscription, error) {

	logs, sub, err := _MiningToken.contract.WatchLogs(opts, "MintFinished")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MiningTokenMintFinished)
				if err := _MiningToken.contract.UnpackLog(event, "MintFinished", log); err != nil {
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

// MiningTokenOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the MiningToken contract.
type MiningTokenOwnershipRenouncedIterator struct {
	Event *MiningTokenOwnershipRenounced // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  kowala.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MiningTokenOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MiningTokenOwnershipRenounced)
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
		it.Event = new(MiningTokenOwnershipRenounced)
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
func (it *MiningTokenOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MiningTokenOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MiningTokenOwnershipRenounced represents a OwnershipRenounced event raised by the MiningToken contract.
type MiningTokenOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_MiningToken *MiningTokenFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*MiningTokenOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _MiningToken.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MiningTokenOwnershipRenouncedIterator{contract: _MiningToken.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_MiningToken *MiningTokenFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *MiningTokenOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _MiningToken.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MiningTokenOwnershipRenounced)
				if err := _MiningToken.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// MiningTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MiningToken contract.
type MiningTokenOwnershipTransferredIterator struct {
	Event *MiningTokenOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  kowala.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MiningTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MiningTokenOwnershipTransferred)
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
		it.Event = new(MiningTokenOwnershipTransferred)
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
func (it *MiningTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MiningTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MiningTokenOwnershipTransferred represents a OwnershipTransferred event raised by the MiningToken contract.
type MiningTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_MiningToken *MiningTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MiningTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MiningToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MiningTokenOwnershipTransferredIterator{contract: _MiningToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_MiningToken *MiningTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MiningTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MiningToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MiningTokenOwnershipTransferred)
				if err := _MiningToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// MiningTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MiningToken contract.
type MiningTokenTransferIterator struct {
	Event *MiningTokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  kowala.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MiningTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MiningTokenTransfer)
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
		it.Event = new(MiningTokenTransfer)
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
func (it *MiningTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MiningTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MiningTokenTransfer represents a Transfer event raised by the MiningToken contract.
type MiningTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Data  []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c16.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256, data bytes)
func (_MiningToken *MiningTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MiningTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MiningToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MiningTokenTransferIterator{contract: _MiningToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c16.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256, data bytes)
func (_MiningToken *MiningTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MiningTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MiningToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MiningTokenTransfer)
				if err := _MiningToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
