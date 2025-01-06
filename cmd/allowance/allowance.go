package allowance

import (
	"github.com/kitanoyoru/wallet/cmd/allowance/get"
	"github.com/kitanoyoru/wallet/cmd/allowance/increase"
	"github.com/kitanoyoru/wallet/cmd/allowance/reduce"
	"github.com/kitanoyoru/wallet/cmd/allowance/set"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use: "allowance",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	cmd.AddCommand(get.Command())
	cmd.AddCommand(set.Command())
	cmd.AddCommand(increase.Command())
	cmd.AddCommand(reduce.Command())

	return cmd
}
