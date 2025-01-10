package main

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/kitanoyoru/wallet/cmd/balance"
	"github.com/kitanoyoru/wallet/cmd/create"
	"github.com/kitanoyoru/wallet/cmd/monitor"
	"github.com/kitanoyoru/wallet/cmd/transfers"
)

var rootCmd = &cobra.Command{
	Use:   "wallet",
	Short: "Run the wallet CLI",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func main() {
	rootCmd.AddCommand(create.Command())
	rootCmd.AddCommand(balance.Command())
	rootCmd.AddCommand(monitor.Command())
	rootCmd.AddCommand(transfers.Command())

	if err := rootCmd.Execute(); err != nil {
		log.Fatal().Err(err).Send()
	}
}
