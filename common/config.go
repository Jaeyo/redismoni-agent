package common

import "sync"

type config struct {
	values map[string]interface{}
}

func (config *config) SetDebug(isDebug bool) {
	config.values["debug"] = isDebug
}

func (config *config) GetDebug() bool {
	if isDebug, ok := config.values["debug"]; ok {
		return bool(isDebug)
	}
	return false
}

func (config *config) GetVersion() string {
	version, _ := config.values["version"]
	return string(version)
}

func newConfig() *config {
	config := &config{
		make(map[string]interface{}),
	}
	config.values["version"] = "0.0.1"
	return config
}

var instance *config
var once sync.Once
func GetConfigInstance() *config {
	once.Do(func() {
		instance = newConfig()
	})
	return instance
}
