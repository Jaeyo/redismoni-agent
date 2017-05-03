package config

var values map[string]interface{}

func init() {
	values = make(map[string]interface{})
	values["version"] = "0.0.1"
}

func getConfigValue(key string, defaultValue interface{}) interface{} {
	if value, ok := values[key]; ok {
		return value
	}
	return defaultValue
}

func SetRedisConfigFilePath(path string) {
	values["redisConfigFilePath"] = path
}

func GetRedisConfigFilePath() string {
	return getConfigValue("redisConfigFilePath", "").(string)
}

func SetAgentKey(agentKey string) {
	values["agentKey"] = agentKey
}

func getAgentKey() string {
	return getConfigValue("agentKey", "").(string)
}

func SetDebug(isDebug bool) {
	values["debug"] = isDebug
}

func GetDebug() bool {
	return getConfigValue("debug", true).(bool) // TODO modify

}

func GetVersion() string {
	return getConfigValue("version", "").(string)
}
