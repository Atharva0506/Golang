// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_nft\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_nftId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"NFT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC721\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"NFT_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OWNER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allBids\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bid\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"end\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"endTime\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ended\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"highestBid\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"highestBidder\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"start\",\"inputs\":[{\"name\":\"_openingBid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_duration\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"started\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Bid\",\"inputs\":[{\"name\":\"bidder\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"End\",\"inputs\":[{\"name\":\"winner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Start\",\"inputs\":[{\"name\":\"startTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"endTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"bidder\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
	Bin: "0x60e060405234801561000f575f5ffd5b50604051611c66380380611c66833981810160405281019061003191906101fd565b600161004f61004461013a60201b60201c565b61016360201b60201c565b5f01819055505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036100c3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100ba90610295565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250508173ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff16815250508060c0818152505050506102b3565b5f7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005f1b905090565b5f819050919050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61019982610170565b9050919050565b6101a98161018f565b81146101b3575f5ffd5b50565b5f815190506101c4816101a0565b92915050565b5f819050919050565b6101dc816101ca565b81146101e6575f5ffd5b50565b5f815190506101f7816101d3565b92915050565b5f5f604083850312156102135761021261016c565b5b5f610220858286016101b6565b9250506020610231858286016101e9565b9150509250929050565b5f82825260208201905092915050565b7f496e76616c6964204e46542061646472657373000000000000000000000000005f82015250565b5f61027f60138361023b565b915061028a8261024b565b602082019050919050565b5f6020820190508181035f8301526102ac81610273565b9050919050565b60805160a05160c0516119436103235f395f81816107aa0152818161095401528181610bc60152610d6001525f81816107ce0152818161091601528181610b670152610d0201525f81816102b1015281816103d401528181610c3301528181610d3f0152610e8401526119435ff3fe6080604052600436106100c1575f3560e01c8063707ddfae1161007e57806391f901571161005857806391f9015714610209578063bd20160714610233578063d57bde791461026f578063efbe1c1c14610299576100c1565b8063707ddfae1461018d5780637c0b8de2146101b75780638fb4b573146101e1576100c1565b8063117803e3146100c557806312fa6feb146100ef5780631998aeef146101195780631f2698ab146101235780633197cbb61461014d5780633ccfd60b14610177575b5f5ffd5b3480156100d0575f5ffd5b506100d96102af565b6040516100e69190610fe0565b60405180910390f35b3480156100fa575f5ffd5b506101036102d3565b6040516101109190611013565b60405180910390f35b6101216102e5565b005b34801561012e575f5ffd5b506101376105c2565b6040516101449190611013565b60405180910390f35b348015610158575f5ffd5b506101616105d3565b60405161016e9190611044565b60405180910390f35b348015610182575f5ffd5b5061018b6105d9565b005b348015610198575f5ffd5b506101a16107a8565b6040516101ae9190611044565b60405180910390f35b3480156101c2575f5ffd5b506101cb6107cc565b6040516101d891906110b8565b60405180910390f35b3480156101ec575f5ffd5b50610207600480360381019061020291906110ff565b6107f0565b005b348015610214575f5ffd5b5061021d6109fd565b60405161022a919061115d565b60405180910390f35b34801561023e575f5ffd5b50610259600480360381019061025491906111a0565b610a22565b6040516102669190611044565b60405180910390f35b34801561027a575f5ffd5b50610283610a37565b6040516102909190611044565b60405180910390f35b3480156102a4575f5ffd5b506102ad610a3d565b005b7f000000000000000000000000000000000000000000000000000000000000000081565b5f60019054906101000a900460ff1681565b5f5f9054906101000a900460ff16801561030b57505f60019054906101000a900460ff16155b61034a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161034190611225565b60405180910390fd5b600154421061038e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103859061128d565b60405180910390fd5b60025434116103d2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103c99061131b565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1603610460576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610457906113a9565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff1660035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461052b5760025460045f60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461052391906113f4565b925050819055505b346002819055503360035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055503373ffffffffffffffffffffffffffffffffffffffff167fe684a55f31b79eca403df938249029212a5925ec6be8012e099b45bc1019e5d2346040516105b89190611044565b60405180910390a2565b5f5f9054906101000a900460ff1681565b60015481565b6105e1610e46565b5f60045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f8111610664576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161065b90611471565b60405180910390fd5b5f60045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055505f3373ffffffffffffffffffffffffffffffffffffffff16826040516106cb906114bc565b5f6040518083038185875af1925050503d805f8114610705576040519150601f19603f3d011682016040523d82523d5f602084013e61070a565b606091505b505090508061074e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107459061151a565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff167f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364836040516107949190611044565b60405180910390a250506107a6610e68565b565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b6107f8610e82565b5f5f9054906101000a900460ff1615801561081f57505f60019054906101000a900460ff16155b61085e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161085590611582565b60405180910390fd5b5f82116108a0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161089790611610565b60405180910390fd5b5f81116108e2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108d990611678565b60405180910390fd5b8160028190555080426108f591906113f4565b60018190555060015f5f6101000a81548160ff0219169083151502179055507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166323b872dd33307f00000000000000000000000000000000000000000000000000000000000000006040518463ffffffff1660e01b815260040161099193929190611696565b5f604051808303815f87803b1580156109a8575f5ffd5b505af11580156109ba573d5f5f3e3d5ffd5b505050507f5a7d5f50ab70a39d193bf53f0fb995377776df93bb6fdcf1cd1868b7e0dd44ee426001546040516109f19291906116cb565b60405180910390a15050565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6004602052805f5260405f205f915090505481565b60025481565b610a45610e82565b610a4d610e46565b5f5f9054906101000a900460ff168015610a7357505f60019054906101000a900460ff16155b610ab2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610aa990611762565b60405180910390fd5b600154421015610af7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610aee906117ca565b60405180910390fd5b60015f60016101000a81548160ff0219169083151502179055505f73ffffffffffffffffffffffffffffffffffffffff1660035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610d00577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166323b872dd3060035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff167f00000000000000000000000000000000000000000000000000000000000000006040518463ffffffff1660e01b8152600401610c0393929190611696565b5f604051808303815f87803b158015610c1a575f5ffd5b505af1158015610c2c573d5f5f3e3d5ffd5b505050505f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16600254604051610c77906114bc565b5f6040518083038185875af1925050503d805f8114610cb1576040519150601f19603f3d011682016040523d82523d5f602084013e610cb6565b606091505b5050905080610cfa576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cf190611832565b60405180910390fd5b50610dcb565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166323b872dd307f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000006040518463ffffffff1660e01b8152600401610d9d93929190611870565b5f604051808303815f87803b158015610db4575f5ffd5b505af1158015610dc6573d5f5f3e3d5ffd5b505050505b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f7d7570b046e5ead3a4f3fbc9fa2445113625a1e79823776146316bafed6b8cc5600254604051610e349190611044565b60405180910390a2610e44610e68565b565b610e4e610f12565b6002610e60610e5b610f53565b610f7c565b5f0181905550565b6001610e7a610e75610f53565b610f7c565b5f0181905550565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610f10576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f07906118ef565b60405180910390fd5b565b610f1a610f85565b15610f51576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005f1b905090565b5f819050919050565b5f6002610f98610f93610f53565b610f7c565b5f015414905090565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610fca82610fa1565b9050919050565b610fda81610fc0565b82525050565b5f602082019050610ff35f830184610fd1565b92915050565b5f8115159050919050565b61100d81610ff9565b82525050565b5f6020820190506110265f830184611004565b92915050565b5f819050919050565b61103e8161102c565b82525050565b5f6020820190506110575f830184611035565b92915050565b5f819050919050565b5f61108061107b61107684610fa1565b61105d565b610fa1565b9050919050565b5f61109182611066565b9050919050565b5f6110a282611087565b9050919050565b6110b281611098565b82525050565b5f6020820190506110cb5f8301846110a9565b92915050565b5f5ffd5b6110de8161102c565b81146110e8575f5ffd5b50565b5f813590506110f9816110d5565b92915050565b5f5f60408385031215611115576111146110d1565b5b5f611122858286016110eb565b9250506020611133858286016110eb565b9150509250929050565b5f61114782610fa1565b9050919050565b6111578161113d565b82525050565b5f6020820190506111705f83018461114e565b92915050565b61117f8161113d565b8114611189575f5ffd5b50565b5f8135905061119a81611176565b92915050565b5f602082840312156111b5576111b46110d1565b5b5f6111c28482850161118c565b91505092915050565b5f82825260208201905092915050565b7f41756374696f6e206e6f742061637469766500000000000000000000000000005f82015250565b5f61120f6012836111cb565b915061121a826111db565b602082019050919050565b5f6020820190508181035f83015261123c81611203565b9050919050565b7f41756374696f6e2068617320656e6465640000000000000000000000000000005f82015250565b5f6112776011836111cb565b915061128282611243565b602082019050919050565b5f6020820190508181035f8301526112a48161126b565b9050919050565b7f426964206d75737420626520686967686572207468616e2063757272656e74205f8201527f6869676865737420626964000000000000000000000000000000000000000000602082015250565b5f611305602b836111cb565b9150611310826112ab565b604082019050919050565b5f6020820190508181035f830152611332816112f9565b9050919050565b7f4f776e65722063616e6e6f7420626964206f6e207468656972206f776e2061755f8201527f6374696f6e000000000000000000000000000000000000000000000000000000602082015250565b5f6113936025836111cb565b915061139e82611339565b604082019050919050565b5f6020820190508181035f8301526113c081611387565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6113fe8261102c565b91506114098361102c565b9250828201905080821115611421576114206113c7565b5b92915050565b7f4e6f7468696e6720746f207769746864726177000000000000000000000000005f82015250565b5f61145b6013836111cb565b915061146682611427565b602082019050919050565b5f6020820190508181035f8301526114888161144f565b9050919050565b5f81905092915050565b50565b5f6114a75f8361148f565b91506114b282611499565b5f82019050919050565b5f6114c68261149c565b9150819050919050565b7f5769746864726177206661696c656400000000000000000000000000000000005f82015250565b5f611504600f836111cb565b915061150f826114d0565b602082019050919050565b5f6020820190508181035f830152611531816114f8565b9050919050565b7f41756374696f6e20616c72656164792073746172746564206f7220656e6465645f82015250565b5f61156c6020836111cb565b915061157782611538565b602082019050919050565b5f6020820190508181035f83015261159981611560565b9050919050565b7f4f70656e696e6720626964206d7573742062652067726561746572207468616e5f8201527f2030000000000000000000000000000000000000000000000000000000000000602082015250565b5f6115fa6022836111cb565b9150611605826115a0565b604082019050919050565b5f6020820190508181035f830152611627816115ee565b9050919050565b7f4475726174696f6e206d7573742062652067726561746572207468616e2030005f82015250565b5f611662601f836111cb565b915061166d8261162e565b602082019050919050565b5f6020820190508181035f83015261168f81611656565b9050919050565b5f6060820190506116a95f83018661114e565b6116b6602083018561114e565b6116c36040830184611035565b949350505050565b5f6040820190506116de5f830185611035565b6116eb6020830184611035565b9392505050565b7f41756374696f6e206e6f742073746172746564206f7220616c726561647920655f8201527f6e64656400000000000000000000000000000000000000000000000000000000602082015250565b5f61174c6024836111cb565b9150611757826116f2565b604082019050919050565b5f6020820190508181035f83015261177981611740565b9050919050565b7f41756374696f6e20686173206e6f7420656e64656420796574000000000000005f82015250565b5f6117b46019836111cb565b91506117bf82611780565b602082019050919050565b5f6020820190508181035f8301526117e1816117a8565b9050919050565b7f5472616e7366657220746f206f776e6572206661696c656400000000000000005f82015250565b5f61181c6018836111cb565b9150611827826117e8565b602082019050919050565b5f6020820190508181035f83015261184981611810565b9050919050565b5f61185a82611087565b9050919050565b61186a81611850565b82525050565b5f6060820190506118835f83018661114e565b6118906020830185611861565b61189d6040830184611035565b949350505050565b7f4e6f7420746865206f776e6572000000000000000000000000000000000000005f82015250565b5f6118d9600d836111cb565b91506118e4826118a5565b602082019050919050565b5f6020820190508181035f830152611906816118cd565b905091905056fea264697066735822122066f6cbc3622778e0b578a6808f2eb06f73e9a070fccb2103f5ccef29f3762c5164736f6c63430008210033",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend, _nft common.Address, _nftId *big.Int) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend, _nft, _nftId)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// NFT is a free data retrieval call binding the contract method 0x7c0b8de2.
