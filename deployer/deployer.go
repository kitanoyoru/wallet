package deployer

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	contracts "github.com/kitanoyoru/wallet/contracts/gen"
	"github.com/kitanoyoru/wallet/pkg/blockchain"
)

func New() *Deployer {
	return &Deployer{}
}

type Deployer struct {
	address  common.Address
	tx       *types.Transaction
	contract *contracts.Contracts
}

func (d *Deployer) Deploy(ctx context.Context, client *ethclient.Client) error {
	signer, err := blockchain.GetSigner(ctx, client)
	if err != nil {
		return err
	}

	address, tx, contract, err := contracts.DeployContracts(signer, client)
	if err != nil {
		return err
	}

	_, err = bind.WaitDeployed(ctx, client, tx)
	if err != nil {
		return err
	}

	d.address = address
	d.tx = tx
	d.contract = contract

	return nil
}

func (d *Deployer) ContractAddress() string {
	return d.address.Hex()
}
