package monitor

import (
	"context"

	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/sync/errgroup"

	contracts "github.com/kitanoyoru/wallet/contracts/interfaces"
	"github.com/kitanoyoru/wallet/pkg/blockchain"
)

type Watcher func(ctx context.Context, contract *contracts.Contracts) error

func NewMonitor(address string, watchers ...Watcher) *Monitor {
	return &Monitor{
		address,
		watchers,
	}
}

type Monitor struct {
	address  string
	watchers []Watcher
}

func (m *Monitor) Start(ctx context.Context, client *ethclient.Client) error {
	err := blockchain.ValidateContractAddress(ctx, client, m.address)
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