//
// Solidity: function NFT() view returns(address)
func (_Contract *ContractCaller) NFT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "NFT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NFT is a free data retrieval call binding the contract method 0x7c0b8de2.
//
// Solidity: function NFT() view returns(address)
func (_Contract *ContractSession) NFT() (common.Address, error) {
	return _Contract.Contract.NFT(&_Contract.CallOpts)
}

// NFT is a free data retrieval call binding the contract method 0x7c0b8de2.
//
// Solidity: function NFT() view returns(address)
func (_Contract *ContractCallerSession) NFT() (common.Address, error) {
	return _Contract.Contract.NFT(&_Contract.CallOpts)
}

// NFTID is a free data retrieval call binding the contract method 0x707ddfae.
//
// Solidity: function NFT_ID() view returns(uint256)
func (_Contract *ContractCaller) NFTID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "NFT_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NFTID is a free data retrieval call binding the contract method 0x707ddfae.
//
// Solidity: function NFT_ID() view returns(uint256)
func (_Contract *ContractSession) NFTID() (*big.Int, error) {
	return _Contract.Contract.NFTID(&_Contract.CallOpts)
}

// NFTID is a free data retrieval call binding the contract method 0x707ddfae.
//
// Solidity: function NFT_ID() view returns(uint256)
func (_Contract *ContractCallerSession) NFTID() (*big.Int, error) {
	return _Contract.Contract.NFTID(&_Contract.CallOpts)
}

