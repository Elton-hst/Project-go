package redis

import (
	"context"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestRedisSaveAndRetrieve(t *testing.T) {
	rdb := redis.NewClient(NewRedis().Client.Options())
	defer rdb.Close()

	key := "teste-chave"
	value := "teste-valor"

	err := rdb.Set(context.Background(), key, value, time.Duration(3000)).Err()
	assert.NoError(t, err, "Erro ao salvar valor no Redis")

	val, err := rdb.Get(context.Background(), key).Result()
	assert.NoError(t, err, "Erro ao recuperar valor do Redis")
	assert.Equal(t, value, val, "Valores recuperado e esperado não são iguais")
}
