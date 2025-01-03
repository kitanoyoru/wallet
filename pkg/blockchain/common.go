package blockchain

import (
	"context"
	"errors"
	"math/big"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
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
