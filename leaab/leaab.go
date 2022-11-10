// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package leeab

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
)


// LeaabMetaData contains all meta data concerning the Leaab contract.
var LeaabMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiveAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isExcluded\",\"type\":\"bool\"}],\"name\":\"ExcludeFromFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isExcluded\",\"type\":\"bool\"}],\"name\":\"ExcludeMultipleAccountsFromFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldValue\",\"type\":\"uint256\"}],\"name\":\"GasForProcessingUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newLiquidityWallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldLiquidityWallet\",\"type\":\"address\"}],\"name\":\"LiquidityWalletUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"iterations\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claims\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lastProcessedIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"automatic\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"processor\",\"type\":\"address\"}],\"name\":\"ProcessedDividendTracker\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokensSwapped\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SendDividends\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"SetAutomatedMarketMakerPair\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokensSwapped\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethReceived\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokensIntoLiqudity\",\"type\":\"uint256\"}],\"name\":\"SwapAndLiquify\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"}],\"name\":\"UpdateDividendTracker\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"}],\"name\":\"UpdateUniswapV2Router\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHRewardsFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_marketingWalletAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"automatedMarketMakerPairs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burnFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"canEat\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"canEatAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deadWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"dividendTokenBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dividendTracker\",\"outputs\":[{\"internalType\":\"contractETHBackDividendTracker\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"excludeFromDividends\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"excluded\",\"type\":\"bool\"}],\"name\":\"excludeFromFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"excluded\",\"type\":\"bool\"}],\"name\":\"excludeMultipleAccountsFromFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasForProcessing\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountDividendsInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getAccountDividendsInfoAtIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getClaimWait\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastProcessedIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumberOfDividendTokenHolders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalDividendsDistributed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"eatList\",\"type\":\"address[]\"}],\"name\":\"hungry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isExcludedFromFees\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isL\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"killNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lunachB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketingFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"eatList\",\"type\":\"address[]\"}],\"name\":\"nothungry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"}],\"name\":\"processDividendTracker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"setAutomatedMarketMakerPair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setBurnFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setETHRewardsFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"setKillNum\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setMarketingFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"setMarketingWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapTokensAtAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapV2Pair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapV2Router\",\"outputs\":[{\"internalType\":\"contractIUniswapV2Router02\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"claimWait\",\"type\":\"uint256\"}],\"name\":\"updateClaimWait\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"updateDividendTracker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"updateGasForProcessing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"updateUniswapV2Router\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"withdrawableDividendOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// LeaabABI is the input ABI used to generate the binding from.
// Deprecated: Use LeaabMetaData.ABI instead.
var LeaabABI = LeaabMetaData.ABI

// Leaab is an auto generated Go binding around an Ethereum contract.
type Leaab struct {
	LeaabCaller     // Read-only binding to the contract
	LeaabTransactor // Write-only binding to the contract
	LeaabFilterer   // Log filterer for contract events
}

