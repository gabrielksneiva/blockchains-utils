package repositories

import (
	"context"

	"blockchains-utils/infra/eventbus"
	"blockchains-utils/infra/rpc"
)

type SolanaRepository struct {
	BaseRepo
}

func NewSolanaRepository(client *rpc.SimulatedClient, bus *eventbus.InMemoryBus) *SolanaRepository {
	return &SolanaRepository{BaseRepo{Chain: "solana", client: client, bus: bus}}
}

func (s *SolanaRepository) Connect(ctx context.Context) error {
	return s.BaseRepo.Connect(ctx)
}
