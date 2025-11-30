package repositories

import (
	"context"

	"github.com/gabrielksneiva/blockchains-utils/infra/eventbus"
	"github.com/gabrielksneiva/blockchains-utils/infra/rpc"
)

type BitcoinRepository struct {
	BaseRepo
}

func NewBitcoinRepository(client *rpc.SimulatedClient, bus *eventbus.InMemoryBus) *BitcoinRepository {
	return &BitcoinRepository{BaseRepo{Chain: "bitcoin", client: client, bus: bus}}
}

func (b *BitcoinRepository) Connect(ctx context.Context) error {
	return b.BaseRepo.Connect(ctx)
}
