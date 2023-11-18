package test

import (
	"app/cmd/config"
	"app/db"
	"app/pkg/util"
	"fmt"
	"log"
	"os"
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

func TestGetEnv(t *testing.T) {
	if os.Getenv("config") != "/home/yohan/devsync/cmd/config" {
		fmt.Println(os.LookupEnv("config"))
		t.Fail()
	}
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

type SourceStruct struct {
	SourceField1 int          `convert:"TargetField1"`
	SourceField2 string       `convert:"TargetField2"`
	NestedStruct NestedStruct `convert:"NestedTargetStruct"`
}

type TargetStruct struct {
	TargetField1 int
	TargetField2 string
	NestedStruct NestedTargetStruct
}

type NestedStruct struct {
	NestedField1 int `convert:"NestedTargetField1"`
	NestedField2 string
}

type NestedTargetStruct struct {
	NestedTargetField1 int
	NestedTargetField2 string
}

func TestStructConverter(t *testing.T) {
	sourceInstance := SourceStruct{
		SourceField1: 42,
		SourceField2: "Hello, World!",
		NestedStruct: NestedStruct{
			NestedField1: 24,
			NestedField2: "Nested Hello!",
		},
	}

	// Convert the source struct to the target struct
	var targetInstance TargetStruct
	util.ConvertStruct(sourceInstance, &targetInstance)
	fmt.Printf("Source Struct: %+v\n", sourceInstance)
	fmt.Println(fmt.Printf("Target Struct: %+v\n", targetInstance))
	log.Fatal(targetInstance)
}
