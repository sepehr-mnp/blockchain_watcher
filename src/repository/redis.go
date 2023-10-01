package repository

import (
	"context"
	"strconv"

	"github.com/go-redis/redis/v8"
)

type RedisRepo interface {
	AddAddressToUnderWatch(address string) error
	RemoveAddress(address string) error
	AddressUnderWatcher(address string) (bool, error)
	SetLatestBlock(uint64)
	GetLatestBlock() uint64
}

type redisRepo struct {
	RedisClient *redis.Client
}

func NewRedisRepo(redisClient *redis.Client) RedisRepo {
	rr := redisRepo{
		RedisClient: redisClient,
	}
	return &rr
}

func (cli *redisRepo) AddressUnderWatcher(address string) (bool, error) {
	ctx := context.Background()
	yes, err := cli.RedisClient.SIsMember(ctx, "WATCHER_ADDRESS", address).Result()
	if err != nil {
		return false, err
	}
	if yes {
		return true, nil
	}
	return false, nil
}

func (cli *redisRepo) AddAddressToUnderWatch(address string) error {
	ctx := context.Background()
	err := cli.RedisClient.SAdd(ctx, "WATCHER_ADDRESS", address).Err()
	if err != nil {
		return err
	}
	return nil
}

func (cli *redisRepo) RemoveAddress(address string) error {
	ctx := context.Background()
	err := cli.RedisClient.SRem(ctx, "WATCHER_ADDRESS", address).Err()
	if err != nil {
		return err
	}
	return nil
}

func (cli *redisRepo) SetLatestBlock(lastCheckedBlock uint64) {
	cli.RedisClient.Set(context.Background(), "lastCheckedBlock", strconv.FormatUint(lastCheckedBlock, 10), 0)
}

func (cli *redisRepo) GetLatestBlock() uint64 {
	redisLastValue, _ := cli.RedisClient.Get(context.Background(), "lastCheckedBlock").Result()
	redisLastBlock, _ := strconv.ParseUint(redisLastValue, 10, 64)
	return redisLastBlock
}
