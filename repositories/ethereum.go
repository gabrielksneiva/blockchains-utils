package repositories

import (
	"context"

	"blockchains-utils/infra/eventbus"
	"blockchains-utils/infra/rpc"
)

type EthereumRepository struct {
	BaseRepo
}

func NewEthereumRepository(client *rpc.SimulatedClient, bus *eventbus.InMemoryBus) *EthereumRepository {
	return &EthereumRepository{BaseRepo{Chain: "ethereum", client: client, bus: bus}}
}

func (e *EthereumRepository) Connect(ctx context.Context) error {
	return e.BaseRepo.Connect(ctx)
}
