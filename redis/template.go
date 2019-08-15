package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/gomodule/redigo/redis"
)

var (
	redisHost = os.Getenv("REDIS_HOST")
	redisPort = os.Getenv("REDIS_PORT")

	key = "KEY"
)

func FLUSH_ALL() error {
	conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", redisHost, redisPort))
	if err != nil {
		return err
	}
	conn.Do("FLUSHALL")
	return nil
}

func getDataFromCache(key string) ([]byte, error) {
	conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", redisHost, redisPort))
	if err != nil {
		return nil, err
	}

	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func setDataToCache(key string, data []byte) error {
	conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", redisHost, redisPort))
	if err != nil {
		return err
	}
	err = conn.Do("SET", key, data)
	if err != nil {
		return err
	}
	return nil
}

func getListDataFromCache(key string) ([]byte, error) {
	conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", redisHost, redisPort))
	if err != nil {
		return err
	}

	data, err := redis.Bytes(conn.Do("LANGE", key, 0, -1))
	if err != nil {
		return nil, err
	}
	return data, nil
}

// RPUSHは最後に追加
func pushListDataToCache(key string, data []byte) error {
	conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", redisHost, redisPort))
	if err != nil {
		return err
	}
	err = conn.Do("RPUSH", key, data)
	if err != nil {
		return err
	}
	return nil
}

// マッチするものを1つ削除
func removeListDataFromCache(key string, data []byte) error {
	conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", redisHost, redisPort))
	if err != nil {
		return err
	}
	err = conn.Do("LREM", key, 1, data)
	if err != nil {
		return err
	}
	return nil
}

func makeKey(id int64) string {
	ID := strconv.Itoa(int(id))
	return Key + ID
}
