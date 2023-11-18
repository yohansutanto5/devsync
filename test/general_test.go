package test

import (
	"app/cmd/config"
	"app/db"
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis"
)

func TestRedisCloud2(t *testing.T) {
	configuration := config.Load("test")
	ds := db.NewDatabase(configuration)
	if ds.Redis == nil {
		t.FailNow()
	}

	a := ds.Redis.Set("lock", "true", 1*time.Second)
	res, err := a.Result()
	if err != nil {
		t.Fail()
	}
	fmt.Println(res)

}

func TestRedisCloud(t *testing.T) {
	redisClientOptions := &redis.Options{
		Addr:     "redis-19069.c277.us-east-1-3.ec2.cloud.redislabs.com:19069",
		Password: "7cotq2Rw1N0B3Z3uoYI3f9zW6no1hWqZ",
		DB:       0,
	}
	redisClient := redis.NewClient(redisClientOptions)
	a := redisClient.Set("lock", "true", 100*time.Second)
	res, err := a.Result()
	if err != nil {
		t.Fail()
	}
	fmt.Println(res)
	defer func() {
		redisClient.Close()
	}()
}

func TestGetRedisCloud(t *testing.T) {
	redisClientOptions := &redis.Options{
		Addr:     "redis-19069.c277.us-east-1-3.ec2.cloud.redislabs.com:19069",
		Password: "7cotq2Rw1N0B3Z3uoYI3f9zW6no1hWqZ",
		DB:       0,
	}
	redisClient := redis.NewClient(redisClientOptions)
	a := redisClient.Get("lock")
	res, err := a.Result()
	if err != nil && res != "true" {
		t.Fail()
	}

	defer func() {
		redisClient.Close()
	}()

}
