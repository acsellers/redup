package main

import (
	"encoding/hex"
	"github.com/simonz05/godis/redis"
)

var Conn *redis.Client

func ConnectToRedis() {
	Conn = redis.New("tcp:127.0.0.1:6379", 0, *RedisPass)
}

func AllKeys() ([]string, error) {
	keys, err := Conn.Keys("*")
	if err == nil {
		return keys, nil
	}
	return nil, err
}

func FindKey(path string) string {
	if len(path) >= 7 {
		decoded_key := HtmlIdToKey(path[6:len(path)])
		keys, err := Conn.Keys("*")
		if err == nil {
			for _, v := range keys {
				if decoded_key == v {
					return v
				}
			}
		}
	}
	return ""
}

func KeyToHtmlId(key string) string {
	return hex.EncodeToString([]byte(key))
}
func HtmlIdToKey(hash string) string {
	val, err := hex.DecodeString(hash)
	if err == nil {
		return string(val)
	}
	return ""
}
