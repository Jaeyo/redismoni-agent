package main

import (
	"redismoni-agent/rdb"
	"fmt"
	"redismoni-agent/common"
)

func main() {
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
