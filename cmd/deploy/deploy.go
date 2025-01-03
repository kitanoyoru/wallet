package deploy

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"

	"github.com/kitanoyoru/wallet/config"
	"github.com/kitanoyoru/wallet/deployer"
)

// Command to deploy smart contracts into the blockchain
func Command() *cobra.Command {
	return &cobra.Command{
		Use: "deploy",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := context.WithTimeout(context.Background(), config.Blockchain.TimeoutIn)
			defer cancel()

			client, err := ethclient.DialContext(ctx, config.Blockchain.Address)
			if err != nil {
				return err
			}

			err = deployer.New().Deploy(ctx, client)
			if err != nil {
				return err
			}

			return nil
		},
	}
}
