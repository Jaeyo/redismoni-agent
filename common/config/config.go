package config

import "sync"

var values map[string]interface{}
var once sync.Once

func initConfigValues() {
	values = make(map[string]interface{})
	values["version"] = "0.0.1"
}

func getConfigValue(key string, defaultValue interface{}) interface{} {
	once.Do(func() {
		initConfigValues()
	})
	if value, ok := values[key]; ok {
		return value
	}
	return defaultValue
}

func setConfigValue(key string, value interface{}) {
	once.Do(func() {
		initConfigValues()
	})
	values[key] = value
}


func SetDebug(isDebug bool) {
	setConfigValue("debug", isDebug)
}

func GetDebug() bool {
	return bool(getConfigValue("debug", false))
}

func GetVersion() string {
	return string(getConfigValue("version", ""))
}
