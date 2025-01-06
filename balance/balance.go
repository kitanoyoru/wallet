package balance

import (
	"context"

	gethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/kitanoyoru/wallet/config"
	"github.com/kitanoyoru/wallet/pkg/blockchain/common"
	ethcontext "github.com/kitanoyoru/wallet/pkg/blockchain/context"
)

func GetContractBalance(ctx context.Context) (int64, error) {
	client := ethcontext.FromContext(ctx)

	value, err := client.BalanceAt(ctx, gethcommon.HexToAddress(config.Contract.Address), nil)
	if err != nil {
		return 0, err
	}

	return common.WeiToEther(value).Int64(), nil
}

func GetAddressBalance(ctx context.Context, address string) (int64, error) {
	client := ethcontext.FromContext(ctx)

	value, err := client.BalanceAt(ctx, gethcommon.HexToAddress(address), nil)
	if err != nil {
		return 0, err
	}

	return common.WeiToEther(value).Int64(), nil
}
