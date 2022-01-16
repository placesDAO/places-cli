// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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



	// IPlacesLocation is an auto generated low-level Go binding around an user-defined struct.
	type IPlacesLocation struct {
	
	LatitudeInt *big.Int
	LongitudeInt *big.Int
	AltitudeInt *big.Int
	HasAltitude bool
	Latitude string
	Longitude string
	Altitude string
	}

	// IPlacesPlace is an auto generated low-level Go binding around an user-defined struct.
	type IPlacesPlace struct {
	
	Name string
	StreetAddress string
	Sublocality string
	Locality string
	SubadministrativeArea string
	AdministrativeArea string
	Country string
	PostalCode string
	CountryCode string
	Location IPlacesLocation
	Attributes [3]string
	}



	// PlacesABI is the input ABI used to generate the binding from.
	const PlacesABI = "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"placesDAO\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"grounders\",\"type\":\"address\"},{\"internalType\":\"contractIPlacesDescriptor\",\"name\":\"placesDescriptor\",\"type\":\"address\"},{\"internalType\":\"contractIPlacesProvider\",\"name\":\"placesProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_proxyAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIPlacesDescriptor\",\"name\":\"descriptor\",\"type\":\"address\"}],\"name\":\"DescriptorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"name\":\"GroundersGiftingEnabledUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"grounders\",\"type\":\"address\"}],\"name\":\"GroundersUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"name\":\"GuestlistEnabledUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintFee\",\"type\":\"uint256\"}],\"name\":\"MintFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"name\":\"NeighborhoodTreasuryEnabledUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"PlaceBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"PlaceCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"placesDao\",\"type\":\"address\"}],\"name\":\"PlacesDAOUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIPlacesProvider\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"PlacesProviderUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GEO_RESOLUTION_INT\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_LATITUDE_INT\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_LONGITUDE_INT\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_LATITUDE_INT\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_LONGITUDE_INT\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guestAddress\",\"type\":\"address\"}],\"name\":\"addGuest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGroundersGiftingEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getLocation\",\"outputs\":[{\"components\":[{\"internalType\":\"int256\",\"name\":\"latitudeInt\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"longitudeInt\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"altitudeInt\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"hasAltitude\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"latitude\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"longitude\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"altitude\",\"type\":\"string\"}],\"internalType\":\"structIPlaces.Location\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMintFeeInWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"currentMintFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNeighborhoodTreasuryEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getPlace\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"streetAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sublocality\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"locality\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"subadministrativeArea\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"administrativeArea\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"country\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"postalCode\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"countryCode\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"latitudeInt\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"longitudeInt\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"altitudeInt\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"hasAltitude\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"latitude\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"longitude\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"altitude\",\"type\":\"string\"}],\"internalType\":\"structIPlaces.Location\",\"name\":\"location\",\"type\":\"tuple\"},{\"internalType\":\"string[3]\",\"name\":\"attributes\",\"type\":\"string[3]\"}],\"internalType\":\"structIPlaces.Place\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPlaceSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"supplyCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"grounderMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxyRegistry\",\"outputs\":[{\"internalType\":\"contractIProxyRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_countValue\",\"type\":\"uint256\"}],\"name\":\"setCounter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPlacesDescriptor\",\"name\":\"placesDescriptor\",\"type\":\"address\"}],\"name\":\"setDescriptor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"grounders\",\"type\":\"address\"}],\"name\":\"setGrounders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isGroundersGiftingEnabled\",\"type\":\"bool\"}],\"name\":\"setGroundersGifting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"guestlist\",\"type\":\"address[]\"}],\"name\":\"setGuestlist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isGuestlistEnabled\",\"type\":\"bool\"}],\"name\":\"setGuestlistEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"mintFee\",\"type\":\"uint256\"}],\"name\":\"setMintFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isNeighborhoodTreasuryEnabled\",\"type\":\"bool\"}],\"name\":\"setNeighborhoodTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"setPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"placesDAO\",\"type\":\"address\"}],\"name\":\"setPlacesDAO\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPlacesProvider\",\"name\":\"placesProvider\",\"type\":\"address\"}],\"name\":\"setPlacesProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

	

	
		// PlacesBin is the compiled bytecode used for deploying new contracts.
		var PlacesBin = "0x# revert with 0, 'ERC721: transfer to non ERC721Receiver implementer'"

		// DeployPlaces deploys a new Ethereum contract, binding an instance of Places to it.
		func DeployPlaces(auth *bind.TransactOpts, backend bind.ContractBackend , placesDAO common.Address, grounders common.Address, placesDescriptor common.Address, placesProvider common.Address, _proxyAddress common.Address) (common.Address, *types.Transaction, *Places, error) {
		  parsed, err := abi.JSON(strings.NewReader(PlacesABI))
		  if err != nil {
		    return common.Address{}, nil, nil, err
		  }
		  
		  address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PlacesBin), backend , placesDAO, grounders, placesDescriptor, placesProvider, _proxyAddress)
		  if err != nil {
		    return common.Address{}, nil, nil, err
		  }
		  return address, tx, &Places{ PlacesCaller: PlacesCaller{contract: contract}, PlacesTransactor: PlacesTransactor{contract: contract}, PlacesFilterer: PlacesFilterer{contract: contract} }, nil
		}
	

	// Places is an auto generated Go binding around an Ethereum contract.
	type Places struct {
	  PlacesCaller     // Read-only binding to the contract
	  PlacesTransactor // Write-only binding to the contract
	  PlacesFilterer   // Log filterer for contract events
	}

	// PlacesCaller is an auto generated read-only Go binding around an Ethereum contract.
	type PlacesCaller struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// PlacesTransactor is an auto generated write-only Go binding around an Ethereum contract.
	type PlacesTransactor struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// PlacesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
	type PlacesFilterer struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// PlacesSession is an auto generated Go binding around an Ethereum contract,
	// with pre-set call and transact options.
	type PlacesSession struct {
	  Contract     *Places        // Generic contract binding to set the session for
	  CallOpts     bind.CallOpts     // Call options to use throughout this session
	  TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
	}

	// PlacesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
	// with pre-set call options.
	type PlacesCallerSession struct {
	  Contract *PlacesCaller // Generic contract caller binding to set the session for
	  CallOpts bind.CallOpts    // Call options to use throughout this session
	}

	// PlacesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
	// with pre-set transact options.
	type PlacesTransactorSession struct {
	  Contract     *PlacesTransactor // Generic contract transactor binding to set the session for
	  TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
	}

	// PlacesRaw is an auto generated low-level Go binding around an Ethereum contract.
	type PlacesRaw struct {
	  Contract *Places // Generic contract binding to access the raw methods on
	}

	// PlacesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
	type PlacesCallerRaw struct {
		Contract *PlacesCaller // Generic read-only contract binding to access the raw methods on
	}

	// PlacesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
	type PlacesTransactorRaw struct {
		Contract *PlacesTransactor // Generic write-only contract binding to access the raw methods on
	}

	// NewPlaces creates a new instance of Places, bound to a specific deployed contract.
	func NewPlaces(address common.Address, backend bind.ContractBackend) (*Places, error) {
	  contract, err := bindPlaces(address, backend, backend, backend)
	  if err != nil {
	    return nil, err
	  }
	  return &Places{ PlacesCaller: PlacesCaller{contract: contract}, PlacesTransactor: PlacesTransactor{contract: contract}, PlacesFilterer: PlacesFilterer{contract: contract} }, nil
	}

	// NewPlacesCaller creates a new read-only instance of Places, bound to a specific deployed contract.
	func NewPlacesCaller(address common.Address, caller bind.ContractCaller) (*PlacesCaller, error) {
	  contract, err := bindPlaces(address, caller, nil, nil)
	  if err != nil {
	    return nil, err
	  }
	  return &PlacesCaller{contract: contract}, nil
	}

	// NewPlacesTransactor creates a new write-only instance of Places, bound to a specific deployed contract.
	func NewPlacesTransactor(address common.Address, transactor bind.ContractTransactor) (*PlacesTransactor, error) {
	  contract, err := bindPlaces(address, nil, transactor, nil)
	  if err != nil {
	    return nil, err
	  }
	  return &PlacesTransactor{contract: contract}, nil
	}

	// NewPlacesFilterer creates a new log filterer instance of Places, bound to a specific deployed contract.
 	func NewPlacesFilterer(address common.Address, filterer bind.ContractFilterer) (*PlacesFilterer, error) {
 	  contract, err := bindPlaces(address, nil, nil, filterer)
 	  if err != nil {
 	    return nil, err
 	  }
 	  return &PlacesFilterer{contract: contract}, nil
 	}

	// bindPlaces binds a generic wrapper to an already deployed contract.
	func bindPlaces(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	  parsed, err := abi.JSON(strings.NewReader(PlacesABI))
	  if err != nil {
	    return nil, err
	  }
	  return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
	}

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_Places *PlacesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
		return _Places.Contract.PlacesCaller.contract.Call(opts, result, method, params...)
	}

	// Transfer initiates a plain transaction to move funds to the contract, calling
	// its default method if one is available.
	func (_Places *PlacesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
		return _Places.Contract.PlacesTransactor.contract.Transfer(opts)
	}

	// Transact invokes the (paid) contract method with params as input values.
	func (_Places *PlacesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
		return _Places.Contract.PlacesTransactor.contract.Transact(opts, method, params...)
	}

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_Places *PlacesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
		return _Places.Contract.contract.Call(opts, result, method, params...)
	}

	// Transfer initiates a plain transaction to move funds to the contract, calling
	// its default method if one is available.
	func (_Places *PlacesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
		return _Places.Contract.contract.Transfer(opts)
	}

	// Transact invokes the (paid) contract method with params as input values.
	func (_Places *PlacesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
		return _Places.Contract.contract.Transact(opts, method, params...)
	}

	
		// GEORESOLUTIONINT is a free data retrieval call binding the contract method 0xac70cc75.
		//
		// Solidity: function GEO_RESOLUTION_INT() view returns(int256)
		func (_Places *PlacesCaller) GEORESOLUTIONINT(opts *bind.CallOpts ) (*big.Int, error) {
			var out []interface{}
			err := _Places.contract.Call(opts, &out, "GEO_RESOLUTION_INT" )
			
			if err != nil {
				return *new(*big.Int),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
			
			return out0,  err
			
		}

		// GEORESOLUTIONINT is a free data retrieval call binding the contract method 0xac70cc75.
		//
		// Solidity: function GEO_RESOLUTION_INT() view returns(int256)
		func (_Places *PlacesSession) GEORESOLUTIONINT() ( *big.Int,  error) {
		  return _Places.Contract.GEORESOLUTIONINT(&_Places.CallOpts )
		}

		// GEORESOLUTIONINT is a free data retrieval call binding the contract method 0xac70cc75.
		//
		// Solidity: function GEO_RESOLUTION_INT() view returns(int256)
		func (_Places *PlacesCallerSession) GEORESOLUTIONINT() ( *big.Int,  error) {
		  return _Places.Contract.GEORESOLUTIONINT(&_Places.CallOpts )
		}
	
		// MAXLATITUDEINT is a free data retrieval call binding the contract method 0xb9485c32.
		//
		// Solidity: function MAX_LATITUDE_INT() view returns(int256)
		func (_Places *PlacesCaller) MAXLATITUDEINT(opts *bind.CallOpts ) (*big.Int, error) {
			var out []interface{}
			err := _Places.contract.Call(opts, &out, "MAX_LATITUDE_INT" )
			
			if err != nil {
				return *new(*big.Int),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
			
			return out0,  err
			
		}

		// MAXLATITUDEINT is a free data retrieval call binding the contract method 0xb9485c32.
		//
		// Solidity: function MAX_LATITUDE_INT() view returns(int256)
		func (_Places *PlacesSession) MAXLATITUDEINT() ( *big.Int,  error) {
		  return _Places.Contract.MAXLATITUDEINT(&_Places.CallOpts )
		}

		// MAXLATITUDEINT is a free data retrieval call binding the contract method 0xb9485c32.
		//
		// Solidity: function MAX_LATITUDE_INT() view returns(int256)
		func (_Places *PlacesCallerSession) MAXLATITUDEINT() ( *big.Int,  error) {
		  return _Places.Contract.MAXLATITUDEINT(&_Places.CallOpts )
		}
	
		// MAXLONGITUDEINT is a free data retrieval call binding the contract method 0xa0b6e3af.
		//
		// Solidity: function MAX_LONGITUDE_INT() view returns(int256)
		func (_Places *PlacesCaller) MAXLONGITUDEINT(opts *bind.CallOpts ) (*big.Int, error) {
			var out []interface{}
			err := _Places.contract.Call(opts, &out, "MAX_LONGITUDE_INT" )
			
			if err != nil {
				return *new(*big.Int),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
			
			return out0,  err
			
		}

		// MAXLONGITUDEINT is a free data retrieval call binding the contract method 0xa0b6e3af.
		//
		// Solidity: function MAX_LONGITUDE_INT() view returns(int256)
		func (_Places *PlacesSession) MAXLONGITUDEINT() ( *big.Int,  error) {
		  return _Places.Contract.MAXLONGITUDEINT(&_Places.CallOpts )
		}

		// MAXLONGITUDEINT is a free data retrieval call binding the contract method 0xa0b6e3af.
		//
		// Solidity: function MAX_LONGITUDE_INT() view returns(int256)
		func (_Places *PlacesCallerSession) MAXLONGITUDEINT() ( *big.Int,  error) {
		  return _Places.Contract.MAXLONGITUDEINT(&_Places.CallOpts )
		}
	
		// MINLATITUDEINT is a free data retrieval call binding the contract method 0x1aac2894.
		//
		// Solidity: function MIN_LATITUDE_INT() view returns(int256)
		func (_Places *PlacesCaller) MINLATITUDEINT(opts *bind.CallOpts ) (*big.Int, error) {
			var out []interface{}
			err := _Places.contract.Call(opts, &out, "MIN_LATITUDE_INT" )
			
			if err != nil {
				return *new(*big.Int),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
			
			return out0,  err
			
		}

		// MINLATITUDEINT is a free data retrieval call binding the contract method 0x1aac2894.
		//
		// Solidity: function MIN_LATITUDE_INT() view returns(int256)
		func (_Places *PlacesSession) MINLATITUDEINT() ( *big.Int,  error) {
		  return _Places.Contract.MINLATITUDEINT(&_Places.CallOpts )
		}

		// MINLATITUDEINT is a free data retrieval call binding the contract method 0x1aac2894.
		//
		// Solidity: function MIN_LATITUDE_INT() view returns(int256)
		func (_Places *PlacesCallerSession) MINLATITUDEINT() ( *big.Int,  error) {
		  return _Places.Contract.MINLATITUDEINT(&_Places.CallOpts )
		}
	
		// MINLONGITUDEINT is a free data retrieval call binding the contract method 0xa0d8a56e.
		//
		// Solidity: function MIN_LONGITUDE_INT() view returns(int256)
		func (_Places *PlacesCaller) MINLONGITUDEINT(opts *bind.CallOpts ) (*big.Int, error) {
			var out []interface{}
			err := _Places.contract.Call(opts, &out, "MIN_LONGITUDE_INT" )
			
			if err != nil {
				return *new(*big.Int),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
			
			return out0,  err
			
		}

		// MINLONGITUDEINT is a free data retrieval call binding the contract method 0xa0d8a56e.
		//
		// Solidity: function MIN_LONGITUDE_INT() view returns(int256)
		func (_Places *PlacesSession) MINLONGITUDEINT() ( *big.Int,  error) {
		  return _Places.Contract.MINLONGITUDEINT(&_Places.CallOpts )
		}

		// MINLONGITUDEINT is a free data retrieval call binding the contract method 0xa0d8a56e.
		//
		// Solidity: function MIN_LONGITUDE_INT() view returns(int256)
		func (_Places *PlacesCallerSession) MINLONGITUDEINT() ( *big.Int,  error) {
		  return _Places.Contract.MINLONGITUDEINT(&_Places.CallOpts )
		}
	
		// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
		//
		// Solidity: function balanceOf(address owner) view returns(uint256)
		func (_Places *PlacesCaller) BalanceOf(opts *bind.CallOpts , owner common.Address ) (*big.Int, error) {
			var out []interface{}
			err := _Places.contract.Call(opts, &out, "balanceOf" , owner)
			
			if err != nil {
				return *new(*big.Int),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
			
			return out0,  err
			
		}

		// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
		//
		// Solidity: function balanceOf(address owner) view returns(uint256)
		func (_Places *PlacesSession) BalanceOf( owner common.Address ) ( *big.Int,  error) {
		  return _Places.Contract.BalanceOf(&_Places.CallOpts , owner)
		}

		// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
		//
		// Solidity: function balanceOf(address owner) view returns(uint256)
		func (_Places *PlacesCallerSession) BalanceOf( owner common.Address ) ( *big.Int,  error) {
		  return _Places.Contract.BalanceOf(&_Places.CallOpts , owner)
		}
	
		// ContractURI is a free data retrieval call binding the contract method 0x7a62b340.
		//
		// Solidity: function contractURI(uint256 tokenId) view returns(string)
		func (_Places *PlacesCaller) ContractURI(opts *bind.CallOpts , tokenId *big.Int ) (string, error) {
			var out []interface{}
			err := _Places.contract.Call(opts, &out, "contractURI" , tokenId)
			
			if err != nil {
				return *new(string),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(string)).(*string)
			
			return out0,  err
			
		}

		// ContractURI is a free data retrieval call binding the contract method 0x7a62b340.
		//
		// Solidity: function contractURI(uint256 tokenId) view returns(string)
		func (_Places *PlacesSession) ContractURI( tokenId *big.Int ) ( string,  error) {
		  return _Places.Contract.ContractURI(&_Places.CallOpts , tokenId)
		}

		// ContractURI is a free data retrieval call binding the contract method 0x7a62b340.
		//
		// Solidity: function contractURI(uint256 tokenId) view returns(string)
		func (_Places *PlacesCallerSession) ContractURI( tokenId *big.Int ) ( string,  error) {
		  return _Places.Contract.ContractURI(&_Places.CallOpts , tokenId)
		}
	
		// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
		//
		// Solidity: function getApproved(uint256 tokenId) view returns(address)
		func (_Places *PlacesCaller) GetApproved(opts *bind.CallOpts , tokenId *big.Int ) (common.Address, error) {
			var out []interface{}
			err := _Places.contract.Call(opts, &out, "getApproved" , tokenId)
			
			if err != nil {
				return *new(common.Address),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
			
			return out0,  err
			
		}

		// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
		//
		// Solidity: function getApproved(uint256 tokenId) view returns(address)
		func (_Places *PlacesSession) GetApproved( tokenId *big.Int ) ( common.Address,  error) {
		  return _Places.Contract.GetApproved(&_Places.CallOpts , tokenId)
		}

		// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
		//
		// Solidity: function getApproved(uint256 tokenId) view returns(address)
		func (_Places *PlacesCallerSession) GetApproved( tokenId *big.Int ) ( common.Address,  error) {
		  return _Places.Contract.GetApproved(&_Places.CallOpts , tokenId)
		}
	
		// GetGroundersGiftingEnabled is a free data retrieval call binding the contract method 0x4482ae9a.
		//
		// Solidity: function getGroundersGiftingEnabled() view returns(bool isEnabled)
		func (_Places *PlacesCaller) GetGroundersGiftingEnabled(opts *bind.CallOpts ) (bool, error) {
			var out []interface{}
			err := _Places.contract.Call(opts, &out, "getGroundersGiftingEnabled" )
			
			if err != nil {
				return *new(bool),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
			
			return out0,  err
			
		}

		// GetGroundersGiftingEnabled is a free data retrieval call binding the contract method 0x4482ae9a.
		//
		// Solidity: function getGroundersGiftingEnabled() view returns(bool isEnabled)
		func (_Places *PlacesSession) GetGroundersGiftingEnabled() ( bool,  error) {
		  return _Places.Contract.GetGroundersGiftingEnabled(&_Places.CallOpts )
		}

		// GetGroundersGiftingEnabled is a free data retrieval call binding the contract method 0x4482ae9a.
		//
		// Solidity: function getGroundersGiftingEnabled() view returns(bool isEnabled)
		func (_Places *PlacesCallerSession) GetGroundersGiftingEnabled() ( bool,  error) {
		  return _Places.Contract.GetGroundersGiftingEnabled(&_Places.CallOpts )
		}
	
		// GetLocation is a free data retrieval call binding the contract method 0x7f7b1393.
		//
		// Solidity: function getLocation(uint256 tokenId) view returns((int256,int256,int256,bool,string,string,string))
		func (_Places *PlacesCaller) GetLocation(opts *bind.CallOpts , tokenId *big.Int ) (IPlacesLocation, error) {
			var out []interface{}
			err := _Places.contract.Call(opts, &out, "getLocation" , tokenId)
			
			if err != nil {
				return *new(IPlacesLocation),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(IPlacesLocation)).(*IPlacesLocation)
			
			return out0,  err
			
		}

		// GetLocation is a free data retrieval call binding the contract method 0x7f7b1393.
		//
		// Solidity: function getLocation(uint256 tokenId) view returns((int256,int256,int256,bool,string,string,string))
		func (_Places *PlacesSession) GetLocation( tokenId *big.Int ) ( IPlacesLocation,  error) {
		  return _Places.Contract.GetLocation(&_Places.CallOpts , tokenId)
		}

		// GetLocation is a free data retrieval call binding the contract method 0x7f7b1393.
		//
		// Solidity: function getLocation(uint256 tokenId) view returns((int256,int256,int256,bool,string,string,string))
		func (_Places *PlacesCallerSession) GetLocation( tokenId *big.Int ) ( IPlacesLocation,  error) {
		  return _Places.Contract.GetLocation(&_Places.CallOpts , tokenId)
		}
	
		// GetMintFeeInWei is a free data retrieval call binding the contract method 0xaef9bc82.
		//
		// Solidity: function getMintFeeInWei() view returns(uint256 currentMintFee)
		func (_Places *PlacesCaller) GetMintFeeInWei(opts *bind.CallOpts ) (*big.Int, error) {
			var out []interface{}
			err := _Places.contract.Call(opts, &out, "getMintFeeInWei" )
			
			if err != nil {
				return *new(*big.Int),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
			
			return out0,  err
			
		}

		// GetMintFeeInWei is a free data retrieval call binding the contract method 0xaef9bc82.
		//
		// Solidity: function getMintFeeInWei() view returns(uint256 currentMintFee)
		func (_Places *PlacesSession) GetMintFeeInWei() ( *big.Int,  error) {
		  return _Places.Contract.GetMintFeeInWei(&_Places.CallOpts )
		}

		// GetMintFeeInWei is a free data retrieval call binding the contract method 0xaef9bc82.
		//
		// Solidity: function getMintFeeInWei() view returns(uint256 currentMintFee)
		func (_Places *PlacesCallerSession) GetMintFeeInWei() ( *big.Int,  error) {
		  return _Places.Contract.GetMintFeeInWei(&_Places.CallOpts )
		}
	
		// GetNeighborhoodTreasuryEnabled is a free data retrieval call binding the contract method 0x6eff3cdf.
		//
		// Solidity: function getNeighborhoodTreasuryEnabled() view returns(bool isEnabled)
		func (_Places *PlacesCaller) GetNeighborhoodTreasuryEnabled(opts *bind.CallOpts ) (bool, error) {
			var out []interface{}
			err := _Places.contract.Call(opts, &out, "getNeighborhoodTreasuryEnabled" )
			
			if err != nil {
				return *new(bool),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
			
			return out0,  err
			
		}

		// GetNeighborhoodTreasuryEnabled is a free data retrieval call binding the contract method 0x6eff3cdf.
		//
		// Solidity: function getNeighborhoodTreasuryEnabled() view returns(bool isEnabled)
		func (_Places *PlacesSession) GetNeighborhoodTreasuryEnabled() ( bool,  error) {
		  return _Places.Contract.GetNeighborhoodTreasuryEnabled(&_Places.CallOpts )
		}

		// GetNeighborhoodTreasuryEnabled is a free data retrieval call binding the contract method 0x6eff3cdf.
		//
		// Solidity: function getNeighborhoodTreasuryEnabled() view returns(bool isEnabled)
		func (_Places *PlacesCallerSession) GetNeighborhoodTreasuryEnabled() ( bool,  error) {
		  return _Places.Contract.GetNeighborhoodTreasuryEnabled(&_Places.CallOpts )
		}
	
		// GetPlace is a free data retrieval call binding the contract method 0x197c3555.
		//
		// Solidity: function getPlace(uint256 tokenId) view returns((string,string,string,string,string,string,string,string,string,(int256,int256,int256,bool,string,string,string),string[3]))
		func (_Places *PlacesCaller) GetPlace(opts *bind.CallOpts , tokenId *big.Int ) (IPlacesPlace, error) {
			var out []interface{}
			err := _Places.contract.Call(opts, &out, "getPlace" , tokenId)
			
			if err != nil {
				return *new(IPlacesPlace),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(IPlacesPlace)).(*IPlacesPlace)
			
			return out0,  err
			
		}

		// GetPlace is a free data retrieval call binding the contract method 0x197c3555.
		//
		// Solidity: function getPlace(uint256 tokenId) view returns((string,string,string,string,string,string,string,string,string,(int256,int256,int256,bool,string,string,string),string[3]))
		func (_Places *PlacesSession) GetPlace( tokenId *big.Int ) ( IPlacesPlace,  error) {
		  return _Places.Contract.GetPlace(&_Places.CallOpts , tokenId)
		}

		// GetPlace is a free data retrieval call binding the contract method 0x197c3555.
		//
		// Solidity: function getPlace(uint256 tokenId) view returns((string,string,string,string,string,string,string,string,string,(int256,int256,int256,bool,string,string,string),string[3]))
		func (_Places *PlacesCallerSession) GetPlace( tokenId *big.Int ) ( IPlacesPlace,  error) {
		  return _Places.Contract.GetPlace(&_Places.CallOpts , tokenId)
		}
	
		// GetPlaceSupply is a free data retrieval call binding the contract method 0x0d1b185d.
		//
		// Solidity: function getPlaceSupply() view returns(uint256 supplyCount)
		func (_Places *PlacesCaller) GetPlaceSupply(opts *bind.CallOpts ) (*big.Int, error) {
			var out []interface{}
			err := _Places.contract.Call(opts, &out, "getPlaceSupply" )
			
			if err != nil {
				return *new(*big.Int),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
			
			return out0,  err
			
		}

		// GetPlaceSupply is a free data retrieval call binding the contract method 0x0d1b185d.
		//
		// Solidity: function getPlaceSupply() view returns(uint256 supplyCount)
		func (_Places *PlacesSession) GetPlaceSupply() ( *big.Int,  error) {
		  return _Places.Contract.GetPlaceSupply(&_Places.CallOpts )
		}

		// GetPlaceSupply is a free data retrieval call binding the contract method 0x0d1b185d.
		//
		// Solidity: function getPlaceSupply() view returns(uint256 supplyCount)
		func (_Places *PlacesCallerSession) GetPlaceSupply() ( *big.Int,  error) {
		  return _Places.Contract.GetPlaceSupply(&_Places.CallOpts )
		}
	
		// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
		//
		// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
		func (_Places *PlacesCaller) IsApprovedForAll(opts *bind.CallOpts , owner common.Address , operator common.Address ) (bool, error) {
			var out []interface{}
			err := _Places.contract.Call(opts, &out, "isApprovedForAll" , owner, operator)
			
			if err != nil {
				return *new(bool),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
			
			return out0,  err
			
		}

		// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
		//
		// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
		func (_Places *PlacesSession) IsApprovedForAll( owner common.Address , operator common.Address ) ( bool,  error) {
		  return _Places.Contract.IsApprovedForAll(&_Places.CallOpts , owner, operator)
		}

		// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
		//
		// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
		func (_Places *PlacesCallerSession) IsApprovedForAll( owner common.Address , operator common.Address ) ( bool,  error) {
		  return _Places.Contract.IsApprovedForAll(&_Places.CallOpts , owner, operator)
		}
	
		// Name is a free data retrieval call binding the contract method 0x06fdde03.
		//
		// Solidity: function name() view returns(string)
		func (_Places *PlacesCaller) Name(opts *bind.CallOpts ) (string, error) {
			var out []interface{}
			err := _Places.contract.Call(opts, &out, "name" )
			
			if err != nil {
				return *new(string),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(string)).(*string)
			
			return out0,  err
			
		}

		// Name is a free data retrieval call binding the contract method 0x06fdde03.
		//
		// Solidity: function name() view returns(string)
		func (_Places *PlacesSession) Name() ( string,  error) {
		  return _Places.Contract.Name(&_Places.CallOpts )
		}

		// Name is a free data retrieval call binding the contract method 0x06fdde03.
		//
		// Solidity: function name() view returns(string)
		func (_Places *PlacesCallerSession) Name() ( string,  error) {
		  return _Places.Contract.Name(&_Places.CallOpts )
		}
	
		// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
		//
		// Solidity: function owner() view returns(address)
		func (_Places *PlacesCaller) Owner(opts *bind.CallOpts ) (common.Address, error) {
			var out []interface{}
			err := _Places.contract.Call(opts, &out, "owner" )
			
			if err != nil {
				return *new(common.Address),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
			
			return out0,  err
			
		}

		// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
		//
		// Solidity: function owner() view returns(address)
		func (_Places *PlacesSession) Owner() ( common.Address,  error) {
		  return _Places.Contract.Owner(&_Places.CallOpts )
		}

		// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
		//
		// Solidity: function owner() view returns(address)
		func (_Places *PlacesCallerSession) Owner() ( common.Address,  error) {
		  return _Places.Contract.Owner(&_Places.CallOpts )
		}
	
		// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
		//
		// Solidity: function ownerOf(uint256 tokenId) view returns(address)
		func (_Places *PlacesCaller) OwnerOf(opts *bind.CallOpts , tokenId *big.Int ) (common.Address, error) {
			var out []interface{}
			err := _Places.contract.Call(opts, &out, "ownerOf" , tokenId)
			
			if err != nil {
				return *new(common.Address),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
			
			return out0,  err
			
		}

		// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
		//
		// Solidity: function ownerOf(uint256 tokenId) view returns(address)
		func (_Places *PlacesSession) OwnerOf( tokenId *big.Int ) ( common.Address,  error) {
		  return _Places.Contract.OwnerOf(&_Places.CallOpts , tokenId)
		}

		// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
		//
		// Solidity: function ownerOf(uint256 tokenId) view returns(address)
		func (_Places *PlacesCallerSession) OwnerOf( tokenId *big.Int ) ( common.Address,  error) {
		  return _Places.Contract.OwnerOf(&_Places.CallOpts , tokenId)
		}
	
		// Paused is a free data retrieval call binding the contract method 0x5c975abb.
		//
		// Solidity: function paused() view returns(bool)
		func (_Places *PlacesCaller) Paused(opts *bind.CallOpts ) (bool, error) {
			var out []interface{}
			err := _Places.contract.Call(opts, &out, "paused" )
			
			if err != nil {
				return *new(bool),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
			
			return out0,  err
			
		}

		// Paused is a free data retrieval call binding the contract method 0x5c975abb.
		//
		// Solidity: function paused() view returns(bool)
		func (_Places *PlacesSession) Paused() ( bool,  error) {
		  return _Places.Contract.Paused(&_Places.CallOpts )
		}

		// Paused is a free data retrieval call binding the contract method 0x5c975abb.
		//
		// Solidity: function paused() view returns(bool)
		func (_Places *PlacesCallerSession) Paused() ( bool,  error) {
		  return _Places.Contract.Paused(&_Places.CallOpts )
		}
	
		// ProxyRegistry is a free data retrieval call binding the contract method 0xb50cbd9f.
		//
		// Solidity: function proxyRegistry() view returns(address)
		func (_Places *PlacesCaller) ProxyRegistry(opts *bind.CallOpts ) (common.Address, error) {
			var out []interface{}
			err := _Places.contract.Call(opts, &out, "proxyRegistry" )
			
			if err != nil {
				return *new(common.Address),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
			
			return out0,  err
			
		}

		// ProxyRegistry is a free data retrieval call binding the contract method 0xb50cbd9f.
		//
		// Solidity: function proxyRegistry() view returns(address)
		func (_Places *PlacesSession) ProxyRegistry() ( common.Address,  error) {
		  return _Places.Contract.ProxyRegistry(&_Places.CallOpts )
		}

		// ProxyRegistry is a free data retrieval call binding the contract method 0xb50cbd9f.
		//
		// Solidity: function proxyRegistry() view returns(address)
		func (_Places *PlacesCallerSession) ProxyRegistry() ( common.Address,  error) {
		  return _Places.Contract.ProxyRegistry(&_Places.CallOpts )
		}
	
		// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
		//
		// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
		func (_Places *PlacesCaller) SupportsInterface(opts *bind.CallOpts , interfaceId [4]byte ) (bool, error) {
			var out []interface{}
			err := _Places.contract.Call(opts, &out, "supportsInterface" , interfaceId)
			
			if err != nil {
				return *new(bool),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
			
			return out0,  err
			
		}

		// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
		//
		// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
		func (_Places *PlacesSession) SupportsInterface( interfaceId [4]byte ) ( bool,  error) {
		  return _Places.Contract.SupportsInterface(&_Places.CallOpts , interfaceId)
		}

		// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
		//
		// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
		func (_Places *PlacesCallerSession) SupportsInterface( interfaceId [4]byte ) ( bool,  error) {
		  return _Places.Contract.SupportsInterface(&_Places.CallOpts , interfaceId)
		}
	
		// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
		//
		// Solidity: function symbol() view returns(string)
		func (_Places *PlacesCaller) Symbol(opts *bind.CallOpts ) (string, error) {
			var out []interface{}
			err := _Places.contract.Call(opts, &out, "symbol" )
			
			if err != nil {
				return *new(string),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(string)).(*string)
			
			return out0,  err
			
		}

		// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
		//
		// Solidity: function symbol() view returns(string)
		func (_Places *PlacesSession) Symbol() ( string,  error) {
		  return _Places.Contract.Symbol(&_Places.CallOpts )
		}

		// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
		//
		// Solidity: function symbol() view returns(string)
		func (_Places *PlacesCallerSession) Symbol() ( string,  error) {
		  return _Places.Contract.Symbol(&_Places.CallOpts )
		}
	
		// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
		//
		// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
		func (_Places *PlacesCaller) TokenByIndex(opts *bind.CallOpts , index *big.Int ) (*big.Int, error) {
			var out []interface{}
			err := _Places.contract.Call(opts, &out, "tokenByIndex" , index)
			
			if err != nil {
				return *new(*big.Int),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
			
			return out0,  err
			
		}

		// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
		//
		// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
		func (_Places *PlacesSession) TokenByIndex( index *big.Int ) ( *big.Int,  error) {
		  return _Places.Contract.TokenByIndex(&_Places.CallOpts , index)
		}

		// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
		//
		// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
		func (_Places *PlacesCallerSession) TokenByIndex( index *big.Int ) ( *big.Int,  error) {
		  return _Places.Contract.TokenByIndex(&_Places.CallOpts , index)
		}
	
		// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
		//
		// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
		func (_Places *PlacesCaller) TokenOfOwnerByIndex(opts *bind.CallOpts , owner common.Address , index *big.Int ) (*big.Int, error) {
			var out []interface{}
			err := _Places.contract.Call(opts, &out, "tokenOfOwnerByIndex" , owner, index)
			
			if err != nil {
				return *new(*big.Int),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
			
			return out0,  err
			
		}

		// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
		//
		// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
		func (_Places *PlacesSession) TokenOfOwnerByIndex( owner common.Address , index *big.Int ) ( *big.Int,  error) {
		  return _Places.Contract.TokenOfOwnerByIndex(&_Places.CallOpts , owner, index)
		}

		// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
		//
		// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
		func (_Places *PlacesCallerSession) TokenOfOwnerByIndex( owner common.Address , index *big.Int ) ( *big.Int,  error) {
		  return _Places.Contract.TokenOfOwnerByIndex(&_Places.CallOpts , owner, index)
		}
	
		// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
		//
		// Solidity: function tokenURI(uint256 tokenId) view returns(string)
		func (_Places *PlacesCaller) TokenURI(opts *bind.CallOpts , tokenId *big.Int ) (string, error) {
			var out []interface{}
			err := _Places.contract.Call(opts, &out, "tokenURI" , tokenId)
			
			if err != nil {
				return *new(string),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(string)).(*string)
			
			return out0,  err
			
		}

		// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
		//
		// Solidity: function tokenURI(uint256 tokenId) view returns(string)
		func (_Places *PlacesSession) TokenURI( tokenId *big.Int ) ( string,  error) {
		  return _Places.Contract.TokenURI(&_Places.CallOpts , tokenId)
		}

		// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
		//
		// Solidity: function tokenURI(uint256 tokenId) view returns(string)
		func (_Places *PlacesCallerSession) TokenURI( tokenId *big.Int ) ( string,  error) {
		  return _Places.Contract.TokenURI(&_Places.CallOpts , tokenId)
		}
	
		// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
		//
		// Solidity: function totalSupply() view returns(uint256)
		func (_Places *PlacesCaller) TotalSupply(opts *bind.CallOpts ) (*big.Int, error) {
			var out []interface{}
			err := _Places.contract.Call(opts, &out, "totalSupply" )
			
			if err != nil {
				return *new(*big.Int),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
			
			return out0,  err
			
		}

		// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
		//
		// Solidity: function totalSupply() view returns(uint256)
		func (_Places *PlacesSession) TotalSupply() ( *big.Int,  error) {
		  return _Places.Contract.TotalSupply(&_Places.CallOpts )
		}

		// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
		//
		// Solidity: function totalSupply() view returns(uint256)
		func (_Places *PlacesCallerSession) TotalSupply() ( *big.Int,  error) {
		  return _Places.Contract.TotalSupply(&_Places.CallOpts )
		}
	

	
		// AddGuest is a paid mutator transaction binding the contract method 0x5136f358.
		//
		// Solidity: function addGuest(address guestAddress) returns()
		func (_Places *PlacesTransactor) AddGuest(opts *bind.TransactOpts , guestAddress common.Address ) (*types.Transaction, error) {
			return _Places.contract.Transact(opts, "addGuest" , guestAddress)
		}

		// AddGuest is a paid mutator transaction binding the contract method 0x5136f358.
		//
		// Solidity: function addGuest(address guestAddress) returns()
		func (_Places *PlacesSession) AddGuest( guestAddress common.Address ) (*types.Transaction, error) {
		  return _Places.Contract.AddGuest(&_Places.TransactOpts , guestAddress)
		}

		// AddGuest is a paid mutator transaction binding the contract method 0x5136f358.
		//
		// Solidity: function addGuest(address guestAddress) returns()
		func (_Places *PlacesTransactorSession) AddGuest( guestAddress common.Address ) (*types.Transaction, error) {
		  return _Places.Contract.AddGuest(&_Places.TransactOpts , guestAddress)
		}
	
		// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
		//
		// Solidity: function approve(address to, uint256 tokenId) returns()
		func (_Places *PlacesTransactor) Approve(opts *bind.TransactOpts , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
			return _Places.contract.Transact(opts, "approve" , to, tokenId)
		}

		// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
		//
		// Solidity: function approve(address to, uint256 tokenId) returns()
		func (_Places *PlacesSession) Approve( to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _Places.Contract.Approve(&_Places.TransactOpts , to, tokenId)
		}

		// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
		//
		// Solidity: function approve(address to, uint256 tokenId) returns()
		func (_Places *PlacesTransactorSession) Approve( to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _Places.Contract.Approve(&_Places.TransactOpts , to, tokenId)
		}
	
		// Burn is a paid mutator transaction binding the contract method 0x42966c68.
		//
		// Solidity: function burn(uint256 tokenId) returns()
		func (_Places *PlacesTransactor) Burn(opts *bind.TransactOpts , tokenId *big.Int ) (*types.Transaction, error) {
			return _Places.contract.Transact(opts, "burn" , tokenId)
		}

		// Burn is a paid mutator transaction binding the contract method 0x42966c68.
		//
		// Solidity: function burn(uint256 tokenId) returns()
		func (_Places *PlacesSession) Burn( tokenId *big.Int ) (*types.Transaction, error) {
		  return _Places.Contract.Burn(&_Places.TransactOpts , tokenId)
		}

		// Burn is a paid mutator transaction binding the contract method 0x42966c68.
		//
		// Solidity: function burn(uint256 tokenId) returns()
		func (_Places *PlacesTransactorSession) Burn( tokenId *big.Int ) (*types.Transaction, error) {
		  return _Places.Contract.Burn(&_Places.TransactOpts , tokenId)
		}
	
		// GrounderMint is a paid mutator transaction binding the contract method 0x1f7bbae1.
		//
		// Solidity: function grounderMint(uint256 tokenId) returns(uint256)
		func (_Places *PlacesTransactor) GrounderMint(opts *bind.TransactOpts , tokenId *big.Int ) (*types.Transaction, error) {
			return _Places.contract.Transact(opts, "grounderMint" , tokenId)
		}

		// GrounderMint is a paid mutator transaction binding the contract method 0x1f7bbae1.
		//
		// Solidity: function grounderMint(uint256 tokenId) returns(uint256)
		func (_Places *PlacesSession) GrounderMint( tokenId *big.Int ) (*types.Transaction, error) {
		  return _Places.Contract.GrounderMint(&_Places.TransactOpts , tokenId)
		}

		// GrounderMint is a paid mutator transaction binding the contract method 0x1f7bbae1.
		//
		// Solidity: function grounderMint(uint256 tokenId) returns(uint256)
		func (_Places *PlacesTransactorSession) GrounderMint( tokenId *big.Int ) (*types.Transaction, error) {
		  return _Places.Contract.GrounderMint(&_Places.TransactOpts , tokenId)
		}
	
		// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
		//
		// Solidity: function mint() payable returns(uint256)
		func (_Places *PlacesTransactor) Mint(opts *bind.TransactOpts ) (*types.Transaction, error) {
			return _Places.contract.Transact(opts, "mint" )
		}

		// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
		//
		// Solidity: function mint() payable returns(uint256)
		func (_Places *PlacesSession) Mint() (*types.Transaction, error) {
		  return _Places.Contract.Mint(&_Places.TransactOpts )
		}

		// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
		//
		// Solidity: function mint() payable returns(uint256)
		func (_Places *PlacesTransactorSession) Mint() (*types.Transaction, error) {
		  return _Places.Contract.Mint(&_Places.TransactOpts )
		}
	
		// OwnerMint is a paid mutator transaction binding the contract method 0xf19e75d4.
		//
		// Solidity: function ownerMint(uint256 tokenId) returns(uint256)
		func (_Places *PlacesTransactor) OwnerMint(opts *bind.TransactOpts , tokenId *big.Int ) (*types.Transaction, error) {
			return _Places.contract.Transact(opts, "ownerMint" , tokenId)
		}

		// OwnerMint is a paid mutator transaction binding the contract method 0xf19e75d4.
		//
		// Solidity: function ownerMint(uint256 tokenId) returns(uint256)
		func (_Places *PlacesSession) OwnerMint( tokenId *big.Int ) (*types.Transaction, error) {
		  return _Places.Contract.OwnerMint(&_Places.TransactOpts , tokenId)
		}

		// OwnerMint is a paid mutator transaction binding the contract method 0xf19e75d4.
		//
		// Solidity: function ownerMint(uint256 tokenId) returns(uint256)
		func (_Places *PlacesTransactorSession) OwnerMint( tokenId *big.Int ) (*types.Transaction, error) {
		  return _Places.Contract.OwnerMint(&_Places.TransactOpts , tokenId)
		}
	
		// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
		//
		// Solidity: function renounceOwnership() returns()
		func (_Places *PlacesTransactor) RenounceOwnership(opts *bind.TransactOpts ) (*types.Transaction, error) {
			return _Places.contract.Transact(opts, "renounceOwnership" )
		}

		// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
		//
		// Solidity: function renounceOwnership() returns()
		func (_Places *PlacesSession) RenounceOwnership() (*types.Transaction, error) {
		  return _Places.Contract.RenounceOwnership(&_Places.TransactOpts )
		}

		// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
		//
		// Solidity: function renounceOwnership() returns()
		func (_Places *PlacesTransactorSession) RenounceOwnership() (*types.Transaction, error) {
		  return _Places.Contract.RenounceOwnership(&_Places.TransactOpts )
		}
	
		// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
		func (_Places *PlacesTransactor) SafeTransferFrom(opts *bind.TransactOpts , from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
			return _Places.contract.Transact(opts, "safeTransferFrom" , from, to, tokenId)
		}

		// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
		func (_Places *PlacesSession) SafeTransferFrom( from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _Places.Contract.SafeTransferFrom(&_Places.TransactOpts , from, to, tokenId)
		}

		// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
		func (_Places *PlacesTransactorSession) SafeTransferFrom( from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _Places.Contract.SafeTransferFrom(&_Places.TransactOpts , from, to, tokenId)
		}
	
		// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
		func (_Places *PlacesTransactor) SafeTransferFrom0(opts *bind.TransactOpts , from common.Address , to common.Address , tokenId *big.Int , _data []byte ) (*types.Transaction, error) {
			return _Places.contract.Transact(opts, "safeTransferFrom0" , from, to, tokenId, _data)
		}

		// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
		func (_Places *PlacesSession) SafeTransferFrom0( from common.Address , to common.Address , tokenId *big.Int , _data []byte ) (*types.Transaction, error) {
		  return _Places.Contract.SafeTransferFrom0(&_Places.TransactOpts , from, to, tokenId, _data)
		}

		// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
		func (_Places *PlacesTransactorSession) SafeTransferFrom0( from common.Address , to common.Address , tokenId *big.Int , _data []byte ) (*types.Transaction, error) {
		  return _Places.Contract.SafeTransferFrom0(&_Places.TransactOpts , from, to, tokenId, _data)
		}
	
		// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
		//
		// Solidity: function setApprovalForAll(address operator, bool approved) returns()
		func (_Places *PlacesTransactor) SetApprovalForAll(opts *bind.TransactOpts , operator common.Address , approved bool ) (*types.Transaction, error) {
			return _Places.contract.Transact(opts, "setApprovalForAll" , operator, approved)
		}

		// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
		//
		// Solidity: function setApprovalForAll(address operator, bool approved) returns()
		func (_Places *PlacesSession) SetApprovalForAll( operator common.Address , approved bool ) (*types.Transaction, error) {
		  return _Places.Contract.SetApprovalForAll(&_Places.TransactOpts , operator, approved)
		}

		// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
		//
		// Solidity: function setApprovalForAll(address operator, bool approved) returns()
		func (_Places *PlacesTransactorSession) SetApprovalForAll( operator common.Address , approved bool ) (*types.Transaction, error) {
		  return _Places.Contract.SetApprovalForAll(&_Places.TransactOpts , operator, approved)
		}
	
		// SetCounter is a paid mutator transaction binding the contract method 0x8bb5d9c3.
		//
		// Solidity: function setCounter(uint256 _countValue) returns()
		func (_Places *PlacesTransactor) SetCounter(opts *bind.TransactOpts , _countValue *big.Int ) (*types.Transaction, error) {
			return _Places.contract.Transact(opts, "setCounter" , _countValue)
		}

		// SetCounter is a paid mutator transaction binding the contract method 0x8bb5d9c3.
		//
		// Solidity: function setCounter(uint256 _countValue) returns()
		func (_Places *PlacesSession) SetCounter( _countValue *big.Int ) (*types.Transaction, error) {
		  return _Places.Contract.SetCounter(&_Places.TransactOpts , _countValue)
		}

		// SetCounter is a paid mutator transaction binding the contract method 0x8bb5d9c3.
		//
		// Solidity: function setCounter(uint256 _countValue) returns()
		func (_Places *PlacesTransactorSession) SetCounter( _countValue *big.Int ) (*types.Transaction, error) {
		  return _Places.Contract.SetCounter(&_Places.TransactOpts , _countValue)
		}
	
		// SetDescriptor is a paid mutator transaction binding the contract method 0x01b9a397.
		//
		// Solidity: function setDescriptor(address placesDescriptor) returns()
		func (_Places *PlacesTransactor) SetDescriptor(opts *bind.TransactOpts , placesDescriptor common.Address ) (*types.Transaction, error) {
			return _Places.contract.Transact(opts, "setDescriptor" , placesDescriptor)
		}

		// SetDescriptor is a paid mutator transaction binding the contract method 0x01b9a397.
		//
		// Solidity: function setDescriptor(address placesDescriptor) returns()
		func (_Places *PlacesSession) SetDescriptor( placesDescriptor common.Address ) (*types.Transaction, error) {
		  return _Places.Contract.SetDescriptor(&_Places.TransactOpts , placesDescriptor)
		}

		// SetDescriptor is a paid mutator transaction binding the contract method 0x01b9a397.
		//
		// Solidity: function setDescriptor(address placesDescriptor) returns()
		func (_Places *PlacesTransactorSession) SetDescriptor( placesDescriptor common.Address ) (*types.Transaction, error) {
		  return _Places.Contract.SetDescriptor(&_Places.TransactOpts , placesDescriptor)
		}
	
		// SetGrounders is a paid mutator transaction binding the contract method 0x0f753e91.
		//
		// Solidity: function setGrounders(address grounders) returns()
		func (_Places *PlacesTransactor) SetGrounders(opts *bind.TransactOpts , grounders common.Address ) (*types.Transaction, error) {
			return _Places.contract.Transact(opts, "setGrounders" , grounders)
		}

		// SetGrounders is a paid mutator transaction binding the contract method 0x0f753e91.
		//
		// Solidity: function setGrounders(address grounders) returns()
		func (_Places *PlacesSession) SetGrounders( grounders common.Address ) (*types.Transaction, error) {
		  return _Places.Contract.SetGrounders(&_Places.TransactOpts , grounders)
		}

		// SetGrounders is a paid mutator transaction binding the contract method 0x0f753e91.
		//
		// Solidity: function setGrounders(address grounders) returns()
		func (_Places *PlacesTransactorSession) SetGrounders( grounders common.Address ) (*types.Transaction, error) {
		  return _Places.Contract.SetGrounders(&_Places.TransactOpts , grounders)
		}
	
		// SetGroundersGifting is a paid mutator transaction binding the contract method 0x19779e28.
		//
		// Solidity: function setGroundersGifting(bool isGroundersGiftingEnabled) returns()
		func (_Places *PlacesTransactor) SetGroundersGifting(opts *bind.TransactOpts , isGroundersGiftingEnabled bool ) (*types.Transaction, error) {
			return _Places.contract.Transact(opts, "setGroundersGifting" , isGroundersGiftingEnabled)
		}

		// SetGroundersGifting is a paid mutator transaction binding the contract method 0x19779e28.
		//
		// Solidity: function setGroundersGifting(bool isGroundersGiftingEnabled) returns()
		func (_Places *PlacesSession) SetGroundersGifting( isGroundersGiftingEnabled bool ) (*types.Transaction, error) {
		  return _Places.Contract.SetGroundersGifting(&_Places.TransactOpts , isGroundersGiftingEnabled)
		}

		// SetGroundersGifting is a paid mutator transaction binding the contract method 0x19779e28.
		//
		// Solidity: function setGroundersGifting(bool isGroundersGiftingEnabled) returns()
		func (_Places *PlacesTransactorSession) SetGroundersGifting( isGroundersGiftingEnabled bool ) (*types.Transaction, error) {
		  return _Places.Contract.SetGroundersGifting(&_Places.TransactOpts , isGroundersGiftingEnabled)
		}
	
		// SetGuestlist is a paid mutator transaction binding the contract method 0x9bc4ff94.
		//
		// Solidity: function setGuestlist(address[] guestlist) returns()
		func (_Places *PlacesTransactor) SetGuestlist(opts *bind.TransactOpts , guestlist []common.Address ) (*types.Transaction, error) {
			return _Places.contract.Transact(opts, "setGuestlist" , guestlist)
		}

		// SetGuestlist is a paid mutator transaction binding the contract method 0x9bc4ff94.
		//
		// Solidity: function setGuestlist(address[] guestlist) returns()
		func (_Places *PlacesSession) SetGuestlist( guestlist []common.Address ) (*types.Transaction, error) {
		  return _Places.Contract.SetGuestlist(&_Places.TransactOpts , guestlist)
		}

		// SetGuestlist is a paid mutator transaction binding the contract method 0x9bc4ff94.
		//
		// Solidity: function setGuestlist(address[] guestlist) returns()
		func (_Places *PlacesTransactorSession) SetGuestlist( guestlist []common.Address ) (*types.Transaction, error) {
		  return _Places.Contract.SetGuestlist(&_Places.TransactOpts , guestlist)
		}
	
		// SetGuestlistEnabled is a paid mutator transaction binding the contract method 0xa2f4bb50.
		//
		// Solidity: function setGuestlistEnabled(bool isGuestlistEnabled) returns()
		func (_Places *PlacesTransactor) SetGuestlistEnabled(opts *bind.TransactOpts , isGuestlistEnabled bool ) (*types.Transaction, error) {
			return _Places.contract.Transact(opts, "setGuestlistEnabled" , isGuestlistEnabled)
		}

		// SetGuestlistEnabled is a paid mutator transaction binding the contract method 0xa2f4bb50.
		//
		// Solidity: function setGuestlistEnabled(bool isGuestlistEnabled) returns()
		func (_Places *PlacesSession) SetGuestlistEnabled( isGuestlistEnabled bool ) (*types.Transaction, error) {
		  return _Places.Contract.SetGuestlistEnabled(&_Places.TransactOpts , isGuestlistEnabled)
		}

		// SetGuestlistEnabled is a paid mutator transaction binding the contract method 0xa2f4bb50.
		//
		// Solidity: function setGuestlistEnabled(bool isGuestlistEnabled) returns()
		func (_Places *PlacesTransactorSession) SetGuestlistEnabled( isGuestlistEnabled bool ) (*types.Transaction, error) {
		  return _Places.Contract.SetGuestlistEnabled(&_Places.TransactOpts , isGuestlistEnabled)
		}
	
		// SetMintFee is a paid mutator transaction binding the contract method 0xeddd0d9c.
		//
		// Solidity: function setMintFee(uint256 mintFee) returns()
		func (_Places *PlacesTransactor) SetMintFee(opts *bind.TransactOpts , mintFee *big.Int ) (*types.Transaction, error) {
			return _Places.contract.Transact(opts, "setMintFee" , mintFee)
		}

		// SetMintFee is a paid mutator transaction binding the contract method 0xeddd0d9c.
		//
		// Solidity: function setMintFee(uint256 mintFee) returns()
		func (_Places *PlacesSession) SetMintFee( mintFee *big.Int ) (*types.Transaction, error) {
		  return _Places.Contract.SetMintFee(&_Places.TransactOpts , mintFee)
		}

		// SetMintFee is a paid mutator transaction binding the contract method 0xeddd0d9c.
		//
		// Solidity: function setMintFee(uint256 mintFee) returns()
		func (_Places *PlacesTransactorSession) SetMintFee( mintFee *big.Int ) (*types.Transaction, error) {
		  return _Places.Contract.SetMintFee(&_Places.TransactOpts , mintFee)
		}
	
		// SetNeighborhoodTreasury is a paid mutator transaction binding the contract method 0xa94193c0.
		//
		// Solidity: function setNeighborhoodTreasury(bool isNeighborhoodTreasuryEnabled) returns()
		func (_Places *PlacesTransactor) SetNeighborhoodTreasury(opts *bind.TransactOpts , isNeighborhoodTreasuryEnabled bool ) (*types.Transaction, error) {
			return _Places.contract.Transact(opts, "setNeighborhoodTreasury" , isNeighborhoodTreasuryEnabled)
		}

		// SetNeighborhoodTreasury is a paid mutator transaction binding the contract method 0xa94193c0.
		//
		// Solidity: function setNeighborhoodTreasury(bool isNeighborhoodTreasuryEnabled) returns()
		func (_Places *PlacesSession) SetNeighborhoodTreasury( isNeighborhoodTreasuryEnabled bool ) (*types.Transaction, error) {
		  return _Places.Contract.SetNeighborhoodTreasury(&_Places.TransactOpts , isNeighborhoodTreasuryEnabled)
		}

		// SetNeighborhoodTreasury is a paid mutator transaction binding the contract method 0xa94193c0.
		//
		// Solidity: function setNeighborhoodTreasury(bool isNeighborhoodTreasuryEnabled) returns()
		func (_Places *PlacesTransactorSession) SetNeighborhoodTreasury( isNeighborhoodTreasuryEnabled bool ) (*types.Transaction, error) {
		  return _Places.Contract.SetNeighborhoodTreasury(&_Places.TransactOpts , isNeighborhoodTreasuryEnabled)
		}
	
		// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
		//
		// Solidity: function setPaused(bool paused) returns()
		func (_Places *PlacesTransactor) SetPaused(opts *bind.TransactOpts , paused bool ) (*types.Transaction, error) {
			return _Places.contract.Transact(opts, "setPaused" , paused)
		}

		// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
		//
		// Solidity: function setPaused(bool paused) returns()
		func (_Places *PlacesSession) SetPaused( paused bool ) (*types.Transaction, error) {
		  return _Places.Contract.SetPaused(&_Places.TransactOpts , paused)
		}

		// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
		//
		// Solidity: function setPaused(bool paused) returns()
		func (_Places *PlacesTransactorSession) SetPaused( paused bool ) (*types.Transaction, error) {
		  return _Places.Contract.SetPaused(&_Places.TransactOpts , paused)
		}
	
		// SetPlacesDAO is a paid mutator transaction binding the contract method 0xa4b69593.
		//
		// Solidity: function setPlacesDAO(address placesDAO) returns()
		func (_Places *PlacesTransactor) SetPlacesDAO(opts *bind.TransactOpts , placesDAO common.Address ) (*types.Transaction, error) {
			return _Places.contract.Transact(opts, "setPlacesDAO" , placesDAO)
		}

		// SetPlacesDAO is a paid mutator transaction binding the contract method 0xa4b69593.
		//
		// Solidity: function setPlacesDAO(address placesDAO) returns()
		func (_Places *PlacesSession) SetPlacesDAO( placesDAO common.Address ) (*types.Transaction, error) {
		  return _Places.Contract.SetPlacesDAO(&_Places.TransactOpts , placesDAO)
		}

		// SetPlacesDAO is a paid mutator transaction binding the contract method 0xa4b69593.
		//
		// Solidity: function setPlacesDAO(address placesDAO) returns()
		func (_Places *PlacesTransactorSession) SetPlacesDAO( placesDAO common.Address ) (*types.Transaction, error) {
		  return _Places.Contract.SetPlacesDAO(&_Places.TransactOpts , placesDAO)
		}
	
		// SetPlacesProvider is a paid mutator transaction binding the contract method 0x9bd486c6.
		//
		// Solidity: function setPlacesProvider(address placesProvider) returns()
		func (_Places *PlacesTransactor) SetPlacesProvider(opts *bind.TransactOpts , placesProvider common.Address ) (*types.Transaction, error) {
			return _Places.contract.Transact(opts, "setPlacesProvider" , placesProvider)
		}

		// SetPlacesProvider is a paid mutator transaction binding the contract method 0x9bd486c6.
		//
		// Solidity: function setPlacesProvider(address placesProvider) returns()
		func (_Places *PlacesSession) SetPlacesProvider( placesProvider common.Address ) (*types.Transaction, error) {
		  return _Places.Contract.SetPlacesProvider(&_Places.TransactOpts , placesProvider)
		}

		// SetPlacesProvider is a paid mutator transaction binding the contract method 0x9bd486c6.
		//
		// Solidity: function setPlacesProvider(address placesProvider) returns()
		func (_Places *PlacesTransactorSession) SetPlacesProvider( placesProvider common.Address ) (*types.Transaction, error) {
		  return _Places.Contract.SetPlacesProvider(&_Places.TransactOpts , placesProvider)
		}
	
		// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
		//
		// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
		func (_Places *PlacesTransactor) TransferFrom(opts *bind.TransactOpts , from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
			return _Places.contract.Transact(opts, "transferFrom" , from, to, tokenId)
		}

		// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
		//
		// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
		func (_Places *PlacesSession) TransferFrom( from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _Places.Contract.TransferFrom(&_Places.TransactOpts , from, to, tokenId)
		}

		// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
		//
		// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
		func (_Places *PlacesTransactorSession) TransferFrom( from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _Places.Contract.TransferFrom(&_Places.TransactOpts , from, to, tokenId)
		}
	
		// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
		//
		// Solidity: function transferOwnership(address newOwner) returns()
		func (_Places *PlacesTransactor) TransferOwnership(opts *bind.TransactOpts , newOwner common.Address ) (*types.Transaction, error) {
			return _Places.contract.Transact(opts, "transferOwnership" , newOwner)
		}

		// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
		//
		// Solidity: function transferOwnership(address newOwner) returns()
		func (_Places *PlacesSession) TransferOwnership( newOwner common.Address ) (*types.Transaction, error) {
		  return _Places.Contract.TransferOwnership(&_Places.TransactOpts , newOwner)
		}

		// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
		//
		// Solidity: function transferOwnership(address newOwner) returns()
		func (_Places *PlacesTransactorSession) TransferOwnership( newOwner common.Address ) (*types.Transaction, error) {
		  return _Places.Contract.TransferOwnership(&_Places.TransactOpts , newOwner)
		}
	

	

	

	
		// PlacesApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Places contract.
		type PlacesApprovalIterator struct {
			Event *PlacesApproval // Event containing the contract specifics and raw log

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
		func (it *PlacesApprovalIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(PlacesApproval)
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
				it.Event = new(PlacesApproval)
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
		func (it *PlacesApprovalIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *PlacesApprovalIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// PlacesApproval represents a Approval event raised by the Places contract.
		type PlacesApproval struct { 
			Owner common.Address; 
			Approved common.Address; 
			TokenId *big.Int; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
		//
		// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
 		func (_Places *PlacesFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*PlacesApprovalIterator, error) {
			
			var ownerRule []interface{}
			for _, ownerItem := range owner {
				ownerRule = append(ownerRule, ownerItem)
			}
			var approvedRule []interface{}
			for _, approvedItem := range approved {
				approvedRule = append(approvedRule, approvedItem)
			}
			var tokenIdRule []interface{}
			for _, tokenIdItem := range tokenId {
				tokenIdRule = append(tokenIdRule, tokenIdItem)
			}

			logs, sub, err := _Places.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
			if err != nil {
				return nil, err
			}
			return &PlacesApprovalIterator{contract: _Places.contract, event: "Approval", logs: logs, sub: sub}, nil
 		}

		// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
		//
		// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
		func (_Places *PlacesFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PlacesApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {
			
			var ownerRule []interface{}
			for _, ownerItem := range owner {
				ownerRule = append(ownerRule, ownerItem)
			}
			var approvedRule []interface{}
			for _, approvedItem := range approved {
				approvedRule = append(approvedRule, approvedItem)
			}
			var tokenIdRule []interface{}
			for _, tokenIdItem := range tokenId {
				tokenIdRule = append(tokenIdRule, tokenIdItem)
			}

			logs, sub, err := _Places.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(PlacesApproval)
						if err := _Places.contract.UnpackLog(event, "Approval", log); err != nil {
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
		// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
		func (_Places *PlacesFilterer) ParseApproval(log types.Log) (*PlacesApproval, error) {
			event := new(PlacesApproval)
			if err := _Places.contract.UnpackLog(event, "Approval", log); err != nil {
				return nil, err
			}
			event.Raw = log
			return event, nil
		}

 	
		// PlacesApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Places contract.
		type PlacesApprovalForAllIterator struct {
			Event *PlacesApprovalForAll // Event containing the contract specifics and raw log

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
		func (it *PlacesApprovalForAllIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(PlacesApprovalForAll)
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
				it.Event = new(PlacesApprovalForAll)
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
		func (it *PlacesApprovalForAllIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *PlacesApprovalForAllIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// PlacesApprovalForAll represents a ApprovalForAll event raised by the Places contract.
		type PlacesApprovalForAll struct { 
			Owner common.Address; 
			Operator common.Address; 
			Approved bool; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
		//
		// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
 		func (_Places *PlacesFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*PlacesApprovalForAllIterator, error) {
			
			var ownerRule []interface{}
			for _, ownerItem := range owner {
				ownerRule = append(ownerRule, ownerItem)
			}
			var operatorRule []interface{}
			for _, operatorItem := range operator {
				operatorRule = append(operatorRule, operatorItem)
			}
			

			logs, sub, err := _Places.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
			if err != nil {
				return nil, err
			}
			return &PlacesApprovalForAllIterator{contract: _Places.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
 		}

		// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
		//
		// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
		func (_Places *PlacesFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *PlacesApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {
			
			var ownerRule []interface{}
			for _, ownerItem := range owner {
				ownerRule = append(ownerRule, ownerItem)
			}
			var operatorRule []interface{}
			for _, operatorItem := range operator {
				operatorRule = append(operatorRule, operatorItem)
			}
			

			logs, sub, err := _Places.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(PlacesApprovalForAll)
						if err := _Places.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

		// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
		//
		// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
		func (_Places *PlacesFilterer) ParseApprovalForAll(log types.Log) (*PlacesApprovalForAll, error) {
			event := new(PlacesApprovalForAll)
			if err := _Places.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
				return nil, err
			}
			event.Raw = log
			return event, nil
		}

 	
		// PlacesDescriptorUpdatedIterator is returned from FilterDescriptorUpdated and is used to iterate over the raw logs and unpacked data for DescriptorUpdated events raised by the Places contract.
		type PlacesDescriptorUpdatedIterator struct {
			Event *PlacesDescriptorUpdated // Event containing the contract specifics and raw log

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
		func (it *PlacesDescriptorUpdatedIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(PlacesDescriptorUpdated)
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
				it.Event = new(PlacesDescriptorUpdated)
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
		func (it *PlacesDescriptorUpdatedIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *PlacesDescriptorUpdatedIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// PlacesDescriptorUpdated represents a DescriptorUpdated event raised by the Places contract.
		type PlacesDescriptorUpdated struct { 
			Descriptor common.Address; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterDescriptorUpdated is a free log retrieval operation binding the contract event 0x6e66ab22238a5471005895947c8f57db923c2a9c9c73180eff80864c21295c1b.
		//
		// Solidity: event DescriptorUpdated(address descriptor)
 		func (_Places *PlacesFilterer) FilterDescriptorUpdated(opts *bind.FilterOpts) (*PlacesDescriptorUpdatedIterator, error) {
			
			

			logs, sub, err := _Places.contract.FilterLogs(opts, "DescriptorUpdated")
			if err != nil {
				return nil, err
			}
			return &PlacesDescriptorUpdatedIterator{contract: _Places.contract, event: "DescriptorUpdated", logs: logs, sub: sub}, nil
 		}

		// WatchDescriptorUpdated is a free log subscription operation binding the contract event 0x6e66ab22238a5471005895947c8f57db923c2a9c9c73180eff80864c21295c1b.
		//
		// Solidity: event DescriptorUpdated(address descriptor)
		func (_Places *PlacesFilterer) WatchDescriptorUpdated(opts *bind.WatchOpts, sink chan<- *PlacesDescriptorUpdated) (event.Subscription, error) {
			
			

			logs, sub, err := _Places.contract.WatchLogs(opts, "DescriptorUpdated")
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(PlacesDescriptorUpdated)
						if err := _Places.contract.UnpackLog(event, "DescriptorUpdated", log); err != nil {
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

		// ParseDescriptorUpdated is a log parse operation binding the contract event 0x6e66ab22238a5471005895947c8f57db923c2a9c9c73180eff80864c21295c1b.
		//
		// Solidity: event DescriptorUpdated(address descriptor)
		func (_Places *PlacesFilterer) ParseDescriptorUpdated(log types.Log) (*PlacesDescriptorUpdated, error) {
			event := new(PlacesDescriptorUpdated)
			if err := _Places.contract.UnpackLog(event, "DescriptorUpdated", log); err != nil {
				return nil, err
			}
			event.Raw = log
			return event, nil
		}

 	
		// PlacesGroundersGiftingEnabledUpdatedIterator is returned from FilterGroundersGiftingEnabledUpdated and is used to iterate over the raw logs and unpacked data for GroundersGiftingEnabledUpdated events raised by the Places contract.
		type PlacesGroundersGiftingEnabledUpdatedIterator struct {
			Event *PlacesGroundersGiftingEnabledUpdated // Event containing the contract specifics and raw log

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
		func (it *PlacesGroundersGiftingEnabledUpdatedIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(PlacesGroundersGiftingEnabledUpdated)
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
				it.Event = new(PlacesGroundersGiftingEnabledUpdated)
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
		func (it *PlacesGroundersGiftingEnabledUpdatedIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *PlacesGroundersGiftingEnabledUpdatedIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// PlacesGroundersGiftingEnabledUpdated represents a GroundersGiftingEnabledUpdated event raised by the Places contract.
		type PlacesGroundersGiftingEnabledUpdated struct { 
			IsEnabled bool; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterGroundersGiftingEnabledUpdated is a free log retrieval operation binding the contract event 0x80585b6bb14f170e6963473df6678471c8b223a1073f77c338ed55245a4830e7.
		//
		// Solidity: event GroundersGiftingEnabledUpdated(bool isEnabled)
 		func (_Places *PlacesFilterer) FilterGroundersGiftingEnabledUpdated(opts *bind.FilterOpts) (*PlacesGroundersGiftingEnabledUpdatedIterator, error) {
			
			

			logs, sub, err := _Places.contract.FilterLogs(opts, "GroundersGiftingEnabledUpdated")
			if err != nil {
				return nil, err
			}
			return &PlacesGroundersGiftingEnabledUpdatedIterator{contract: _Places.contract, event: "GroundersGiftingEnabledUpdated", logs: logs, sub: sub}, nil
 		}

		// WatchGroundersGiftingEnabledUpdated is a free log subscription operation binding the contract event 0x80585b6bb14f170e6963473df6678471c8b223a1073f77c338ed55245a4830e7.
		//
		// Solidity: event GroundersGiftingEnabledUpdated(bool isEnabled)
		func (_Places *PlacesFilterer) WatchGroundersGiftingEnabledUpdated(opts *bind.WatchOpts, sink chan<- *PlacesGroundersGiftingEnabledUpdated) (event.Subscription, error) {
			
			

			logs, sub, err := _Places.contract.WatchLogs(opts, "GroundersGiftingEnabledUpdated")
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(PlacesGroundersGiftingEnabledUpdated)
						if err := _Places.contract.UnpackLog(event, "GroundersGiftingEnabledUpdated", log); err != nil {
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

		// ParseGroundersGiftingEnabledUpdated is a log parse operation binding the contract event 0x80585b6bb14f170e6963473df6678471c8b223a1073f77c338ed55245a4830e7.
		//
		// Solidity: event GroundersGiftingEnabledUpdated(bool isEnabled)
		func (_Places *PlacesFilterer) ParseGroundersGiftingEnabledUpdated(log types.Log) (*PlacesGroundersGiftingEnabledUpdated, error) {
			event := new(PlacesGroundersGiftingEnabledUpdated)
			if err := _Places.contract.UnpackLog(event, "GroundersGiftingEnabledUpdated", log); err != nil {
				return nil, err
			}
			event.Raw = log
			return event, nil
		}

 	
		// PlacesGroundersUpdatedIterator is returned from FilterGroundersUpdated and is used to iterate over the raw logs and unpacked data for GroundersUpdated events raised by the Places contract.
		type PlacesGroundersUpdatedIterator struct {
			Event *PlacesGroundersUpdated // Event containing the contract specifics and raw log

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
		func (it *PlacesGroundersUpdatedIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(PlacesGroundersUpdated)
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
				it.Event = new(PlacesGroundersUpdated)
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
		func (it *PlacesGroundersUpdatedIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *PlacesGroundersUpdatedIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// PlacesGroundersUpdated represents a GroundersUpdated event raised by the Places contract.
		type PlacesGroundersUpdated struct { 
			Grounders common.Address; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterGroundersUpdated is a free log retrieval operation binding the contract event 0xbacd45cc8ada63864f7a9f3480bcf83a4cee8bb9af58eaa40b32b9f039124956.
		//
		// Solidity: event GroundersUpdated(address grounders)
 		func (_Places *PlacesFilterer) FilterGroundersUpdated(opts *bind.FilterOpts) (*PlacesGroundersUpdatedIterator, error) {
			
			

			logs, sub, err := _Places.contract.FilterLogs(opts, "GroundersUpdated")
			if err != nil {
				return nil, err
			}
			return &PlacesGroundersUpdatedIterator{contract: _Places.contract, event: "GroundersUpdated", logs: logs, sub: sub}, nil
 		}

		// WatchGroundersUpdated is a free log subscription operation binding the contract event 0xbacd45cc8ada63864f7a9f3480bcf83a4cee8bb9af58eaa40b32b9f039124956.
		//
		// Solidity: event GroundersUpdated(address grounders)
		func (_Places *PlacesFilterer) WatchGroundersUpdated(opts *bind.WatchOpts, sink chan<- *PlacesGroundersUpdated) (event.Subscription, error) {
			
			

			logs, sub, err := _Places.contract.WatchLogs(opts, "GroundersUpdated")
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(PlacesGroundersUpdated)
						if err := _Places.contract.UnpackLog(event, "GroundersUpdated", log); err != nil {
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

		// ParseGroundersUpdated is a log parse operation binding the contract event 0xbacd45cc8ada63864f7a9f3480bcf83a4cee8bb9af58eaa40b32b9f039124956.
		//
		// Solidity: event GroundersUpdated(address grounders)
		func (_Places *PlacesFilterer) ParseGroundersUpdated(log types.Log) (*PlacesGroundersUpdated, error) {
			event := new(PlacesGroundersUpdated)
			if err := _Places.contract.UnpackLog(event, "GroundersUpdated", log); err != nil {
				return nil, err
			}
			event.Raw = log
			return event, nil
		}

 	
		// PlacesGuestlistEnabledUpdatedIterator is returned from FilterGuestlistEnabledUpdated and is used to iterate over the raw logs and unpacked data for GuestlistEnabledUpdated events raised by the Places contract.
		type PlacesGuestlistEnabledUpdatedIterator struct {
			Event *PlacesGuestlistEnabledUpdated // Event containing the contract specifics and raw log

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
		func (it *PlacesGuestlistEnabledUpdatedIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(PlacesGuestlistEnabledUpdated)
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
				it.Event = new(PlacesGuestlistEnabledUpdated)
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
		func (it *PlacesGuestlistEnabledUpdatedIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *PlacesGuestlistEnabledUpdatedIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// PlacesGuestlistEnabledUpdated represents a GuestlistEnabledUpdated event raised by the Places contract.
		type PlacesGuestlistEnabledUpdated struct { 
			IsEnabled bool; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterGuestlistEnabledUpdated is a free log retrieval operation binding the contract event 0x8f7a913faf0f8b0c6a28c72de0f1931aad9e81e8f6ea3b73e185159559950237.
		//
		// Solidity: event GuestlistEnabledUpdated(bool isEnabled)
 		func (_Places *PlacesFilterer) FilterGuestlistEnabledUpdated(opts *bind.FilterOpts) (*PlacesGuestlistEnabledUpdatedIterator, error) {
			
			

			logs, sub, err := _Places.contract.FilterLogs(opts, "GuestlistEnabledUpdated")
			if err != nil {
				return nil, err
			}
			return &PlacesGuestlistEnabledUpdatedIterator{contract: _Places.contract, event: "GuestlistEnabledUpdated", logs: logs, sub: sub}, nil
 		}

		// WatchGuestlistEnabledUpdated is a free log subscription operation binding the contract event 0x8f7a913faf0f8b0c6a28c72de0f1931aad9e81e8f6ea3b73e185159559950237.
		//
		// Solidity: event GuestlistEnabledUpdated(bool isEnabled)
		func (_Places *PlacesFilterer) WatchGuestlistEnabledUpdated(opts *bind.WatchOpts, sink chan<- *PlacesGuestlistEnabledUpdated) (event.Subscription, error) {
			
			

			logs, sub, err := _Places.contract.WatchLogs(opts, "GuestlistEnabledUpdated")
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(PlacesGuestlistEnabledUpdated)
						if err := _Places.contract.UnpackLog(event, "GuestlistEnabledUpdated", log); err != nil {
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

		// ParseGuestlistEnabledUpdated is a log parse operation binding the contract event 0x8f7a913faf0f8b0c6a28c72de0f1931aad9e81e8f6ea3b73e185159559950237.
		//
		// Solidity: event GuestlistEnabledUpdated(bool isEnabled)
		func (_Places *PlacesFilterer) ParseGuestlistEnabledUpdated(log types.Log) (*PlacesGuestlistEnabledUpdated, error) {
			event := new(PlacesGuestlistEnabledUpdated)
			if err := _Places.contract.UnpackLog(event, "GuestlistEnabledUpdated", log); err != nil {
				return nil, err
			}
			event.Raw = log
			return event, nil
		}

 	
		// PlacesMintFeeUpdatedIterator is returned from FilterMintFeeUpdated and is used to iterate over the raw logs and unpacked data for MintFeeUpdated events raised by the Places contract.
		type PlacesMintFeeUpdatedIterator struct {
			Event *PlacesMintFeeUpdated // Event containing the contract specifics and raw log

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
		func (it *PlacesMintFeeUpdatedIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(PlacesMintFeeUpdated)
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
				it.Event = new(PlacesMintFeeUpdated)
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
		func (it *PlacesMintFeeUpdatedIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *PlacesMintFeeUpdatedIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// PlacesMintFeeUpdated represents a MintFeeUpdated event raised by the Places contract.
		type PlacesMintFeeUpdated struct { 
			MintFee *big.Int; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterMintFeeUpdated is a free log retrieval operation binding the contract event 0x38fbb1c8b109c430f0c030e7ed076cf5611a307773a4e8e365601e8f8bceaec6.
		//
		// Solidity: event MintFeeUpdated(uint256 mintFee)
 		func (_Places *PlacesFilterer) FilterMintFeeUpdated(opts *bind.FilterOpts) (*PlacesMintFeeUpdatedIterator, error) {
			
			

			logs, sub, err := _Places.contract.FilterLogs(opts, "MintFeeUpdated")
			if err != nil {
				return nil, err
			}
			return &PlacesMintFeeUpdatedIterator{contract: _Places.contract, event: "MintFeeUpdated", logs: logs, sub: sub}, nil
 		}

		// WatchMintFeeUpdated is a free log subscription operation binding the contract event 0x38fbb1c8b109c430f0c030e7ed076cf5611a307773a4e8e365601e8f8bceaec6.
		//
		// Solidity: event MintFeeUpdated(uint256 mintFee)
		func (_Places *PlacesFilterer) WatchMintFeeUpdated(opts *bind.WatchOpts, sink chan<- *PlacesMintFeeUpdated) (event.Subscription, error) {
			
			

			logs, sub, err := _Places.contract.WatchLogs(opts, "MintFeeUpdated")
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(PlacesMintFeeUpdated)
						if err := _Places.contract.UnpackLog(event, "MintFeeUpdated", log); err != nil {
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

		// ParseMintFeeUpdated is a log parse operation binding the contract event 0x38fbb1c8b109c430f0c030e7ed076cf5611a307773a4e8e365601e8f8bceaec6.
		//
		// Solidity: event MintFeeUpdated(uint256 mintFee)
		func (_Places *PlacesFilterer) ParseMintFeeUpdated(log types.Log) (*PlacesMintFeeUpdated, error) {
			event := new(PlacesMintFeeUpdated)
			if err := _Places.contract.UnpackLog(event, "MintFeeUpdated", log); err != nil {
				return nil, err
			}
			event.Raw = log
			return event, nil
		}

 	
		// PlacesNeighborhoodTreasuryEnabledUpdatedIterator is returned from FilterNeighborhoodTreasuryEnabledUpdated and is used to iterate over the raw logs and unpacked data for NeighborhoodTreasuryEnabledUpdated events raised by the Places contract.
		type PlacesNeighborhoodTreasuryEnabledUpdatedIterator struct {
			Event *PlacesNeighborhoodTreasuryEnabledUpdated // Event containing the contract specifics and raw log

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
		func (it *PlacesNeighborhoodTreasuryEnabledUpdatedIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(PlacesNeighborhoodTreasuryEnabledUpdated)
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
				it.Event = new(PlacesNeighborhoodTreasuryEnabledUpdated)
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
		func (it *PlacesNeighborhoodTreasuryEnabledUpdatedIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *PlacesNeighborhoodTreasuryEnabledUpdatedIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// PlacesNeighborhoodTreasuryEnabledUpdated represents a NeighborhoodTreasuryEnabledUpdated event raised by the Places contract.
		type PlacesNeighborhoodTreasuryEnabledUpdated struct { 
			IsEnabled bool; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterNeighborhoodTreasuryEnabledUpdated is a free log retrieval operation binding the contract event 0x6cee79bc3cf75bcfff040337779718773d202789b27662681efccf53d649e9d6.
		//
		// Solidity: event NeighborhoodTreasuryEnabledUpdated(bool isEnabled)
 		func (_Places *PlacesFilterer) FilterNeighborhoodTreasuryEnabledUpdated(opts *bind.FilterOpts) (*PlacesNeighborhoodTreasuryEnabledUpdatedIterator, error) {
			
			

			logs, sub, err := _Places.contract.FilterLogs(opts, "NeighborhoodTreasuryEnabledUpdated")
			if err != nil {
				return nil, err
			}
			return &PlacesNeighborhoodTreasuryEnabledUpdatedIterator{contract: _Places.contract, event: "NeighborhoodTreasuryEnabledUpdated", logs: logs, sub: sub}, nil
 		}

		// WatchNeighborhoodTreasuryEnabledUpdated is a free log subscription operation binding the contract event 0x6cee79bc3cf75bcfff040337779718773d202789b27662681efccf53d649e9d6.
		//
		// Solidity: event NeighborhoodTreasuryEnabledUpdated(bool isEnabled)
		func (_Places *PlacesFilterer) WatchNeighborhoodTreasuryEnabledUpdated(opts *bind.WatchOpts, sink chan<- *PlacesNeighborhoodTreasuryEnabledUpdated) (event.Subscription, error) {
			
			

			logs, sub, err := _Places.contract.WatchLogs(opts, "NeighborhoodTreasuryEnabledUpdated")
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(PlacesNeighborhoodTreasuryEnabledUpdated)
						if err := _Places.contract.UnpackLog(event, "NeighborhoodTreasuryEnabledUpdated", log); err != nil {
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

		// ParseNeighborhoodTreasuryEnabledUpdated is a log parse operation binding the contract event 0x6cee79bc3cf75bcfff040337779718773d202789b27662681efccf53d649e9d6.
		//
		// Solidity: event NeighborhoodTreasuryEnabledUpdated(bool isEnabled)
		func (_Places *PlacesFilterer) ParseNeighborhoodTreasuryEnabledUpdated(log types.Log) (*PlacesNeighborhoodTreasuryEnabledUpdated, error) {
			event := new(PlacesNeighborhoodTreasuryEnabledUpdated)
			if err := _Places.contract.UnpackLog(event, "NeighborhoodTreasuryEnabledUpdated", log); err != nil {
				return nil, err
			}
			event.Raw = log
			return event, nil
		}

 	
		// PlacesOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Places contract.
		type PlacesOwnershipTransferredIterator struct {
			Event *PlacesOwnershipTransferred // Event containing the contract specifics and raw log

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
		func (it *PlacesOwnershipTransferredIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(PlacesOwnershipTransferred)
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
				it.Event = new(PlacesOwnershipTransferred)
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
		func (it *PlacesOwnershipTransferredIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *PlacesOwnershipTransferredIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// PlacesOwnershipTransferred represents a OwnershipTransferred event raised by the Places contract.
		type PlacesOwnershipTransferred struct { 
			PreviousOwner common.Address; 
			NewOwner common.Address; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
		//
		// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
 		func (_Places *PlacesFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PlacesOwnershipTransferredIterator, error) {
			
			var previousOwnerRule []interface{}
			for _, previousOwnerItem := range previousOwner {
				previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
			}
			var newOwnerRule []interface{}
			for _, newOwnerItem := range newOwner {
				newOwnerRule = append(newOwnerRule, newOwnerItem)
			}

			logs, sub, err := _Places.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
			if err != nil {
				return nil, err
			}
			return &PlacesOwnershipTransferredIterator{contract: _Places.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
 		}

		// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
		//
		// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
		func (_Places *PlacesFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PlacesOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {
			
			var previousOwnerRule []interface{}
			for _, previousOwnerItem := range previousOwner {
				previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
			}
			var newOwnerRule []interface{}
			for _, newOwnerItem := range newOwner {
				newOwnerRule = append(newOwnerRule, newOwnerItem)
			}

			logs, sub, err := _Places.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(PlacesOwnershipTransferred)
						if err := _Places.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
		func (_Places *PlacesFilterer) ParseOwnershipTransferred(log types.Log) (*PlacesOwnershipTransferred, error) {
			event := new(PlacesOwnershipTransferred)
			if err := _Places.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
				return nil, err
			}
			event.Raw = log
			return event, nil
		}

 	
		// PlacesPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Places contract.
		type PlacesPausedIterator struct {
			Event *PlacesPaused // Event containing the contract specifics and raw log

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
		func (it *PlacesPausedIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(PlacesPaused)
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
				it.Event = new(PlacesPaused)
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
		func (it *PlacesPausedIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *PlacesPausedIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// PlacesPaused represents a Paused event raised by the Places contract.
		type PlacesPaused struct { 
			Account common.Address; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
		//
		// Solidity: event Paused(address account)
 		func (_Places *PlacesFilterer) FilterPaused(opts *bind.FilterOpts) (*PlacesPausedIterator, error) {
			
			

			logs, sub, err := _Places.contract.FilterLogs(opts, "Paused")
			if err != nil {
				return nil, err
			}
			return &PlacesPausedIterator{contract: _Places.contract, event: "Paused", logs: logs, sub: sub}, nil
 		}

		// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
		//
		// Solidity: event Paused(address account)
		func (_Places *PlacesFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PlacesPaused) (event.Subscription, error) {
			
			

			logs, sub, err := _Places.contract.WatchLogs(opts, "Paused")
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(PlacesPaused)
						if err := _Places.contract.UnpackLog(event, "Paused", log); err != nil {
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

		// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
		//
		// Solidity: event Paused(address account)
		func (_Places *PlacesFilterer) ParsePaused(log types.Log) (*PlacesPaused, error) {
			event := new(PlacesPaused)
			if err := _Places.contract.UnpackLog(event, "Paused", log); err != nil {
				return nil, err
			}
			event.Raw = log
			return event, nil
		}

 	
		// PlacesPlaceBurnedIterator is returned from FilterPlaceBurned and is used to iterate over the raw logs and unpacked data for PlaceBurned events raised by the Places contract.
		type PlacesPlaceBurnedIterator struct {
			Event *PlacesPlaceBurned // Event containing the contract specifics and raw log

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
		func (it *PlacesPlaceBurnedIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(PlacesPlaceBurned)
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
				it.Event = new(PlacesPlaceBurned)
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
		func (it *PlacesPlaceBurnedIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *PlacesPlaceBurnedIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// PlacesPlaceBurned represents a PlaceBurned event raised by the Places contract.
		type PlacesPlaceBurned struct { 
			TokenId *big.Int; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterPlaceBurned is a free log retrieval operation binding the contract event 0x8bb0e8d87ef33e1cabb621aa61fa4f55b273d0360e21eb194536a020c4a027a1.
		//
		// Solidity: event PlaceBurned(uint256 tokenId)
 		func (_Places *PlacesFilterer) FilterPlaceBurned(opts *bind.FilterOpts) (*PlacesPlaceBurnedIterator, error) {
			
			

			logs, sub, err := _Places.contract.FilterLogs(opts, "PlaceBurned")
			if err != nil {
				return nil, err
			}
			return &PlacesPlaceBurnedIterator{contract: _Places.contract, event: "PlaceBurned", logs: logs, sub: sub}, nil
 		}

		// WatchPlaceBurned is a free log subscription operation binding the contract event 0x8bb0e8d87ef33e1cabb621aa61fa4f55b273d0360e21eb194536a020c4a027a1.
		//
		// Solidity: event PlaceBurned(uint256 tokenId)
		func (_Places *PlacesFilterer) WatchPlaceBurned(opts *bind.WatchOpts, sink chan<- *PlacesPlaceBurned) (event.Subscription, error) {
			
			

			logs, sub, err := _Places.contract.WatchLogs(opts, "PlaceBurned")
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(PlacesPlaceBurned)
						if err := _Places.contract.UnpackLog(event, "PlaceBurned", log); err != nil {
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

		// ParsePlaceBurned is a log parse operation binding the contract event 0x8bb0e8d87ef33e1cabb621aa61fa4f55b273d0360e21eb194536a020c4a027a1.
		//
		// Solidity: event PlaceBurned(uint256 tokenId)
		func (_Places *PlacesFilterer) ParsePlaceBurned(log types.Log) (*PlacesPlaceBurned, error) {
			event := new(PlacesPlaceBurned)
			if err := _Places.contract.UnpackLog(event, "PlaceBurned", log); err != nil {
				return nil, err
			}
			event.Raw = log
			return event, nil
		}

 	
		// PlacesPlaceCreatedIterator is returned from FilterPlaceCreated and is used to iterate over the raw logs and unpacked data for PlaceCreated events raised by the Places contract.
		type PlacesPlaceCreatedIterator struct {
			Event *PlacesPlaceCreated // Event containing the contract specifics and raw log

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
		func (it *PlacesPlaceCreatedIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(PlacesPlaceCreated)
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
				it.Event = new(PlacesPlaceCreated)
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
		func (it *PlacesPlaceCreatedIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *PlacesPlaceCreatedIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// PlacesPlaceCreated represents a PlaceCreated event raised by the Places contract.
		type PlacesPlaceCreated struct { 
			TokenId *big.Int; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterPlaceCreated is a free log retrieval operation binding the contract event 0x9e54f6a45e7b311f3418d55268846484cfff4a9b737ecd127179d18469396a00.
		//
		// Solidity: event PlaceCreated(uint256 tokenId)
 		func (_Places *PlacesFilterer) FilterPlaceCreated(opts *bind.FilterOpts) (*PlacesPlaceCreatedIterator, error) {
			
			

			logs, sub, err := _Places.contract.FilterLogs(opts, "PlaceCreated")
			if err != nil {
				return nil, err
			}
			return &PlacesPlaceCreatedIterator{contract: _Places.contract, event: "PlaceCreated", logs: logs, sub: sub}, nil
 		}

		// WatchPlaceCreated is a free log subscription operation binding the contract event 0x9e54f6a45e7b311f3418d55268846484cfff4a9b737ecd127179d18469396a00.
		//
		// Solidity: event PlaceCreated(uint256 tokenId)
		func (_Places *PlacesFilterer) WatchPlaceCreated(opts *bind.WatchOpts, sink chan<- *PlacesPlaceCreated) (event.Subscription, error) {
			
			

			logs, sub, err := _Places.contract.WatchLogs(opts, "PlaceCreated")
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(PlacesPlaceCreated)
						if err := _Places.contract.UnpackLog(event, "PlaceCreated", log); err != nil {
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

		// ParsePlaceCreated is a log parse operation binding the contract event 0x9e54f6a45e7b311f3418d55268846484cfff4a9b737ecd127179d18469396a00.
		//
		// Solidity: event PlaceCreated(uint256 tokenId)
		func (_Places *PlacesFilterer) ParsePlaceCreated(log types.Log) (*PlacesPlaceCreated, error) {
			event := new(PlacesPlaceCreated)
			if err := _Places.contract.UnpackLog(event, "PlaceCreated", log); err != nil {
				return nil, err
			}
			event.Raw = log
			return event, nil
		}

 	
		// PlacesPlacesDAOUpdatedIterator is returned from FilterPlacesDAOUpdated and is used to iterate over the raw logs and unpacked data for PlacesDAOUpdated events raised by the Places contract.
		type PlacesPlacesDAOUpdatedIterator struct {
			Event *PlacesPlacesDAOUpdated // Event containing the contract specifics and raw log

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
		func (it *PlacesPlacesDAOUpdatedIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(PlacesPlacesDAOUpdated)
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
				it.Event = new(PlacesPlacesDAOUpdated)
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
		func (it *PlacesPlacesDAOUpdatedIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *PlacesPlacesDAOUpdatedIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// PlacesPlacesDAOUpdated represents a PlacesDAOUpdated event raised by the Places contract.
		type PlacesPlacesDAOUpdated struct { 
			PlacesDao common.Address; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterPlacesDAOUpdated is a free log retrieval operation binding the contract event 0xc0a55d29aab2534e16123c7af21cc636ba3c5f34f7d2885dc479b3f2b049d70f.
		//
		// Solidity: event PlacesDAOUpdated(address placesDao)
 		func (_Places *PlacesFilterer) FilterPlacesDAOUpdated(opts *bind.FilterOpts) (*PlacesPlacesDAOUpdatedIterator, error) {
			
			

			logs, sub, err := _Places.contract.FilterLogs(opts, "PlacesDAOUpdated")
			if err != nil {
				return nil, err
			}
			return &PlacesPlacesDAOUpdatedIterator{contract: _Places.contract, event: "PlacesDAOUpdated", logs: logs, sub: sub}, nil
 		}

		// WatchPlacesDAOUpdated is a free log subscription operation binding the contract event 0xc0a55d29aab2534e16123c7af21cc636ba3c5f34f7d2885dc479b3f2b049d70f.
		//
		// Solidity: event PlacesDAOUpdated(address placesDao)
		func (_Places *PlacesFilterer) WatchPlacesDAOUpdated(opts *bind.WatchOpts, sink chan<- *PlacesPlacesDAOUpdated) (event.Subscription, error) {
			
			

			logs, sub, err := _Places.contract.WatchLogs(opts, "PlacesDAOUpdated")
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(PlacesPlacesDAOUpdated)
						if err := _Places.contract.UnpackLog(event, "PlacesDAOUpdated", log); err != nil {
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

		// ParsePlacesDAOUpdated is a log parse operation binding the contract event 0xc0a55d29aab2534e16123c7af21cc636ba3c5f34f7d2885dc479b3f2b049d70f.
		//
		// Solidity: event PlacesDAOUpdated(address placesDao)
		func (_Places *PlacesFilterer) ParsePlacesDAOUpdated(log types.Log) (*PlacesPlacesDAOUpdated, error) {
			event := new(PlacesPlacesDAOUpdated)
			if err := _Places.contract.UnpackLog(event, "PlacesDAOUpdated", log); err != nil {
				return nil, err
			}
			event.Raw = log
			return event, nil
		}

 	
		// PlacesPlacesProviderUpdatedIterator is returned from FilterPlacesProviderUpdated and is used to iterate over the raw logs and unpacked data for PlacesProviderUpdated events raised by the Places contract.
		type PlacesPlacesProviderUpdatedIterator struct {
			Event *PlacesPlacesProviderUpdated // Event containing the contract specifics and raw log

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
		func (it *PlacesPlacesProviderUpdatedIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(PlacesPlacesProviderUpdated)
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
				it.Event = new(PlacesPlacesProviderUpdated)
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
		func (it *PlacesPlacesProviderUpdatedIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *PlacesPlacesProviderUpdatedIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// PlacesPlacesProviderUpdated represents a PlacesProviderUpdated event raised by the Places contract.
		type PlacesPlacesProviderUpdated struct { 
			Provider common.Address; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterPlacesProviderUpdated is a free log retrieval operation binding the contract event 0x6e1e740c0a936d02d70a8289b193c598908d94ba3b2dc60efbeece183ec3140c.
		//
		// Solidity: event PlacesProviderUpdated(address provider)
 		func (_Places *PlacesFilterer) FilterPlacesProviderUpdated(opts *bind.FilterOpts) (*PlacesPlacesProviderUpdatedIterator, error) {
			
			

			logs, sub, err := _Places.contract.FilterLogs(opts, "PlacesProviderUpdated")
			if err != nil {
				return nil, err
			}
			return &PlacesPlacesProviderUpdatedIterator{contract: _Places.contract, event: "PlacesProviderUpdated", logs: logs, sub: sub}, nil
 		}

		// WatchPlacesProviderUpdated is a free log subscription operation binding the contract event 0x6e1e740c0a936d02d70a8289b193c598908d94ba3b2dc60efbeece183ec3140c.
		//
		// Solidity: event PlacesProviderUpdated(address provider)
		func (_Places *PlacesFilterer) WatchPlacesProviderUpdated(opts *bind.WatchOpts, sink chan<- *PlacesPlacesProviderUpdated) (event.Subscription, error) {
			
			

			logs, sub, err := _Places.contract.WatchLogs(opts, "PlacesProviderUpdated")
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(PlacesPlacesProviderUpdated)
						if err := _Places.contract.UnpackLog(event, "PlacesProviderUpdated", log); err != nil {
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

		// ParsePlacesProviderUpdated is a log parse operation binding the contract event 0x6e1e740c0a936d02d70a8289b193c598908d94ba3b2dc60efbeece183ec3140c.
		//
		// Solidity: event PlacesProviderUpdated(address provider)
		func (_Places *PlacesFilterer) ParsePlacesProviderUpdated(log types.Log) (*PlacesPlacesProviderUpdated, error) {
			event := new(PlacesPlacesProviderUpdated)
			if err := _Places.contract.UnpackLog(event, "PlacesProviderUpdated", log); err != nil {
				return nil, err
			}
			event.Raw = log
			return event, nil
		}

 	
		// PlacesTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Places contract.
		type PlacesTransferIterator struct {
			Event *PlacesTransfer // Event containing the contract specifics and raw log

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
		func (it *PlacesTransferIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(PlacesTransfer)
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
				it.Event = new(PlacesTransfer)
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
		func (it *PlacesTransferIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *PlacesTransferIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// PlacesTransfer represents a Transfer event raised by the Places contract.
		type PlacesTransfer struct { 
			From common.Address; 
			To common.Address; 
			TokenId *big.Int; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
		//
		// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
 		func (_Places *PlacesFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*PlacesTransferIterator, error) {
			
			var fromRule []interface{}
			for _, fromItem := range from {
				fromRule = append(fromRule, fromItem)
			}
			var toRule []interface{}
			for _, toItem := range to {
				toRule = append(toRule, toItem)
			}
			var tokenIdRule []interface{}
			for _, tokenIdItem := range tokenId {
				tokenIdRule = append(tokenIdRule, tokenIdItem)
			}

			logs, sub, err := _Places.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
			if err != nil {
				return nil, err
			}
			return &PlacesTransferIterator{contract: _Places.contract, event: "Transfer", logs: logs, sub: sub}, nil
 		}

		// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
		//
		// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
		func (_Places *PlacesFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PlacesTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {
			
			var fromRule []interface{}
			for _, fromItem := range from {
				fromRule = append(fromRule, fromItem)
			}
			var toRule []interface{}
			for _, toItem := range to {
				toRule = append(toRule, toItem)
			}
			var tokenIdRule []interface{}
			for _, tokenIdItem := range tokenId {
				tokenIdRule = append(tokenIdRule, tokenIdItem)
			}

			logs, sub, err := _Places.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(PlacesTransfer)
						if err := _Places.contract.UnpackLog(event, "Transfer", log); err != nil {
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
		// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
		func (_Places *PlacesFilterer) ParseTransfer(log types.Log) (*PlacesTransfer, error) {
			event := new(PlacesTransfer)
			if err := _Places.contract.UnpackLog(event, "Transfer", log); err != nil {
				return nil, err
			}
			event.Raw = log
			return event, nil
		}

 	
		// PlacesUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Places contract.
		type PlacesUnpausedIterator struct {
			Event *PlacesUnpaused // Event containing the contract specifics and raw log

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
		func (it *PlacesUnpausedIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(PlacesUnpaused)
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
				it.Event = new(PlacesUnpaused)
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
		func (it *PlacesUnpausedIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *PlacesUnpausedIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// PlacesUnpaused represents a Unpaused event raised by the Places contract.
		type PlacesUnpaused struct { 
			Account common.Address; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
		//
		// Solidity: event Unpaused(address account)
 		func (_Places *PlacesFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PlacesUnpausedIterator, error) {
			
			

			logs, sub, err := _Places.contract.FilterLogs(opts, "Unpaused")
			if err != nil {
				return nil, err
			}
			return &PlacesUnpausedIterator{contract: _Places.contract, event: "Unpaused", logs: logs, sub: sub}, nil
 		}

		// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
		//
		// Solidity: event Unpaused(address account)
		func (_Places *PlacesFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PlacesUnpaused) (event.Subscription, error) {
			
			

			logs, sub, err := _Places.contract.WatchLogs(opts, "Unpaused")
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(PlacesUnpaused)
						if err := _Places.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

		// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
		//
		// Solidity: event Unpaused(address account)
		func (_Places *PlacesFilterer) ParseUnpaused(log types.Log) (*PlacesUnpaused, error) {
			event := new(PlacesUnpaused)
			if err := _Places.contract.UnpackLog(event, "Unpaused", log); err != nil {
				return nil, err
			}
			event.Raw = log
			return event, nil
		}

 	


