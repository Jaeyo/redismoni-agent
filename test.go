package main

import (
	"fmt"
	"redismoni-agent/redisConfig"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	config, err := redisConfig.NewConfiguration("./redis-master.conf")
	checkError(err)
	port, err := config.GetInt("port", -1)
	checkError(err)
	fmt.Println(port)
}