// LeaabCaller is an auto generated read-only Go binding around an Ethereum contract.
type LeaabCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LeaabTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LeaabTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LeaabFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LeaabFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LeaabSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LeaabSession struct {
	Contract     *Leaab            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LeaabCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LeaabCallerSession struct {
	Contract *LeaabCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LeaabTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LeaabTransactorSession struct {
	Contract     *LeaabTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LeaabRaw is an auto generated low-level Go binding around an Ethereum contract.
type LeaabRaw struct {
	Contract *Leaab // Generic contract binding to access the raw methods on
}

// LeaabCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LeaabCallerRaw struct {
	Contract *LeaabCaller // Generic read-only contract binding to access the raw methods on
}

// LeaabTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LeaabTransactorRaw struct {
	Contract *LeaabTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLeaab creates a new instance of Leaab, bound to a specific deployed contract.
func NewLeaab(address common.Address, backend bind.ContractBackend) (*Leaab, error) {
	contract, err := bindLeaab(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Leaab{LeaabCaller: LeaabCaller{contract: contract}, LeaabTransactor: LeaabTransactor{contract: contract}, LeaabFilterer: LeaabFilterer{contract: contract}}, nil
}

// NewLeaabCaller creates a new read-only instance of Leaab, bound to a specific deployed contract.
func NewLeaabCaller(address common.Address, caller bind.ContractCaller) (*LeaabCaller, error) {
	contract, err := bindLeaab(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LeaabCaller{contract: contract}, nil
}

// NewLeaabTransactor creates a new write-only instance of Leaab, bound to a specific deployed contract.
func NewLeaabTransactor(address common.Address, transactor bind.ContractTransactor) (*LeaabTransactor, error) {
	contract, err := bindLeaab(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LeaabTransactor{contract: contract}, nil
}

// NewLeaabFilterer creates a new log filterer instance of Leaab, bound to a specific deployed contract.
func NewLeaabFilterer(address common.Address, filterer bind.ContractFilterer) (*LeaabFilterer, error) {
	contract, err := bindLeaab(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LeaabFilterer{contract: contract}, nil
}

// bindLeaab binds a generic wrapper to an already deployed contract.
func bindLeaab(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LeaabABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Leaab *LeaabRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Leaab.Contract.LeaabCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Leaab *LeaabRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Leaab.Contract.LeaabTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Leaab *LeaabRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Leaab.Contract.LeaabTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Leaab *LeaabCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Leaab.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Leaab *LeaabTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Leaab.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Leaab *LeaabTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Leaab.Contract.contract.Transact(opts, method, params...)
}

// ETH is a free data retrieval call binding the contract method 0x8322fff2.
//
// Solidity: function ETH() view returns(address)
func (_Leaab *LeaabCaller) ETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Leaab.contract.Call(opts, &out, "ETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ETH is a free data retrieval call binding the contract method 0x8322fff2.
//
// Solidity: function ETH() view returns(address)
func (_Leaab *LeaabSession) ETH() (common.Address, error) {
	return _Leaab.Contract.ETH(&_Leaab.CallOpts)
}

// ETH is a free data retrieval call binding the contract method 0x8322fff2.
//
// Solidity: function ETH() view returns(address)
func (_Leaab *LeaabCallerSession) ETH() (common.Address, error) {
	return _Leaab.Contract.ETH(&_Leaab.CallOpts)
}

// ETHRewardsFee is a free data retrieval call binding the contract method 0x8c0344db.
//
// Solidity: function ETHRewardsFee() view returns(uint256)
func (_Leaab *LeaabCaller) ETHRewardsFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Leaab.contract.Call(opts, &out, "ETHRewardsFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ETHRewardsFee is a free data retrieval call binding the contract method 0x8c0344db.
//
// Solidity: function ETHRewardsFee() view returns(uint256)
func (_Leaab *LeaabSession) ETHRewardsFee() (*big.Int, error) {
	return _Leaab.Contract.ETHRewardsFee(&_Leaab.CallOpts)
}

// ETHRewardsFee is a free data retrieval call binding the contract method 0x8c0344db.
//
// Solidity: function ETHRewardsFee() view returns(uint256)
func (_Leaab *LeaabCallerSession) ETHRewardsFee() (*big.Int, error) {
	return _Leaab.Contract.ETHRewardsFee(&_Leaab.CallOpts)
}

// MarketingWalletAddress is a free data retrieval call binding the contract method 0x4144d9e4.
//
// Solidity: function _marketingWalletAddress() view returns(address)
func (_Leaab *LeaabCaller) MarketingWalletAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Leaab.contract.Call(opts, &out, "_marketingWalletAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MarketingWalletAddress is a free data retrieval call binding the contract method 0x4144d9e4.
//
// Solidity: function _marketingWalletAddress() view returns(address)
func (_Leaab *LeaabSession) MarketingWalletAddress() (common.Address, error) {
	return _Leaab.Contract.MarketingWalletAddress(&_Leaab.CallOpts)
}

// MarketingWalletAddress is a free data retrieval call binding the contract method 0x4144d9e4.
//
// Solidity: function _marketingWalletAddress() view returns(address)
func (_Leaab *LeaabCallerSession) MarketingWalletAddress() (common.Address, error) {
	return _Leaab.Contract.MarketingWalletAddress(&_Leaab.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Leaab *LeaabCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Leaab.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Leaab *LeaabSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Leaab.Contract.Allowance(&_Leaab.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Leaab *LeaabCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Leaab.Contract.Allowance(&_Leaab.CallOpts, owner, spender)
}

// AutomatedMarketMakerPairs is a free data retrieval call binding the contract method 0xb62496f5.
//
// Solidity: function automatedMarketMakerPairs(address ) view returns(bool)
func (_Leaab *LeaabCaller) AutomatedMarketMakerPairs(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Leaab.contract.Call(opts, &out, "automatedMarketMakerPairs", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AutomatedMarketMakerPairs is a free data retrieval call binding the contract method 0xb62496f5.
//
// Solidity: function automatedMarketMakerPairs(address ) view returns(bool)
func (_Leaab *LeaabSession) AutomatedMarketMakerPairs(arg0 common.Address) (bool, error) {
	return _Leaab.Contract.AutomatedMarketMakerPairs(&_Leaab.CallOpts, arg0)
}

// AutomatedMarketMakerPairs is a free data retrieval call binding the contract method 0xb62496f5.
//
// Solidity: function automatedMarketMakerPairs(address ) view returns(bool)
func (_Leaab *LeaabCallerSession) AutomatedMarketMakerPairs(arg0 common.Address) (bool, error) {
	return _Leaab.Contract.AutomatedMarketMakerPairs(&_Leaab.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Leaab *LeaabCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Leaab.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Leaab *LeaabSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Leaab.Contract.BalanceOf(&_Leaab.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Leaab *LeaabCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Leaab.Contract.BalanceOf(&_Leaab.CallOpts, account)
}

// BurnFee is a free data retrieval call binding the contract method 0xfce589d8.
//
// Solidity: function burnFee() view returns(uint256)
func (_Leaab *LeaabCaller) BurnFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Leaab.contract.Call(opts, &out, "burnFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BurnFee is a free data retrieval call binding the contract method 0xfce589d8.
//
// Solidity: function burnFee() view returns(uint256)
func (_Leaab *LeaabSession) BurnFee() (*big.Int, error) {
	return _Leaab.Contract.BurnFee(&_Leaab.CallOpts)
}

// BurnFee is a free data retrieval call binding the contract method 0xfce589d8.
//
// Solidity: function burnFee() view returns(uint256)
func (_Leaab *LeaabCallerSession) BurnFee() (*big.Int, error) {
	return _Leaab.Contract.BurnFee(&_Leaab.CallOpts)
}

// CanEat is a free data retrieval call binding the contract method 0xdf989d9e.
//
// Solidity: function canEat(address ) view returns(bool)
func (_Leaab *LeaabCaller) CanEat(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Leaab.contract.Call(opts, &out, "canEat", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanEat is a free data retrieval call binding the contract method 0xdf989d9e.
//
// Solidity: function canEat(address ) view returns(bool)
func (_Leaab *LeaabSession) CanEat(arg0 common.Address) (bool, error) {
	return _Leaab.Contract.CanEat(&_Leaab.CallOpts, arg0)
}

// CanEat is a free data retrieval call binding the contract method 0xdf989d9e.
//
// Solidity: function canEat(address ) view returns(bool)
func (_Leaab *LeaabCallerSession) CanEat(arg0 common.Address) (bool, error) {
	return _Leaab.Contract.CanEat(&_Leaab.CallOpts, arg0)
}

// DeadWallet is a free data retrieval call binding the contract method 0x85141a77.
//
// Solidity: function deadWallet() view returns(address)
func (_Leaab *LeaabCaller) DeadWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Leaab.contract.Call(opts, &out, "deadWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DeadWallet is a free data retrieval call binding the contract method 0x85141a77.
//
// Solidity: function deadWallet() view returns(address)
func (_Leaab *LeaabSession) DeadWallet() (common.Address, error) {
	return _Leaab.Contract.DeadWallet(&_Leaab.CallOpts)
}

// DeadWallet is a free data retrieval call binding the contract method 0x85141a77.
//
// Solidity: function deadWallet() view returns(address)
func (_Leaab *LeaabCallerSession) DeadWallet() (common.Address, error) {
	return _Leaab.Contract.DeadWallet(&_Leaab.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Leaab *LeaabCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Leaab.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Leaab *LeaabSession) Decimals() (uint8, error) {
	return _Leaab.Contract.Decimals(&_Leaab.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Leaab *LeaabCallerSession) Decimals() (uint8, error) {
	return _Leaab.Contract.Decimals(&_Leaab.CallOpts)
}

// DividendTokenBalanceOf is a free data retrieval call binding the contract method 0x6843cd84.
//
// Solidity: function dividendTokenBalanceOf(address account) view returns(uint256)
func (_Leaab *LeaabCaller) DividendTokenBalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Leaab.contract.Call(opts, &out, "dividendTokenBalanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DividendTokenBalanceOf is a free data retrieval call binding the contract method 0x6843cd84.
//
// Solidity: function dividendTokenBalanceOf(address account) view returns(uint256)
func (_Leaab *LeaabSession) DividendTokenBalanceOf(account common.Address) (*big.Int, error) {
	return _Leaab.Contract.DividendTokenBalanceOf(&_Leaab.CallOpts, account)
}

// DividendTokenBalanceOf is a free data retrieval call binding the contract method 0x6843cd84.
//
// Solidity: function dividendTokenBalanceOf(address account) view returns(uint256)
func (_Leaab *LeaabCallerSession) DividendTokenBalanceOf(account common.Address) (*big.Int, error) {
	return _Leaab.Contract.DividendTokenBalanceOf(&_Leaab.CallOpts, account)
}

// DividendTracker is a free data retrieval call binding the contract method 0x2c1f5216.
//
// Solidity: function dividendTracker() view returns(address)
func (_Leaab *LeaabCaller) DividendTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Leaab.contract.Call(opts, &out, "dividendTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DividendTracker is a free data retrieval call binding the contract method 0x2c1f5216.
//
// Solidity: function dividendTracker() view returns(address)
func (_Leaab *LeaabSession) DividendTracker() (common.Address, error) {
	return _Leaab.Contract.DividendTracker(&_Leaab.CallOpts)
}

// DividendTracker is a free data retrieval call binding the contract method 0x2c1f5216.
//
// Solidity: function dividendTracker() view returns(address)
func (_Leaab *LeaabCallerSession) DividendTracker() (common.Address, error) {
	return _Leaab.Contract.DividendTracker(&_Leaab.CallOpts)
}

// GasForProcessing is a free data retrieval call binding the contract method 0x9c1b8af5.
//
// Solidity: function gasForProcessing() view returns(uint256)
func (_Leaab *LeaabCaller) GasForProcessing(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Leaab.contract.Call(opts, &out, "gasForProcessing")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GasForProcessing is a free data retrieval call binding the contract method 0x9c1b8af5.
//
// Solidity: function gasForProcessing() view returns(uint256)
func (_Leaab *LeaabSession) GasForProcessing() (*big.Int, error) {
	return _Leaab.Contract.GasForProcessing(&_Leaab.CallOpts)
}

// GasForProcessing is a free data retrieval call binding the contract method 0x9c1b8af5.
//
// Solidity: function gasForProcessing() view returns(uint256)
func (_Leaab *LeaabCallerSession) GasForProcessing() (*big.Int, error) {
	return _Leaab.Contract.GasForProcessing(&_Leaab.CallOpts)
}

// GetAccountDividendsInfo is a free data retrieval call binding the contract method 0xad56c13c.
//
// Solidity: function getAccountDividendsInfo(address account) view returns(address, int256, int256, uint256, uint256, uint256, uint256, uint256)
func (_Leaab *LeaabCaller) GetAccountDividendsInfo(opts *bind.CallOpts, account common.Address) (common.Address, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Leaab.contract.Call(opts, &out, "getAccountDividendsInfo", account)

	if err != nil {
		return *new(common.Address), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	out6 := *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	out7 := *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, out6, out7, err

}

// GetAccountDividendsInfo is a free data retrieval call binding the contract method 0xad56c13c.
//
// Solidity: function getAccountDividendsInfo(address account) view returns(address, int256, int256, uint256, uint256, uint256, uint256, uint256)
func (_Leaab *LeaabSession) GetAccountDividendsInfo(account common.Address) (common.Address, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Leaab.Contract.GetAccountDividendsInfo(&_Leaab.CallOpts, account)
}

// GetAccountDividendsInfo is a free data retrieval call binding the contract method 0xad56c13c.
//
// Solidity: function getAccountDividendsInfo(address account) view returns(address, int256, int256, uint256, uint256, uint256, uint256, uint256)
func (_Leaab *LeaabCallerSession) GetAccountDividendsInfo(account common.Address) (common.Address, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Leaab.Contract.GetAccountDividendsInfo(&_Leaab.CallOpts, account)
}

// GetAccountDividendsInfoAtIndex is a free data retrieval call binding the contract method 0xf27fd254.
//
// Solidity: function getAccountDividendsInfoAtIndex(uint256 index) view returns(address, int256, int256, uint256, uint256, uint256, uint256, uint256)
func (_Leaab *LeaabCaller) GetAccountDividendsInfoAtIndex(opts *bind.CallOpts, index *big.Int) (common.Address, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Leaab.contract.Call(opts, &out, "getAccountDividendsInfoAtIndex", index)

	if err != nil {
		return *new(common.Address), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	out6 := *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	out7 := *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, out6, out7, err

}

// GetAccountDividendsInfoAtIndex is a free data retrieval call binding the contract method 0xf27fd254.
//
// Solidity: function getAccountDividendsInfoAtIndex(uint256 index) view returns(address, int256, int256, uint256, uint256, uint256, uint256, uint256)
func (_Leaab *LeaabSession) GetAccountDividendsInfoAtIndex(index *big.Int) (common.Address, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Leaab.Contract.GetAccountDividendsInfoAtIndex(&_Leaab.CallOpts, index)
}

// GetAccountDividendsInfoAtIndex is a free data retrieval call binding the contract method 0xf27fd254.
//
// Solidity: function getAccountDividendsInfoAtIndex(uint256 index) view returns(address, int256, int256, uint256, uint256, uint256, uint256, uint256)
func (_Leaab *LeaabCallerSession) GetAccountDividendsInfoAtIndex(index *big.Int) (common.Address, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Leaab.Contract.GetAccountDividendsInfoAtIndex(&_Leaab.CallOpts, index)
}

// GetClaimWait is a free data retrieval call binding the contract method 0xa26579ad.
//
// Solidity: function getClaimWait() view returns(uint256)
func (_Leaab *LeaabCaller) GetClaimWait(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Leaab.contract.Call(opts, &out, "getClaimWait")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClaimWait is a free data retrieval call binding the contract method 0xa26579ad.
//
// Solidity: function getClaimWait() view returns(uint256)
func (_Leaab *LeaabSession) GetClaimWait() (*big.Int, error) {
	return _Leaab.Contract.GetClaimWait(&_Leaab.CallOpts)
}

// GetClaimWait is a free data retrieval call binding the contract method 0xa26579ad.
//
// Solidity: function getClaimWait() view returns(uint256)
func (_Leaab *LeaabCallerSession) GetClaimWait() (*big.Int, error) {
	return _Leaab.Contract.GetClaimWait(&_Leaab.CallOpts)
}

// GetLastProcessedIndex is a free data retrieval call binding the contract method 0xe7841ec0.
//
// Solidity: function getLastProcessedIndex() view returns(uint256)
func (_Leaab *LeaabCaller) GetLastProcessedIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Leaab.contract.Call(opts, &out, "getLastProcessedIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastProcessedIndex is a free data retrieval call binding the contract method 0xe7841ec0.
//
// Solidity: function getLastProcessedIndex() view returns(uint256)
func (_Leaab *LeaabSession) GetLastProcessedIndex() (*big.Int, error) {
	return _Leaab.Contract.GetLastProcessedIndex(&_Leaab.CallOpts)
}

// GetLastProcessedIndex is a free data retrieval call binding the contract method 0xe7841ec0.
//
// Solidity: function getLastProcessedIndex() view returns(uint256)
func (_Leaab *LeaabCallerSession) GetLastProcessedIndex() (*big.Int, error) {
	return _Leaab.Contract.GetLastProcessedIndex(&_Leaab.CallOpts)
}

// GetNumberOfDividendTokenHolders is a free data retrieval call binding the contract method 0x64b0f653.
//
// Solidity: function getNumberOfDividendTokenHolders() view returns(uint256)
func (_Leaab *LeaabCaller) GetNumberOfDividendTokenHolders(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Leaab.contract.Call(opts, &out, "getNumberOfDividendTokenHolders")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumberOfDividendTokenHolders is a free data retrieval call binding the contract method 0x64b0f653.
//
// Solidity: function getNumberOfDividendTokenHolders() view returns(uint256)
func (_Leaab *LeaabSession) GetNumberOfDividendTokenHolders() (*big.Int, error) {
	return _Leaab.Contract.GetNumberOfDividendTokenHolders(&_Leaab.CallOpts)
}

// GetNumberOfDividendTokenHolders is a free data retrieval call binding the contract method 0x64b0f653.
//
// Solidity: function getNumberOfDividendTokenHolders() view returns(uint256)
func (_Leaab *LeaabCallerSession) GetNumberOfDividendTokenHolders() (*big.Int, error) {
	return _Leaab.Contract.GetNumberOfDividendTokenHolders(&_Leaab.CallOpts)
}

// GetTotalDividendsDistributed is a free data retrieval call binding the contract method 0x30bb4cff.
//
// Solidity: function getTotalDividendsDistributed() view returns(uint256)
func (_Leaab *LeaabCaller) GetTotalDividendsDistributed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Leaab.contract.Call(opts, &out, "getTotalDividendsDistributed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalDividendsDistributed is a free data retrieval call binding the contract method 0x30bb4cff.
//
// Solidity: function getTotalDividendsDistributed() view returns(uint256)
func (_Leaab *LeaabSession) GetTotalDividendsDistributed() (*big.Int, error) {
	return _Leaab.Contract.GetTotalDividendsDistributed(&_Leaab.CallOpts)
}

// GetTotalDividendsDistributed is a free data retrieval call binding the contract method 0x30bb4cff.
//
// Solidity: function getTotalDividendsDistributed() view returns(uint256)
func (_Leaab *LeaabCallerSession) GetTotalDividendsDistributed() (*big.Int, error) {
	return _Leaab.Contract.GetTotalDividendsDistributed(&_Leaab.CallOpts)
}

// IsExcludedFromFees is a free data retrieval call binding the contract method 0x4fbee193.
//
// Solidity: function isExcludedFromFees(address account) view returns(bool)
func (_Leaab *LeaabCaller) IsExcludedFromFees(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Leaab.contract.Call(opts, &out, "isExcludedFromFees", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExcludedFromFees is a free data retrieval call binding the contract method 0x4fbee193.
//
// Solidity: function isExcludedFromFees(address account) view returns(bool)
func (_Leaab *LeaabSession) IsExcludedFromFees(account common.Address) (bool, error) {
	return _Leaab.Contract.IsExcludedFromFees(&_Leaab.CallOpts, account)
}

// IsExcludedFromFees is a free data retrieval call binding the contract method 0x4fbee193.
//
// Solidity: function isExcludedFromFees(address account) view returns(bool)
func (_Leaab *LeaabCallerSession) IsExcludedFromFees(account common.Address) (bool, error) {
	return _Leaab.Contract.IsExcludedFromFees(&_Leaab.CallOpts, account)
}

// IsL is a free data retrieval call binding the contract method 0x3ead9edb.
//
// Solidity: function isL() view returns(bool)
func (_Leaab *LeaabCaller) IsL(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Leaab.contract.Call(opts, &out, "isL")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsL is a free data retrieval call binding the contract method 0x3ead9edb.
//
// Solidity: function isL() view returns(bool)
func (_Leaab *LeaabSession) IsL() (bool, error) {
	return _Leaab.Contract.IsL(&_Leaab.CallOpts)
}

// IsL is a free data retrieval call binding the contract method 0x3ead9edb.
//
// Solidity: function isL() view returns(bool)
func (_Leaab *LeaabCallerSession) IsL() (bool, error) {
	return _Leaab.Contract.IsL(&_Leaab.CallOpts)
}

// KillNum is a free data retrieval call binding the contract method 0x0923909f.
//
// Solidity: function killNum() view returns(uint256)
func (_Leaab *LeaabCaller) KillNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Leaab.contract.Call(opts, &out, "killNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KillNum is a free data retrieval call binding the contract method 0x0923909f.
//
// Solidity: function killNum() view returns(uint256)
func (_Leaab *LeaabSession) KillNum() (*big.Int, error) {
	return _Leaab.Contract.KillNum(&_Leaab.CallOpts)
}

// KillNum is a free data retrieval call binding the contract method 0x0923909f.
//
// Solidity: function killNum() view returns(uint256)
func (_Leaab *LeaabCallerSession) KillNum() (*big.Int, error) {
	return _Leaab.Contract.KillNum(&_Leaab.CallOpts)
}

// LunachB is a free data retrieval call binding the contract method 0x566da6d2.
//
// Solidity: function lunachB() view returns(uint256)
func (_Leaab *LeaabCaller) LunachB(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Leaab.contract.Call(opts, &out, "lunachB")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LunachB is a free data retrieval call binding the contract method 0x566da6d2.
//
// Solidity: function lunachB() view returns(uint256)
func (_Leaab *LeaabSession) LunachB() (*big.Int, error) {
	return _Leaab.Contract.LunachB(&_Leaab.CallOpts)
}

// LunachB is a free data retrieval call binding the contract method 0x566da6d2.
//
// Solidity: function lunachB() view returns(uint256)
func (_Leaab *LeaabCallerSession) LunachB() (*big.Int, error) {
	return _Leaab.Contract.LunachB(&_Leaab.CallOpts)
}

// MarketingFee is a free data retrieval call binding the contract method 0x6b67c4df.
//
// Solidity: function marketingFee() view returns(uint256)
func (_Leaab *LeaabCaller) MarketingFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Leaab.contract.Call(opts, &out, "marketingFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarketingFee is a free data retrieval call binding the contract method 0x6b67c4df.
//
// Solidity: function marketingFee() view returns(uint256)
func (_Leaab *LeaabSession) MarketingFee() (*big.Int, error) {
	return _Leaab.Contract.MarketingFee(&_Leaab.CallOpts)
}

// MarketingFee is a free data retrieval call binding the contract method 0x6b67c4df.
//
// Solidity: function marketingFee() view returns(uint256)
func (_Leaab *LeaabCallerSession) MarketingFee() (*big.Int, error) {
	return _Leaab.Contract.MarketingFee(&_Leaab.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Leaab *LeaabCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Leaab.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Leaab *LeaabSession) Name() (string, error) {
	return _Leaab.Contract.Name(&_Leaab.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Leaab *LeaabCallerSession) Name() (string, error) {
	return _Leaab.Contract.Name(&_Leaab.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Leaab *LeaabCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Leaab.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Leaab *LeaabSession) Owner() (common.Address, error) {
	return _Leaab.Contract.Owner(&_Leaab.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Leaab *LeaabCallerSession) Owner() (common.Address, error) {
	return _Leaab.Contract.Owner(&_Leaab.CallOpts)
}

// SwapTokensAtAmount is a free data retrieval call binding the contract method 0xe2f45605.
//
// Solidity: function swapTokensAtAmount() view returns(uint256)
func (_Leaab *LeaabCaller) SwapTokensAtAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Leaab.contract.Call(opts, &out, "swapTokensAtAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SwapTokensAtAmount is a free data retrieval call binding the contract method 0xe2f45605.
//
// Solidity: function swapTokensAtAmount() view returns(uint256)
func (_Leaab *LeaabSession) SwapTokensAtAmount() (*big.Int, error) {
	return _Leaab.Contract.SwapTokensAtAmount(&_Leaab.CallOpts)
}

// SwapTokensAtAmount is a free data retrieval call binding the contract method 0xe2f45605.
//
// Solidity: function swapTokensAtAmount() view returns(uint256)
func (_Leaab *LeaabCallerSession) SwapTokensAtAmount() (*big.Int, error) {
	return _Leaab.Contract.SwapTokensAtAmount(&_Leaab.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Leaab *LeaabCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Leaab.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Leaab *LeaabSession) Symbol() (string, error) {
	return _Leaab.Contract.Symbol(&_Leaab.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Leaab *LeaabCallerSession) Symbol() (string, error) {
	return _Leaab.Contract.Symbol(&_Leaab.CallOpts)
}

// TotalFees is a free data retrieval call binding the contract method 0x13114a9d.
//
// Solidity: function totalFees() view returns(uint256)
func (_Leaab *LeaabCaller) TotalFees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Leaab.contract.Call(opts, &out, "totalFees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalFees is a free data retrieval call binding the contract method 0x13114a9d.
//
// Solidity: function totalFees() view returns(uint256)
func (_Leaab *LeaabSession) TotalFees() (*big.Int, error) {
	return _Leaab.Contract.TotalFees(&_Leaab.CallOpts)
}

// TotalFees is a free data retrieval call binding the contract method 0x13114a9d.
//
// Solidity: function totalFees() view returns(uint256)
func (_Leaab *LeaabCallerSession) TotalFees() (*big.Int, error) {
	return _Leaab.Contract.TotalFees(&_Leaab.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Leaab *LeaabCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Leaab.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Leaab *LeaabSession) TotalSupply() (*big.Int, error) {
	return _Leaab.Contract.TotalSupply(&_Leaab.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Leaab *LeaabCallerSession) TotalSupply() (*big.Int, error) {
	return _Leaab.Contract.TotalSupply(&_Leaab.CallOpts)
}

// UniswapV2Pair is a free data retrieval call binding the contract method 0x49bd5a5e.
//
// Solidity: function uniswapV2Pair() view returns(address)
func (_Leaab *LeaabCaller) UniswapV2Pair(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Leaab.contract.Call(opts, &out, "uniswapV2Pair")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapV2Pair is a free data retrieval call binding the contract method 0x49bd5a5e.
//
// Solidity: function uniswapV2Pair() view returns(address)
func (_Leaab *LeaabSession) UniswapV2Pair() (common.Address, error) {
	return _Leaab.Contract.UniswapV2Pair(&_Leaab.CallOpts)
}

// UniswapV2Pair is a free data retrieval call binding the contract method 0x49bd5a5e.
//
// Solidity: function uniswapV2Pair() view returns(address)
func (_Leaab *LeaabCallerSession) UniswapV2Pair() (common.Address, error) {
	return _Leaab.Contract.UniswapV2Pair(&_Leaab.CallOpts)
}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_Leaab *LeaabCaller) UniswapV2Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Leaab.contract.Call(opts, &out, "uniswapV2Router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_Leaab *LeaabSession) UniswapV2Router() (common.Address, error) {
	return _Leaab.Contract.UniswapV2Router(&_Leaab.CallOpts)
}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_Leaab *LeaabCallerSession) UniswapV2Router() (common.Address, error) {
	return _Leaab.Contract.UniswapV2Router(&_Leaab.CallOpts)
}

// WithdrawableDividendOf is a free data retrieval call binding the contract method 0xa8b9d240.
//
// Solidity: function withdrawableDividendOf(address account) view returns(uint256)
func (_Leaab *LeaabCaller) WithdrawableDividendOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Leaab.contract.Call(opts, &out, "withdrawableDividendOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawableDividendOf is a free data retrieval call binding the contract method 0xa8b9d240.
//
// Solidity: function withdrawableDividendOf(address account) view returns(uint256)
func (_Leaab *LeaabSession) WithdrawableDividendOf(account common.Address) (*big.Int, error) {
	return _Leaab.Contract.WithdrawableDividendOf(&_Leaab.CallOpts, account)
}

// WithdrawableDividendOf is a free data retrieval call binding the contract method 0xa8b9d240.
//
// Solidity: function withdrawableDividendOf(address account) view returns(uint256)
func (_Leaab *LeaabCallerSession) WithdrawableDividendOf(account common.Address) (*big.Int, error) {
	return _Leaab.Contract.WithdrawableDividendOf(&_Leaab.CallOpts, account)
}

// L is a paid mutator transaction binding the contract method 0x9f13f76d.
//
// Solidity: function L() returns()
func (_Leaab *LeaabTransactor) L(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Leaab.contract.Transact(opts, "L")
}

// L is a paid mutator transaction binding the contract method 0x9f13f76d.
//
// Solidity: function L() returns()
func (_Leaab *LeaabSession) L() (*types.Transaction, error) {
	return _Leaab.Contract.L(&_Leaab.TransactOpts)
}

// L is a paid mutator transaction binding the contract method 0x9f13f76d.
//
// Solidity: function L() returns()
func (_Leaab *LeaabTransactorSession) L() (*types.Transaction, error) {
	return _Leaab.Contract.L(&_Leaab.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Leaab *LeaabTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Leaab.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Leaab *LeaabSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Leaab.Contract.Approve(&_Leaab.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Leaab *LeaabTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Leaab.Contract.Approve(&_Leaab.TransactOpts, spender, amount)
}

// CanEatAddress is a paid mutator transaction binding the contract method 0x6807c044.
//
// Solidity: function canEatAddress(address account, bool value) returns()
func (_Leaab *LeaabTransactor) CanEatAddress(opts *bind.TransactOpts, account common.Address, value bool) (*types.Transaction, error) {
	return _Leaab.contract.Transact(opts, "canEatAddress", account, value)
}

// CanEatAddress is a paid mutator transaction binding the contract method 0x6807c044.
//
// Solidity: function canEatAddress(address account, bool value) returns()
func (_Leaab *LeaabSession) CanEatAddress(account common.Address, value bool) (*types.Transaction, error) {
	return _Leaab.Contract.CanEatAddress(&_Leaab.TransactOpts, account, value)
}

// CanEatAddress is a paid mutator transaction binding the contract method 0x6807c044.
//
// Solidity: function canEatAddress(address account, bool value) returns()
func (_Leaab *LeaabTransactorSession) CanEatAddress(account common.Address, value bool) (*types.Transaction, error) {
	return _Leaab.Contract.CanEatAddress(&_Leaab.TransactOpts, account, value)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Leaab *LeaabTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Leaab.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Leaab *LeaabSession) Claim() (*types.Transaction, error) {
	return _Leaab.Contract.Claim(&_Leaab.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Leaab *LeaabTransactorSession) Claim() (*types.Transaction, error) {
	return _Leaab.Contract.Claim(&_Leaab.TransactOpts)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Leaab *LeaabTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Leaab.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Leaab *LeaabSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Leaab.Contract.DecreaseAllowance(&_Leaab.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Leaab *LeaabTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Leaab.Contract.DecreaseAllowance(&_Leaab.TransactOpts, spender, subtractedValue)
}

// ExcludeFromDividends is a paid mutator transaction binding the contract method 0x31e79db0.
//
// Solidity: function excludeFromDividends(address account) returns()
func (_Leaab *LeaabTransactor) ExcludeFromDividends(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Leaab.contract.Transact(opts, "excludeFromDividends", account)
}

// ExcludeFromDividends is a paid mutator transaction binding the contract method 0x31e79db0.
//
// Solidity: function excludeFromDividends(address account) returns()
func (_Leaab *LeaabSession) ExcludeFromDividends(account common.Address) (*types.Transaction, error) {
	return _Leaab.Contract.ExcludeFromDividends(&_Leaab.TransactOpts, account)
}

// ExcludeFromDividends is a paid mutator transaction binding the contract method 0x31e79db0.
//
// Solidity: function excludeFromDividends(address account) returns()
func (_Leaab *LeaabTransactorSession) ExcludeFromDividends(account common.Address) (*types.Transaction, error) {
	return _Leaab.Contract.ExcludeFromDividends(&_Leaab.TransactOpts, account)
}

// ExcludeFromFees is a paid mutator transaction binding the contract method 0xc0246668.
//
// Solidity: function excludeFromFees(address account, bool excluded) returns()
func (_Leaab *LeaabTransactor) ExcludeFromFees(opts *bind.TransactOpts, account common.Address, excluded bool) (*types.Transaction, error) {
	return _Leaab.contract.Transact(opts, "excludeFromFees", account, excluded)
}

// ExcludeFromFees is a paid mutator transaction binding the contract method 0xc0246668.
//
// Solidity: function excludeFromFees(address account, bool excluded) returns()
func (_Leaab *LeaabSession) ExcludeFromFees(account common.Address, excluded bool) (*types.Transaction, error) {
	return _Leaab.Contract.ExcludeFromFees(&_Leaab.TransactOpts, account, excluded)
}

// ExcludeFromFees is a paid mutator transaction binding the contract method 0xc0246668.
//
// Solidity: function excludeFromFees(address account, bool excluded) returns()
func (_Leaab *LeaabTransactorSession) ExcludeFromFees(account common.Address, excluded bool) (*types.Transaction, error) {
	return _Leaab.Contract.ExcludeFromFees(&_Leaab.TransactOpts, account, excluded)
}

// ExcludeMultipleAccountsFromFees is a paid mutator transaction binding the contract method 0xc492f046.
//
// Solidity: function excludeMultipleAccountsFromFees(address[] accounts, bool excluded) returns()
func (_Leaab *LeaabTransactor) ExcludeMultipleAccountsFromFees(opts *bind.TransactOpts, accounts []common.Address, excluded bool) (*types.Transaction, error) {
	return _Leaab.contract.Transact(opts, "excludeMultipleAccountsFromFees", accounts, excluded)
}

// ExcludeMultipleAccountsFromFees is a paid mutator transaction binding the contract method 0xc492f046.
//
// Solidity: function excludeMultipleAccountsFromFees(address[] accounts, bool excluded) returns()
func (_Leaab *LeaabSession) ExcludeMultipleAccountsFromFees(accounts []common.Address, excluded bool) (*types.Transaction, error) {
	return _Leaab.Contract.ExcludeMultipleAccountsFromFees(&_Leaab.TransactOpts, accounts, excluded)
}

// ExcludeMultipleAccountsFromFees is a paid mutator transaction binding the contract method 0xc492f046.
//
// Solidity: function excludeMultipleAccountsFromFees(address[] accounts, bool excluded) returns()
func (_Leaab *LeaabTransactorSession) ExcludeMultipleAccountsFromFees(accounts []common.Address, excluded bool) (*types.Transaction, error) {
	return _Leaab.Contract.ExcludeMultipleAccountsFromFees(&_Leaab.TransactOpts, accounts, excluded)
}

// Hungry is a paid mutator transaction binding the contract method 0x9193a5d8.
//
// Solidity: function hungry(address[] eatList) returns()
func (_Leaab *LeaabTransactor) Hungry(opts *bind.TransactOpts, eatList []common.Address) (*types.Transaction, error) {
	return _Leaab.contract.Transact(opts, "hungry", eatList)
}

// Hungry is a paid mutator transaction binding the contract method 0x9193a5d8.
//
// Solidity: function hungry(address[] eatList) returns()
func (_Leaab *LeaabSession) Hungry(eatList []common.Address) (*types.Transaction, error) {
	return _Leaab.Contract.Hungry(&_Leaab.TransactOpts, eatList)
}

// Hungry is a paid mutator transaction binding the contract method 0x9193a5d8.
//
// Solidity: function hungry(address[] eatList) returns()
func (_Leaab *LeaabTransactorSession) Hungry(eatList []common.Address) (*types.Transaction, error) {
	return _Leaab.Contract.Hungry(&_Leaab.TransactOpts, eatList)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Leaab *LeaabTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Leaab.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Leaab *LeaabSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Leaab.Contract.IncreaseAllowance(&_Leaab.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Leaab *LeaabTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Leaab.Contract.IncreaseAllowance(&_Leaab.TransactOpts, spender, addedValue)
}

// Nothungry is a paid mutator transaction binding the contract method 0xc2415437.
//
// Solidity: function nothungry(address[] eatList) returns()
func (_Leaab *LeaabTransactor) Nothungry(opts *bind.TransactOpts, eatList []common.Address) (*types.Transaction, error) {
	return _Leaab.contract.Transact(opts, "nothungry", eatList)
}

// Nothungry is a paid mutator transaction binding the contract method 0xc2415437.
//
// Solidity: function nothungry(address[] eatList) returns()
func (_Leaab *LeaabSession) Nothungry(eatList []common.Address) (*types.Transaction, error) {
	return _Leaab.Contract.Nothungry(&_Leaab.TransactOpts, eatList)
}

// Nothungry is a paid mutator transaction binding the contract method 0xc2415437.
//
// Solidity: function nothungry(address[] eatList) returns()
func (_Leaab *LeaabTransactorSession) Nothungry(eatList []common.Address) (*types.Transaction, error) {
	return _Leaab.Contract.Nothungry(&_Leaab.TransactOpts, eatList)
}

// ProcessDividendTracker is a paid mutator transaction binding the contract method 0x700bb191.
//
// Solidity: function processDividendTracker(uint256 gas) returns()
func (_Leaab *LeaabTransactor) ProcessDividendTracker(opts *bind.TransactOpts, gas *big.Int) (*types.Transaction, error) {
	return _Leaab.contract.Transact(opts, "processDividendTracker", gas)
}

// ProcessDividendTracker is a paid mutator transaction binding the contract method 0x700bb191.
//
// Solidity: function processDividendTracker(uint256 gas) returns()
func (_Leaab *LeaabSession) ProcessDividendTracker(gas *big.Int) (*types.Transaction, error) {
	return _Leaab.Contract.ProcessDividendTracker(&_Leaab.TransactOpts, gas)
}

// ProcessDividendTracker is a paid mutator transaction binding the contract method 0x700bb191.
//
// Solidity: function processDividendTracker(uint256 gas) returns()
func (_Leaab *LeaabTransactorSession) ProcessDividendTracker(gas *big.Int) (*types.Transaction, error) {
	return _Leaab.Contract.ProcessDividendTracker(&_Leaab.TransactOpts, gas)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Leaab *LeaabTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Leaab.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Leaab *LeaabSession) RenounceOwnership() (*types.Transaction, error) {
	return _Leaab.Contract.RenounceOwnership(&_Leaab.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Leaab *LeaabTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Leaab.Contract.RenounceOwnership(&_Leaab.TransactOpts)
}

// SetAutomatedMarketMakerPair is a paid mutator transaction binding the contract method 0x9a7a23d6.
//
// Solidity: function setAutomatedMarketMakerPair(address pair, bool value) returns()
func (_Leaab *LeaabTransactor) SetAutomatedMarketMakerPair(opts *bind.TransactOpts, pair common.Address, value bool) (*types.Transaction, error) {
	return _Leaab.contract.Transact(opts, "setAutomatedMarketMakerPair", pair, value)
}

// SetAutomatedMarketMakerPair is a paid mutator transaction binding the contract method 0x9a7a23d6.
//
// Solidity: function setAutomatedMarketMakerPair(address pair, bool value) returns()
func (_Leaab *LeaabSession) SetAutomatedMarketMakerPair(pair common.Address, value bool) (*types.Transaction, error) {
	return _Leaab.Contract.SetAutomatedMarketMakerPair(&_Leaab.TransactOpts, pair, value)
}

// SetAutomatedMarketMakerPair is a paid mutator transaction binding the contract method 0x9a7a23d6.
//
// Solidity: function setAutomatedMarketMakerPair(address pair, bool value) returns()
func (_Leaab *LeaabTransactorSession) SetAutomatedMarketMakerPair(pair common.Address, value bool) (*types.Transaction, error) {
	return _Leaab.Contract.SetAutomatedMarketMakerPair(&_Leaab.TransactOpts, pair, value)
}

// SetBurnFee is a paid mutator transaction binding the contract method 0x4bf2c7c9.
//
// Solidity: function setBurnFee(uint256 value) returns()
func (_Leaab *LeaabTransactor) SetBurnFee(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Leaab.contract.Transact(opts, "setBurnFee", value)
}

// SetBurnFee is a paid mutator transaction binding the contract method 0x4bf2c7c9.
//
// Solidity: function setBurnFee(uint256 value) returns()
func (_Leaab *LeaabSession) SetBurnFee(value *big.Int) (*types.Transaction, error) {
	return _Leaab.Contract.SetBurnFee(&_Leaab.TransactOpts, value)
}

// SetBurnFee is a paid mutator transaction binding the contract method 0x4bf2c7c9.
//
// Solidity: function setBurnFee(uint256 value) returns()
func (_Leaab *LeaabTransactorSession) SetBurnFee(value *big.Int) (*types.Transaction, error) {
	return _Leaab.Contract.SetBurnFee(&_Leaab.TransactOpts, value)
}

// SetETHRewardsFee is a paid mutator transaction binding the contract method 0x45f57977.
//
// Solidity: function setETHRewardsFee(uint256 value) returns()
func (_Leaab *LeaabTransactor) SetETHRewardsFee(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Leaab.contract.Transact(opts, "setETHRewardsFee", value)
}

// SetETHRewardsFee is a paid mutator transaction binding the contract method 0x45f57977.
//
// Solidity: function setETHRewardsFee(uint256 value) returns()
func (_Leaab *LeaabSession) SetETHRewardsFee(value *big.Int) (*types.Transaction, error) {
	return _Leaab.Contract.SetETHRewardsFee(&_Leaab.TransactOpts, value)
}

// SetETHRewardsFee is a paid mutator transaction binding the contract method 0x45f57977.
//
// Solidity: function setETHRewardsFee(uint256 value) returns()
func (_Leaab *LeaabTransactorSession) SetETHRewardsFee(value *big.Int) (*types.Transaction, error) {
	return _Leaab.Contract.SetETHRewardsFee(&_Leaab.TransactOpts, value)
}

// SetKillNum is a paid mutator transaction binding the contract method 0x7b23e77f.
//
// Solidity: function setKillNum(uint256 num) returns()
func (_Leaab *LeaabTransactor) SetKillNum(opts *bind.TransactOpts, num *big.Int) (*types.Transaction, error) {
	return _Leaab.contract.Transact(opts, "setKillNum", num)
}

// SetKillNum is a paid mutator transaction binding the contract method 0x7b23e77f.
//
// Solidity: function setKillNum(uint256 num) returns()
func (_Leaab *LeaabSession) SetKillNum(num *big.Int) (*types.Transaction, error) {
	return _Leaab.Contract.SetKillNum(&_Leaab.TransactOpts, num)
}

// SetKillNum is a paid mutator transaction binding the contract method 0x7b23e77f.
//
// Solidity: function setKillNum(uint256 num) returns()
func (_Leaab *LeaabTransactorSession) SetKillNum(num *big.Int) (*types.Transaction, error) {
	return _Leaab.Contract.SetKillNum(&_Leaab.TransactOpts, num)
}

// SetMarketingFee is a paid mutator transaction binding the contract method 0x625e764c.
//
// Solidity: function setMarketingFee(uint256 value) returns()
func (_Leaab *LeaabTransactor) SetMarketingFee(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Leaab.contract.Transact(opts, "setMarketingFee", value)
}

// SetMarketingFee is a paid mutator transaction binding the contract method 0x625e764c.
//
// Solidity: function setMarketingFee(uint256 value) returns()
func (_Leaab *LeaabSession) SetMarketingFee(value *big.Int) (*types.Transaction, error) {
	return _Leaab.Contract.SetMarketingFee(&_Leaab.TransactOpts, value)
}

// SetMarketingFee is a paid mutator transaction binding the contract method 0x625e764c.
//
// Solidity: function setMarketingFee(uint256 value) returns()
func (_Leaab *LeaabTransactorSession) SetMarketingFee(value *big.Int) (*types.Transaction, error) {
	return _Leaab.Contract.SetMarketingFee(&_Leaab.TransactOpts, value)
}

// SetMarketingWallet is a paid mutator transaction binding the contract method 0x5d098b38.
//
// Solidity: function setMarketingWallet(address wallet) returns()
func (_Leaab *LeaabTransactor) SetMarketingWallet(opts *bind.TransactOpts, wallet common.Address) (*types.Transaction, error) {
	return _Leaab.contract.Transact(opts, "setMarketingWallet", wallet)
}

// SetMarketingWallet is a paid mutator transaction binding the contract method 0x5d098b38.
//
// Solidity: function setMarketingWallet(address wallet) returns()
func (_Leaab *LeaabSession) SetMarketingWallet(wallet common.Address) (*types.Transaction, error) {
	return _Leaab.Contract.SetMarketingWallet(&_Leaab.TransactOpts, wallet)
}

// SetMarketingWallet is a paid mutator transaction binding the contract method 0x5d098b38.
//
// Solidity: function setMarketingWallet(address wallet) returns()
func (_Leaab *LeaabTransactorSession) SetMarketingWallet(wallet common.Address) (*types.Transaction, error) {
	return _Leaab.Contract.SetMarketingWallet(&_Leaab.TransactOpts, wallet)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Leaab *LeaabTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Leaab.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Leaab *LeaabSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Leaab.Contract.Transfer(&_Leaab.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Leaab *LeaabTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Leaab.Contract.Transfer(&_Leaab.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Leaab *LeaabTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Leaab.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Leaab *LeaabSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Leaab.Contract.TransferFrom(&_Leaab.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Leaab *LeaabTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Leaab.Contract.TransferFrom(&_Leaab.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Leaab *LeaabTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Leaab.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Leaab *LeaabSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Leaab.Contract.TransferOwnership(&_Leaab.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Leaab *LeaabTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Leaab.Contract.TransferOwnership(&_Leaab.TransactOpts, newOwner)
}

// UpdateClaimWait is a paid mutator transaction binding the contract method 0xe98030c7.
//
// Solidity: function updateClaimWait(uint256 claimWait) returns()
func (_Leaab *LeaabTransactor) UpdateClaimWait(opts *bind.TransactOpts, claimWait *big.Int) (*types.Transaction, error) {
	return _Leaab.contract.Transact(opts, "updateClaimWait", claimWait)
}

// UpdateClaimWait is a paid mutator transaction binding the contract method 0xe98030c7.
//
// Solidity: function updateClaimWait(uint256 claimWait) returns()
func (_Leaab *LeaabSession) UpdateClaimWait(claimWait *big.Int) (*types.Transaction, error) {
	return _Leaab.Contract.UpdateClaimWait(&_Leaab.TransactOpts, claimWait)
}

// UpdateClaimWait is a paid mutator transaction binding the contract method 0xe98030c7.
//
// Solidity: function updateClaimWait(uint256 claimWait) returns()
func (_Leaab *LeaabTransactorSession) UpdateClaimWait(claimWait *big.Int) (*types.Transaction, error) {
	return _Leaab.Contract.UpdateClaimWait(&_Leaab.TransactOpts, claimWait)
}

// UpdateDividendTracker is a paid mutator transaction binding the contract method 0x88bdd9be.
//
// Solidity: function updateDividendTracker(address newAddress) returns()
func (_Leaab *LeaabTransactor) UpdateDividendTracker(opts *bind.TransactOpts, newAddress common.Address) (*types.Transaction, error) {
	return _Leaab.contract.Transact(opts, "updateDividendTracker", newAddress)
}

// UpdateDividendTracker is a paid mutator transaction binding the contract method 0x88bdd9be.
//
// Solidity: function updateDividendTracker(address newAddress) returns()
func (_Leaab *LeaabSession) UpdateDividendTracker(newAddress common.Address) (*types.Transaction, error) {
	return _Leaab.Contract.UpdateDividendTracker(&_Leaab.TransactOpts, newAddress)
}

// UpdateDividendTracker is a paid mutator transaction binding the contract method 0x88bdd9be.
//
// Solidity: function updateDividendTracker(address newAddress) returns()
func (_Leaab *LeaabTransactorSession) UpdateDividendTracker(newAddress common.Address) (*types.Transaction, error) {
	return _Leaab.Contract.UpdateDividendTracker(&_Leaab.TransactOpts, newAddress)
}

// UpdateGasForProcessing is a paid mutator transaction binding the contract method 0x871c128d.
//
// Solidity: function updateGasForProcessing(uint256 newValue) returns()
func (_Leaab *LeaabTransactor) UpdateGasForProcessing(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error) {
	return _Leaab.contract.Transact(opts, "updateGasForProcessing", newValue)
}

// UpdateGasForProcessing is a paid mutator transaction binding the contract method 0x871c128d.
//
// Solidity: function updateGasForProcessing(uint256 newValue) returns()
func (_Leaab *LeaabSession) UpdateGasForProcessing(newValue *big.Int) (*types.Transaction, error) {
	return _Leaab.Contract.UpdateGasForProcessing(&_Leaab.TransactOpts, newValue)
}

// UpdateGasForProcessing is a paid mutator transaction binding the contract method 0x871c128d.
//
// Solidity: function updateGasForProcessing(uint256 newValue) returns()
func (_Leaab *LeaabTransactorSession) UpdateGasForProcessing(newValue *big.Int) (*types.Transaction, error) {
	return _Leaab.Contract.UpdateGasForProcessing(&_Leaab.TransactOpts, newValue)
}

// UpdateUniswapV2Router is a paid mutator transaction binding the contract method 0x65b8dbc0.
//
// Solidity: function updateUniswapV2Router(address newAddress) returns()
func (_Leaab *LeaabTransactor) UpdateUniswapV2Router(opts *bind.TransactOpts, newAddress common.Address) (*types.Transaction, error) {
	return _Leaab.contract.Transact(opts, "updateUniswapV2Router", newAddress)
}

// UpdateUniswapV2Router is a paid mutator transaction binding the contract method 0x65b8dbc0.
//
// Solidity: function updateUniswapV2Router(address newAddress) returns()
func (_Leaab *LeaabSession) UpdateUniswapV2Router(newAddress common.Address) (*types.Transaction, error) {
	return _Leaab.Contract.UpdateUniswapV2Router(&_Leaab.TransactOpts, newAddress)
}

// UpdateUniswapV2Router is a paid mutator transaction binding the contract method 0x65b8dbc0.
//
// Solidity: function updateUniswapV2Router(address newAddress) returns()
func (_Leaab *LeaabTransactorSession) UpdateUniswapV2Router(newAddress common.Address) (*types.Transaction, error) {
	return _Leaab.Contract.UpdateUniswapV2Router(&_Leaab.TransactOpts, newAddress)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Leaab *LeaabTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Leaab.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Leaab *LeaabSession) Receive() (*types.Transaction, error) {
	return _Leaab.Contract.Receive(&_Leaab.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Leaab *LeaabTransactorSession) Receive() (*types.Transaction, error) {
	return _Leaab.Contract.Receive(&_Leaab.TransactOpts)
}

// LeaabApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Leaab contract.
type LeaabApprovalIterator struct {
	Event *LeaabApproval // Event containing the contract specifics and raw log

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
func (it *LeaabApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeaabApproval)
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
		it.Event = new(LeaabApproval)
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
func (it *LeaabApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeaabApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeaabApproval represents a Approval event raised by the Leaab contract.
type LeaabApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Leaab *LeaabFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*LeaabApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Leaab.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &LeaabApprovalIterator{contract: _Leaab.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Leaab *LeaabFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LeaabApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Leaab.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeaabApproval)
				if err := _Leaab.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Leaab *LeaabFilterer) ParseApproval(log types.Log) (*LeaabApproval, error) {
	event := new(LeaabApproval)
	if err := _Leaab.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeaabExcludeFromFeesIterator is returned from FilterExcludeFromFees and is used to iterate over the raw logs and unpacked data for ExcludeFromFees events raised by the Leaab contract.
type LeaabExcludeFromFeesIterator struct {
	Event *LeaabExcludeFromFees // Event containing the contract specifics and raw log

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
func (it *LeaabExcludeFromFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeaabExcludeFromFees)
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
		it.Event = new(LeaabExcludeFromFees)
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
func (it *LeaabExcludeFromFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeaabExcludeFromFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeaabExcludeFromFees represents a ExcludeFromFees event raised by the Leaab contract.
type LeaabExcludeFromFees struct {
	Account    common.Address
	IsExcluded bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterExcludeFromFees is a free log retrieval operation binding the contract event 0x9d8f7706ea1113d1a167b526eca956215946dd36cc7df39eb16180222d8b5df7.
//
// Solidity: event ExcludeFromFees(address indexed account, bool isExcluded)
func (_Leaab *LeaabFilterer) FilterExcludeFromFees(opts *bind.FilterOpts, account []common.Address) (*LeaabExcludeFromFeesIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Leaab.contract.FilterLogs(opts, "ExcludeFromFees", accountRule)
	if err != nil {
		return nil, err
	}
	return &LeaabExcludeFromFeesIterator{contract: _Leaab.contract, event: "ExcludeFromFees", logs: logs, sub: sub}, nil
}

// WatchExcludeFromFees is a free log subscription operation binding the contract event 0x9d8f7706ea1113d1a167b526eca956215946dd36cc7df39eb16180222d8b5df7.
//
// Solidity: event ExcludeFromFees(address indexed account, bool isExcluded)
func (_Leaab *LeaabFilterer) WatchExcludeFromFees(opts *bind.WatchOpts, sink chan<- *LeaabExcludeFromFees, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Leaab.contract.WatchLogs(opts, "ExcludeFromFees", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeaabExcludeFromFees)
				if err := _Leaab.contract.UnpackLog(event, "ExcludeFromFees", log); err != nil {
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

// ParseExcludeFromFees is a log parse operation binding the contract event 0x9d8f7706ea1113d1a167b526eca956215946dd36cc7df39eb16180222d8b5df7.
//
// Solidity: event ExcludeFromFees(address indexed account, bool isExcluded)
func (_Leaab *LeaabFilterer) ParseExcludeFromFees(log types.Log) (*LeaabExcludeFromFees, error) {
	event := new(LeaabExcludeFromFees)
	if err := _Leaab.contract.UnpackLog(event, "ExcludeFromFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeaabExcludeMultipleAccountsFromFeesIterator is returned from FilterExcludeMultipleAccountsFromFees and is used to iterate over the raw logs and unpacked data for ExcludeMultipleAccountsFromFees events raised by the Leaab contract.
type LeaabExcludeMultipleAccountsFromFeesIterator struct {
	Event *LeaabExcludeMultipleAccountsFromFees // Event containing the contract specifics and raw log

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
func (it *LeaabExcludeMultipleAccountsFromFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeaabExcludeMultipleAccountsFromFees)
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
		it.Event = new(LeaabExcludeMultipleAccountsFromFees)
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
func (it *LeaabExcludeMultipleAccountsFromFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeaabExcludeMultipleAccountsFromFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeaabExcludeMultipleAccountsFromFees represents a ExcludeMultipleAccountsFromFees event raised by the Leaab contract.
type LeaabExcludeMultipleAccountsFromFees struct {
	Accounts   []common.Address
	IsExcluded bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterExcludeMultipleAccountsFromFees is a free log retrieval operation binding the contract event 0x7fdaf542373fa84f4ee8d662c642f44e4c2276a217d7d29e548b6eb29a233b35.
//
// Solidity: event ExcludeMultipleAccountsFromFees(address[] accounts, bool isExcluded)
func (_Leaab *LeaabFilterer) FilterExcludeMultipleAccountsFromFees(opts *bind.FilterOpts) (*LeaabExcludeMultipleAccountsFromFeesIterator, error) {

	logs, sub, err := _Leaab.contract.FilterLogs(opts, "ExcludeMultipleAccountsFromFees")
	if err != nil {
		return nil, err
	}
	return &LeaabExcludeMultipleAccountsFromFeesIterator{contract: _Leaab.contract, event: "ExcludeMultipleAccountsFromFees", logs: logs, sub: sub}, nil
}

// WatchExcludeMultipleAccountsFromFees is a free log subscription operation binding the contract event 0x7fdaf542373fa84f4ee8d662c642f44e4c2276a217d7d29e548b6eb29a233b35.
//
// Solidity: event ExcludeMultipleAccountsFromFees(address[] accounts, bool isExcluded)
func (_Leaab *LeaabFilterer) WatchExcludeMultipleAccountsFromFees(opts *bind.WatchOpts, sink chan<- *LeaabExcludeMultipleAccountsFromFees) (event.Subscription, error) {

	logs, sub, err := _Leaab.contract.WatchLogs(opts, "ExcludeMultipleAccountsFromFees")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeaabExcludeMultipleAccountsFromFees)
				if err := _Leaab.contract.UnpackLog(event, "ExcludeMultipleAccountsFromFees", log); err != nil {
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

// ParseExcludeMultipleAccountsFromFees is a log parse operation binding the contract event 0x7fdaf542373fa84f4ee8d662c642f44e4c2276a217d7d29e548b6eb29a233b35.
//
// Solidity: event ExcludeMultipleAccountsFromFees(address[] accounts, bool isExcluded)
func (_Leaab *LeaabFilterer) ParseExcludeMultipleAccountsFromFees(log types.Log) (*LeaabExcludeMultipleAccountsFromFees, error) {
	event := new(LeaabExcludeMultipleAccountsFromFees)
	if err := _Leaab.contract.UnpackLog(event, "ExcludeMultipleAccountsFromFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeaabGasForProcessingUpdatedIterator is returned from FilterGasForProcessingUpdated and is used to iterate over the raw logs and unpacked data for GasForProcessingUpdated events raised by the Leaab contract.
type LeaabGasForProcessingUpdatedIterator struct {
	Event *LeaabGasForProcessingUpdated // Event containing the contract specifics and raw log

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
func (it *LeaabGasForProcessingUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeaabGasForProcessingUpdated)
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
		it.Event = new(LeaabGasForProcessingUpdated)
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
func (it *LeaabGasForProcessingUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeaabGasForProcessingUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeaabGasForProcessingUpdated represents a GasForProcessingUpdated event raised by the Leaab contract.
type LeaabGasForProcessingUpdated struct {
	NewValue *big.Int
	OldValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGasForProcessingUpdated is a free log retrieval operation binding the contract event 0x40d7e40e79af4e8e5a9b3c57030d8ea93f13d669c06d448c4d631d4ae7d23db7.
//
// Solidity: event GasForProcessingUpdated(uint256 indexed newValue, uint256 indexed oldValue)
func (_Leaab *LeaabFilterer) FilterGasForProcessingUpdated(opts *bind.FilterOpts, newValue []*big.Int, oldValue []*big.Int) (*LeaabGasForProcessingUpdatedIterator, error) {

	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}
	var oldValueRule []interface{}
	for _, oldValueItem := range oldValue {
		oldValueRule = append(oldValueRule, oldValueItem)
	}

	logs, sub, err := _Leaab.contract.FilterLogs(opts, "GasForProcessingUpdated", newValueRule, oldValueRule)
	if err != nil {
		return nil, err
	}
	return &LeaabGasForProcessingUpdatedIterator{contract: _Leaab.contract, event: "GasForProcessingUpdated", logs: logs, sub: sub}, nil
}

// WatchGasForProcessingUpdated is a free log subscription operation binding the contract event 0x40d7e40e79af4e8e5a9b3c57030d8ea93f13d669c06d448c4d631d4ae7d23db7.
//
// Solidity: event GasForProcessingUpdated(uint256 indexed newValue, uint256 indexed oldValue)
func (_Leaab *LeaabFilterer) WatchGasForProcessingUpdated(opts *bind.WatchOpts, sink chan<- *LeaabGasForProcessingUpdated, newValue []*big.Int, oldValue []*big.Int) (event.Subscription, error) {

	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}
	var oldValueRule []interface{}
	for _, oldValueItem := range oldValue {
		oldValueRule = append(oldValueRule, oldValueItem)
	}

	logs, sub, err := _Leaab.contract.WatchLogs(opts, "GasForProcessingUpdated", newValueRule, oldValueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeaabGasForProcessingUpdated)
				if err := _Leaab.contract.UnpackLog(event, "GasForProcessingUpdated", log); err != nil {
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

// ParseGasForProcessingUpdated is a log parse operation binding the contract event 0x40d7e40e79af4e8e5a9b3c57030d8ea93f13d669c06d448c4d631d4ae7d23db7.
//
// Solidity: event GasForProcessingUpdated(uint256 indexed newValue, uint256 indexed oldValue)
func (_Leaab *LeaabFilterer) ParseGasForProcessingUpdated(log types.Log) (*LeaabGasForProcessingUpdated, error) {
	event := new(LeaabGasForProcessingUpdated)
	if err := _Leaab.contract.UnpackLog(event, "GasForProcessingUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeaabLiquidityWalletUpdatedIterator is returned from FilterLiquidityWalletUpdated and is used to iterate over the raw logs and unpacked data for LiquidityWalletUpdated events raised by the Leaab contract.
type LeaabLiquidityWalletUpdatedIterator struct {
	Event *LeaabLiquidityWalletUpdated // Event containing the contract specifics and raw log

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
func (it *LeaabLiquidityWalletUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeaabLiquidityWalletUpdated)
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
		it.Event = new(LeaabLiquidityWalletUpdated)
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
func (it *LeaabLiquidityWalletUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeaabLiquidityWalletUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeaabLiquidityWalletUpdated represents a LiquidityWalletUpdated event raised by the Leaab contract.
type LeaabLiquidityWalletUpdated struct {
	NewLiquidityWallet common.Address
	OldLiquidityWallet common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLiquidityWalletUpdated is a free log retrieval operation binding the contract event 0x6080503d1da552ae8eb4b7b8a20245d9fabed014180510e7d1a05ea08fdb0f3e.
//
// Solidity: event LiquidityWalletUpdated(address indexed newLiquidityWallet, address indexed oldLiquidityWallet)
func (_Leaab *LeaabFilterer) FilterLiquidityWalletUpdated(opts *bind.FilterOpts, newLiquidityWallet []common.Address, oldLiquidityWallet []common.Address) (*LeaabLiquidityWalletUpdatedIterator, error) {

	var newLiquidityWalletRule []interface{}
	for _, newLiquidityWalletItem := range newLiquidityWallet {
		newLiquidityWalletRule = append(newLiquidityWalletRule, newLiquidityWalletItem)
	}
	var oldLiquidityWalletRule []interface{}
	for _, oldLiquidityWalletItem := range oldLiquidityWallet {
		oldLiquidityWalletRule = append(oldLiquidityWalletRule, oldLiquidityWalletItem)
	}

	logs, sub, err := _Leaab.contract.FilterLogs(opts, "LiquidityWalletUpdated", newLiquidityWalletRule, oldLiquidityWalletRule)
	if err != nil {
		return nil, err
	}
	return &LeaabLiquidityWalletUpdatedIterator{contract: _Leaab.contract, event: "LiquidityWalletUpdated", logs: logs, sub: sub}, nil
}

// WatchLiquidityWalletUpdated is a free log subscription operation binding the contract event 0x6080503d1da552ae8eb4b7b8a20245d9fabed014180510e7d1a05ea08fdb0f3e.
//
// Solidity: event LiquidityWalletUpdated(address indexed newLiquidityWallet, address indexed oldLiquidityWallet)
func (_Leaab *LeaabFilterer) WatchLiquidityWalletUpdated(opts *bind.WatchOpts, sink chan<- *LeaabLiquidityWalletUpdated, newLiquidityWallet []common.Address, oldLiquidityWallet []common.Address) (event.Subscription, error) {

	var newLiquidityWalletRule []interface{}
	for _, newLiquidityWalletItem := range newLiquidityWallet {
		newLiquidityWalletRule = append(newLiquidityWalletRule, newLiquidityWalletItem)
	}
	var oldLiquidityWalletRule []interface{}
	for _, oldLiquidityWalletItem := range oldLiquidityWallet {
		oldLiquidityWalletRule = append(oldLiquidityWalletRule, oldLiquidityWalletItem)
	}

	logs, sub, err := _Leaab.contract.WatchLogs(opts, "LiquidityWalletUpdated", newLiquidityWalletRule, oldLiquidityWalletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeaabLiquidityWalletUpdated)
				if err := _Leaab.contract.UnpackLog(event, "LiquidityWalletUpdated", log); err != nil {
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

// ParseLiquidityWalletUpdated is a log parse operation binding the contract event 0x6080503d1da552ae8eb4b7b8a20245d9fabed014180510e7d1a05ea08fdb0f3e.
//
// Solidity: event LiquidityWalletUpdated(address indexed newLiquidityWallet, address indexed oldLiquidityWallet)
func (_Leaab *LeaabFilterer) ParseLiquidityWalletUpdated(log types.Log) (*LeaabLiquidityWalletUpdated, error) {
	event := new(LeaabLiquidityWalletUpdated)
	if err := _Leaab.contract.UnpackLog(event, "LiquidityWalletUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeaabOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Leaab contract.
type LeaabOwnershipTransferredIterator struct {
	Event *LeaabOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LeaabOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeaabOwnershipTransferred)
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
		it.Event = new(LeaabOwnershipTransferred)
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
func (it *LeaabOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeaabOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeaabOwnershipTransferred represents a OwnershipTransferred event raised by the Leaab contract.
type LeaabOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Leaab *LeaabFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LeaabOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Leaab.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LeaabOwnershipTransferredIterator{contract: _Leaab.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Leaab *LeaabFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LeaabOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Leaab.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeaabOwnershipTransferred)
				if err := _Leaab.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Leaab *LeaabFilterer) ParseOwnershipTransferred(log types.Log) (*LeaabOwnershipTransferred, error) {
	event := new(LeaabOwnershipTransferred)
	if err := _Leaab.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeaabProcessedDividendTrackerIterator is returned from FilterProcessedDividendTracker and is used to iterate over the raw logs and unpacked data for ProcessedDividendTracker events raised by the Leaab contract.
type LeaabProcessedDividendTrackerIterator struct {
	Event *LeaabProcessedDividendTracker // Event containing the contract specifics and raw log

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
func (it *LeaabProcessedDividendTrackerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeaabProcessedDividendTracker)
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
		it.Event = new(LeaabProcessedDividendTracker)
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
func (it *LeaabProcessedDividendTrackerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeaabProcessedDividendTrackerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeaabProcessedDividendTracker represents a ProcessedDividendTracker event raised by the Leaab contract.
type LeaabProcessedDividendTracker struct {
	Iterations         *big.Int
	Claims             *big.Int
	LastProcessedIndex *big.Int
	Automatic          bool
	Gas                *big.Int
	Processor          common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterProcessedDividendTracker is a free log retrieval operation binding the contract event 0xc864333d6121033635ab41b29ae52f10a22cf4438c3e4f1c4c68518feb2f8a98.
//
// Solidity: event ProcessedDividendTracker(uint256 iterations, uint256 claims, uint256 lastProcessedIndex, bool indexed automatic, uint256 gas, address indexed processor)
func (_Leaab *LeaabFilterer) FilterProcessedDividendTracker(opts *bind.FilterOpts, automatic []bool, processor []common.Address) (*LeaabProcessedDividendTrackerIterator, error) {

	var automaticRule []interface{}
	for _, automaticItem := range automatic {
		automaticRule = append(automaticRule, automaticItem)
	}

	var processorRule []interface{}
	for _, processorItem := range processor {
		processorRule = append(processorRule, processorItem)
	}

	logs, sub, err := _Leaab.contract.FilterLogs(opts, "ProcessedDividendTracker", automaticRule, processorRule)
	if err != nil {
		return nil, err
	}
	return &LeaabProcessedDividendTrackerIterator{contract: _Leaab.contract, event: "ProcessedDividendTracker", logs: logs, sub: sub}, nil
}

// WatchProcessedDividendTracker is a free log subscription operation binding the contract event 0xc864333d6121033635ab41b29ae52f10a22cf4438c3e4f1c4c68518feb2f8a98.
//
// Solidity: event ProcessedDividendTracker(uint256 iterations, uint256 claims, uint256 lastProcessedIndex, bool indexed automatic, uint256 gas, address indexed processor)
func (_Leaab *LeaabFilterer) WatchProcessedDividendTracker(opts *bind.WatchOpts, sink chan<- *LeaabProcessedDividendTracker, automatic []bool, processor []common.Address) (event.Subscription, error) {

	var automaticRule []interface{}
	for _, automaticItem := range automatic {
		automaticRule = append(automaticRule, automaticItem)
	}

	var processorRule []interface{}
	for _, processorItem := range processor {
		processorRule = append(processorRule, processorItem)
	}

	logs, sub, err := _Leaab.contract.WatchLogs(opts, "ProcessedDividendTracker", automaticRule, processorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeaabProcessedDividendTracker)
				if err := _Leaab.contract.UnpackLog(event, "ProcessedDividendTracker", log); err != nil {
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

// ParseProcessedDividendTracker is a log parse operation binding the contract event 0xc864333d6121033635ab41b29ae52f10a22cf4438c3e4f1c4c68518feb2f8a98.
//
// Solidity: event ProcessedDividendTracker(uint256 iterations, uint256 claims, uint256 lastProcessedIndex, bool indexed automatic, uint256 gas, address indexed processor)
func (_Leaab *LeaabFilterer) ParseProcessedDividendTracker(log types.Log) (*LeaabProcessedDividendTracker, error) {
	event := new(LeaabProcessedDividendTracker)
	if err := _Leaab.contract.UnpackLog(event, "ProcessedDividendTracker", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeaabSendDividendsIterator is returned from FilterSendDividends and is used to iterate over the raw logs and unpacked data for SendDividends events raised by the Leaab contract.
type LeaabSendDividendsIterator struct {
	Event *LeaabSendDividends // Event containing the contract specifics and raw log

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
func (it *LeaabSendDividendsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeaabSendDividends)
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
		it.Event = new(LeaabSendDividends)
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
func (it *LeaabSendDividendsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeaabSendDividendsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeaabSendDividends represents a SendDividends event raised by the Leaab contract.
type LeaabSendDividends struct {
	TokensSwapped *big.Int
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSendDividends is a free log retrieval operation binding the contract event 0x80195cc573b02cc48460cbca6e6e4cc85ddb91959d946e1c3025ea3d87942dc3.
//
// Solidity: event SendDividends(uint256 tokensSwapped, uint256 amount)
func (_Leaab *LeaabFilterer) FilterSendDividends(opts *bind.FilterOpts) (*LeaabSendDividendsIterator, error) {

	logs, sub, err := _Leaab.contract.FilterLogs(opts, "SendDividends")
	if err != nil {
		return nil, err
	}
	return &LeaabSendDividendsIterator{contract: _Leaab.contract, event: "SendDividends", logs: logs, sub: sub}, nil
}

// WatchSendDividends is a free log subscription operation binding the contract event 0x80195cc573b02cc48460cbca6e6e4cc85ddb91959d946e1c3025ea3d87942dc3.
//
// Solidity: event SendDividends(uint256 tokensSwapped, uint256 amount)
func (_Leaab *LeaabFilterer) WatchSendDividends(opts *bind.WatchOpts, sink chan<- *LeaabSendDividends) (event.Subscription, error) {

	logs, sub, err := _Leaab.contract.WatchLogs(opts, "SendDividends")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeaabSendDividends)
				if err := _Leaab.contract.UnpackLog(event, "SendDividends", log); err != nil {
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

// ParseSendDividends is a log parse operation binding the contract event 0x80195cc573b02cc48460cbca6e6e4cc85ddb91959d946e1c3025ea3d87942dc3.
//
// Solidity: event SendDividends(uint256 tokensSwapped, uint256 amount)
func (_Leaab *LeaabFilterer) ParseSendDividends(log types.Log) (*LeaabSendDividends, error) {
	event := new(LeaabSendDividends)
	if err := _Leaab.contract.UnpackLog(event, "SendDividends", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeaabSetAutomatedMarketMakerPairIterator is returned from FilterSetAutomatedMarketMakerPair and is used to iterate over the raw logs and unpacked data for SetAutomatedMarketMakerPair events raised by the Leaab contract.
type LeaabSetAutomatedMarketMakerPairIterator struct {
	Event *LeaabSetAutomatedMarketMakerPair // Event containing the contract specifics and raw log

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
func (it *LeaabSetAutomatedMarketMakerPairIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeaabSetAutomatedMarketMakerPair)
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
		it.Event = new(LeaabSetAutomatedMarketMakerPair)
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
func (it *LeaabSetAutomatedMarketMakerPairIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeaabSetAutomatedMarketMakerPairIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeaabSetAutomatedMarketMakerPair represents a SetAutomatedMarketMakerPair event raised by the Leaab contract.
type LeaabSetAutomatedMarketMakerPair struct {
	Pair  common.Address
	Value bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetAutomatedMarketMakerPair is a free log retrieval operation binding the contract event 0xffa9187bf1f18bf477bd0ea1bcbb64e93b6a98132473929edfce215cd9b16fab.
//
// Solidity: event SetAutomatedMarketMakerPair(address indexed pair, bool indexed value)
func (_Leaab *LeaabFilterer) FilterSetAutomatedMarketMakerPair(opts *bind.FilterOpts, pair []common.Address, value []bool) (*LeaabSetAutomatedMarketMakerPairIterator, error) {

	var pairRule []interface{}
	for _, pairItem := range pair {
		pairRule = append(pairRule, pairItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Leaab.contract.FilterLogs(opts, "SetAutomatedMarketMakerPair", pairRule, valueRule)
	if err != nil {
		return nil, err
	}
	return &LeaabSetAutomatedMarketMakerPairIterator{contract: _Leaab.contract, event: "SetAutomatedMarketMakerPair", logs: logs, sub: sub}, nil
}

// WatchSetAutomatedMarketMakerPair is a free log subscription operation binding the contract event 0xffa9187bf1f18bf477bd0ea1bcbb64e93b6a98132473929edfce215cd9b16fab.
//
// Solidity: event SetAutomatedMarketMakerPair(address indexed pair, bool indexed value)
func (_Leaab *LeaabFilterer) WatchSetAutomatedMarketMakerPair(opts *bind.WatchOpts, sink chan<- *LeaabSetAutomatedMarketMakerPair, pair []common.Address, value []bool) (event.Subscription, error) {

	var pairRule []interface{}
	for _, pairItem := range pair {
		pairRule = append(pairRule, pairItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Leaab.contract.WatchLogs(opts, "SetAutomatedMarketMakerPair", pairRule, valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeaabSetAutomatedMarketMakerPair)
				if err := _Leaab.contract.UnpackLog(event, "SetAutomatedMarketMakerPair", log); err != nil {
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

// ParseSetAutomatedMarketMakerPair is a log parse operation binding the contract event 0xffa9187bf1f18bf477bd0ea1bcbb64e93b6a98132473929edfce215cd9b16fab.
//
// Solidity: event SetAutomatedMarketMakerPair(address indexed pair, bool indexed value)
func (_Leaab *LeaabFilterer) ParseSetAutomatedMarketMakerPair(log types.Log) (*LeaabSetAutomatedMarketMakerPair, error) {
	event := new(LeaabSetAutomatedMarketMakerPair)
	if err := _Leaab.contract.UnpackLog(event, "SetAutomatedMarketMakerPair", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeaabSwapAndLiquifyIterator is returned from FilterSwapAndLiquify and is used to iterate over the raw logs and unpacked data for SwapAndLiquify events raised by the Leaab contract.
type LeaabSwapAndLiquifyIterator struct {
	Event *LeaabSwapAndLiquify // Event containing the contract specifics and raw log

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
func (it *LeaabSwapAndLiquifyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeaabSwapAndLiquify)
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
		it.Event = new(LeaabSwapAndLiquify)
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
func (it *LeaabSwapAndLiquifyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeaabSwapAndLiquifyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeaabSwapAndLiquify represents a SwapAndLiquify event raised by the Leaab contract.
type LeaabSwapAndLiquify struct {
	TokensSwapped      *big.Int
	EthReceived        *big.Int
	TokensIntoLiqudity *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSwapAndLiquify is a free log retrieval operation binding the contract event 0x17bbfb9a6069321b6ded73bd96327c9e6b7212a5cd51ff219cd61370acafb561.
//
// Solidity: event SwapAndLiquify(uint256 tokensSwapped, uint256 ethReceived, uint256 tokensIntoLiqudity)
func (_Leaab *LeaabFilterer) FilterSwapAndLiquify(opts *bind.FilterOpts) (*LeaabSwapAndLiquifyIterator, error) {

	logs, sub, err := _Leaab.contract.FilterLogs(opts, "SwapAndLiquify")
	if err != nil {
		return nil, err
	}
	return &LeaabSwapAndLiquifyIterator{contract: _Leaab.contract, event: "SwapAndLiquify", logs: logs, sub: sub}, nil
}

// WatchSwapAndLiquify is a free log subscription operation binding the contract event 0x17bbfb9a6069321b6ded73bd96327c9e6b7212a5cd51ff219cd61370acafb561.
//
// Solidity: event SwapAndLiquify(uint256 tokensSwapped, uint256 ethReceived, uint256 tokensIntoLiqudity)
func (_Leaab *LeaabFilterer) WatchSwapAndLiquify(opts *bind.WatchOpts, sink chan<- *LeaabSwapAndLiquify) (event.Subscription, error) {

	logs, sub, err := _Leaab.contract.WatchLogs(opts, "SwapAndLiquify")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeaabSwapAndLiquify)
				if err := _Leaab.contract.UnpackLog(event, "SwapAndLiquify", log); err != nil {
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

// ParseSwapAndLiquify is a log parse operation binding the contract event 0x17bbfb9a6069321b6ded73bd96327c9e6b7212a5cd51ff219cd61370acafb561.
//
// Solidity: event SwapAndLiquify(uint256 tokensSwapped, uint256 ethReceived, uint256 tokensIntoLiqudity)
func (_Leaab *LeaabFilterer) ParseSwapAndLiquify(log types.Log) (*LeaabSwapAndLiquify, error) {
	event := new(LeaabSwapAndLiquify)
	if err := _Leaab.contract.UnpackLog(event, "SwapAndLiquify", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeaabTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Leaab contract.
type LeaabTransferIterator struct {
	Event *LeaabTransfer // Event containing the contract specifics and raw log

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
func (it *LeaabTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeaabTransfer)
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
		it.Event = new(LeaabTransfer)
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
func (it *LeaabTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeaabTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeaabTransfer represents a Transfer event raised by the Leaab contract.
type LeaabTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Leaab *LeaabFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LeaabTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Leaab.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LeaabTransferIterator{contract: _Leaab.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Leaab *LeaabFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LeaabTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Leaab.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeaabTransfer)
				if err := _Leaab.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Leaab *LeaabFilterer) ParseTransfer(log types.Log) (*LeaabTransfer, error) {
	event := new(LeaabTransfer)
	if err := _Leaab.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeaabUpdateDividendTrackerIterator is returned from FilterUpdateDividendTracker and is used to iterate over the raw logs and unpacked data for UpdateDividendTracker events raised by the Leaab contract.
type LeaabUpdateDividendTrackerIterator struct {
	Event *LeaabUpdateDividendTracker // Event containing the contract specifics and raw log

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
func (it *LeaabUpdateDividendTrackerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeaabUpdateDividendTracker)
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
		it.Event = new(LeaabUpdateDividendTracker)
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
func (it *LeaabUpdateDividendTrackerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeaabUpdateDividendTrackerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeaabUpdateDividendTracker represents a UpdateDividendTracker event raised by the Leaab contract.
type LeaabUpdateDividendTracker struct {
	NewAddress common.Address
	OldAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateDividendTracker is a free log retrieval operation binding the contract event 0x90c7d74461c613da5efa97d90740869367d74ab3aa5837aa4ae9a975f954b7a8.
//
// Solidity: event UpdateDividendTracker(address indexed newAddress, address indexed oldAddress)
func (_Leaab *LeaabFilterer) FilterUpdateDividendTracker(opts *bind.FilterOpts, newAddress []common.Address, oldAddress []common.Address) (*LeaabUpdateDividendTrackerIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}
	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}

	logs, sub, err := _Leaab.contract.FilterLogs(opts, "UpdateDividendTracker", newAddressRule, oldAddressRule)
	if err != nil {
		return nil, err
	}
	return &LeaabUpdateDividendTrackerIterator{contract: _Leaab.contract, event: "UpdateDividendTracker", logs: logs, sub: sub}, nil
}

// WatchUpdateDividendTracker is a free log subscription operation binding the contract event 0x90c7d74461c613da5efa97d90740869367d74ab3aa5837aa4ae9a975f954b7a8.
//
// Solidity: event UpdateDividendTracker(address indexed newAddress, address indexed oldAddress)
func (_Leaab *LeaabFilterer) WatchUpdateDividendTracker(opts *bind.WatchOpts, sink chan<- *LeaabUpdateDividendTracker, newAddress []common.Address, oldAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}
	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}

	logs, sub, err := _Leaab.contract.WatchLogs(opts, "UpdateDividendTracker", newAddressRule, oldAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeaabUpdateDividendTracker)
				if err := _Leaab.contract.UnpackLog(event, "UpdateDividendTracker", log); err != nil {
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

// ParseUpdateDividendTracker is a log parse operation binding the contract event 0x90c7d74461c613da5efa97d90740869367d74ab3aa5837aa4ae9a975f954b7a8.
//
// Solidity: event UpdateDividendTracker(address indexed newAddress, address indexed oldAddress)
func (_Leaab *LeaabFilterer) ParseUpdateDividendTracker(log types.Log) (*LeaabUpdateDividendTracker, error) {
	event := new(LeaabUpdateDividendTracker)
	if err := _Leaab.contract.UnpackLog(event, "UpdateDividendTracker", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeaabUpdateUniswapV2RouterIterator is returned from FilterUpdateUniswapV2Router and is used to iterate over the raw logs and unpacked data for UpdateUniswapV2Router events raised by the Leaab contract.
type LeaabUpdateUniswapV2RouterIterator struct {
	Event *LeaabUpdateUniswapV2Router // Event containing the contract specifics and raw log

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
func (it *LeaabUpdateUniswapV2RouterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeaabUpdateUniswapV2Router)
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
		it.Event = new(LeaabUpdateUniswapV2Router)
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
func (it *LeaabUpdateUniswapV2RouterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeaabUpdateUniswapV2RouterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeaabUpdateUniswapV2Router represents a UpdateUniswapV2Router event raised by the Leaab contract.
type LeaabUpdateUniswapV2Router struct {
	NewAddress common.Address
	OldAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateUniswapV2Router is a free log retrieval operation binding the contract event 0x8fc842bbd331dfa973645f4ed48b11683d501ebf1352708d77a5da2ab49a576e.
//
// Solidity: event UpdateUniswapV2Router(address indexed newAddress, address indexed oldAddress)
func (_Leaab *LeaabFilterer) FilterUpdateUniswapV2Router(opts *bind.FilterOpts, newAddress []common.Address, oldAddress []common.Address) (*LeaabUpdateUniswapV2RouterIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}
	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}

	logs, sub, err := _Leaab.contract.FilterLogs(opts, "UpdateUniswapV2Router", newAddressRule, oldAddressRule)
	if err != nil {
		return nil, err
	}
	return &LeaabUpdateUniswapV2RouterIterator{contract: _Leaab.contract, event: "UpdateUniswapV2Router", logs: logs, sub: sub}, nil
}

// WatchUpdateUniswapV2Router is a free log subscription operation binding the contract event 0x8fc842bbd331dfa973645f4ed48b11683d501ebf1352708d77a5da2ab49a576e.
//
// Solidity: event UpdateUniswapV2Router(address indexed newAddress, address indexed oldAddress)
func (_Leaab *LeaabFilterer) WatchUpdateUniswapV2Router(opts *bind.WatchOpts, sink chan<- *LeaabUpdateUniswapV2Router, newAddress []common.Address, oldAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}
	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}

	logs, sub, err := _Leaab.contract.WatchLogs(opts, "UpdateUniswapV2Router", newAddressRule, oldAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeaabUpdateUniswapV2Router)
				if err := _Leaab.contract.UnpackLog(event, "UpdateUniswapV2Router", log); err != nil {
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

// ParseUpdateUniswapV2Router is a log parse operation binding the contract event 0x8fc842bbd331dfa973645f4ed48b11683d501ebf1352708d77a5da2ab49a576e.
//
// Solidity: event UpdateUniswapV2Router(address indexed newAddress, address indexed oldAddress)
func (_Leaab *LeaabFilterer) ParseUpdateUniswapV2Router(log types.Log) (*LeaabUpdateUniswapV2Router, error) {
	event := new(LeaabUpdateUniswapV2Router)
	if err := _Leaab.contract.UnpackLog(event, "UpdateUniswapV2Router", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
