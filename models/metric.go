package models

type Metric struct {
	key1 string
	key2 string
	key3 string
	key4 string
	key5 string
	value interface{}
	ver int
}

func NewMetric(key1, key2, key3, key4, key5 string, value interface{}) *Metric {
	return &Metric{key1, key2, key3, key4, key5, value, 1}
}

func NewRedisMetric(key2, key3, key4, key5 string, value interface{}) *Metric {
	return NewMetric("redis", key2, key3, key4, key5, value)
}