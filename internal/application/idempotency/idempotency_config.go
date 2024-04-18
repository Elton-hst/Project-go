package idempotency

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Elton-hst/internal/infrastructure/redis"
	rd "github.com/redis/go-redis/v9"
)

type IdempotentHandler[T any] struct {
	redis *redis.Redis
}

func (i *IdempotentHandler[T]) Startt(ctx context.Context, idempotencyKey string) (T, bool, error) {
	var t T
	tr := i.redis.Client.HSetNX(ctx, "idempotencyKey:"+idempotencyKey, "status", "start")
	if tr.Err() != nil {
		return t, false, tr.Err()
	}

	if tr.Val() {
		return t, false, nil
	}

	b := i.redis.Client.HGet(ctx, "idempotencyKey:"+idempotencyKey, "status")
	if b == nil {
		return t, false, nil
	}

	return t, true, nil
}

func (i *IdempotentHandler[T]) Store(ctx context.Context, idempotencyKey string, value T) error {
	b, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return i.redis.Client.HSet(ctx, "idempotencyKey:"+idempotencyKey, "value", b).Err()
}

func (h *IdempotentHandler[T]) Start(ctx context.Context, idempotencyKey string, value interface{}) (interface{}, error) {
	val, err := h.redis.Client.Get(ctx, idempotencyKey).Result()
	if err == nil {
		return val, nil
	} else if err != rd.Nil {
		return nil, fmt.Errorf("error checking key in Redis: %v", err)
	}

	result := value

	if err := h.redis.Client.Set(ctx, idempotencyKey, result, 0).Err(); err != nil {
		return nil, fmt.Errorf("error setting key in Redis: %v", err)
	}

	return result, nil
}
