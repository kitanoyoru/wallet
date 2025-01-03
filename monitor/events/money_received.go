package events

import (
	"context"
	"encoding/json"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/rs/zerolog/log"

	contracts "github.com/kitanoyoru/wallet/contracts/gen"
	"github.com/kitanoyoru/wallet/pkg/blockchain/common"
)

const MoneyReceivedEvent = "MoneyReceived"

type MoneyReceived struct {
	Sender      string    `json:"sender"`
	BlockNumber uint64    `json:"block_number"`
	Amount      *big.Int  `json:"amount"`
	Timestamp   time.Time `json:"timestamp"`
}

func WatchMoneyReceived(ctx context.Context, contract *contracts.Contracts) error {
	events := make(chan *contracts.ContractsMoneyReceived)
	opts := &bind.WatchOpts{
		Context: ctx,
	}

	subscription, err := contract.WatchMoneyReceived(opts, events, nil)
	if err != nil {
		return err
	}
	defer subscription.Unsubscribe()

	for {
		select {
		case <-ctx.Done():
			return nil
		case err := <-subscription.Err():
			return err
		case event := <-events:
			e := MoneyReceived{
				Sender:      event.Sender.Hex(),
				BlockNumber: event.Raw.BlockNumber,
				Amount:      common.WeiToEther(event.Amount),
				Timestamp:   time.Now(),
			}

			raw, err := json.MarshalIndent(e, "", "  ")
			if err != nil {
				log.Error().Str("watch", MoneyReceivedEvent).Err(err).Msg("failed to marshal event")
				continue
			}

			log.Info().Str("watch", MoneyReceivedEvent).Msg(string(raw))
		}
	}
}
