package acache

import (
	"context"
	"time"
)

type Cache[K ~string] interface {
	Get(ctx context.Context, key K) (string, error)
	GetInt64(ctx context.Context, key K) (int64, error)
	GetUint64(ctx context.Context, key K) (uint64, error)
	GetBytes(ctx context.Context, key K) ([]byte, error)
	GetJSON(ctx context.Context, key K, dest any) error
	GetBool(ctx context.Context, key K) (bool, error)
	Set(ctx context.Context, key K, value any, exp time.Duration) error
	SetJSON(ctx context.Context, key K, value any, exp time.Duration) error
	IncrBy(ctx context.Context, key K, value int64) error
	Delete(ctx context.Context, key K) error
	Reset(ctx context.Context) error
	Close() error
}
