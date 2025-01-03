package main

import (
	"github.com/kitanoyoru/wallet/cmd/deploy"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const AppName = "wallet"

var rootCmd = &cobra.Command{
	Short: AppName,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func main() {
	rootCmd.AddCommand(deploy.Command())

	if err := rootCmd.Execute(); err != nil {
		log.Fatal().Err(err).Send()
	}
}
