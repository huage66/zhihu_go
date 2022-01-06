package database

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/huage66/zhihu_go/zhihu_component/logger"
	"github.com/huage66/zhihu_go/zhihu_creator/config"
	"net"
)

func NewRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Setting.Redis.Host, config.Setting.Redis.Port),
		Username: config.Setting.Redis.Username,
		Password: config.Setting.Redis.Password,
		DB:       config.Setting.Redis.DB,
		OnConnect: func(ctx context.Context, cn *redis.Conn) error {
			_, err := cn.Ping(ctx).Result()
			if err != nil {
				logger.InfoF("ok")
			}
			return err
		},
		Dialer: func(ctx context.Context, network, addr string) (net.Conn, error) {

		},
	})
}
