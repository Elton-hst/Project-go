package idempotency

import (
	"context"
	"encoding/json"

	"github.com/Elton-hst/internal/application/logger"
	"github.com/Elton-hst/internal/infrastructure/redis"
	rd "github.com/redis/go-redis/v9"
)

func Start(ctx context.Context, idempotencyKey string, value interface{}) ([]byte, error) {
	rdb := rd.NewClient(redis.NewRedis().Client.Options())
	defer rdb.Close()

	val, err := rdb.Get(ctx, idempotencyKey).Result()
	if err == nil {
		logger.Info.Printf("Success on find %s", val)
		return []byte(val), nil
	} else if err != rd.Nil {
		logger.Error.Printf("Error checking key in Redis: %v", err)
		return nil, err
	}

	result, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}

	if err := rdb.Set(ctx, idempotencyKey, result, 0).Err(); err != nil {
		logger.Error.Printf("Error setting key in Redis: %v", err)
		return nil, err
	}

	return result, nil
}
