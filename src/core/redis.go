package core

import (
	"fmt"

	"github.com/go-redis/redis/v8"

	log "github.com/sirupsen/logrus"
)

func OpenRedis(address, pass string, db int) *redis.Client {
	cl := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: pass,
		DB:       db,
	})
	log.Info(fmt.Sprintf("Redis is connected: %s", address))
	return cl
}
