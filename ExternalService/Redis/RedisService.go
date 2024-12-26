package redis

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
)

var client *redis.Client

func Init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

}

func SetValue(key string, value interface{}) bool {
	data, err := json.Marshal(value)
	if err != nil {
		return false
	}
	err = client.Set(context.Background(), key, data, 0).Err()
	if err != nil {
		return false
	}
	return true
}

func GetValue(key string) (string, error) {
	res, err := client.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}
	return res, err
}

func RemoveValue(key string) error {
	err := client.Del(context.Background(), key).Err()
	if err != nil {
		return err
	}
	return nil
}
