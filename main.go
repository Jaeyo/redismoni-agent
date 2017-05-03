package main

import (
	"redismoni-agent/rdb"
	"fmt"
	"flag"
	"redismoni-agent/common/config"
	"os"
	"redismoni-agent/redisConfig"
	"redismoni-agent/models"
	"redismoni-agent/common/util"
)

func init() {
	isDebug := flag.Bool("d", true, "debug mode")
	agentKey := flag.String("k", "", "agent key")
	redisConfigFilePath := flag.String("c", "", "redis config file path")
	needPrintVersion := flag.Bool("v", false, "print version")

	flag.Parse()

	if len(*redisConfigFilePath) == 0 {
		util.ExitWithError("redis ocnfig file path not specified")
	}
	if len(*agentKey) == 0 {
		util.ExitWithError("agent key not specified")
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
		util.ExitWithError(err)
	}
	if len(rdbDumpPath) == 0 {
		util.ExitWithError("rdb file path not specified in redis configuration")
	}

	return rdbDumpPath
}

func getInfoMetrics(c chan int) {
	// TODO IMME
}

func getRdbMetrics(rdbDumpPath string, c chan []*models.Metric) {
	profiler := rdb.NewProfiler()
	memUsages, err := profiler.StartProfile(rdbDumpPath)
	if err != nil {
		util.ExitWithError(err)
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
