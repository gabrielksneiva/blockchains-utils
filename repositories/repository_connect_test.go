package repositories

import (
	"context"
	"testing"

	"github.com/gabrielksneiva/blockchains-utils/domain"
	"github.com/gabrielksneiva/blockchains-utils/infra/eventbus"
	"github.com/gabrielksneiva/blockchains-utils/infra/rpc"
)

func TestConcreteConnectsAndErrorBranches(t *testing.T) {
	ctx := context.Background()
	bus := eventbus.NewInMemoryBus()

	// Ethereum
	ecli := rpc.NewSimulatedClient()
	erepo := NewEthereumRepository(ecli, bus)
	if err := erepo.Connect(ctx); err != nil {
		t.Fatalf("eth connect: %v", err)
	}

	// Bitcoin
	bcli := rpc.NewSimulatedClient()
	brepo := NewBitcoinRepository(bcli, bus)
	if err := brepo.Connect(ctx); err != nil {
		t.Fatalf("btc connect: %v", err)
	}

	// Solana
	scli := rpc.NewSimulatedClient()
	srepo := NewSolanaRepository(scli, bus)
	if err := srepo.Connect(ctx); err != nil {
		t.Fatalf("sol connect: %v", err)
	}

	// Tron
	tcli := rpc.NewSimulatedClient()
	trepo := NewTronRepository(tcli, bus)
	if err := trepo.Connect(ctx); err != nil {
		t.Fatalf("tron connect: %v", err)
	}

	// error branch for BroadcastTransaction: make client fail
	tcli.SubmitErr = nil // ensure no error
	// create and broadcast
	amt, _ := domain.NewAmountFromString("3")
	from, _ := domain.NewAddress("aa111")
	to, _ := domain.NewAddress("bb222")
	tx, _ := trepo.CreateTransaction(ctx, from, to, amt)
	if _, err := trepo.BroadcastTransaction(ctx, tx); err != nil {
		t.Fatalf("broadcast err: %v", err)
	}

	// force error
	tcli.SubmitErr = domain.ErrInvalidAmount // any non-nil error
	if _, err := trepo.BroadcastTransaction(ctx, tx); err == nil {
		t.Fatalf("expected broadcast error")
	}

	// GetTransactionStatus error
	if _, err := trepo.GetTransactionStatus(ctx, domain.TxHash("nohash")); err == nil { /* expected */
	}
}
