package allowance

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

func GetAllowance(ctx context.Context, beneficiaryAddress string) (int64, error) {
	client := ethcontext.FromContext(ctx)

	contract, err := common.GetContract(ctx, client, config.Contract.Address)
	if err != nil {
		return 0, err
	}

	opts := &bind.CallOpts{
		Pending: false,
		Context: ctx,
	}
	address := gethcommon.HexToAddress(beneficiaryAddress)

	amount, err := contract.Allowances(opts, address)
	if err != nil {
		return 0, err
	}

	return amount.Int64(), nil
}

func SetAllowance(ctx context.Context, beneficiaryAddress string, amount int64) error {
	client := ethcontext.FromContext(ctx)

	contract, err := common.GetContract(ctx, client, config.Contract.Address)
	if err != nil {
		return err
	}

	signer, err := common.GetSigner(ctx, client)
	if err != nil {
		return err
	}

	address := gethcommon.HexToAddress(beneficiaryAddress)

	tx, err := contract.SetAllowance(signer, address, common.EtherToWei(big.NewInt(amount)))
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

func IncreaseAllowance(ctx context.Context, beneficiaryAddress string, amount int64) error {
	client := ethcontext.FromContext(ctx)

	contract, err := common.GetContract(ctx, client, config.Contract.Address)
	if err != nil {
		return err
	}

	signer, err := common.GetSigner(ctx, client)
	if err != nil {
		return err
	}

	address := gethcommon.HexToAddress(beneficiaryAddress)

	tx, err := contract.IncreaseAllowance(signer, address, common.EtherToWei(big.NewInt(amount)))
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

func ReduceAllowance(ctx context.Context, beneficiaryAddress string, amount int64) error {
	client := ethcontext.FromContext(ctx)

	contract, err := common.GetContract(ctx, client, config.Contract.Address)
	if err != nil {
		return err
	}

	signer, err := common.GetSigner(ctx, client)
	if err != nil {
		return err
	}

	address := gethcommon.HexToAddress(beneficiaryAddress)

	tx, err := contract.ReduceAllowance(signer, address, common.EtherToWei(big.NewInt(amount)))
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
