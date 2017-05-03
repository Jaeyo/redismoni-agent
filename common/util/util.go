package util

import (
	"redismoni-agent/common/logger"
	"os"
	"fmt"
)

func ExitWithError(errMsg ...interface{}) {
	fmt.Println(errMsg) // TODO remove
	logger.Error(errMsg...)
	os.Exit(1)
}
