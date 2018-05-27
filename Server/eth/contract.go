// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// DocumentRegisterABI is the input ABI used to generate the binding from.
const DocumentRegisterABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"_address\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numberOfDocuments\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_uuid\",\"type\":\"string\"}],\"name\":\"consultByUuid\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"DestroyContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_points\",\"type\":\"uint256\"},{\"name\":\"_trajeto\",\"type\":\"string\"},{\"name\":\"_uuid\",\"type\":\"string\"},{\"name\":\"_km\",\"type\":\"uint256\"}],\"name\":\"registryTrajectory\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// DocumentRegisterBin is the compiled bytecode used for deploying new contracts.
const DocumentRegisterBin = `0x6080604052600060035534801561001557600080fd5b5060008054600160a060020a03191633179055610710806100376000396000f30060806040526004361061006c5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166318bad21781146100715780633fa13d00146100af5780634472f031146100d657806383075e3d1461017f578063aeb8d7ff14610196575b600080fd5b34801561007d57600080fd5b50610086610227565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b3480156100bb57600080fd5b506100c4610243565b60408051918252519081900360200190f35b3480156100e257600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261012f9436949293602493928401919081908401838280828437509497506102499650505050505050565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561016b578181015183820152602001610153565b505050509050019250505060405180910390f35b34801561018b57600080fd5b50610194610329565b005b60408051602060046024803582810135601f810185900485028601850190965285855261019495833595369560449491939091019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375094975050933594506103449350505050565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60035481565b60606102536105c4565b6001836040518082805190602001908083835b602083106102855780518252601f199092019160209182019101610266565b51815160209384036101000a60001901801990921691161790529201948552506040805194859003820185206060860182528054865260018101548684015260028101805483518186028101860185528181529296509287019450909290919083018282801561031457602002820191906000526020600020905b815481526020019060010190808311610300575b50505091909252505050604001519392505050565b60005473ffffffffffffffffffffffffffffffffffffffff16ff5b61034c6105c4565b6103546105e6565b606061035e6105c4565b6003546001016003819055506001866040518082805190602001908083835b6020831061039c5780518252601f19909201916020918201910161037d565b51815160209384036101000a60001901801990921691161790529201948552506040805194859003820185206060860182528054865260018101548684015260028101805483518186028101860185528181529296509287019450909290919083018282801561042b57602002820191906000526020600020905b815481526020019060010190808311610417575b5050509190925250506040805160808101825289815260208082018a90528183018c905242606083015260035460009081526002825292909220815180519599509197508794909350610481928492019061060f565b506020828101516001830155604083015180516104a4926002850192019061060f565b50606082015181600301559050508360400151516001016040519080825280602002602001820160405280156104e4578160200160208202803883390190505b5091508260600151828560400151518151811015156104ff57fe5b9060200190602002018181525050606060405190810160405280898660000151018152602001428152602001838152509050806001876040518082805190602001908083835b602083106105645780518252601f199092019160209182019101610545565b51815160209384036101000a600019018019909216911617905292019485525060408051948590038201909420855181558582015160018201559385015180516105b7945060028601935091019061068d565b5050505050505050505050565b6060604051908101604052806000815260200160008152602001606081525090565b608060405190810160405280606081526020016000815260200160608152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061065057805160ff191683800117855561067d565b8280016001018555821561067d579182015b8281111561067d578251825591602001919060010190610662565b506106899291506106c7565b5090565b82805482825590600052602060002090810192821561067d579160200282018281111561067d578251825591602001919060010190610662565b6106e191905b8082111561068957600081556001016106cd565b905600a165627a7a723058204104ca17f8a4aa2e582fab0bd42cf55e3e253595d6df7693376f1dce17deac170029`

