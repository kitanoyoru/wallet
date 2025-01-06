package transfers

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"

	"github.com/kitanoyoru/wallet/config"
	"github.com/kitanoyoru/wallet/pkg/blockchain/common"
	ethcontext "github.com/kitanoyoru/wallet/pkg/blockchain/context"
)

func Receive(ctx context.Context, amount int64) error {
	client := ethcontext.FromContext(ctx)

	contract, err := common.GetContract(ctx, client, config.Contract.Address)
	if err != nil {
		return err
	}

	signer, err := common.GetSigner(ctx, client)
	if err != nil {
		return err
	}

	signer.Value = common.EtherToWei(big.NewInt(amount))
	tx, err := contract.Receive(signer)
	if err != nil {
		return err
	}

	receipt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		return err
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		return errors.New("receipt status is not successful")
	}

	return nil
}
