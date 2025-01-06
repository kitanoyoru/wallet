package get

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/kitanoyoru/wallet/allowance"
	"github.com/kitanoyoru/wallet/config"
	ethcontext "github.com/kitanoyoru/wallet/pkg/blockchain/context"
)

func Command() *cobra.Command {
	var beneficiaryAddress string

	cmd := &cobra.Command{
		Use: "get",
		RunE: func(cmd *cobra.Command, args []string) error {
			dialCtx, cancel := context.WithTimeout(context.Background(), config.Blockchain.TimeoutIn)
			defer cancel()

			client, err := ethclient.DialContext(dialCtx, config.Blockchain.WS)
			if err != nil {
				return err
			}

			ctx := ethcontext.WrapToContext(context.Background(), client)

			allowance, err := allowance.GetAllowance(ctx, beneficiaryAddress)
			if err != nil {
				return err
			}

			log.Info().Int64("allowance", allowance).Send()

			return nil
		},
	}

	cmd.Flags().StringVarP(&beneficiaryAddress, "beneficiary.address", "b", "", "Beneficiary address")
	_ = cmd.MarkFlagRequired("beneficiary.address")

	return cmd
}
