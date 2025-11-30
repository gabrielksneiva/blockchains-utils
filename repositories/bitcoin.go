package repositories

import (
	"context"

	"blockchains-utils/infra/eventbus"
	"blockchains-utils/infra/rpc"
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
