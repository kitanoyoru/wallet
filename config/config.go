package config

import (
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

const (
	filename = "config"
	filetype = "yaml"
	filepath = "./config"
)

var (
	Blockchain BlockchainConfig
	Conctract  ContractConfig
)

// BlockchainConfig struct
type BlockchainConfig struct {
	Address    string `mapstructure:"address"`
	WS         string `mapstructure:"ws"`
	PrivateKey string `mapstructure:"pk"`
	Timeout    string `mapstructure:"timeout"`
	TimeoutIn  time.Duration
}

// ContractConfig struct
type ContractConfig struct {
	Address   string `mapstructure:"address"`
	GasLimit  int64  `mapstructure:"gas_limit"`
	GasPrice  int64  `mapstructure:"gas_price"`
	WeiFounds int64  `mapstructure:"default_wei_founds"`
}

func init() {
	v := viper.New()

	v.SetConfigFile(filename)
	v.SetConfigType(filetype)
	v.AddConfigPath(filepath)

	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	err = v.Unmarshal(&Blockchain)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	err = v.Unmarshal(&Conctract)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	Blockchain.TimeoutIn, err = time.ParseDuration(Blockchain.Timeout)
	if err != nil {
		log.Fatal().Err(err).Send()
	}
}
