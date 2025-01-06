package monitor

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/kitanoyoru/wallet/config"
	"github.com/kitanoyoru/wallet/monitor"
	"github.com/kitanoyoru/wallet/monitor/events"
	ethcontext "github.com/kitanoyoru/wallet/pkg/blockchain/context"
)

func Command() *cobra.Command {
	return &cobra.Command{
		Use:   "monitor",
		Short: "Monitor events from EVM",
		RunE: func(cmd *cobra.Command, args []string) error {
			dialCtx, cancel := context.WithTimeout(context.Background(), config.Blockchain.TimeoutIn)
			defer cancel()

			client, err := ethclient.DialContext(dialCtx, config.Blockchain.WS)
			if err != nil {
				return err
			}

			ctx := ethcontext.WrapToContext(context.Background(), client)

			m := monitor.New(
				events.WatchAllowanceChanged,
				events.WatchMoneyReceived,
				events.WatchMoneySent,
				events.WatchOwnershipTransferred,
			)

			ctx, cancel = context.WithCancel(context.Background())
			defer cancel()

			signalChan := make(chan os.Signal, 1)
			signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT)

			go func() {
				_ = <-signalChan
				cancel()
			}()

			return m.Start(ctx)
		},
	}
}
