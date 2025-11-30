package rpc

import (
	"context"
	"testing"

	"github.com/gabrielksneiva/blockchains-utils/domain"
)

func TestSimulatedClient_BlockAdvanceAndMissingBalance(t *testing.T) {
	c := NewSimulatedClient()
	ctx := context.Background()
	c.AdvanceBlock(42)
	blk, err := c.GetLatestBlock(ctx)
	if err != nil {
		t.Fatalf("get latest block err: %v", err)
	}
	if blk.Number != 42 {
		t.Fatalf("expected block 42 got %d", blk.Number)
	}

	// missing balance should return ErrNotFound
	if _, err := c.GetBalance(ctx, "noaddr"); err == nil {
		t.Fatalf("expected not found")
	}
	// coverage: submit transaction with empty hash
	amt, _ := domain.NewAmountFromString("1")
	tx := domain.Transaction{From: "a", To: "b", Amount: amt}
	h, _ := c.SubmitTransaction(ctx, tx)
	if h == "" {
		t.Fatalf("expected generated hash")
	}
}
