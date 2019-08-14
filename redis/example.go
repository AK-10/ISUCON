package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gomodule/redigo/redis"
)

var (
	redisHost = os.Getenv("REDIS_HOST")
	redisPort = os.Getenv("REDIS_PORT")

	userKey = "USER-ID-"
)

type User struct {
	ID   int64
	Name string
}

func main() {
}

func setUserToCache(user User) {
	conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", redisHost, redisPort))
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	key := makeUserKey(user.ID)
	encodedUser, err := json.Marshal(user)
	conn.Do("SET", key, encodedUser)
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

func getUserFromCache(id int64) (*User, error) {
	key := makeUserKey(id)

	data, err := getDataFromCache(key)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var u User
	err = json.Unmarshal(data, &u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func getAllUserFromCache() ([]User, error) {
	conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", redisHost, redisPort))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	key := userKey + "*"
	keys, err := redis.Strings(conn.Do("Keys", key))

	users := make([]User, 0, len(keys))
	for _, v := range keys {
		data, err := getDataFromCache(v)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		var u User
		err = json.Unmarshal(data, &u)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func makeUserKey(id int64) string {
	userID := strconv.Itoa(int(id))
	return userKey + userID
}
