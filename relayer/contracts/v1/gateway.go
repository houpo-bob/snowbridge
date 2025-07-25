// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// InboundMessage is an auto generated low-level Go binding around an user-defined struct.
type InboundMessage struct {
	ChannelID      [32]byte
	Nonce          uint64
	Command        uint8
	Params         []byte
	MaxDispatchGas uint64
	MaxFeePerGas   *big.Int
	Reward         *big.Int
	Id             [32]byte
}

// MultiAddress is an auto generated low-level Go binding around an user-defined struct.
type MultiAddress struct {
	Kind uint8
	Data []byte
}

// VerificationDigestItem is an auto generated low-level Go binding around an user-defined struct.
type VerificationDigestItem struct {
	Kind              *big.Int
	ConsensusEngineID [4]byte
	Data              []byte
}

// VerificationHeadProof is an auto generated low-level Go binding around an user-defined struct.
type VerificationHeadProof struct {
	Pos   *big.Int
	Width *big.Int
	Proof [][32]byte
}

// VerificationMMRLeafPartial is an auto generated low-level Go binding around an user-defined struct.
type VerificationMMRLeafPartial struct {
	Version              uint8
	ParentNumber         uint32
	ParentHash           [32]byte
	NextAuthoritySetID   uint64
	NextAuthoritySetLen  uint32
	NextAuthoritySetRoot [32]byte
}

// VerificationParachainHeader is an auto generated low-level Go binding around an user-defined struct.
type VerificationParachainHeader struct {
	ParentHash     [32]byte
	Number         *big.Int
	StateRoot      [32]byte
	ExtrinsicsRoot [32]byte
	DigestItems    []VerificationDigestItem
}

// VerificationProof is an auto generated low-level Go binding around an user-defined struct.
type VerificationProof struct {
	Header         VerificationParachainHeader
	HeadProof      VerificationHeadProof
	LeafPartial    VerificationMMRLeafPartial
	LeafProof      [][32]byte
	LeafProofOrder *big.Int
}

// GatewayMetaData contains all meta data concerning the Gateway contract.
var GatewayMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"agentOf\",\"inputs\":[{\"name\":\"agentID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"channelNoncesOf\",\"inputs\":[{\"name\":\"channelID\",\"type\":\"bytes32\",\"internalType\":\"ChannelID\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"channelOperatingModeOf\",\"inputs\":[{\"name\":\"channelID\",\"type\":\"bytes32\",\"internalType\":\"ChannelID\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumOperatingMode\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositEther\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"isTokenRegistered\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatingMode\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumOperatingMode\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pricingParameters\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"UD60x18\"},{\"name\":\"\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"queryForeignTokenID\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quoteRegisterTokenFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quoteSendTokenFee\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destinationChain\",\"type\":\"uint32\",\"internalType\":\"ParaID\"},{\"name\":\"destinationFee\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"sendToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destinationChain\",\"type\":\"uint32\",\"internalType\":\"ParaID\"},{\"name\":\"destinationAddress\",\"type\":\"tuple\",\"internalType\":\"structMultiAddress\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"enumKind\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destinationFee\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"amount\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"submitV1\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structInboundMessage\",\"components\":[{\"name\":\"channelID\",\"type\":\"bytes32\",\"internalType\":\"ChannelID\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"command\",\"type\":\"uint8\",\"internalType\":\"enumCommand\"},{\"name\":\"params\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"maxDispatchGas\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reward\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"leafProof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"headerProof\",\"type\":\"tuple\",\"internalType\":\"structVerification.Proof\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structVerification.ParachainHeader\",\"components\":[{\"name\":\"parentHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"number\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"stateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"extrinsicsRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digestItems\",\"type\":\"tuple[]\",\"internalType\":\"structVerification.DigestItem[]\",\"components\":[{\"name\":\"kind\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"consensusEngineID\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"headProof\",\"type\":\"tuple\",\"internalType\":\"structVerification.HeadProof\",\"components\":[{\"name\":\"pos\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"width\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]},{\"name\":\"leafPartial\",\"type\":\"tuple\",\"internalType\":\"structVerification.MMRLeafPartial\",\"components\":[{\"name\":\"version\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"parentNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"parentHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nextAuthoritySetID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nextAuthoritySetLen\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nextAuthoritySetRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"leafProof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"leafProofOrder\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"tokenAddressOf\",\"inputs\":[{\"name\":\"tokenID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AgentFundsWithdrawn\",\"inputs\":[{\"name\":\"agentID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundMessageDispatched\",\"inputs\":[{\"name\":\"channelID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"ChannelID\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"messageID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"success\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundMessageAccepted\",\"inputs\":[{\"name\":\"channelID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"ChannelID\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"messageID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PricingParametersChanged\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenSent\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"destinationChain\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"ParaID\"},{\"name\":\"destinationAddress\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structMultiAddress\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"enumKind\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"amount\",\"type\":\"uint128\",\"indexed\":false,\"internalType\":\"uint128\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeesChanged\",\"inputs\":[],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ChannelDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidChannelUpdate\",\"inputs\":[]}]",
}

// GatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayMetaData.ABI instead.
var GatewayABI = GatewayMetaData.ABI