// OWNER is a free data retrieval call binding the contract method 0x117803e3.
//
// Solidity: function OWNER() view returns(address)
func (_Contract *ContractCaller) OWNER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "OWNER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OWNER is a free data retrieval call binding the contract method 0x117803e3.
//
// Solidity: function OWNER() view returns(address)
func (_Contract *ContractSession) OWNER() (common.Address, error) {
	return _Contract.Contract.OWNER(&_Contract.CallOpts)
}

// OWNER is a free data retrieval call binding the contract method 0x117803e3.
//
// Solidity: function OWNER() view returns(address)
func (_Contract *ContractCallerSession) OWNER() (common.Address, error) {
	return _Contract.Contract.OWNER(&_Contract.CallOpts)
}

// AllBids is a free data retrieval call binding the contract method 0xbd201607.
//
// Solidity: function allBids(address ) view returns(uint256)
func (_Contract *ContractCaller) AllBids(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "allBids", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllBids is a free data retrieval call binding the contract method 0xbd201607.
//
// Solidity: function allBids(address ) view returns(uint256)
func (_Contract *ContractSession) AllBids(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.AllBids(&_Contract.CallOpts, arg0)
}

// AllBids is a free data retrieval call binding the contract method 0xbd201607.
//
// Solidity: function allBids(address ) view returns(uint256)
func (_Contract *ContractCallerSession) AllBids(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.AllBids(&_Contract.CallOpts, arg0)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_Contract *ContractCaller) EndTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "endTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_Contract *ContractSession) EndTime() (*big.Int, error) {
	return _Contract.Contract.EndTime(&_Contract.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_Contract *ContractCallerSession) EndTime() (*big.Int, error) {
	return _Contract.Contract.EndTime(&_Contract.CallOpts)
}

// Ended is a free data retrieval call binding the contract method 0x12fa6feb.
//
// Solidity: function ended() view returns(bool)
func (_Contract *ContractCaller) Ended(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "ended")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Ended is a free data retrieval call binding the contract method 0x12fa6feb.
//
// Solidity: function ended() view returns(bool)
func (_Contract *ContractSession) Ended() (bool, error) {
	return _Contract.Contract.Ended(&_Contract.CallOpts)
}

// Ended is a free data retrieval call binding the contract method 0x12fa6feb.
//
// Solidity: function ended() view returns(bool)
func (_Contract *ContractCallerSession) Ended() (bool, error) {
	return _Contract.Contract.Ended(&_Contract.CallOpts)
}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_Contract *ContractCaller) HighestBid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "highestBid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_Contract *ContractSession) HighestBid() (*big.Int, error) {
	return _Contract.Contract.HighestBid(&_Contract.CallOpts)
}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_Contract *ContractCallerSession) HighestBid() (*big.Int, error) {
	return _Contract.Contract.HighestBid(&_Contract.CallOpts)
}

