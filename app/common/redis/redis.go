package redis

import (
	"encoding/json"
	"log"

	"github.com/gomodule/redigo/redis"
)

func Connect() redis.Conn {
	rds, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatal(err)
	}

	return rds
}

func GetData[T any](RDS redis.Conn, key string) ([]T, error) {
	var data []T
	cache, err := redis.String(RDS.Do("GET", key))
	if err != nil {
		return data, err
	}

	err = json.Unmarshal([]byte(cache), &data)

	return data, err
}

func SetData(RDS redis.Conn, key string, data any) error {
	mstr, err := json.Marshal(data)
	_, err = RDS.Do("SET", key, mstr)
	return err
}

func ClearData(RDS redis.Conn, data_key string) error {
	_, err := RDS.Do("GETDEL", data_key)
	return err
}
