package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/gomodule/redigo/redis"
)

var (
	redisHost = os.Getenv("REDIS_HOST")
	redisPort = os.Getenv("REDIS_PORT")

	key       = "KEY"
	EVENT_KEY = "EVENT"
)

type User struct {
	ID int64
}

type Event struct {
	ID       int64  `json:"id,omitempty"`
	Title    string `json:"title,omitempty"`
	PublicFg bool   `json:"public,omitempty"`
	ClosedFg bool   `json:"closed,omitempty"`
	Price    int64  `json:"price,omitempty"`

	Total   int `json:"total"`
	Remains int `json:"remains"`
}

func main() {
	event, err := getEventsFromCache()
	fmt.Println(event)
	fmt.Println(err)
}

func getEventsFromCache() ([]Event, error) {
	data, err := getAllHashDataFromCache(EVENT_KEY)
	fmt.Println(data)
	if err != nil {
		return nil, err
	}
	var events []Event
	err = json.Unmarshal(data, &events)

	return events, err
}

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
