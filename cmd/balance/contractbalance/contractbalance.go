package contractbalance

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/kitanoyoru/wallet/balance"
	"github.com/kitanoyoru/wallet/config"
	ethcontext "github.com/kitanoyoru/wallet/pkg/blockchain/context"
)

func Command() *cobra.Command {
	return &cobra.Command{
		Use: "contractbalance",
		RunE: func(cmd *cobra.Command, args []string) error {
			dialCtx, cancel := context.WithTimeout(context.Background(), config.Blockchain.TimeoutIn)
			defer cancel()

			client, err := ethclient.DialContext(dialCtx, config.Blockchain.WS)
			if err != nil {
				log.Fatal().Err(err).Send()
			}

			ctx := ethcontext.WrapToContext(context.Background(), client)

			amount, err := balance.GetContractBalance(ctx)
			if err != nil {
				log.Fatal().Err(err).Send()
			}

			log.Info().Int64("amount", amount).Send()

			return nil
		},
	}
}
