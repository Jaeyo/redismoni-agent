package redisInfo

import (
	"redismoni-agent/redisConfig"
	"redismoni-agent/common/util"
	"github.com/go-redis/redis"
	"fmt"
)

type RedisInfo struct {

}

func GetInfo() string {
	redisClient := getRedisClient()
	info, err := redisClient.Info().Result()
	if err != nil {
		util.ExitWithError(err)
	}

	return info
}

func getRedisClient() *redis.Client {
	port, err := redisConfig.GetInt("port", -1)
	if err != nil {
		util.ExitWithError(err)
	}
	if port == -1 {
		util.ExitWithError("failed to find redis port")
	}

	password, err := redisConfig.GetString("requirepass", "")
	if err != nil {
		util.ExitWithError(err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("localhost:%d", port),
		Password: password,
		DB: 0,
	})
	_, err = redisClient.Ping().Result()
	if err != nil {
		util.ExitWithError(err)
	}

	return redisClient
}