// Gateway is an auto generated Go binding around an Ethereum contract.
type Gateway struct {
	GatewayCaller     // Read-only binding to the contract
	GatewayTransactor // Write-only binding to the contract
	GatewayFilterer   // Log filterer for contract events
}

// GatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewaySession struct {
	Contract     *Gateway          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayCallerSession struct {
	Contract *GatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// GatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayTransactorSession struct {
	Contract     *GatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// GatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayRaw struct {
	Contract *Gateway // Generic contract binding to access the raw methods on
}

// GatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayCallerRaw struct {
	Contract *GatewayCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayTransactorRaw struct {
	Contract *GatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGateway creates a new instance of Gateway, bound to a specific deployed contract.
func NewGateway(address common.Address, backend bind.ContractBackend) (*Gateway, error) {
	contract, err := bindGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Gateway{GatewayCaller: GatewayCaller{contract: contract}, GatewayTransactor: GatewayTransactor{contract: contract}, GatewayFilterer: GatewayFilterer{contract: contract}}, nil
}

// NewGatewayCaller creates a new read-only instance of Gateway, bound to a specific deployed contract.
func NewGatewayCaller(address common.Address, caller bind.ContractCaller) (*GatewayCaller, error) {
	contract, err := bindGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayCaller{contract: contract}, nil
}

// NewGatewayTransactor creates a new write-only instance of Gateway, bound to a specific deployed contract.
func NewGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayTransactor, error) {
	contract, err := bindGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayTransactor{contract: contract}, nil
}

// NewGatewayFilterer creates a new log filterer instance of Gateway, bound to a specific deployed contract.
func NewGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayFilterer, error) {
	contract, err := bindGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayFilterer{contract: contract}, nil
}

// bindGateway binds a generic wrapper to an already deployed contract.
func bindGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gateway *GatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gateway.Contract.GatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gateway *GatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gateway.Contract.GatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gateway *GatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gateway.Contract.GatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gateway *GatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gateway *GatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gateway *GatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gateway.Contract.contract.Transact(opts, method, params...)
}

