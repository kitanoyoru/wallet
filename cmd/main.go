package main

import (
	"github.com/kitanoyoru/wallet/cmd/balance"
	"github.com/kitanoyoru/wallet/cmd/deploy"
	"github.com/kitanoyoru/wallet/cmd/monitor"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const AppName = "wallet"

var rootCmd = &cobra.Command{
	Use:   "wallet",
	Short: "Run the wallet CLI",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func main() {
	rootCmd.AddCommand(deploy.Command())
	rootCmd.AddCommand(balance.Command())
	rootCmd.AddCommand(monitor.Command())

	if err := rootCmd.Execute(); err != nil {
		log.Fatal().Err(err).Send()
	}
}
