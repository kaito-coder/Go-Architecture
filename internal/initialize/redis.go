package initialize

import (
	"context"
	"fmt"

	"github.com/kaito-coder/go-ecommerce-architecture/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	r := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", r.Host, r.Port),
		Password: r.Password,
		DB:       r.DB,
		PoolSize: 10,
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Failed to connect to Redis", zap.Error(err))
	}

	fmt.Println("Init Redis is running")
	global.Rdb = rdb
	redisExample()
}

func redisExample(){
	// Set a value
	err := global.Rdb.Set(ctx, "test", "test kaito", 0).Err()
	if err != nil {
		fmt.Println("Error redis setting", zap.Error(err))
		return
	}
	value, err := global.Rdb.Get(ctx, "test").Result()
	if err != nil {
		fmt.Println("Error redis getting", zap.Error(err))
		return
	}
	global.Logger.Info("key", zap.String("test", value))
}
