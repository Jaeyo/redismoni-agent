package main

import (
	"fmt"
	"redismoni-agent/redisInfo"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	info := redisInfo.GetInfo()
	fmt.Println(info)
}
