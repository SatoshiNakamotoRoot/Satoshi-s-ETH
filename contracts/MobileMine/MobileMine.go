// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MobileMine

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// MobileMineABI is the input ABI used to generate the binding from.
const MobileMineABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"Mine\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"Miners\",\"outputs\":[{\"name\":\"Registry\",\"type\":\"bool\"},{\"name\":\"TotalPay\",\"type\":\"uint256\"},{\"name\":\"PayTime\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Manager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ReceiveFoundation\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"MobileMiner\",\"type\":\"address\"}],\"name\":\"MinerSetting\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ActiveUsers\",\"outputs\":[{\"name\":\"LastTime\",\"type\":\"uint256\"},{\"name\":\"ActiveNum\",\"type\":\"uint256\"},{\"name\":\"RegistryUsers\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newManager\",\"type\":\"address\"}],\"name\":\"transferManagement\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// MobileMineBin is the compiled bytecode used for deploying new contracts.
const MobileMineBin = `{
	"linkReferences": {},
	"object": "6060604052341561000f57600080fd5b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506107468061005e6000396000f300606060405260043610610083576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806342c8705b146100955780637501ce2d146100c257806378357e5314610121578063a22eef5614610176578063b62aea321461019f578063ccf7fd8d146101d8578063e4edf8521461020f575b34600560008282540192505081905550005b34156100a057600080fd5b6100a8610248565b604051808215151515815260200191505060405180910390f35b34156100cd57600080fd5b6100f9600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506104d9565b6040518084151515158152602001838152602001828152602001935050505060405180910390f35b341561012c57600080fd5b610134610510565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561018157600080fd5b610189610535565b6040518082815260200191505060405180910390f35b34156101aa57600080fd5b6101d6600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190505061053b565b005b34156101e357600080fd5b6101eb610664565b60405180848152602001838152602001828152602001935050505060405180910390f35b341561021a57600080fd5b610246600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190505061067c565b005b600060011515600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff1615151415806102f357504262015180600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206002015401115b1561030157600090506104d6565b60016002800154013073ffffffffffffffffffffffffffffffffffffffff163181151561032a57fe5b046006819055503373ffffffffffffffffffffffffffffffffffffffff166108fc6006549081150290604051600060405180830381858888f19350505050151561037357600080fd5b600654600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001016000828254019250508190555042600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206002018190555042620151806002600001540110156104bc574260026000018190555060016002600101819055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc60643073ffffffffffffffffffffffffffffffffffffffff163181151561049157fe5b049081150290604051600060405180830381858888f1935050505015156104b757600080fd5b6104d1565b60016002600101600082825401925050819055505b600190505b90565b60016020528060005260406000206000915090508060000160009054906101000a900460ff16908060010154908060020154905083565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60055481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561059657600080fd5b60011515600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff1615151415156106615760018060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160006101000a81548160ff021916908315150217905550600160028001600082825401925050819055505b50565b60028060000154908060010154908060020154905083565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156106d757600080fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505600a165627a7a72305820c48773fde838fd7c5f3a8072b44a51fb64f21c26ed8f93a74a8bbaa60f1b32840029",
	"opcodes": "PUSH1 0x60 PUSH1 0x40 MSTORE CALLVALUE ISZERO PUSH2 0xF JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST CALLER PUSH1 0x0 DUP1 PUSH2 0x100 EXP DUP2 SLOAD DUP2 PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF MUL NOT AND SWAP1 DUP4 PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND MUL OR SWAP1 SSTORE POP PUSH2 0x746 DUP1 PUSH2 0x5E PUSH1 0x0 CODECOPY PUSH1 0x0 RETURN STOP PUSH1 0x60 PUSH1 0x40 MSTORE PUSH1 0x4 CALLDATASIZE LT PUSH2 0x83 JUMPI PUSH1 0x0 CALLDATALOAD PUSH29 0x100000000000000000000000000000000000000000000000000000000 SWAP1 DIV PUSH4 0xFFFFFFFF AND DUP1 PUSH4 0x42C8705B EQ PUSH2 0x95 JUMPI DUP1 PUSH4 0x7501CE2D EQ PUSH2 0xC2 JUMPI DUP1 PUSH4 0x78357E53 EQ PUSH2 0x121 JUMPI DUP1 PUSH4 0xA22EEF56 EQ PUSH2 0x176 JUMPI DUP1 PUSH4 0xB62AEA32 EQ PUSH2 0x19F JUMPI DUP1 PUSH4 0xCCF7FD8D EQ PUSH2 0x1D8 JUMPI DUP1 PUSH4 0xE4EDF852 EQ PUSH2 0x20F JUMPI JUMPDEST CALLVALUE PUSH1 0x5 PUSH1 0x0 DUP3 DUP3 SLOAD ADD SWAP3 POP POP DUP2 SWAP1 SSTORE POP STOP JUMPDEST CALLVALUE ISZERO PUSH2 0xA0 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH2 0xA8 PUSH2 0x248 JUMP JUMPDEST PUSH1 0x40 MLOAD DUP1 DUP3 ISZERO ISZERO ISZERO ISZERO DUP2 MSTORE PUSH1 0x20 ADD SWAP2 POP POP PUSH1 0x40 MLOAD DUP1 SWAP2 SUB SWAP1 RETURN JUMPDEST CALLVALUE ISZERO PUSH2 0xCD JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH2 0xF9 PUSH1 0x4 DUP1 DUP1 CALLDATALOAD PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND SWAP1 PUSH1 0x20 ADD SWAP1 SWAP2 SWAP1 POP POP PUSH2 0x4D9 JUMP JUMPDEST PUSH1 0x40 MLOAD DUP1 DUP5 ISZERO ISZERO ISZERO ISZERO DUP2 MSTORE PUSH1 0x20 ADD DUP4 DUP2 MSTORE PUSH1 0x20 ADD DUP3 DUP2 MSTORE PUSH1 0x20 ADD SWAP4 POP POP POP POP PUSH1 0x40 MLOAD DUP1 SWAP2 SUB SWAP1 RETURN JUMPDEST CALLVALUE ISZERO PUSH2 0x12C JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH2 0x134 PUSH2 0x510 JUMP JUMPDEST PUSH1 0x40 MLOAD DUP1 DUP3 PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND DUP2 MSTORE PUSH1 0x20 ADD SWAP2 POP POP PUSH1 0x40 MLOAD DUP1 SWAP2 SUB SWAP1 RETURN JUMPDEST CALLVALUE ISZERO PUSH2 0x181 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH2 0x189 PUSH2 0x535 JUMP JUMPDEST PUSH1 0x40 MLOAD DUP1 DUP3 DUP2 MSTORE PUSH1 0x20 ADD SWAP2 POP POP PUSH1 0x40 MLOAD DUP1 SWAP2 SUB SWAP1 RETURN JUMPDEST CALLVALUE ISZERO PUSH2 0x1AA JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH2 0x1D6 PUSH1 0x4 DUP1 DUP1 CALLDATALOAD PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND SWAP1 PUSH1 0x20 ADD SWAP1 SWAP2 SWAP1 POP POP PUSH2 0x53B JUMP JUMPDEST STOP JUMPDEST CALLVALUE ISZERO PUSH2 0x1E3 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH2 0x1EB PUSH2 0x664 JUMP JUMPDEST PUSH1 0x40 MLOAD DUP1 DUP5 DUP2 MSTORE PUSH1 0x20 ADD DUP4 DUP2 MSTORE PUSH1 0x20 ADD DUP3 DUP2 MSTORE PUSH1 0x20 ADD SWAP4 POP POP POP POP PUSH1 0x40 MLOAD DUP1 SWAP2 SUB SWAP1 RETURN JUMPDEST CALLVALUE ISZERO PUSH2 0x21A JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH2 0x246 PUSH1 0x4 DUP1 DUP1 CALLDATALOAD PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND SWAP1 PUSH1 0x20 ADD SWAP1 SWAP2 SWAP1 POP POP PUSH2 0x67C JUMP JUMPDEST STOP JUMPDEST PUSH1 0x0 PUSH1 0x1 ISZERO ISZERO PUSH1 0x1 PUSH1 0x0 CALLER PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND DUP2 MSTORE PUSH1 0x20 ADD SWAP1 DUP2 MSTORE PUSH1 0x20 ADD PUSH1 0x0 KECCAK256 PUSH1 0x0 ADD PUSH1 0x0 SWAP1 SLOAD SWAP1 PUSH2 0x100 EXP SWAP1 DIV PUSH1 0xFF AND ISZERO ISZERO EQ ISZERO DUP1 PUSH2 0x2F3 JUMPI POP TIMESTAMP PUSH3 0x15180 PUSH1 0x1 PUSH1 0x0 CALLER PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND DUP2 MSTORE PUSH1 0x20 ADD SWAP1 DUP2 MSTORE PUSH1 0x20 ADD PUSH1 0x0 KECCAK256 PUSH1 0x2 ADD SLOAD ADD GT JUMPDEST ISZERO PUSH2 0x301 JUMPI PUSH1 0x0 SWAP1 POP PUSH2 0x4D6 JUMP JUMPDEST PUSH1 0x1 PUSH1 0x2 DUP1 ADD SLOAD ADD ADDRESS PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND BALANCE DUP2 ISZERO ISZERO PUSH2 0x32A JUMPI INVALID JUMPDEST DIV PUSH1 0x6 DUP2 SWAP1 SSTORE POP CALLER PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH2 0x8FC PUSH1 0x6 SLOAD SWAP1 DUP2 ISZERO MUL SWAP1 PUSH1 0x40 MLOAD PUSH1 0x0 PUSH1 0x40 MLOAD DUP1 DUP4 SUB DUP2 DUP6 DUP9 DUP9 CALL SWAP4 POP POP POP POP ISZERO ISZERO PUSH2 0x373 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH1 0x6 SLOAD PUSH1 0x1 PUSH1 0x0 CALLER PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND DUP2 MSTORE PUSH1 0x20 ADD SWAP1 DUP2 MSTORE PUSH1 0x20 ADD PUSH1 0x0 KECCAK256 PUSH1 0x1 ADD PUSH1 0x0 DUP3 DUP3 SLOAD ADD SWAP3 POP POP DUP2 SWAP1 SSTORE POP TIMESTAMP PUSH1 0x1 PUSH1 0x0 CALLER PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND DUP2 MSTORE PUSH1 0x20 ADD SWAP1 DUP2 MSTORE PUSH1 0x20 ADD PUSH1 0x0 KECCAK256 PUSH1 0x2 ADD DUP2 SWAP1 SSTORE POP TIMESTAMP PUSH3 0x15180 PUSH1 0x2 PUSH1 0x0 ADD SLOAD ADD LT ISZERO PUSH2 0x4BC JUMPI TIMESTAMP PUSH1 0x2 PUSH1 0x0 ADD DUP2 SWAP1 SSTORE POP PUSH1 0x1 PUSH1 0x2 PUSH1 0x1 ADD DUP2 SWAP1 SSTORE POP PUSH1 0x0 DUP1 SWAP1 SLOAD SWAP1 PUSH2 0x100 EXP SWAP1 DIV PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH2 0x8FC PUSH1 0x64 ADDRESS PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND BALANCE DUP2 ISZERO ISZERO PUSH2 0x491 JUMPI INVALID JUMPDEST DIV SWAP1 DUP2 ISZERO MUL SWAP1 PUSH1 0x40 MLOAD PUSH1 0x0 PUSH1 0x40 MLOAD DUP1 DUP4 SUB DUP2 DUP6 DUP9 DUP9 CALL SWAP4 POP POP POP POP ISZERO ISZERO PUSH2 0x4B7 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH2 0x4D1 JUMP JUMPDEST PUSH1 0x1 PUSH1 0x2 PUSH1 0x1 ADD PUSH1 0x0 DUP3 DUP3 SLOAD ADD SWAP3 POP POP DUP2 SWAP1 SSTORE POP JUMPDEST PUSH1 0x1 SWAP1 POP JUMPDEST SWAP1 JUMP JUMPDEST PUSH1 0x1 PUSH1 0x20 MSTORE DUP1 PUSH1 0x0 MSTORE PUSH1 0x40 PUSH1 0x0 KECCAK256 PUSH1 0x0 SWAP2 POP SWAP1 POP DUP1 PUSH1 0x0 ADD PUSH1 0x0 SWAP1 SLOAD SWAP1 PUSH2 0x100 EXP SWAP1 DIV PUSH1 0xFF AND SWAP1 DUP1 PUSH1 0x1 ADD SLOAD SWAP1 DUP1 PUSH1 0x2 ADD SLOAD SWAP1 POP DUP4 JUMP JUMPDEST PUSH1 0x0 DUP1 SWAP1 SLOAD SWAP1 PUSH2 0x100 EXP SWAP1 DIV PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND DUP2 JUMP JUMPDEST PUSH1 0x5 SLOAD DUP2 JUMP JUMPDEST PUSH1 0x0 DUP1 SWAP1 SLOAD SWAP1 PUSH2 0x100 EXP SWAP1 DIV PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND CALLER PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND EQ ISZERO ISZERO PUSH2 0x596 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH1 0x1 ISZERO ISZERO PUSH1 0x1 PUSH1 0x0 DUP4 PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND DUP2 MSTORE PUSH1 0x20 ADD SWAP1 DUP2 MSTORE PUSH1 0x20 ADD PUSH1 0x0 KECCAK256 PUSH1 0x0 ADD PUSH1 0x0 SWAP1 SLOAD SWAP1 PUSH2 0x100 EXP SWAP1 DIV PUSH1 0xFF AND ISZERO ISZERO EQ ISZERO ISZERO PUSH2 0x661 JUMPI PUSH1 0x1 DUP1 PUSH1 0x0 DUP4 PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND DUP2 MSTORE PUSH1 0x20 ADD SWAP1 DUP2 MSTORE PUSH1 0x20 ADD PUSH1 0x0 KECCAK256 PUSH1 0x0 ADD PUSH1 0x0 PUSH2 0x100 EXP DUP2 SLOAD DUP2 PUSH1 0xFF MUL NOT AND SWAP1 DUP4 ISZERO ISZERO MUL OR SWAP1 SSTORE POP PUSH1 0x1 PUSH1 0x2 DUP1 ADD PUSH1 0x0 DUP3 DUP3 SLOAD ADD SWAP3 POP POP DUP2 SWAP1 SSTORE POP JUMPDEST POP JUMP JUMPDEST PUSH1 0x2 DUP1 PUSH1 0x0 ADD SLOAD SWAP1 DUP1 PUSH1 0x1 ADD SLOAD SWAP1 DUP1 PUSH1 0x2 ADD SLOAD SWAP1 POP DUP4 JUMP JUMPDEST PUSH1 0x0 DUP1 SWAP1 SLOAD SWAP1 PUSH2 0x100 EXP SWAP1 DIV PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND CALLER PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND EQ ISZERO ISZERO PUSH2 0x6D7 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST DUP1 PUSH1 0x0 DUP1 PUSH2 0x100 EXP DUP2 SLOAD DUP2 PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF MUL NOT AND SWAP1 DUP4 PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND MUL OR SWAP1 SSTORE POP POP JUMP STOP LOG1 PUSH6 0x627A7A723058 KECCAK256 0xc4 DUP8 PUSH20 0xFDE838FD7C5F3A8072B44A51FB64F21C26ED8F93 0xa7 0x4a DUP12 0xba 0xa6 0xf 0x1b ORIGIN DUP5 STOP 0x29 ",
	"sourceMap": "28:2546:0:-;;;958:71;;;;;;;;1010:10;1002:7;;:18;;;;;;;;;;;;;;;;;;28:2546;;;;;;"
}`