// HighestBidder is a free data retrieval call binding the contract method 0x91f90157.
//
// Solidity: function highestBidder() view returns(address)
func (_Contract *ContractCaller) HighestBidder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "highestBidder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HighestBidder is a free data retrieval call binding the contract method 0x91f90157.
//
// Solidity: function highestBidder() view returns(address)
func (_Contract *ContractSession) HighestBidder() (common.Address, error) {
	return _Contract.Contract.HighestBidder(&_Contract.CallOpts)
}

// HighestBidder is a free data retrieval call binding the contract method 0x91f90157.
//
// Solidity: function highestBidder() view returns(address)
func (_Contract *ContractCallerSession) HighestBidder() (common.Address, error) {
	return _Contract.Contract.HighestBidder(&_Contract.CallOpts)
}

// Started is a free data retrieval call binding the contract method 0x1f2698ab.
//
// Solidity: function started() view returns(bool)
func (_Contract *ContractCaller) Started(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "started")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Started is a free data retrieval call binding the contract method 0x1f2698ab.
//
// Solidity: function started() view returns(bool)
func (_Contract *ContractSession) Started() (bool, error) {
	return _Contract.Contract.Started(&_Contract.CallOpts)
}

// Started is a free data retrieval call binding the contract method 0x1f2698ab.
//
// Solidity: function started() view returns(bool)
func (_Contract *ContractCallerSession) Started() (bool, error) {
	return _Contract.Contract.Started(&_Contract.CallOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Contract *ContractTransactor) Bid(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "bid")
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Contract *ContractSession) Bid() (*types.Transaction, error) {
	return _Contract.Contract.Bid(&_Contract.TransactOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Contract *ContractTransactorSession) Bid() (*types.Transaction, error) {
	return _Contract.Contract.Bid(&_Contract.TransactOpts)
}

// End is a paid mutator transaction binding the contract method 0xefbe1c1c.
//
// Solidity: function end() returns()
func (_Contract *ContractTransactor) End(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "end")
}

