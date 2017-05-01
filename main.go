package main

import "redismoni-agent/rdb"

func main() {
	profiler := rdb.NewProfiler()
	profiler.StartProfile()
}
