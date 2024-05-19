package redis_adapter

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/sachatarba/course-db/internal/config"
)

type IRedisConnector interface {
	Connect() *redis.Client
}

type RedisConnector struct {
	Conf *config.RedisConfig
}

func (connector *RedisConnector) Connect() *redis.Client {
	conf := connector.Conf

	addr := fmt.Sprintf("%s:%s", conf.Host, conf.Port)

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: conf.Password,
	})

	return client
}