// AgentOf is a free data retrieval call binding the contract method 0x5e6dae26.
//
// Solidity: function agentOf(bytes32 agentID) view returns(address)
func (_Gateway *GatewayCaller) AgentOf(opts *bind.CallOpts, agentID [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "agentOf", agentID)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AgentOf is a free data retrieval call binding the contract method 0x5e6dae26.
//
// Solidity: function agentOf(bytes32 agentID) view returns(address)
func (_Gateway *GatewaySession) AgentOf(agentID [32]byte) (common.Address, error) {
	return _Gateway.Contract.AgentOf(&_Gateway.CallOpts, agentID)
}

// AgentOf is a free data retrieval call binding the contract method 0x5e6dae26.
//
// Solidity: function agentOf(bytes32 agentID) view returns(address)
func (_Gateway *GatewayCallerSession) AgentOf(agentID [32]byte) (common.Address, error) {
	return _Gateway.Contract.AgentOf(&_Gateway.CallOpts, agentID)
}

// ChannelNoncesOf is a free data retrieval call binding the contract method 0x2a6c3229.
//
// Solidity: function channelNoncesOf(bytes32 channelID) view returns(uint64, uint64)
func (_Gateway *GatewayCaller) ChannelNoncesOf(opts *bind.CallOpts, channelID [32]byte) (uint64, uint64, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "channelNoncesOf", channelID)

	if err != nil {
		return *new(uint64), *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	out1 := *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return out0, out1, err

}

// ChannelNoncesOf is a free data retrieval call binding the contract method 0x2a6c3229.
//
// Solidity: function channelNoncesOf(bytes32 channelID) view returns(uint64, uint64)
func (_Gateway *GatewaySession) ChannelNoncesOf(channelID [32]byte) (uint64, uint64, error) {
	return _Gateway.Contract.ChannelNoncesOf(&_Gateway.CallOpts, channelID)
}

// ChannelNoncesOf is a free data retrieval call binding the contract method 0x2a6c3229.
//
// Solidity: function channelNoncesOf(bytes32 channelID) view returns(uint64, uint64)
func (_Gateway *GatewayCallerSession) ChannelNoncesOf(channelID [32]byte) (uint64, uint64, error) {
	return _Gateway.Contract.ChannelNoncesOf(&_Gateway.CallOpts, channelID)
}

// ChannelOperatingModeOf is a free data retrieval call binding the contract method 0x0705f465.
//
// Solidity: function channelOperatingModeOf(bytes32 channelID) view returns(uint8)
func (_Gateway *GatewayCaller) ChannelOperatingModeOf(opts *bind.CallOpts, channelID [32]byte) (uint8, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "channelOperatingModeOf", channelID)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ChannelOperatingModeOf is a free data retrieval call binding the contract method 0x0705f465.
//
// Solidity: function channelOperatingModeOf(bytes32 channelID) view returns(uint8)
func (_Gateway *GatewaySession) ChannelOperatingModeOf(channelID [32]byte) (uint8, error) {
	return _Gateway.Contract.ChannelOperatingModeOf(&_Gateway.CallOpts, channelID)
}

// ChannelOperatingModeOf is a free data retrieval call binding the contract method 0x0705f465.
//
// Solidity: function channelOperatingModeOf(bytes32 channelID) view returns(uint8)
func (_Gateway *GatewayCallerSession) ChannelOperatingModeOf(channelID [32]byte) (uint8, error) {
	return _Gateway.Contract.ChannelOperatingModeOf(&_Gateway.CallOpts, channelID)
}

// IsTokenRegistered is a free data retrieval call binding the contract method 0x26aa101f.
//
// Solidity: function isTokenRegistered(address token) view returns(bool)
func (_Gateway *GatewayCaller) IsTokenRegistered(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "isTokenRegistered", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenRegistered is a free data retrieval call binding the contract method 0x26aa101f.
//
// Solidity: function isTokenRegistered(address token) view returns(bool)
func (_Gateway *GatewaySession) IsTokenRegistered(token common.Address) (bool, error) {
	return _Gateway.Contract.IsTokenRegistered(&_Gateway.CallOpts, token)
}

// IsTokenRegistered is a free data retrieval call binding the contract method 0x26aa101f.
//
// Solidity: function isTokenRegistered(address token) view returns(bool)
func (_Gateway *GatewayCallerSession) IsTokenRegistered(token common.Address) (bool, error) {
	return _Gateway.Contract.IsTokenRegistered(&_Gateway.CallOpts, token)
}

// OperatingMode is a free data retrieval call binding the contract method 0x38004f69.
//
// Solidity: function operatingMode() view returns(uint8)
func (_Gateway *GatewayCaller) OperatingMode(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "operatingMode")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OperatingMode is a free data retrieval call binding the contract method 0x38004f69.
//
// Solidity: function operatingMode() view returns(uint8)
func (_Gateway *GatewaySession) OperatingMode() (uint8, error) {
	return _Gateway.Contract.OperatingMode(&_Gateway.CallOpts)
}

// OperatingMode is a free data retrieval call binding the contract method 0x38004f69.
//
// Solidity: function operatingMode() view returns(uint8)
func (_Gateway *GatewayCallerSession) OperatingMode() (uint8, error) {
	return _Gateway.Contract.OperatingMode(&_Gateway.CallOpts)
}

// PricingParameters is a free data retrieval call binding the contract method 0x0b617646.
//
// Solidity: function pricingParameters() view returns(uint256, uint128)
func (_Gateway *GatewayCaller) PricingParameters(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "pricingParameters")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// PricingParameters is a free data retrieval call binding the contract method 0x0b617646.
//
// Solidity: function pricingParameters() view returns(uint256, uint128)
func (_Gateway *GatewaySession) PricingParameters() (*big.Int, *big.Int, error) {
	return _Gateway.Contract.PricingParameters(&_Gateway.CallOpts)
}

// PricingParameters is a free data retrieval call binding the contract method 0x0b617646.
//
// Solidity: function pricingParameters() view returns(uint256, uint128)
func (_Gateway *GatewayCallerSession) PricingParameters() (*big.Int, *big.Int, error) {
	return _Gateway.Contract.PricingParameters(&_Gateway.CallOpts)
}

// QueryForeignTokenID is a free data retrieval call binding the contract method 0xbe8d42c0.
//
// Solidity: function queryForeignTokenID(address token) view returns(bytes32)
func (_Gateway *GatewayCaller) QueryForeignTokenID(opts *bind.CallOpts, token common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "queryForeignTokenID", token)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// QueryForeignTokenID is a free data retrieval call binding the contract method 0xbe8d42c0.
//
// Solidity: function queryForeignTokenID(address token) view returns(bytes32)
func (_Gateway *GatewaySession) QueryForeignTokenID(token common.Address) ([32]byte, error) {
	return _Gateway.Contract.QueryForeignTokenID(&_Gateway.CallOpts, token)
}

// QueryForeignTokenID is a free data retrieval call binding the contract method 0xbe8d42c0.
//
// Solidity: function queryForeignTokenID(address token) view returns(bytes32)
func (_Gateway *GatewayCallerSession) QueryForeignTokenID(token common.Address) ([32]byte, error) {
	return _Gateway.Contract.QueryForeignTokenID(&_Gateway.CallOpts, token)
}

// QuoteRegisterTokenFee is a free data retrieval call binding the contract method 0x805ce31d.
//
// Solidity: function quoteRegisterTokenFee() view returns(uint256)
func (_Gateway *GatewayCaller) QuoteRegisterTokenFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "quoteRegisterTokenFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuoteRegisterTokenFee is a free data retrieval call binding the contract method 0x805ce31d.
//
// Solidity: function quoteRegisterTokenFee() view returns(uint256)
func (_Gateway *GatewaySession) QuoteRegisterTokenFee() (*big.Int, error) {
	return _Gateway.Contract.QuoteRegisterTokenFee(&_Gateway.CallOpts)
}

// QuoteRegisterTokenFee is a free data retrieval call binding the contract method 0x805ce31d.
//
// Solidity: function quoteRegisterTokenFee() view returns(uint256)
func (_Gateway *GatewayCallerSession) QuoteRegisterTokenFee() (*big.Int, error) {
	return _Gateway.Contract.QuoteRegisterTokenFee(&_Gateway.CallOpts)
}

// QuoteSendTokenFee is a free data retrieval call binding the contract method 0x928bc49d.
//
// Solidity: function quoteSendTokenFee(address token, uint32 destinationChain, uint128 destinationFee) view returns(uint256)
func (_Gateway *GatewayCaller) QuoteSendTokenFee(opts *bind.CallOpts, token common.Address, destinationChain uint32, destinationFee *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "quoteSendTokenFee", token, destinationChain, destinationFee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuoteSendTokenFee is a free data retrieval call binding the contract method 0x928bc49d.
//
// Solidity: function quoteSendTokenFee(address token, uint32 destinationChain, uint128 destinationFee) view returns(uint256)
func (_Gateway *GatewaySession) QuoteSendTokenFee(token common.Address, destinationChain uint32, destinationFee *big.Int) (*big.Int, error) {
	return _Gateway.Contract.QuoteSendTokenFee(&_Gateway.CallOpts, token, destinationChain, destinationFee)
}

// QuoteSendTokenFee is a free data retrieval call binding the contract method 0x928bc49d.
//
// Solidity: function quoteSendTokenFee(address token, uint32 destinationChain, uint128 destinationFee) view returns(uint256)
func (_Gateway *GatewayCallerSession) QuoteSendTokenFee(token common.Address, destinationChain uint32, destinationFee *big.Int) (*big.Int, error) {
	return _Gateway.Contract.QuoteSendTokenFee(&_Gateway.CallOpts, token, destinationChain, destinationFee)
}

// TokenAddressOf is a free data retrieval call binding the contract method 0xfe61cc49.
//
// Solidity: function tokenAddressOf(bytes32 tokenID) view returns(address)
func (_Gateway *GatewayCaller) TokenAddressOf(opts *bind.CallOpts, tokenID [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "tokenAddressOf", tokenID)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenAddressOf is a free data retrieval call binding the contract method 0xfe61cc49.
//
// Solidity: function tokenAddressOf(bytes32 tokenID) view returns(address)
func (_Gateway *GatewaySession) TokenAddressOf(tokenID [32]byte) (common.Address, error) {
	return _Gateway.Contract.TokenAddressOf(&_Gateway.CallOpts, tokenID)
}

// TokenAddressOf is a free data retrieval call binding the contract method 0xfe61cc49.
//
// Solidity: function tokenAddressOf(bytes32 tokenID) view returns(address)
func (_Gateway *GatewayCallerSession) TokenAddressOf(tokenID [32]byte) (common.Address, error) {
	return _Gateway.Contract.TokenAddressOf(&_Gateway.CallOpts, tokenID)
}

// DepositEther is a paid mutator transaction binding the contract method 0x98ea5fca.
//
// Solidity: function depositEther() payable returns()
func (_Gateway *GatewayTransactor) DepositEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "depositEther")
}

// DepositEther is a paid mutator transaction binding the contract method 0x98ea5fca.
//
// Solidity: function depositEther() payable returns()
func (_Gateway *GatewaySession) DepositEther() (*types.Transaction, error) {
	return _Gateway.Contract.DepositEther(&_Gateway.TransactOpts)
}

// DepositEther is a paid mutator transaction binding the contract method 0x98ea5fca.
//
// Solidity: function depositEther() payable returns()
func (_Gateway *GatewayTransactorSession) DepositEther() (*types.Transaction, error) {
	return _Gateway.Contract.DepositEther(&_Gateway.TransactOpts)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x09824a80.
//
// Solidity: function registerToken(address token) payable returns()
func (_Gateway *GatewayTransactor) RegisterToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "registerToken", token)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x09824a80.
//
// Solidity: function registerToken(address token) payable returns()
func (_Gateway *GatewaySession) RegisterToken(token common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.RegisterToken(&_Gateway.TransactOpts, token)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x09824a80.
//
// Solidity: function registerToken(address token) payable returns()
func (_Gateway *GatewayTransactorSession) RegisterToken(token common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.RegisterToken(&_Gateway.TransactOpts, token)
}

// SendToken is a paid mutator transaction binding the contract method 0x52054834.
//
// Solidity: function sendToken(address token, uint32 destinationChain, (uint8,bytes) destinationAddress, uint128 destinationFee, uint128 amount) payable returns()
func (_Gateway *GatewayTransactor) SendToken(opts *bind.TransactOpts, token common.Address, destinationChain uint32, destinationAddress MultiAddress, destinationFee *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "sendToken", token, destinationChain, destinationAddress, destinationFee, amount)
}

// SendToken is a paid mutator transaction binding the contract method 0x52054834.
//
// Solidity: function sendToken(address token, uint32 destinationChain, (uint8,bytes) destinationAddress, uint128 destinationFee, uint128 amount) payable returns()
func (_Gateway *GatewaySession) SendToken(token common.Address, destinationChain uint32, destinationAddress MultiAddress, destinationFee *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.SendToken(&_Gateway.TransactOpts, token, destinationChain, destinationAddress, destinationFee, amount)
}

// SendToken is a paid mutator transaction binding the contract method 0x52054834.
//
// Solidity: function sendToken(address token, uint32 destinationChain, (uint8,bytes) destinationAddress, uint128 destinationFee, uint128 amount) payable returns()
func (_Gateway *GatewayTransactorSession) SendToken(token common.Address, destinationChain uint32, destinationAddress MultiAddress, destinationFee *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.SendToken(&_Gateway.TransactOpts, token, destinationChain, destinationAddress, destinationFee, amount)
}

// SubmitV1 is a paid mutator transaction binding the contract method 0xdf4ed829.
//
// Solidity: function submitV1((bytes32,uint64,uint8,bytes,uint64,uint256,uint256,bytes32) message, bytes32[] leafProof, ((bytes32,uint256,bytes32,bytes32,(uint256,bytes4,bytes)[]),(uint256,uint256,bytes32[]),(uint8,uint32,bytes32,uint64,uint32,bytes32),bytes32[],uint256) headerProof) returns()
func (_Gateway *GatewayTransactor) SubmitV1(opts *bind.TransactOpts, message InboundMessage, leafProof [][32]byte, headerProof VerificationProof) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "submitV1", message, leafProof, headerProof)
}

