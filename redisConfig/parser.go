package redisConfig

import (
	"path"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func NewConfiguration(filePath string) (*Configuration, error) {
	config := &Configuration{
		make(map[string]string),
	}
	err := config.load(filePath)
	if err != nil {
		return nil, err
	}
	return config, nil
}

type Configuration struct {
	data map[string]string
}

func (config *Configuration) set(key, value string) {
	config.data[key] = value
}

func (config *Configuration) GetBoolean(key string, defaultValue bool) (bool, error) {
	value, exists := config.data[key]
	if !exists {
		return defaultValue, nil
	}
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return false, err
	}
	return boolValue, nil
}

func (config *Configuration) GetInt(key string, defaultValue int) (int, error) {
	value, exists := config.data[key]
	if !exists {
		return defaultValue, nil
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return -1, err
	}
	return intValue, nil
}

func (config *Configuration) GetString(key, defaultValue string) (string, error) {
	value, exists := config.data[key]
	if !exists {
		return defaultValue, nil
	}
	return value, nil
}

func (config *Configuration) load(filePath string) error {
	filePath = path.Clean(filePath)
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(bufio.NewReader(file))
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}

		index := strings.Index(line, " ")
		if index < 0 {
			config.set(line, "true")
			continue
		}

		key := line[:index]
		value := line[index+1:]

		if strings.HasPrefix(value, "\"") &&
			strings.HasSuffix(value, "\"") &&
			!strings.Contains(value[1:len(value)-1], "\""){
			value = value[1:len(value)-1]
		}

		config.set(key, value)
	}

	return nil
}
