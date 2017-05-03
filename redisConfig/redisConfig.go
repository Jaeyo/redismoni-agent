package redisConfig

import (
	"path"
	"os"
	"bufio"
	"strings"
	"strconv"
	"redismoni-agent/common/config"
	"redismoni-agent/common/util"
)

var redisConfValues map[string]string

func init() {
	redisConfValues = make(map[string]string)

	redisConfigFilePath := config.GetRedisConfigFilePath()
	redisConfigFilePath = path.Clean(redisConfigFilePath)
	redisConfigFile, err := os.Open(redisConfigFilePath)
	if err != nil {
		util.ExitWithError(err)

	}
	defer redisConfigFile.Close()

	scanner := bufio.NewScanner(bufio.NewReader(redisConfigFile))
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}

		index := strings.Index(line, " ")
		if index < 0 {
			redisConfValues[line] = "true"
			continue
		}

		key := line[:index]
		value := line[index+1:]

		if strings.HasPrefix(value, "\"") &&
			strings.HasSuffix(value, "\"") &&
			!strings.Contains(value[1:len(value)-1], "\""){
			value = value[1:len(value)-1]
		}

		redisConfValues[key] = value
	}
}

func GetBoolean(key string, defaultValue bool) (bool, error) {
	value, exists := redisConfValues[key]
	if !exists {
		return defaultValue, nil
	}
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return false, err
	}
	return boolValue, nil
}

func GetInt(key string, defaultValue int) (int, error) {
	value, exists := redisConfValues[key]
	if !exists {
		return defaultValue, nil
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return -1, err
	}
	return intValue, nil
}

func GetString(key, defaultValue string) (string, error) {
	value, exists := redisConfValues[key]
	if !exists {
		return defaultValue, nil
	}
	return value, nil
}