// DeployMobileMine deploys a new Ethereum contract, binding an instance of MobileMine to it.
func DeployMobileMine(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MobileMine, error) {
	parsed, err := abi.JSON(strings.NewReader(MobileMineABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MobileMineBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MobileMine{MobileMineCaller: MobileMineCaller{contract: contract}, MobileMineTransactor: MobileMineTransactor{contract: contract}}, nil
}

// MobileMine is an auto generated Go binding around an Ethereum contract.
type MobileMine struct {
	MobileMineCaller     // Read-only binding to the contract
	MobileMineTransactor // Write-only binding to the contract
}

// MobileMineCaller is an auto generated read-only Go binding around an Ethereum contract.
type MobileMineCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MobileMineTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MobileMineTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MobileMineSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MobileMineSession struct {
	Contract     *MobileMine       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MobileMineCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MobileMineCallerSession struct {
	Contract *MobileMineCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MobileMineTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MobileMineTransactorSession struct {
	Contract     *MobileMineTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MobileMineRaw is an auto generated low-level Go binding around an Ethereum contract.
type MobileMineRaw struct {
	Contract *MobileMine // Generic contract binding to access the raw methods on
}

// MobileMineCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MobileMineCallerRaw struct {
	Contract *MobileMineCaller // Generic read-only contract binding to access the raw methods on
}

// MobileMineTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MobileMineTransactorRaw struct {
	Contract *MobileMineTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMobileMine creates a new instance of MobileMine, bound to a specific deployed contract.
func NewMobileMine(address common.Address, backend bind.ContractBackend) (*MobileMine, error) {
	contract, err := bindMobileMine(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MobileMine{MobileMineCaller: MobileMineCaller{contract: contract}, MobileMineTransactor: MobileMineTransactor{contract: contract}}, nil
}

// NewMobileMineCaller creates a new read-only instance of MobileMine, bound to a specific deployed contract.
func NewMobileMineCaller(address common.Address, caller bind.ContractCaller) (*MobileMineCaller, error) {
	contract, err := bindMobileMine(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &MobileMineCaller{contract: contract}, nil
}

// NewMobileMineTransactor creates a new write-only instance of MobileMine, bound to a specific deployed contract.
func NewMobileMineTransactor(address common.Address, transactor bind.ContractTransactor) (*MobileMineTransactor, error) {
	contract, err := bindMobileMine(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &MobileMineTransactor{contract: contract}, nil
}

// bindMobileMine binds a generic wrapper to an already deployed contract.
func bindMobileMine(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MobileMineABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MobileMine *MobileMineRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MobileMine.Contract.MobileMineCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MobileMine *MobileMineRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MobileMine.Contract.MobileMineTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MobileMine *MobileMineRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MobileMine.Contract.MobileMineTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MobileMine *MobileMineCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MobileMine.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MobileMine *MobileMineTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MobileMine.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MobileMine *MobileMineTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MobileMine.Contract.contract.Transact(opts, method, params...)
}

// ActiveUsers is a free data retrieval call binding the contract method 0xccf7fd8d.
//
// Solidity: function ActiveUsers() constant returns(LastTime uint256, ActiveNum uint256, RegistryUsers uint256)
func (_MobileMine *MobileMineCaller) ActiveUsers(opts *bind.CallOpts) (struct {
	LastTime      *big.Int
	ActiveNum     *big.Int
	RegistryUsers *big.Int
}, error) {
	ret := new(struct {
		LastTime      *big.Int
		ActiveNum     *big.Int
		RegistryUsers *big.Int
	})
	out := ret
	err := _MobileMine.contract.Call(opts, out, "ActiveUsers")
	return *ret, err
}

// ActiveUsers is a free data retrieval call binding the contract method 0xccf7fd8d.
//
// Solidity: function ActiveUsers() constant returns(LastTime uint256, ActiveNum uint256, RegistryUsers uint256)
func (_MobileMine *MobileMineSession) ActiveUsers() (struct {
	LastTime      *big.Int
	ActiveNum     *big.Int
	RegistryUsers *big.Int
}, error) {
	return _MobileMine.Contract.ActiveUsers(&_MobileMine.CallOpts)
}

// ActiveUsers is a free data retrieval call binding the contract method 0xccf7fd8d.
//
// Solidity: function ActiveUsers() constant returns(LastTime uint256, ActiveNum uint256, RegistryUsers uint256)
func (_MobileMine *MobileMineCallerSession) ActiveUsers() (struct {
	LastTime      *big.Int
	ActiveNum     *big.Int
	RegistryUsers *big.Int
}, error) {
	return _MobileMine.Contract.ActiveUsers(&_MobileMine.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x78357e53.
//
// Solidity: function Manager() constant returns(address)
func (_MobileMine *MobileMineCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MobileMine.contract.Call(opts, out, "Manager")
	return *ret0, err
}

// Manager is a free data retrieval call binding the contract method 0x78357e53.
//
// Solidity: function Manager() constant returns(address)
func (_MobileMine *MobileMineSession) Manager() (common.Address, error) {
	return _MobileMine.Contract.Manager(&_MobileMine.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x78357e53.
//
// Solidity: function Manager() constant returns(address)
func (_MobileMine *MobileMineCallerSession) Manager() (common.Address, error) {
	return _MobileMine.Contract.Manager(&_MobileMine.CallOpts)
}

// Miners is a free data retrieval call binding the contract method 0x7501ce2d.
//
// Solidity: function Miners( address) constant returns(Registry bool, TotalPay uint256, PayTime uint256)
func (_MobileMine *MobileMineCaller) Miners(opts *bind.CallOpts, arg0 common.Address) (struct {
	Registry bool
	TotalPay *big.Int
	PayTime  *big.Int
}, error) {
	ret := new(struct {
		Registry bool
		TotalPay *big.Int
		PayTime  *big.Int
	})
	out := ret
	err := _MobileMine.contract.Call(opts, out, "Miners", arg0)
	return *ret, err
}

// Miners is a free data retrieval call binding the contract method 0x7501ce2d.
//
// Solidity: function Miners( address) constant returns(Registry bool, TotalPay uint256, PayTime uint256)
func (_MobileMine *MobileMineSession) Miners(arg0 common.Address) (struct {
	Registry bool
	TotalPay *big.Int
	PayTime  *big.Int
}, error) {
	return _MobileMine.Contract.Miners(&_MobileMine.CallOpts, arg0)
}

// Miners is a free data retrieval call binding the contract method 0x7501ce2d.
//
// Solidity: function Miners( address) constant returns(Registry bool, TotalPay uint256, PayTime uint256)
func (_MobileMine *MobileMineCallerSession) Miners(arg0 common.Address) (struct {
	Registry bool
	TotalPay *big.Int
	PayTime  *big.Int
}, error) {
	return _MobileMine.Contract.Miners(&_MobileMine.CallOpts, arg0)
}

// ReceiveFoundation is a free data retrieval call binding the contract method 0xa22eef56.
//
// Solidity: function ReceiveFoundation() constant returns(uint256)
func (_MobileMine *MobileMineCaller) ReceiveFoundation(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MobileMine.contract.Call(opts, out, "ReceiveFoundation")
	return *ret0, err
}

// ReceiveFoundation is a free data retrieval call binding the contract method 0xa22eef56.
//
// Solidity: function ReceiveFoundation() constant returns(uint256)
func (_MobileMine *MobileMineSession) ReceiveFoundation() (*big.Int, error) {
	return _MobileMine.Contract.ReceiveFoundation(&_MobileMine.CallOpts)
}

// ReceiveFoundation is a free data retrieval call binding the contract method 0xa22eef56.
//
// Solidity: function ReceiveFoundation() constant returns(uint256)
func (_MobileMine *MobileMineCallerSession) ReceiveFoundation() (*big.Int, error) {
	return _MobileMine.Contract.ReceiveFoundation(&_MobileMine.CallOpts)
}

// Mine is a paid mutator transaction binding the contract method 0x42c8705b.
//
// Solidity: function Mine() returns(success bool)
func (_MobileMine *MobileMineTransactor) Mine(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MobileMine.contract.Transact(opts, "Mine")
}

// Mine is a paid mutator transaction binding the contract method 0x42c8705b.
//
// Solidity: function Mine() returns(success bool)
func (_MobileMine *MobileMineSession) Mine() (*types.Transaction, error) {
	return _MobileMine.Contract.Mine(&_MobileMine.TransactOpts)
}

// Mine is a paid mutator transaction binding the contract method 0x42c8705b.
//
// Solidity: function Mine() returns(success bool)
func (_MobileMine *MobileMineTransactorSession) Mine() (*types.Transaction, error) {
	return _MobileMine.Contract.Mine(&_MobileMine.TransactOpts)
}

// MinerSetting is a paid mutator transaction binding the contract method 0xb62aea32.
//
// Solidity: function MinerSetting(MobileMiner address) returns()
func (_MobileMine *MobileMineTransactor) MinerSetting(opts *bind.TransactOpts, MobileMiner common.Address) (*types.Transaction, error) {
	return _MobileMine.contract.Transact(opts, "MinerSetting", MobileMiner)
}

// MinerSetting is a paid mutator transaction binding the contract method 0xb62aea32.
//
// Solidity: function MinerSetting(MobileMiner address) returns()
func (_MobileMine *MobileMineSession) MinerSetting(MobileMiner common.Address) (*types.Transaction, error) {
	return _MobileMine.Contract.MinerSetting(&_MobileMine.TransactOpts, MobileMiner)
}

// MinerSetting is a paid mutator transaction binding the contract method 0xb62aea32.
//
// Solidity: function MinerSetting(MobileMiner address) returns()
func (_MobileMine *MobileMineTransactorSession) MinerSetting(MobileMiner common.Address) (*types.Transaction, error) {
	return _MobileMine.Contract.MinerSetting(&_MobileMine.TransactOpts, MobileMiner)
}

// TransferManagement is a paid mutator transaction binding the contract method 0xe4edf852.
//
// Solidity: function transferManagement(newManager address) returns()
func (_MobileMine *MobileMineTransactor) TransferManagement(opts *bind.TransactOpts, newManager common.Address) (*types.Transaction, error) {
	return _MobileMine.contract.Transact(opts, "transferManagement", newManager)
}

// TransferManagement is a paid mutator transaction binding the contract method 0xe4edf852.
//
// Solidity: function transferManagement(newManager address) returns()
func (_MobileMine *MobileMineSession) TransferManagement(newManager common.Address) (*types.Transaction, error) {
	return _MobileMine.Contract.TransferManagement(&_MobileMine.TransactOpts, newManager)
}

// TransferManagement is a paid mutator transaction binding the contract method 0xe4edf852.
//
// Solidity: function transferManagement(newManager address) returns()
func (_MobileMine *MobileMineTransactorSession) TransferManagement(newManager common.Address) (*types.Transaction, error) {
	return _MobileMine.Contract.TransferManagement(&_MobileMine.TransactOpts, newManager)
}
