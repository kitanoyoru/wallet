package transfers

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"

	"github.com/kitanoyoru/wallet/config"
	"github.com/kitanoyoru/wallet/pkg/blockchain/common"
	ethcontext "github.com/kitanoyoru/wallet/pkg/blockchain/context"
)

func Send(ctx context.Context, to string, amount int64) error {
	client := ethcontext.FromContext(ctx)

	contract, err := common.GetContract(ctx, client, config.Contract.Address)
	if err != nil {
		return err
	}

	signer, err := common.GetSigner(ctx, client)
	if err != nil {
		return err
	}

	toAddress := gethcommon.HexToAddress(to)

	tx, err := contract.SendMoney(signer, toAddress, common.EtherToWei(big.NewInt(amount)))
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
