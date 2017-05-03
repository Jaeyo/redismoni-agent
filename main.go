package main

import (
	"redismoni-agent/rdb"
	"fmt"
	"redismoni-agent/common"
	"flag"
)

func initConfig() {
	isDebug := flag.Bool("d", false, "debug mode")
	needPrintVersion := flag.Bool("v", false, "print version")

	flag.Parse()

	config := common.GetConfigInstance()
	config.SetDebug(*isDebug)

	if needPrintVersion {
		fmt.Println(config.GetVersion())
	}
}

func main() {
	initConfig()

	profiler := rdb.NewProfiler()
	profiler.StartProfile()

	infoChan := make(chan int)
	rdbChan := make(chan int)

	go getInfoMetrics(infoChan)
	go getRdbMetrics(rdbChan)

	infoVal := <- infoChan
	rdbVal := <- rdbChan

	fmt.Println(infoVal, rdbVal)
}

func getInfoMetrics(c chan int) {
	c <- 1
}

func getRdbMetrics(c chan int) interface {
	c <- 2
}
