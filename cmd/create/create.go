package create 

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"

	"github.com/kitanoyoru/wallet/config"
	"github.com/kitanoyoru/wallet/deployer"
	ethcontext "github.com/kitanoyoru/wallet/pkg/blockchain/context"
)

func Command() *cobra.Command {
	return &cobra.Command{
		Use: "create",
		RunE: func(cmd *cobra.Command, args []string) error {
			dialCtx, cancel := context.WithTimeout(context.Background(), config.Blockchain.TimeoutIn)
			defer cancel()

			client, err := ethclient.DialContext(dialCtx, config.Blockchain.WS)
			if err != nil {
				return err
			}

			ctx := ethcontext.WrapToContext(context.Background(), client)

			err = deployer.New().Deploy(ctx)
			if err != nil {
				return err
			}

			return nil
		},
	}
}
