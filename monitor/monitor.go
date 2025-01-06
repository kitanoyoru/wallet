package monitor

import (
	"context"

	"golang.org/x/sync/errgroup"

	"github.com/kitanoyoru/wallet/config"
	contracts "github.com/kitanoyoru/wallet/contracts/gen"
	"github.com/kitanoyoru/wallet/pkg/blockchain/common"
	ethcontext "github.com/kitanoyoru/wallet/pkg/blockchain/context"
)

type Watcher func(ctx context.Context, contract *contracts.Contracts) error

func New(watchers ...Watcher) *Monitor {
	return &Monitor{
		watchers,
	}
}

type Monitor struct {
	watchers []Watcher
}

func (m *Monitor) Start(ctx context.Context) error {
	client := ethcontext.FromContext(ctx)

	contract, err := common.GetContract(ctx, client, config.Contract.Address)
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
