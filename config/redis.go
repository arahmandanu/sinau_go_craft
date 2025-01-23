package config

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
)

func InitRedis() (*redis.Options, error) {
	redisOptions, err := redis.ParseURL(GetRedisDsn())
	if err != nil {
		fmt.Println("redis parse err:", err)
		return nil, err
	} else {
		fmt.Println("redis connect success")
	}

	return redisOptions, nil
}

func CallRedis(redisOptions *redis.Options) (*redis.Client, error) {
	RedisClient := redis.NewClient(redisOptions)
	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("redis connect err:", err)
		return nil, err
	}

	return RedisClient, nil
}
