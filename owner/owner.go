package owner

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/kitanoyoru/wallet/pkg/blockchain/common"
	ethcontext "github.com/kitanoyoru/wallet/pkg/blockchain/context"
)

func GetOwner(ctx context.Context, contractAddress string) (string, error) {
	client := ethcontext.FromContext(ctx)

	ownerAddress, err := common.GetOwner(ctx, client, contractAddress)
	if err != nil {
		return "", nil
	}

	return ownerAddress.Hex(), nil
}

func TransferOwner(ctx context.Context, contractAddress, targetAddress string) error {
	client := ethcontext.FromContext(ctx)

	contract, err := common.GetContract(ctx, client, contractAddress)
	if err != nil {
		return err
	}

	signer, err := common.GetSigner(ctx, client)
	if err != nil {
		return err
	}

	tx, err := contract.TransferOwnership(signer, gethcommon.HexToAddress(targetAddress))
	if err != nil {
		return err
	}

	receipt, err := bind.WaitMined(ctx, client, tx)
	if receipt.Status != types.ReceiptStatusSuccessful || err != nil {
		return err
	}

	// process transaction

	return nil
}
