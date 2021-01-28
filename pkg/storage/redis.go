package storage

import (
	"fmt"

	"github.com/go-redis/redis"
)

// Config ...
type ConfigRedis struct {
	Host     string
	Port     string
	Password string
	DBName   int
}

// NewRedisClient ...
func NewRedisClient(cfg ConfigRedis) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DBName,
	})

	pong, err := client.Ping().Result()
	if err != nil || pong != "PONG" {
		return nil, err
	}

	return client, nil
}
