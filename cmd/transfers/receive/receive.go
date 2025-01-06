package receive

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"

	"github.com/kitanoyoru/wallet/config"
	ethcontext "github.com/kitanoyoru/wallet/pkg/blockchain/context"
	"github.com/kitanoyoru/wallet/transfers"
)

func Command() *cobra.Command {
	var amount int64

	cmd := &cobra.Command{
		Use: "receive",
		RunE: func(cmd *cobra.Command, args []string) error {
			dialCtx, cancel := context.WithTimeout(context.Background(), config.Blockchain.TimeoutIn)
			defer cancel()

			client, err := ethclient.DialContext(dialCtx, config.Blockchain.WS)
			if err != nil {
				return err
			}

			ctx := ethcontext.WrapToContext(context.Background(), client)

			return transfers.Receive(ctx, amount)
		},
	}

	cmd.Flags().Int64VarP(&amount, "amount", "a", 0, "ETH amount")
	_ = cmd.MarkFlagRequired("amount")

	return cmd
}
