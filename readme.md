# blockchains-utils

Go module providing a small, testable abstraction over multiple blockchains (BTC, ETH, SOL, TRON) following DDD / Clean Architecture principles.

Module path: `github.com/gabrielksneiva/blockchains-utils`

Quickstart example

```go
package main

import (
    "context"
    "fmt"
    "time"

    bu "github.com/gabrielksneiva/blockchains-utils"
    "github.com/gabrielksneiva/blockchains-utils/infra/eventbus"
    "github.com/gabrielksneiva/blockchains-utils/infra/rpc"
    "github.com/gabrielksneiva/blockchains-utils/repositories"
)

func main() {
    ctx := context.Background()

    // create infra pieces
    sim := rpc.NewSimulatedClient()
    bus := eventbus.NewInMemoryBus()

    // create repository (example for ethereum)
    repo := &repositories.BaseRepo{
        Chain:  "ethereum",
        client: sim,
        bus:    bus,
    }

    _ = repo.Connect(ctx)

    // subscriber for new transactions
    ch, _ := bus.Subscribe(ctx, bu.EventNewTransaction)
    go func() {
        for v := range ch {
            fmt.Printf("event received: %#v\n", v)
        }
    }()

    // use repo to create tx
    from := bu.domain.Address("0xabc")
    to := bu.domain.Address("0xdef")
    amount, _ := bu.domain.NewAmountFromString("100")
    tx, _ := repo.CreateTransaction(ctx, from, to, amount)

    fmt.Println("created tx:", tx.Hash)

    // advance block and confirm transaction in simulated client
    sim.AdvanceBlock(1)
    repo.ConfirmTransaction(tx.Hash, 1)

    // wait a bit for events
    time.Sleep(200 * time.Millisecond)
}
```

Notes

- This repository is designed to be imported as a module. Create a semantic tag (example: `v0.1.0`) to allow other modules to depend on a stable version.
- For real production use replace `infra/rpc.SimulatedClient` with real chain clients and wire a persistent event logger/handler.

Contributing

- Run tests: `go test ./... -cover`
- Add PRs against `main` and include tests for new behavior.

License: MIT (add appropriate LICENSE file if needed)
