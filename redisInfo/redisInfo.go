package redisInfo

import (
	"redismoni-agent/redisConfig"
	"redismoni-agent/common/util"
	"github.com/go-redis/redis"
	"fmt"
	"strings"
	"strconv"
)

type RedisInfo struct {
	ProcessId int
	UptimeInSec int64
	ConnectedClients int
	UsedMemory int64
	UsedCpu int
	KeyCounts map[int]int
}

func newRedisInfo() *RedisInfo {
	return &RedisInfo{
		KeyCounts: make(map[int]int),
	}
}

func GetInfo() *RedisInfo {
	redisClient := getRedisClient()
	info, err := redisClient.Info().Result()
	if err != nil {
		util.ExitWithError(err)
	}

	redisInfo := &RedisInfo{}
	for _, line := range strings.Split(info, "\n") {
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}
		index := strings.Index(line, ":")
		key := line[:index]
		value := line[index+1:]

		switch key {
		case "process_id":
			processId, _ := strconv.Atoi(value)
			redisInfo.ProcessId = processId
		case "uptime_in_seconds":
			uptimeInSec, _ := strconv.Atoi(value)
			redisInfo.UptimeInSec = int64(uptimeInSec)
		case "connected_clients":
			connectedClients , _ := strconv.Atoi(value)
			redisInfo.ConnectedClients = connectedClients
		case "used_memory":
			usedMemory, _ := strconv.Atoi(value)
			redisInfo.UsedMemory = int64(usedMemory)
		case "used_cpu_sys":
			usedCpu, _ := strconv.Atoi(value)
			redisInfo.UsedCpu = usedCpu
		default:
			if strings.HasPrefix(key, "db") {
				db, _ := strconv.Atoi(key[len("db"):])
				keyCount, _ := strconv.Atoi(value[len("keys="):strings.Index(value, ",")])
				redisInfo.KeyCounts[db] = keyCount
			}
		}
	}

	return redisInfo
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
