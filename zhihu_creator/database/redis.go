package database

import (
	"context"
	"fmt"
	"github.com/huage66/zhihu_go/zhihu_component/logger"

	"github.com/go-redis/redis/v8"
	"github.com/huage66/zhihu_go/zhihu_creator/config"
)

func NewRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Setting.Redis.Host, config.Setting.Redis.Port),
		Username: config.Setting.Redis.Username,
		Password: config.Setting.Redis.Password,
		DB:       config.Setting.Redis.DB,
	})
	_, err := client.Ping(context.TODO()).Result()
	if err != nil {
		panic(err)
	}
	return client
}
