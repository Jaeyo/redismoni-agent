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

func newConfig() *config {
	return &config{
		make(map[string]interface{}),
	}
}

var instance *config
var once sync.Once
func GetConfigInstance() *config {
	once.Do(func() {
		instance = newConfig()
	})
	return instance
}
