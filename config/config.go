package config

import (
	"log"
	"time"

	"github.com/caarlos0/env/v8"
)

var (
	Blockchain BlockchainConfig
	Contract   ContractConfig
)

type BlockchainConfig struct {
	Address    string        `env:"BLOCKCHAIN_ADDRESS,required"`
	WS         string        `env:"BLOCKCHAIN_WS,required"`
	PrivateKey string        `env:"BLOCKCHAIN_PK,required"`
	TimeoutIn    time.Duration `env:"BLOCKCHAIN_TIMEOUT" envDefault:"1s"`
}

type ContractConfig struct {
	Address   string `env:"CONTRACT_ADDRESS,required"`
	GasLimit  int64  `env:"CONTRACT_GAS_LIMIT" envDefault:"3000000"`
	GasPrice  int64  `env:"CONTRACT_GAS_PRICE" envDefault:"1000000"`
	WeiFounds int64  `env:"CONTRACT_DEFAULT_WEI_FUNDS" envDefault:"0"`
}

func init() {
	err := env.Parse(&Blockchain)
	if err != nil {
		log.Fatalf("Failed to parse blockchain config: %v", err)
	}

	err = env.Parse(&Contract)
	if err != nil {
		log.Fatalf("Failed to parse contract config: %v", err)
	}
}