// End is a paid mutator transaction binding the contract method 0xefbe1c1c.
//
// Solidity: function end() returns()
func (_Contract *ContractSession) End() (*types.Transaction, error) {
	return _Contract.Contract.End(&_Contract.TransactOpts)
}

// End is a paid mutator transaction binding the contract method 0xefbe1c1c.
//
// Solidity: function end() returns()
func (_Contract *ContractTransactorSession) End() (*types.Transaction, error) {
	return _Contract.Contract.End(&_Contract.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0x8fb4b573.
//
// Solidity: function start(uint256 _openingBid, uint256 _duration) returns()
func (_Contract *ContractTransactor) Start(opts *bind.TransactOpts, _openingBid *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "start", _openingBid, _duration)
}

// Start is a paid mutator transaction binding the contract method 0x8fb4b573.
//
// Solidity: function start(uint256 _openingBid, uint256 _duration) returns()
func (_Contract *ContractSession) Start(_openingBid *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Start(&_Contract.TransactOpts, _openingBid, _duration)
}

// Start is a paid mutator transaction binding the contract method 0x8fb4b573.
//
// Solidity: function start(uint256 _openingBid, uint256 _duration) returns()
func (_Contract *ContractTransactorSession) Start(_openingBid *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Start(&_Contract.TransactOpts, _openingBid, _duration)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Contract *ContractTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Contract *ContractSession) Withdraw() (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Contract *ContractTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts)
}

