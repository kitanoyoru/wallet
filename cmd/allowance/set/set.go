package set

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"

	"github.com/kitanoyoru/wallet/allowance"
	"github.com/kitanoyoru/wallet/config"
	ethcontext "github.com/kitanoyoru/wallet/pkg/blockchain/context"
)

func Command() *cobra.Command {
	var (
		beneficiaryAddress string
		amount             int64
	)

	cmd := &cobra.Command{
		Use: "set",
		RunE: func(cmd *cobra.Command, args []string) error {
			dialCtx, cancel := context.WithTimeout(context.Background(), config.Blockchain.TimeoutIn)
			defer cancel()

			client, err := ethclient.DialContext(dialCtx, config.Blockchain.WS)
			if err != nil {
				return err
			}

			ctx := ethcontext.WrapToContext(context.Background(), client)

			return allowance.SetAllowance(ctx, beneficiaryAddress, amount)
		},
	}

	cmd.Flags().StringVarP(&beneficiaryAddress, "beneficiary.address", "b", "", "Beneficiary address")
	_ = cmd.MarkFlagRequired("beneficiary.address")

	cmd.Flags().Int64VarP(&amount, "amount", "a", 0, "Allowance amount")
	_ = cmd.MarkFlagRequired("amount")

	return cmd
}
