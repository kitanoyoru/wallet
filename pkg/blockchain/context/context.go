package context

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/kitanoyoru/wallet/config"
	"github.com/rs/zerolog/log"
)

var ethClientContextKey = struct{}{}

func WrapToContext(ctx context.Context, client *ethclient.Client) context.Context {
	return context.WithValue(ctx, ethClientContextKey, client)
}

func FromContext(ctx context.Context) *ethclient.Client {
	value := ctx.Value(ethClientContextKey)
	if value == nil {
		client, err := ethclient.Dial(config.Blockchain.Address)
		if err != nil {
			log.Fatal().Err(err).Send()
		}

		return client
	}

	client, ok := value.(*ethclient.Client)
	if !ok {
		client, err := ethclient.Dial(config.Blockchain.Address)
		if err != nil {
			log.Fatal().Err(err).Send()
		}

		return client
	}

	return client
}
