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

// ===============================
//  既存のDBのすべてのキーを削除
//  必ず成功する
// ===============================

func FLUSH_ALL() error {
	conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", redisHost, redisPort))
	if err != nil {
		return err
	}
	conn.Do("FLUSHALL")
	return nil
}

// ==============================
//
//  構造体を格納しやすくするため
//  			の関数
//
// ==============================

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

// =================================
//
//  構造体をKey: Valueの組み合わせで
//   格納しやすくするための関数
//
// =================================

// fieldが存在する場合更新
func setHashDataToCache(key, field string, data []byte) error {
	conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", redisHost, redisPort))
	if err != nil {
		return err
	}

	_, err = conn.Do("HSET", key, field, data)
	if err != nil {
		return err
	}
	return nil
}

func getHashDataFromCache(key, field string) ([]byte, error) {
	conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", redisHost, redisPort))
	if err != nil {
		return nil, err
	}

	data, err := redis.Bytes(conn.Do("HGET", key, field))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func removeHashDataFromCache(key, field string) error {
	conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", redisHost, redisPort))
	if err != nil {
		return err
	}

	_, err = conn.Do("HDEL", key, field)
	if err != nil {
		return err
	}
	return nil
}

// func getAllHashDataFromCache(key string) ([]byte, error) {
// 	conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", redisHost, redisPort))
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	strs, err := redis.Strings(conn.Do("HGETALL", key))
// 	if err != nil {
// 		return nil, err
// 	}
// 	var values []string
// 	values = make([]string, 0, len(strs)/2+1)
// 	for i := 0; i < len(strs); i += 2 {
// 		values = append(values, strs[i+1])
// 	}
// 	str := strings.Join(values[:], ",")
// 	str = "[" + str + "]"
//
// 	return []byte(str), err
// }

// 入力された順
func getAllHashDataFromCache(key string) ([]byte, error) {
	conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", redisHost, redisPort))
	if err != nil {
		return nil, err
	}

	strs, err := redis.Strings(conn.Do("HVALS", key))
	if err != nil {
		return nil, err
	}
	str := strings.Join(strs[:], ",")
	str = "[" + str + "]"

	return []byte(str), err
}

// =================================
//
//   構造体をListで
//   格納しやすくするための関数
//
// =================================

func getListDataFromCache(key string) ([]byte, error) {
	conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", redisHost, redisPort))
	if err != nil {
		return nil, err
	}

	strs, err := redis.Strings(conn.Do("LRANGE", key, 0, -1))
	if err != nil {
		return nil, err
	}
	str := strings.Join(strs[:], ",")
	str = "[" + str + "]"

	return []byte(str), err
}

// RPUSHは最後に追加
func pushListDataToCache(key string, data []byte) error {
	conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", redisHost, redisPort))
	if err != nil {
		return err
	}
	_, err = conn.Do("RPUSH", key, data)
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
	_, err = conn.Do("LREM", key, 1, data)
	if err != nil {
		return err
	}
	return nil
}
