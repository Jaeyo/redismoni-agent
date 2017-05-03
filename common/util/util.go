package util

import (
	"redismoni-agent/common/logger"
	"os"
)

func ExitWithError(errMsg ...interface{}) {
	logger.Error(errMsg...)
	os.Exit(1)
}
