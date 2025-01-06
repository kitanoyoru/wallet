package common

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"
	"regexp"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"

	"github.com/kitanoyoru/wallet/config"
	contracts "github.com/kitanoyoru/wallet/contracts/gen"
)

var (
	ErrInvalidKey             = errors.New("invalid key")
	ErrInvalidAddress         = errors.New("invalid address")
	ErrInvalidContractAddress = errors.New("invalid contract address")
)

// validateContractAddress validate the contract address checking if the contract is deployed
func ValidateContractAddress(ctx context.Context, client *ethclient.Client, address string) error {
	if err := ValidateAddress(address); err != nil {
		return err
	}
	contractAddress := common.HexToAddress(address)
	bytecode, err := client.CodeAt(ctx, contractAddress, nil)
	if err != nil {
		return err
	}
	if len(bytecode) > 0 {
		return nil
	}
	return ErrInvalidContractAddress
}

// GetSigner get the signer for sign transactions
func GetSigner(ctx context.Context, client *ethclient.Client) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(config.Blockchain.PrivateKey)
	if err != nil {
		return nil, err
	}
	publicKey, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		return nil, ErrInvalidKey
	}

	address := crypto.PubkeyToAddress(*publicKey)
	nonce, err := client.PendingNonceAt(ctx, address)
	if err != nil {
		return nil, err
	}
	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, err
	}
	signer, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}

	signer.Nonce = big.NewInt(int64(nonce))
	signer.Value = big.NewInt(config.Contract.WeiFounds)
	signer.GasLimit = uint64(config.Contract.GasLimit)
	signer.GasPrice = big.NewInt(config.Contract.GasPrice)

	return signer, nil
}

func GetContract(ctx context.Context, client *ethclient.Client, contractAddress string) (*contracts.Contracts, error) {
	err := ValidateContractAddress(ctx, client, contractAddress)
	if err != nil {
		return nil, err
	}

	contract, err := contracts.NewContracts(common.HexToAddress(contractAddress), client)
	if err != nil {
		return nil, err
	}

	return contract, nil
}

func GetOwner(ctx context.Context, client *ethclient.Client, contractAddress string) (*common.Address, error) {
	contract, err := GetContract(ctx, client, contractAddress)
	if err != nil {
		return nil, err
	}

	ownerAddress, err := contract.Owner(&bind.CallOpts{Context: ctx, Pending: false})
	if err != nil {
		return nil, err
	}

	return &ownerAddress, nil
}

// validateAddress validate address format
func ValidateAddress(address string) error {
	regex := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	if ok := regex.MatchString(address); !ok {
		return ErrInvalidAddress
	}
	return nil
}

// etherToWei convert Ether to Wei
func EtherToWei(eth *big.Int) *big.Int {
	return new(big.Int).Mul(eth, big.NewInt(params.Ether))
}

// weiToEther convert Wei to Ether
func WeiToEther(wei *big.Int) *big.Int {
	return new(big.Int).Div(wei, big.NewInt(params.Ether))
}
