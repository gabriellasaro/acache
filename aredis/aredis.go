package aredis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/gabriellasaro/acache"
	"github.com/redis/go-redis/v9"
)

type redisCache[K ~string] struct {
	client *redis.Client
}

func NewARedis[K ~string](client *redis.Client) acache.Cache[K] {
	return &redisCache[K]{
		client: client,
	}
}

func (r *redisCache[K]) Get(ctx context.Context, key K) (string, error) {
	return r.client.Get(ctx, string(key)).Result()
}

func (r *redisCache[K]) GetInt64(ctx context.Context, key K) (int64, error) {
	return r.client.Get(ctx, string(key)).Int64()
}

func (r *redisCache[K]) GetUint64(ctx context.Context, key K) (uint64, error) {
	return r.client.Get(ctx, string(key)).Uint64()
}

func (r *redisCache[K]) GetBytes(ctx context.Context, key K) ([]byte, error) {
	return r.client.Get(ctx, string(key)).Bytes()
}

func (r *redisCache[K]) GetJSON(ctx context.Context, key K, dest any) error {
	data, err := r.client.Get(ctx, string(key)).Bytes()
	if err != nil {
		return err
	}

	return json.Unmarshal(data, dest)
}

func (r *redisCache[K]) GetBool(ctx context.Context, key K) (bool, error) {
	return r.client.Get(ctx, string(key)).Bool()
}

func (r *redisCache[K]) Set(ctx context.Context, key K, value any, exp time.Duration) error {
	return r.client.Set(ctx, string(key), value, exp).Err()
}

func (r *redisCache[K]) SetJSON(ctx context.Context, key K, value any, exp time.Duration) error {
	v, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return r.client.Set(ctx, string(key), v, exp).Err()
}

func (r *redisCache[K]) IncrBy(ctx context.Context, key K, val int64) error {
	return r.client.IncrBy(ctx, string(key), val).Err()
}

func (r *redisCache[K]) Delete(ctx context.Context, key K) error {
	return r.client.Del(ctx, string(key)).Err()
}

func (r *redisCache[K]) Reset(ctx context.Context) error {
	return r.client.FlushDB(ctx).Err()
}

func (r *redisCache[K]) Close() error {
	return r.client.Close()
}
