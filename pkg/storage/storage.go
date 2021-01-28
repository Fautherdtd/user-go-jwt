package storage

import (
	"os"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

func InitRedis() (redis.Client, error) {
	dbName, err := strconv.Atoi(os.Getenv("REDIS_DBName"))
	if err == nil {
		logrus.Fatalf(err.Error())
	}

	client, err := NewRedisClient(ConfigRedis{
		Host:     os.Getenv("REDIS_HOST"),
		Port:     os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DBName:   dbName,
	})
	if err != nil {
		logrus.Fatalf("error initializing redis client: %s", err.Error())
	}

	return *client, nil
}
