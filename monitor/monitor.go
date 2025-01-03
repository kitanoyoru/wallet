package monitor

import (
	"context"

	gethcommon "github.com/ethereum/go-ethereum/common"
	"golang.org/x/sync/errgroup"

	contracts "github.com/kitanoyoru/wallet/contracts/gen"
	"github.com/kitanoyoru/wallet/pkg/blockchain/common"
	ethcontext "github.com/kitanoyoru/wallet/pkg/blockchain/context"
)

type Watcher func(ctx context.Context, contract *contracts.Contracts) error

func New(address string, watchers ...Watcher) *Monitor {
	return &Monitor{
		address,
		watchers,
	}
}

type Monitor struct {
	address  string
	watchers []Watcher
}

func (m *Monitor) Start(ctx context.Context) error {
	client := ethcontext.FromContext(ctx)

	err := common.ValidateContractAddress(ctx, client, m.address)
	if err != nil {
		return err
	}

	contract, err := contracts.NewContracts(gethcommon.HexToAddress(m.address), client)
	if err != nil {
		return err
	}

	var eg errgroup.Group
	for _, watcher := range m.watchers {
		eg.Go(func() error {
			return watcher(ctx, contract)
		})
	}

	if err := eg.Wait(); err != nil {
		return err
	}

	return nil
}
