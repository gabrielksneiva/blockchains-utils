package repositories

import (
	"context"

	"blockchains-utils/infra/eventbus"
	"blockchains-utils/infra/rpc"
)

type TronRepository struct {
	BaseRepo
}

func NewTronRepository(client *rpc.SimulatedClient, bus *eventbus.InMemoryBus) *TronRepository {
	return &TronRepository{BaseRepo{Chain: "tron", client: client, bus: bus}}
}

func (t *TronRepository) Connect(ctx context.Context) error {
	return t.BaseRepo.Connect(ctx)
}
