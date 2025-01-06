package send

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"

	"github.com/kitanoyoru/wallet/config"
	ethcontext "github.com/kitanoyoru/wallet/pkg/blockchain/context"
	"github.com/kitanoyoru/wallet/transfers"
)

func Command() *cobra.Command {
	var (
		targetAddress string
		amount        int64
	)

	cmd := &cobra.Command{
		Use: "send",
		RunE: func(cmd *cobra.Command, args []string) error {
			dialCtx, cancel := context.WithTimeout(context.Background(), config.Blockchain.TimeoutIn)
			defer cancel()

			client, err := ethclient.DialContext(dialCtx, config.Blockchain.WS)
			if err != nil {
				return err
			}

			ctx := ethcontext.WrapToContext(context.Background(), client)

			return transfers.Send(ctx, targetAddress, amount)
		},
	}

	cmd.Flags().StringVarP(&targetAddress, "target.address", "t", "", "Target address")
	cmd.Flags().Int64VarP(&amount, "amount", "a", 0, "ETH amount")

	_ = cmd.MarkFlagRequired("target.address")
	_ = cmd.MarkFlagRequired("amount")

	return cmd
}
