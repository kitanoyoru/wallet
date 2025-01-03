package command

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
)

func NewDeployCommand(ctx context.Context) *cobra.Command {
	deployCommand := &cobra.Command{
		Use:   "deploy",
		Short: "Deploy contract to blockchain",
		RunE: func(cmd *cobra.Command, args []string) error {
			return deploy(ctx)
		},
	}

	return deployCommand
}

func deploy(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, config.App.Blockchain.TimeoutIn)
	defer cancel()

	client, err := ethclient.DialContext(ctx, config.App.Blockchain.Address)
	if err != nil {
		return err
	}

	deployer := blockchain.NewDeployer()
	err = deployer.Deploy(ctx, client)
	if err != nil {
		return err
	}

	log.Printf("contract deployed at address %s\n", deployer.ContractAddress())
	return nil
}
