package main

import (
	"encoding/json"
	"fmt"
	"os"
	// "strconv"
	"strings"

	"github.com/gomodule/redigo/redis"
)

var (
	redisHost = os.Getenv("REDIS_HOST")
	redisPort = os.Getenv("REDIS_PORT")

	key = "KEY"
)

func getSetDataFromCache(key string) ([]byte, error) {
	conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", redisHost, redisPort))
	if err != nil {
		return nil, err
	}

	strs, err := redis.Strings(conn.Do("SINTER", key))
	if err != nil {
		return nil, err
	}
	str := strings.Join(strs[:], ",")
	str = "[" + str + "]"

	return []byte(str), err
}

func pushSetDataToCache(key string, data []byte) error {
	conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", redisHost, redisPort))
	if err != nil {
		return err
	}
	_, err = conn.Do("SADD", key, data)
	if err != nil {
		return err
	}
	return nil
}

// マッチするものを1つ削除
func removeSetDataFromCache(key string, data []byte) error {
	conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", redisHost, redisPort))
	if err != nil {
		return err
	}
	_, err = conn.Do("SREM", key, data)
	if err != nil {
		return err
	}
	return nil
}
