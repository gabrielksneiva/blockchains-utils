package postgres

import (
	"context"
	"testing"

	"github.com/jmoiron/sqlx"
)

func TestPGLogger_NilDB(t *testing.T) {
	p := NewPGLogger(nil)
	if err := p.EnsureSchema(context.Background()); err == nil {
		t.Fatalf("expected error when db nil")
	}
	if err := p.LogEvent(context.Background(), "t", "p"); err == nil {
		t.Fatalf("expected error when db nil")
	}
	// nil is handled properly
	_ = sqlx.NewDb(nil, "postgres")
}
