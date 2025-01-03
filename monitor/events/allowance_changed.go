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

const AllowanceChangedEvent = "AllowanceChanged"

type AllowanceChanged struct {
	Sender      string    `json:"sender"`
	Beneficiary string    `json:"beneficiary"`
	PrevAmount  *big.Int  `json:"prev_amount"`
	NewAmount   *big.Int  `json:"new_amount"`
	Timestamp   time.Time `json:"timestamp"`
}

func WatchAllowanceChanged(ctx context.Context, contract *contracts.Contracts) error {
	events := make(chan *contracts.ContractsAllowanceChanged)
	opts := &bind.WatchOpts{
		Context: ctx,
	}

	subscription, err := contract.WatchAllowanceChanged(opts, events, nil, nil)
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
			e := AllowanceChanged{
				Sender:      event.Sender.Hex(),
				Beneficiary: event.Beneficiary.Hex(),
				PrevAmount:  common.WeiToEther(event.PrevAmount),
				NewAmount:   common.WeiToEther(event.NewAmount),
				Timestamp:   time.Now(),
			}

			raw, err := json.MarshalIndent(e, "", "  ")
			if err != nil {
				log.Error().Str("watch", AllowanceChangedEvent).Err(err).Msg("failed to marshal event")
				continue
			}

			log.Info().Str("watch", AllowanceChangedEvent).Msg(string(raw))
		}
	}
}
