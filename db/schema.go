package db

import (
	"goport/abi"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type FulfilledOrder struct {
	Hash          [32]byte           `bun:"type:bytea,notnull"`
	Offerer       common.Address     `bun:"type:bytea,notnull"`
	Zone          common.Address     `bun:"type:bytea,notnull"`
	Recipient     common.Address     `bun:"type:bytea,notnull"`
	Offer         []abi.SpentItem    `bun:"type:jsonb,notnull"`
	Consideration []abi.ReceivedItem `bun:"type:jsonb,notnull"`
}

type CancelledOrder struct {
	Hash    [32]byte       `bun:"type:bytea,notnull"`
	Offerer common.Address `bun:"type:bytea,notnull"`
	Zone    common.Address `bun:"type:bytea,notnull"`
	Raw     types.Log      `bun:"type:jsonb,notnull"`
}

type ValidatedOrder struct {
	Hash    [32]byte       `bun:"type:bytea,notnull"`
	Offerer common.Address `bun:"type:bytea,notnull"`
	Zone    common.Address `bun:"type:bytea,notnull"`
	Raw     types.Log      `bun:"type:jsonb,notnull"`
}

type CounterIncremented struct {
	Counter *big.Int       `bun:"type:numeric,notnull"`
	Offerer common.Address `bun:"type:bytea,notnull"`
	Raw     types.Log      `bun:"type:jsonb,notnull"`
}