// SubmitV1 is a paid mutator transaction binding the contract method 0xdf4ed829.
//
// Solidity: function submitV1((bytes32,uint64,uint8,bytes,uint64,uint256,uint256,bytes32) message, bytes32[] leafProof, ((bytes32,uint256,bytes32,bytes32,(uint256,bytes4,bytes)[]),(uint256,uint256,bytes32[]),(uint8,uint32,bytes32,uint64,uint32,bytes32),bytes32[],uint256) headerProof) returns()
func (_Gateway *GatewaySession) SubmitV1(message InboundMessage, leafProof [][32]byte, headerProof VerificationProof) (*types.Transaction, error) {
	return _Gateway.Contract.SubmitV1(&_Gateway.TransactOpts, message, leafProof, headerProof)
}

// SubmitV1 is a paid mutator transaction binding the contract method 0xdf4ed829.
//
// Solidity: function submitV1((bytes32,uint64,uint8,bytes,uint64,uint256,uint256,bytes32) message, bytes32[] leafProof, ((bytes32,uint256,bytes32,bytes32,(uint256,bytes4,bytes)[]),(uint256,uint256,bytes32[]),(uint8,uint32,bytes32,uint64,uint32,bytes32),bytes32[],uint256) headerProof) returns()
func (_Gateway *GatewayTransactorSession) SubmitV1(message InboundMessage, leafProof [][32]byte, headerProof VerificationProof) (*types.Transaction, error) {
	return _Gateway.Contract.SubmitV1(&_Gateway.TransactOpts, message, leafProof, headerProof)
}

