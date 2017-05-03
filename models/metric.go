package models

import "strconv"

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

func NewRedisRdbMetric(db, type_ string, value interface{}) *Metric {
	return NewMetric("redis", "rdb", db, type_, "", value)
}

func NewRedisInfoMetric(type_ string, value interface{}) *Metric {
	return NewMetric("redis", "info", type_, "", "", value)
}

func NewRedisInfoKeyCountMetric(db, keyCount int) *Metric {
	_db := strconv.Itoa(db)
	_keyCount := strconv.Itoa(keyCount)
	return NewMetric("redis", "info", "key_count", _db, "", _keyCount)
}