// ContractBidIterator is returned from FilterBid and is used to iterate over the raw logs and unpacked data for Bid events raised by the Contract contract.
type ContractBidIterator struct {
	Event *ContractBid // Event containing the contract specifics and raw log

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
func (it *ContractBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractBid)
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
		it.Event = new(ContractBid)
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
func (it *ContractBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractBid represents a Bid event raised by the Contract contract.
type ContractBid struct {
	Bidder common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBid is a free log retrieval operation binding the contract event 0xe684a55f31b79eca403df938249029212a5925ec6be8012e099b45bc1019e5d2.
//
// Solidity: event Bid(address indexed bidder, uint256 amount)
func (_Contract *ContractFilterer) FilterBid(opts *bind.FilterOpts, bidder []common.Address) (*ContractBidIterator, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Bid", bidderRule)
	if err != nil {
		return nil, err
	}
	return &ContractBidIterator{contract: _Contract.contract, event: "Bid", logs: logs, sub: sub}, nil
}

// WatchBid is a free log subscription operation binding the contract event 0xe684a55f31b79eca403df938249029212a5925ec6be8012e099b45bc1019e5d2.
//
// Solidity: event Bid(address indexed bidder, uint256 amount)
func (_Contract *ContractFilterer) WatchBid(opts *bind.WatchOpts, sink chan<- *ContractBid, bidder []common.Address) (event.Subscription, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Bid", bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractBid)
				if err := _Contract.contract.UnpackLog(event, "Bid", log); err != nil {
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

// ParseBid is a log parse operation binding the contract event 0xe684a55f31b79eca403df938249029212a5925ec6be8012e099b45bc1019e5d2.
//
// Solidity: event Bid(address indexed bidder, uint256 amount)
func (_Contract *ContractFilterer) ParseBid(log types.Log) (*ContractBid, error) {
	event := new(ContractBid)
	if err := _Contract.contract.UnpackLog(event, "Bid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractEndIterator is returned from FilterEnd and is used to iterate over the raw logs and unpacked data for End events raised by the Contract contract.
type ContractEndIterator struct {
	Event *ContractEnd // Event containing the contract specifics and raw log

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
func (it *ContractEndIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEnd)
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
		it.Event = new(ContractEnd)
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
func (it *ContractEndIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEndIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEnd represents a End event raised by the Contract contract.
type ContractEnd struct {
	Winner common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEnd is a free log retrieval operation binding the contract event 0x7d7570b046e5ead3a4f3fbc9fa2445113625a1e79823776146316bafed6b8cc5.
//
// Solidity: event End(address indexed winner, uint256 amount)
func (_Contract *ContractFilterer) FilterEnd(opts *bind.FilterOpts, winner []common.Address) (*ContractEndIterator, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "End", winnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractEndIterator{contract: _Contract.contract, event: "End", logs: logs, sub: sub}, nil
}

// WatchEnd is a free log subscription operation binding the contract event 0x7d7570b046e5ead3a4f3fbc9fa2445113625a1e79823776146316bafed6b8cc5.
//
// Solidity: event End(address indexed winner, uint256 amount)
func (_Contract *ContractFilterer) WatchEnd(opts *bind.WatchOpts, sink chan<- *ContractEnd, winner []common.Address) (event.Subscription, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "End", winnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEnd)
				if err := _Contract.contract.UnpackLog(event, "End", log); err != nil {
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

// ParseEnd is a log parse operation binding the contract event 0x7d7570b046e5ead3a4f3fbc9fa2445113625a1e79823776146316bafed6b8cc5.
//
// Solidity: event End(address indexed winner, uint256 amount)
func (_Contract *ContractFilterer) ParseEnd(log types.Log) (*ContractEnd, error) {
	event := new(ContractEnd)
	if err := _Contract.contract.UnpackLog(event, "End", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractStartIterator is returned from FilterStart and is used to iterate over the raw logs and unpacked data for Start events raised by the Contract contract.
type ContractStartIterator struct {
	Event *ContractStart // Event containing the contract specifics and raw log

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
func (it *ContractStartIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractStart)
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
		it.Event = new(ContractStart)
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
func (it *ContractStartIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractStartIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractStart represents a Start event raised by the Contract contract.
type ContractStart struct {
	StartTime *big.Int
	EndTime   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStart is a free log retrieval operation binding the contract event 0x5a7d5f50ab70a39d193bf53f0fb995377776df93bb6fdcf1cd1868b7e0dd44ee.
//
// Solidity: event Start(uint256 startTime, uint256 endTime)
func (_Contract *ContractFilterer) FilterStart(opts *bind.FilterOpts) (*ContractStartIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Start")
	if err != nil {
		return nil, err
	}
	return &ContractStartIterator{contract: _Contract.contract, event: "Start", logs: logs, sub: sub}, nil
}

// WatchStart is a free log subscription operation binding the contract event 0x5a7d5f50ab70a39d193bf53f0fb995377776df93bb6fdcf1cd1868b7e0dd44ee.
//
// Solidity: event Start(uint256 startTime, uint256 endTime)
func (_Contract *ContractFilterer) WatchStart(opts *bind.WatchOpts, sink chan<- *ContractStart) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Start")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractStart)
				if err := _Contract.contract.UnpackLog(event, "Start", log); err != nil {
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

// ParseStart is a log parse operation binding the contract event 0x5a7d5f50ab70a39d193bf53f0fb995377776df93bb6fdcf1cd1868b7e0dd44ee.
//
// Solidity: event Start(uint256 startTime, uint256 endTime)
func (_Contract *ContractFilterer) ParseStart(log types.Log) (*ContractStart, error) {
	event := new(ContractStart)
	if err := _Contract.contract.UnpackLog(event, "Start", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Contract contract.
type ContractWithdrawIterator struct {
	Event *ContractWithdraw // Event containing the contract specifics and raw log

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
func (it *ContractWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractWithdraw)
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
		it.Event = new(ContractWithdraw)
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
func (it *ContractWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractWithdraw represents a Withdraw event raised by the Contract contract.
type ContractWithdraw struct {
	Bidder common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed bidder, uint256 amount)
func (_Contract *ContractFilterer) FilterWithdraw(opts *bind.FilterOpts, bidder []common.Address) (*ContractWithdrawIterator, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Withdraw", bidderRule)
	if err != nil {
		return nil, err
	}
	return &ContractWithdrawIterator{contract: _Contract.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed bidder, uint256 amount)
func (_Contract *ContractFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *ContractWithdraw, bidder []common.Address) (event.Subscription, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Withdraw", bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractWithdraw)
				if err := _Contract.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed bidder, uint256 amount)
func (_Contract *ContractFilterer) ParseWithdraw(log types.Log) (*ContractWithdraw, error) {
	event := new(ContractWithdraw)
	if err := _Contract.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
