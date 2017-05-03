package main

import (
	"redismoni-agent/rdb"
	"fmt"
	"flag"
	"redismoni-agent/common/config"
	"os"
	"redismoni-agent/common/logger"
	"redismoni-agent/redisConfig"
	"redismoni-agent/models"
)

func init() {
	isDebug := flag.Bool("d", false, "debug mode")
	agentKey := flag.String("k", "", "agent key")
	redisConfigFilePath := flag.String("c", "", "redis config file path")
	needPrintVersion := flag.Bool("v", false, "print version")

	flag.Parse()

	if len(*redisConfigFilePath) == 0 {
		logger.Error("redis config file path not specified")
		os.Exit(1)
	}
	if len(*agentKey) == 0 {
		logger.Error("agent key not specified")
		os.Exit(1)
	}

	config.SetRedisConfigFilePath(*redisConfigFilePath)
	config.SetAgentKey(*agentKey)
	config.SetDebug(*isDebug)

	if needPrintVersion {
		fmt.Println(config.GetVersion())
		os.Exit(0)
	}
}

func main() {
	rdbDumpPath := getRdbDumpPath()

	infoChan := make(chan []*models.Metric)
	rdbChan := make(chan []*models.Metric)

	go getInfoMetrics(infoChan)
	go getRdbMetrics(rdbDumpPath, rdbChan)

	infoVal := <- infoChan
	rdbVal := <- rdbChan

	fmt.Println(infoVal, rdbVal)
}

func getRdbDumpPath() string {
	rdbDumpPath, err := redisConfig.GetString("dbfilename", "")
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	if len(rdbDumpPath) == 0 {
		logger.Error("rdb file path not specified in redis configution")
		os.Exit(1)
	}

	return rdbDumpPath
}

func getInfoMetrics(c chan int) {
	c <- 1
}

func getRdbMetrics(rdbDumpPath string, c chan []*models.Metric) {
	profiler := rdb.NewProfiler()
	memUsages, err := profiler.StartProfile(rdbDumpPath)
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}

	metrics := []*models.Metric{}
	for dbNum, memUsage := range memUsages {
		db := fmt.Sprintf("db_%d", dbNum)
		metrics = append(metrics, models.NewRedisMetric(db, "total", "", "", memUsage.GetTotal()))
		metrics = append(metrics, models.NewRedisMetric(db, "string", "", "", memUsage.StringUsage))
		metrics = append(metrics, models.NewRedisMetric(db, "hash", "", "", memUsage.HashUsage))
		metrics = append(metrics, models.NewRedisMetric(db, "set", "", "", memUsage.SetUsage))
		metrics = append(metrics, models.NewRedisMetric(db, "list", "", "", memUsage.ListUsage))
		metrics = append(metrics, models.NewRedisMetric(db, "sorted_set", "", "", memUsage.SortedSetUsage))
	}

	c <- metrics
}
