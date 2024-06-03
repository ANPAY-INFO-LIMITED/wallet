// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package sqlc

import (
	"context"
)

type Querier interface {
	FetchLiquidityParams(ctx context.Context) ([]byte, error)
	GetLoopInSwap(ctx context.Context, swapHash []byte) (GetLoopInSwapRow, error)
	GetLoopInSwaps(ctx context.Context) ([]GetLoopInSwapsRow, error)
	GetLoopOutSwap(ctx context.Context, swapHash []byte) (GetLoopOutSwapRow, error)
	GetLoopOutSwaps(ctx context.Context) ([]GetLoopOutSwapsRow, error)
	GetSwapUpdates(ctx context.Context, swapHash []byte) ([]SwapUpdate, error)
	InsertHtlcKeys(ctx context.Context, arg InsertHtlcKeysParams) error
	InsertLoopIn(ctx context.Context, arg InsertLoopInParams) error
	InsertLoopOut(ctx context.Context, arg InsertLoopOutParams) error
	InsertSwap(ctx context.Context, arg InsertSwapParams) error
	InsertSwapUpdate(ctx context.Context, arg InsertSwapUpdateParams) error
	UpsertLiquidityParams(ctx context.Context, params []byte) error
}

var _ Querier = (*Queries)(nil)
