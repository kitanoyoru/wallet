package deployer

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	contracts "github.com/kitanoyoru/wallet/contracts/gen"
	"github.com/kitanoyoru/wallet/pkg/blockchain/common"
	ethcontext "github.com/kitanoyoru/wallet/pkg/blockchain/context"
)

func New() *Deployer {
	return &Deployer{}
}

type Deployer struct {
	address  gethcommon.Address
	contract *contracts.Contracts
	tx       *types.Transaction
}

func (d *Deployer) Deploy(ctx context.Context) error {
	client := ethcontext.FromContext(ctx)

	signer, err := common.GetSigner(ctx, client)
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
	d.contract = contract
	d.tx = tx

	return nil
}

func (d *Deployer) ContractAddress() string {
	return d.address.Hex()
}