// DeployDocumentRegister deploys a new Ethereum contract, binding an instance of DocumentRegister to it.
func DeployDocumentRegister(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DocumentRegister, error) {
	parsed, err := abi.JSON(strings.NewReader(DocumentRegisterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DocumentRegisterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DocumentRegister{DocumentRegisterCaller: DocumentRegisterCaller{contract: contract}, DocumentRegisterTransactor: DocumentRegisterTransactor{contract: contract}, DocumentRegisterFilterer: DocumentRegisterFilterer{contract: contract}}, nil
}

// DocumentRegister is an auto generated Go binding around an Ethereum contract.
type DocumentRegister struct {
	DocumentRegisterCaller     // Read-only binding to the contract
	DocumentRegisterTransactor // Write-only binding to the contract
	DocumentRegisterFilterer   // Log filterer for contract events
}

// DocumentRegisterCaller is an auto generated read-only Go binding around an Ethereum contract.
type DocumentRegisterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DocumentRegisterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DocumentRegisterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DocumentRegisterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DocumentRegisterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DocumentRegisterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DocumentRegisterSession struct {
	Contract     *DocumentRegister // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DocumentRegisterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DocumentRegisterCallerSession struct {
	Contract *DocumentRegisterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// DocumentRegisterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DocumentRegisterTransactorSession struct {
	Contract     *DocumentRegisterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// DocumentRegisterRaw is an auto generated low-level Go binding around an Ethereum contract.
type DocumentRegisterRaw struct {
	Contract *DocumentRegister // Generic contract binding to access the raw methods on
}

// DocumentRegisterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DocumentRegisterCallerRaw struct {
	Contract *DocumentRegisterCaller // Generic read-only contract binding to access the raw methods on
}

// DocumentRegisterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DocumentRegisterTransactorRaw struct {
	Contract *DocumentRegisterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDocumentRegister creates a new instance of DocumentRegister, bound to a specific deployed contract.
func NewDocumentRegister(address common.Address, backend bind.ContractBackend) (*DocumentRegister, error) {
	contract, err := bindDocumentRegister(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DocumentRegister{DocumentRegisterCaller: DocumentRegisterCaller{contract: contract}, DocumentRegisterTransactor: DocumentRegisterTransactor{contract: contract}, DocumentRegisterFilterer: DocumentRegisterFilterer{contract: contract}}, nil
}

// NewDocumentRegisterCaller creates a new read-only instance of DocumentRegister, bound to a specific deployed contract.
func NewDocumentRegisterCaller(address common.Address, caller bind.ContractCaller) (*DocumentRegisterCaller, error) {
	contract, err := bindDocumentRegister(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DocumentRegisterCaller{contract: contract}, nil
}

// NewDocumentRegisterTransactor creates a new write-only instance of DocumentRegister, bound to a specific deployed contract.
func NewDocumentRegisterTransactor(address common.Address, transactor bind.ContractTransactor) (*DocumentRegisterTransactor, error) {
	contract, err := bindDocumentRegister(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DocumentRegisterTransactor{contract: contract}, nil
}

// NewDocumentRegisterFilterer creates a new log filterer instance of DocumentRegister, bound to a specific deployed contract.
func NewDocumentRegisterFilterer(address common.Address, filterer bind.ContractFilterer) (*DocumentRegisterFilterer, error) {
	contract, err := bindDocumentRegister(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DocumentRegisterFilterer{contract: contract}, nil
}

// bindDocumentRegister binds a generic wrapper to an already deployed contract.
func bindDocumentRegister(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DocumentRegisterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DocumentRegister *DocumentRegisterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DocumentRegister.Contract.DocumentRegisterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DocumentRegister *DocumentRegisterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DocumentRegister.Contract.DocumentRegisterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DocumentRegister *DocumentRegisterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DocumentRegister.Contract.DocumentRegisterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DocumentRegister *DocumentRegisterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DocumentRegister.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DocumentRegister *DocumentRegisterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DocumentRegister.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DocumentRegister *DocumentRegisterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DocumentRegister.Contract.contract.Transact(opts, method, params...)
}

// Address is a free data retrieval call binding the contract method 0x18bad217.
//
// Solidity: function _address() constant returns(address)
func (_DocumentRegister *DocumentRegisterCaller) Address(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DocumentRegister.contract.Call(opts, out, "_address")
	return *ret0, err
}

// Address is a free data retrieval call binding the contract method 0x18bad217.
//
// Solidity: function _address() constant returns(address)
func (_DocumentRegister *DocumentRegisterSession) Address() (common.Address, error) {
	return _DocumentRegister.Contract.Address(&_DocumentRegister.CallOpts)
}

// Address is a free data retrieval call binding the contract method 0x18bad217.
//
// Solidity: function _address() constant returns(address)
func (_DocumentRegister *DocumentRegisterCallerSession) Address() (common.Address, error) {
	return _DocumentRegister.Contract.Address(&_DocumentRegister.CallOpts)
}

// ConsultByUuid is a free data retrieval call binding the contract method 0x4472f031.
//
// Solidity: function consultByUuid(_uuid string) constant returns(uint256[])
func (_DocumentRegister *DocumentRegisterCaller) ConsultByUuid(opts *bind.CallOpts, _uuid string) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _DocumentRegister.contract.Call(opts, out, "consultByUuid", _uuid)
	return *ret0, err
}

// ConsultByUuid is a free data retrieval call binding the contract method 0x4472f031.
//
// Solidity: function consultByUuid(_uuid string) constant returns(uint256[])
func (_DocumentRegister *DocumentRegisterSession) ConsultByUuid(_uuid string) ([]*big.Int, error) {
	return _DocumentRegister.Contract.ConsultByUuid(&_DocumentRegister.CallOpts, _uuid)
}

// ConsultByUuid is a free data retrieval call binding the contract method 0x4472f031.
//
// Solidity: function consultByUuid(_uuid string) constant returns(uint256[])
func (_DocumentRegister *DocumentRegisterCallerSession) ConsultByUuid(_uuid string) ([]*big.Int, error) {
	return _DocumentRegister.Contract.ConsultByUuid(&_DocumentRegister.CallOpts, _uuid)
}

// NumberOfDocuments is a free data retrieval call binding the contract method 0x3fa13d00.
//
// Solidity: function numberOfDocuments() constant returns(uint256)
func (_DocumentRegister *DocumentRegisterCaller) NumberOfDocuments(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DocumentRegister.contract.Call(opts, out, "numberOfDocuments")
	return *ret0, err
}

// NumberOfDocuments is a free data retrieval call binding the contract method 0x3fa13d00.
//
// Solidity: function numberOfDocuments() constant returns(uint256)
func (_DocumentRegister *DocumentRegisterSession) NumberOfDocuments() (*big.Int, error) {
	return _DocumentRegister.Contract.NumberOfDocuments(&_DocumentRegister.CallOpts)
}

// NumberOfDocuments is a free data retrieval call binding the contract method 0x3fa13d00.
//
// Solidity: function numberOfDocuments() constant returns(uint256)
func (_DocumentRegister *DocumentRegisterCallerSession) NumberOfDocuments() (*big.Int, error) {
	return _DocumentRegister.Contract.NumberOfDocuments(&_DocumentRegister.CallOpts)
}

// DestroyContract is a paid mutator transaction binding the contract method 0x83075e3d.
//
// Solidity: function DestroyContract() returns()
func (_DocumentRegister *DocumentRegisterTransactor) DestroyContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DocumentRegister.contract.Transact(opts, "DestroyContract")
}

// DestroyContract is a paid mutator transaction binding the contract method 0x83075e3d.
//
// Solidity: function DestroyContract() returns()
func (_DocumentRegister *DocumentRegisterSession) DestroyContract() (*types.Transaction, error) {
	return _DocumentRegister.Contract.DestroyContract(&_DocumentRegister.TransactOpts)
}

// DestroyContract is a paid mutator transaction binding the contract method 0x83075e3d.
//
// Solidity: function DestroyContract() returns()
func (_DocumentRegister *DocumentRegisterTransactorSession) DestroyContract() (*types.Transaction, error) {
	return _DocumentRegister.Contract.DestroyContract(&_DocumentRegister.TransactOpts)
}

// RegistryTrajectory is a paid mutator transaction binding the contract method 0xaeb8d7ff.
//
// Solidity: function registryTrajectory(_points uint256, _trajeto string, _uuid string, _km uint256) returns()
func (_DocumentRegister *DocumentRegisterTransactor) RegistryTrajectory(opts *bind.TransactOpts, _points *big.Int, _trajeto string, _uuid string, _km *big.Int) (*types.Transaction, error) {
	return _DocumentRegister.contract.Transact(opts, "registryTrajectory", _points, _trajeto, _uuid, _km)
}

// RegistryTrajectory is a paid mutator transaction binding the contract method 0xaeb8d7ff.
//
// Solidity: function registryTrajectory(_points uint256, _trajeto string, _uuid string, _km uint256) returns()
func (_DocumentRegister *DocumentRegisterSession) RegistryTrajectory(_points *big.Int, _trajeto string, _uuid string, _km *big.Int) (*types.Transaction, error) {
	return _DocumentRegister.Contract.RegistryTrajectory(&_DocumentRegister.TransactOpts, _points, _trajeto, _uuid, _km)
}

// RegistryTrajectory is a paid mutator transaction binding the contract method 0xaeb8d7ff.
//
// Solidity: function registryTrajectory(_points uint256, _trajeto string, _uuid string, _km uint256) returns()
func (_DocumentRegister *DocumentRegisterTransactorSession) RegistryTrajectory(_points *big.Int, _trajeto string, _uuid string, _km *big.Int) (*types.Transaction, error) {
	return _DocumentRegister.Contract.RegistryTrajectory(&_DocumentRegister.TransactOpts, _points, _trajeto, _uuid, _km)
}