// GatewayAgentFundsWithdrawnIterator is returned from FilterAgentFundsWithdrawn and is used to iterate over the raw logs and unpacked data for AgentFundsWithdrawn events raised by the Gateway contract.
type GatewayAgentFundsWithdrawnIterator struct {
	Event *GatewayAgentFundsWithdrawn // Event containing the contract specifics and raw log

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
func (it *GatewayAgentFundsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayAgentFundsWithdrawn)
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
		it.Event = new(GatewayAgentFundsWithdrawn)
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
func (it *GatewayAgentFundsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayAgentFundsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayAgentFundsWithdrawn represents a AgentFundsWithdrawn event raised by the Gateway contract.
type GatewayAgentFundsWithdrawn struct {
	AgentID   [32]byte
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAgentFundsWithdrawn is a free log retrieval operation binding the contract event 0xf953871855f78d5ccdd6268f2d9d69fc67f26542a35d2bba1c615521aed57054.
//
// Solidity: event AgentFundsWithdrawn(bytes32 indexed agentID, address indexed recipient, uint256 amount)
func (_Gateway *GatewayFilterer) FilterAgentFundsWithdrawn(opts *bind.FilterOpts, agentID [][32]byte, recipient []common.Address) (*GatewayAgentFundsWithdrawnIterator, error) {

	var agentIDRule []interface{}
	for _, agentIDItem := range agentID {
		agentIDRule = append(agentIDRule, agentIDItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "AgentFundsWithdrawn", agentIDRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &GatewayAgentFundsWithdrawnIterator{contract: _Gateway.contract, event: "AgentFundsWithdrawn", logs: logs, sub: sub}, nil
}

// WatchAgentFundsWithdrawn is a free log subscription operation binding the contract event 0xf953871855f78d5ccdd6268f2d9d69fc67f26542a35d2bba1c615521aed57054.
//
// Solidity: event AgentFundsWithdrawn(bytes32 indexed agentID, address indexed recipient, uint256 amount)
func (_Gateway *GatewayFilterer) WatchAgentFundsWithdrawn(opts *bind.WatchOpts, sink chan<- *GatewayAgentFundsWithdrawn, agentID [][32]byte, recipient []common.Address) (event.Subscription, error) {

	var agentIDRule []interface{}
	for _, agentIDItem := range agentID {
		agentIDRule = append(agentIDRule, agentIDItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "AgentFundsWithdrawn", agentIDRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayAgentFundsWithdrawn)
				if err := _Gateway.contract.UnpackLog(event, "AgentFundsWithdrawn", log); err != nil {
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

// ParseAgentFundsWithdrawn is a log parse operation binding the contract event 0xf953871855f78d5ccdd6268f2d9d69fc67f26542a35d2bba1c615521aed57054.
//
// Solidity: event AgentFundsWithdrawn(bytes32 indexed agentID, address indexed recipient, uint256 amount)
func (_Gateway *GatewayFilterer) ParseAgentFundsWithdrawn(log types.Log) (*GatewayAgentFundsWithdrawn, error) {
	event := new(GatewayAgentFundsWithdrawn)
	if err := _Gateway.contract.UnpackLog(event, "AgentFundsWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Gateway contract.
type GatewayDepositedIterator struct {
	Event *GatewayDeposited // Event containing the contract specifics and raw log

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
func (it *GatewayDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayDeposited)
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
		it.Event = new(GatewayDeposited)
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
func (it *GatewayDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayDeposited represents a Deposited event raised by the Gateway contract.
type GatewayDeposited struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address sender, uint256 amount)
func (_Gateway *GatewayFilterer) FilterDeposited(opts *bind.FilterOpts) (*GatewayDepositedIterator, error) {

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return &GatewayDepositedIterator{contract: _Gateway.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address sender, uint256 amount)
func (_Gateway *GatewayFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *GatewayDeposited) (event.Subscription, error) {

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayDeposited)
				if err := _Gateway.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address sender, uint256 amount)
func (_Gateway *GatewayFilterer) ParseDeposited(log types.Log) (*GatewayDeposited, error) {
	event := new(GatewayDeposited)
	if err := _Gateway.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayInboundMessageDispatchedIterator is returned from FilterInboundMessageDispatched and is used to iterate over the raw logs and unpacked data for InboundMessageDispatched events raised by the Gateway contract.
type GatewayInboundMessageDispatchedIterator struct {
	Event *GatewayInboundMessageDispatched // Event containing the contract specifics and raw log

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
func (it *GatewayInboundMessageDispatchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayInboundMessageDispatched)
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
		it.Event = new(GatewayInboundMessageDispatched)
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
func (it *GatewayInboundMessageDispatchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayInboundMessageDispatchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayInboundMessageDispatched represents a InboundMessageDispatched event raised by the Gateway contract.
type GatewayInboundMessageDispatched struct {
	ChannelID [32]byte
	Nonce     uint64
	MessageID [32]byte
	Success   bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterInboundMessageDispatched is a free log retrieval operation binding the contract event 0x617fdb0cb78f01551a192a3673208ec5eb09f20a90acf673c63a0dcb11745a7a.
//
// Solidity: event InboundMessageDispatched(bytes32 indexed channelID, uint64 nonce, bytes32 indexed messageID, bool success)
func (_Gateway *GatewayFilterer) FilterInboundMessageDispatched(opts *bind.FilterOpts, channelID [][32]byte, messageID [][32]byte) (*GatewayInboundMessageDispatchedIterator, error) {

	var channelIDRule []interface{}
	for _, channelIDItem := range channelID {
		channelIDRule = append(channelIDRule, channelIDItem)
	}

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "InboundMessageDispatched", channelIDRule, messageIDRule)
	if err != nil {
		return nil, err
	}
	return &GatewayInboundMessageDispatchedIterator{contract: _Gateway.contract, event: "InboundMessageDispatched", logs: logs, sub: sub}, nil
}

// WatchInboundMessageDispatched is a free log subscription operation binding the contract event 0x617fdb0cb78f01551a192a3673208ec5eb09f20a90acf673c63a0dcb11745a7a.
//
// Solidity: event InboundMessageDispatched(bytes32 indexed channelID, uint64 nonce, bytes32 indexed messageID, bool success)
func (_Gateway *GatewayFilterer) WatchInboundMessageDispatched(opts *bind.WatchOpts, sink chan<- *GatewayInboundMessageDispatched, channelID [][32]byte, messageID [][32]byte) (event.Subscription, error) {

	var channelIDRule []interface{}
	for _, channelIDItem := range channelID {
		channelIDRule = append(channelIDRule, channelIDItem)
	}

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "InboundMessageDispatched", channelIDRule, messageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayInboundMessageDispatched)
				if err := _Gateway.contract.UnpackLog(event, "InboundMessageDispatched", log); err != nil {
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

// ParseInboundMessageDispatched is a log parse operation binding the contract event 0x617fdb0cb78f01551a192a3673208ec5eb09f20a90acf673c63a0dcb11745a7a.
//
// Solidity: event InboundMessageDispatched(bytes32 indexed channelID, uint64 nonce, bytes32 indexed messageID, bool success)
func (_Gateway *GatewayFilterer) ParseInboundMessageDispatched(log types.Log) (*GatewayInboundMessageDispatched, error) {
	event := new(GatewayInboundMessageDispatched)
	if err := _Gateway.contract.UnpackLog(event, "InboundMessageDispatched", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayOutboundMessageAcceptedIterator is returned from FilterOutboundMessageAccepted and is used to iterate over the raw logs and unpacked data for OutboundMessageAccepted events raised by the Gateway contract.
type GatewayOutboundMessageAcceptedIterator struct {
	Event *GatewayOutboundMessageAccepted // Event containing the contract specifics and raw log

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
func (it *GatewayOutboundMessageAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayOutboundMessageAccepted)
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
		it.Event = new(GatewayOutboundMessageAccepted)
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
func (it *GatewayOutboundMessageAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayOutboundMessageAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayOutboundMessageAccepted represents a OutboundMessageAccepted event raised by the Gateway contract.
type GatewayOutboundMessageAccepted struct {
	ChannelID [32]byte
	Nonce     uint64
	MessageID [32]byte
	Payload   []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOutboundMessageAccepted is a free log retrieval operation binding the contract event 0x7153f9357c8ea496bba60bf82e67143e27b64462b49041f8e689e1b05728f84f.
//
// Solidity: event OutboundMessageAccepted(bytes32 indexed channelID, uint64 nonce, bytes32 indexed messageID, bytes payload)
func (_Gateway *GatewayFilterer) FilterOutboundMessageAccepted(opts *bind.FilterOpts, channelID [][32]byte, messageID [][32]byte) (*GatewayOutboundMessageAcceptedIterator, error) {

	var channelIDRule []interface{}
	for _, channelIDItem := range channelID {
		channelIDRule = append(channelIDRule, channelIDItem)
	}

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "OutboundMessageAccepted", channelIDRule, messageIDRule)
	if err != nil {
		return nil, err
	}
	return &GatewayOutboundMessageAcceptedIterator{contract: _Gateway.contract, event: "OutboundMessageAccepted", logs: logs, sub: sub}, nil
}

// WatchOutboundMessageAccepted is a free log subscription operation binding the contract event 0x7153f9357c8ea496bba60bf82e67143e27b64462b49041f8e689e1b05728f84f.
//
// Solidity: event OutboundMessageAccepted(bytes32 indexed channelID, uint64 nonce, bytes32 indexed messageID, bytes payload)
func (_Gateway *GatewayFilterer) WatchOutboundMessageAccepted(opts *bind.WatchOpts, sink chan<- *GatewayOutboundMessageAccepted, channelID [][32]byte, messageID [][32]byte) (event.Subscription, error) {

	var channelIDRule []interface{}
	for _, channelIDItem := range channelID {
		channelIDRule = append(channelIDRule, channelIDItem)
	}

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "OutboundMessageAccepted", channelIDRule, messageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayOutboundMessageAccepted)
				if err := _Gateway.contract.UnpackLog(event, "OutboundMessageAccepted", log); err != nil {
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

// ParseOutboundMessageAccepted is a log parse operation binding the contract event 0x7153f9357c8ea496bba60bf82e67143e27b64462b49041f8e689e1b05728f84f.
//
// Solidity: event OutboundMessageAccepted(bytes32 indexed channelID, uint64 nonce, bytes32 indexed messageID, bytes payload)
func (_Gateway *GatewayFilterer) ParseOutboundMessageAccepted(log types.Log) (*GatewayOutboundMessageAccepted, error) {
	event := new(GatewayOutboundMessageAccepted)
	if err := _Gateway.contract.UnpackLog(event, "OutboundMessageAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayPricingParametersChangedIterator is returned from FilterPricingParametersChanged and is used to iterate over the raw logs and unpacked data for PricingParametersChanged events raised by the Gateway contract.
type GatewayPricingParametersChangedIterator struct {
	Event *GatewayPricingParametersChanged // Event containing the contract specifics and raw log

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
func (it *GatewayPricingParametersChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayPricingParametersChanged)
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
		it.Event = new(GatewayPricingParametersChanged)
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
func (it *GatewayPricingParametersChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayPricingParametersChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayPricingParametersChanged represents a PricingParametersChanged event raised by the Gateway contract.
type GatewayPricingParametersChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPricingParametersChanged is a free log retrieval operation binding the contract event 0x5e3c25378b5946068b94aa2ea10c4c1e215cc975f994322b159ddc9237a973d4.
//
// Solidity: event PricingParametersChanged()
func (_Gateway *GatewayFilterer) FilterPricingParametersChanged(opts *bind.FilterOpts) (*GatewayPricingParametersChangedIterator, error) {

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "PricingParametersChanged")
	if err != nil {
		return nil, err
	}
	return &GatewayPricingParametersChangedIterator{contract: _Gateway.contract, event: "PricingParametersChanged", logs: logs, sub: sub}, nil
}

// WatchPricingParametersChanged is a free log subscription operation binding the contract event 0x5e3c25378b5946068b94aa2ea10c4c1e215cc975f994322b159ddc9237a973d4.
//
// Solidity: event PricingParametersChanged()
func (_Gateway *GatewayFilterer) WatchPricingParametersChanged(opts *bind.WatchOpts, sink chan<- *GatewayPricingParametersChanged) (event.Subscription, error) {

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "PricingParametersChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayPricingParametersChanged)
				if err := _Gateway.contract.UnpackLog(event, "PricingParametersChanged", log); err != nil {
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

// ParsePricingParametersChanged is a log parse operation binding the contract event 0x5e3c25378b5946068b94aa2ea10c4c1e215cc975f994322b159ddc9237a973d4.
//
// Solidity: event PricingParametersChanged()
func (_Gateway *GatewayFilterer) ParsePricingParametersChanged(log types.Log) (*GatewayPricingParametersChanged, error) {
	event := new(GatewayPricingParametersChanged)
	if err := _Gateway.contract.UnpackLog(event, "PricingParametersChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayTokenSentIterator is returned from FilterTokenSent and is used to iterate over the raw logs and unpacked data for TokenSent events raised by the Gateway contract.
type GatewayTokenSentIterator struct {
	Event *GatewayTokenSent // Event containing the contract specifics and raw log

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
func (it *GatewayTokenSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayTokenSent)
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
		it.Event = new(GatewayTokenSent)
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
func (it *GatewayTokenSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayTokenSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayTokenSent represents a TokenSent event raised by the Gateway contract.
type GatewayTokenSent struct {
	Token              common.Address
	Sender             common.Address
	DestinationChain   uint32
	DestinationAddress MultiAddress
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterTokenSent is a free log retrieval operation binding the contract event 0x24c5d2de620c6e25186ae16f6919eba93b6e2c1a33857cc419d9f3a00d6967e9.
//
// Solidity: event TokenSent(address indexed token, address indexed sender, uint32 indexed destinationChain, (uint8,bytes) destinationAddress, uint128 amount)
func (_Gateway *GatewayFilterer) FilterTokenSent(opts *bind.FilterOpts, token []common.Address, sender []common.Address, destinationChain []uint32) (*GatewayTokenSentIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var destinationChainRule []interface{}
	for _, destinationChainItem := range destinationChain {
		destinationChainRule = append(destinationChainRule, destinationChainItem)
	}

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "TokenSent", tokenRule, senderRule, destinationChainRule)
	if err != nil {
		return nil, err
	}
	return &GatewayTokenSentIterator{contract: _Gateway.contract, event: "TokenSent", logs: logs, sub: sub}, nil
}

// WatchTokenSent is a free log subscription operation binding the contract event 0x24c5d2de620c6e25186ae16f6919eba93b6e2c1a33857cc419d9f3a00d6967e9.
//
// Solidity: event TokenSent(address indexed token, address indexed sender, uint32 indexed destinationChain, (uint8,bytes) destinationAddress, uint128 amount)
func (_Gateway *GatewayFilterer) WatchTokenSent(opts *bind.WatchOpts, sink chan<- *GatewayTokenSent, token []common.Address, sender []common.Address, destinationChain []uint32) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var destinationChainRule []interface{}
	for _, destinationChainItem := range destinationChain {
		destinationChainRule = append(destinationChainRule, destinationChainItem)
	}

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "TokenSent", tokenRule, senderRule, destinationChainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayTokenSent)
				if err := _Gateway.contract.UnpackLog(event, "TokenSent", log); err != nil {
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

// ParseTokenSent is a log parse operation binding the contract event 0x24c5d2de620c6e25186ae16f6919eba93b6e2c1a33857cc419d9f3a00d6967e9.
//
// Solidity: event TokenSent(address indexed token, address indexed sender, uint32 indexed destinationChain, (uint8,bytes) destinationAddress, uint128 amount)
func (_Gateway *GatewayFilterer) ParseTokenSent(log types.Log) (*GatewayTokenSent, error) {
	event := new(GatewayTokenSent)
	if err := _Gateway.contract.UnpackLog(event, "TokenSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayTokenTransferFeesChangedIterator is returned from FilterTokenTransferFeesChanged and is used to iterate over the raw logs and unpacked data for TokenTransferFeesChanged events raised by the Gateway contract.
type GatewayTokenTransferFeesChangedIterator struct {
	Event *GatewayTokenTransferFeesChanged // Event containing the contract specifics and raw log

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
func (it *GatewayTokenTransferFeesChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayTokenTransferFeesChanged)
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
		it.Event = new(GatewayTokenTransferFeesChanged)
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
func (it *GatewayTokenTransferFeesChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayTokenTransferFeesChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayTokenTransferFeesChanged represents a TokenTransferFeesChanged event raised by the Gateway contract.
type GatewayTokenTransferFeesChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTokenTransferFeesChanged is a free log retrieval operation binding the contract event 0x4793c0cb5bef4b1fdbbfbcf17e06991844eb881088b012442af17a12ff38d5cd.
//
// Solidity: event TokenTransferFeesChanged()
func (_Gateway *GatewayFilterer) FilterTokenTransferFeesChanged(opts *bind.FilterOpts) (*GatewayTokenTransferFeesChangedIterator, error) {

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "TokenTransferFeesChanged")
	if err != nil {
		return nil, err
	}
	return &GatewayTokenTransferFeesChangedIterator{contract: _Gateway.contract, event: "TokenTransferFeesChanged", logs: logs, sub: sub}, nil
}

// WatchTokenTransferFeesChanged is a free log subscription operation binding the contract event 0x4793c0cb5bef4b1fdbbfbcf17e06991844eb881088b012442af17a12ff38d5cd.
//
// Solidity: event TokenTransferFeesChanged()
func (_Gateway *GatewayFilterer) WatchTokenTransferFeesChanged(opts *bind.WatchOpts, sink chan<- *GatewayTokenTransferFeesChanged) (event.Subscription, error) {

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "TokenTransferFeesChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayTokenTransferFeesChanged)
				if err := _Gateway.contract.UnpackLog(event, "TokenTransferFeesChanged", log); err != nil {
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

// ParseTokenTransferFeesChanged is a log parse operation binding the contract event 0x4793c0cb5bef4b1fdbbfbcf17e06991844eb881088b012442af17a12ff38d5cd.
//
// Solidity: event TokenTransferFeesChanged()
func (_Gateway *GatewayFilterer) ParseTokenTransferFeesChanged(log types.Log) (*GatewayTokenTransferFeesChanged, error) {
	event := new(GatewayTokenTransferFeesChanged)
	if err := _Gateway.contract.UnpackLog(event, "TokenTransferFeesChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
