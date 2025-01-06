package balance

import (
	"github.com/spf13/cobra"

	"github.com/kitanoyoru/wallet/cmd/balance/addressbalance"
	"github.com/kitanoyoru/wallet/cmd/balance/contractbalance"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use: "balance",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	cmd.AddCommand(contractbalance.Command())
	cmd.AddCommand(addressbalance.Command())

	return cmd
}
