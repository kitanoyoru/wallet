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

const MoneySentEvent = "MoneySent"

type MoneySent struct {
	Beneficiary string    `json:"beneficiary"`
	BlockNumber uint64    `json:"block_number"`
	Amount      *big.Int  `json:"amount"`
	Timestamp   time.Time `json:"timestamp"`
}

func WatchMoneySent(ctx context.Context, contract *contracts.Contracts) error {
	events := make(chan *contracts.ContractsMoneySent)
	opts := &bind.WatchOpts{
		Context: ctx,
	}

	subscription, err := contract.WatchMoneySent(opts, events, nil)
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
			e := MoneySent{
				Beneficiary: event.Beneficiary.Hex(),
				BlockNumber: event.Raw.BlockNumber,
				Amount:      common.WeiToEther(event.Amount),
				Timestamp:   time.Now(),
			}

			raw, err := json.MarshalIndent(e, "", "  ")
			if err != nil {
				log.Error().Str("watch", MoneySentEvent).Err(err).Msg("failed to marshal event")
				continue
			}

			log.Info().Str("watch", MoneySentEvent).Msg(string(raw))
		}
	}
}
