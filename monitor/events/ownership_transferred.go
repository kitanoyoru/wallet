package events

import (
	"context"
	"encoding/json"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/rs/zerolog/log"

	contracts "github.com/kitanoyoru/wallet/contracts/gen"
)

const OwnershipTransferredEvent = "OwnershipTransferred"

type OwnershipTransferred struct {
	PreviousOwner string    `json:"previous_owner"`
	NewOwner      string    `json:"new_owner"`
	BlockNumber   uint64    `json:"block_number"`
	Timestamp     time.Time `json:"timestamp"`
}

func WatchOwnershipTransferred(ctx context.Context, contract *contracts.Contracts) error {
	events := make(chan *contracts.ContractsOwnershipTransferred)
	opts := &bind.WatchOpts{
		Context: ctx,
	}

	subscription, err := contract.WatchOwnershipTransferred(opts, events, nil, nil)
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
			e := OwnershipTransferred{
				PreviousOwner: event.PreviousOwner.Hex(),
				NewOwner:      event.NewOwner.Hex(),
				BlockNumber:   event.Raw.BlockNumber,
				Timestamp:     time.Now(),
			}

			raw, err := json.MarshalIndent(e, "", "  ")
			if err != nil {
				log.Error().Str("watch", OwnershipTransferredEvent).Err(err).Msg("failed to marshal event")
				continue
			}

			log.Info().Str("watch", OwnershipTransferredEvent).Msg(string(raw))
		}
	}
}
