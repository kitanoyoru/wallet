package transfers

import (
	"github.com/kitanoyoru/wallet/cmd/transfers/receive"
	"github.com/kitanoyoru/wallet/cmd/transfers/send"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use: "transfers",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	cmd.AddCommand(send.Command())
	cmd.AddCommand(receive.Command())

	return cmd
}
