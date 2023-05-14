package configuration

import (
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/hilmiikhsan/go_rest_api/exception"
)

func NewRedis(config Config) *redis.Client {
	host := config.Get("REDIS_HOST")
	port := config.Get("REDIS_PORT")
	maxPoolSize, err := strconv.Atoi(config.Get("REDIS_POOL_MAX_SIZE"))
	minIdlePoolSize, err := strconv.Atoi(config.Get("REDIS_POOL_MIN_IDLE_SIZE"))
	exception.PanicLogging(err)

	redisStore := redis.NewClient(&redis.Options{
		Addr:         host + ":" + port,
		PoolSize:     maxPoolSize,
		MinIdleConns: minIdlePoolSize,
	})

	return redisStore
}
