package db

import (
	"app/cmd/config"
	"app/constanta"
	"app/pkg/log"
	"app/pkg/util"
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	_ "github.com/lib/pq"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DataStore struct {
	Db       *gorm.DB
	DbView   *gorm.DB
	Redis    *redis.Client
	Producer *kafka.Producer
	// add elasticsearch here
}

func NewDatabase(config config.Configuration) *DataStore {
	var err error
	// Initiate Primary SQL Database
	sqlConnection := ConnectPrimaryDatabase(config)

	// Initiate Secondary SQL Database

	// Initiate Redis Connection
	var redisClient *redis.Client
	if config.Mode.Redis {
		redisClient, err = ConnectRedis(config)
		if err != nil {
			log.Fatal("Failed to Initiate Redis Connection : " + err.Error())
			return nil
		}
	}

	// Initiate Kafka Producer Connection
	var producerClient *kafka.Producer
	if config.Mode.Kafka {
		producerClient, err = ConnectKafkaProducer(config)
		if err != nil {
			log.Fatal("Failed to Initiate Kafka Producer Connection : " + err.Error())
			return nil
		}
	}

	return &DataStore{
		Db:       sqlConnection,
		Redis:    redisClient,
		Producer: producerClient,
	}
}

func GetContext(c *gin.Context) *DataStore {
	dbService, exists := c.Get("db")
	if exists {
		// Check if dbService is of the expected type
		if db, ok := dbService.(*DataStore); ok {
			return db
		} else {
			log.Error(util.GetTransactionID(c), "Failed to connect to DB", constanta.FailToConnectCode, nil)
		}
	} else {
		// Handle the case where the key "db" does not exist in the Gin context
		log.Error(util.GetTransactionID(c), "Failed to connect to DB", constanta.FailToConnectCode, nil)
	}
	return nil
}

func ConnectRedis(config config.Configuration) (*redis.Client, error) {
	redisClientOptions := &redis.Options{
		Addr:     config.Redis.Address,
		Password: config.Redis.Password,
		DB:       config.Redis.DB,
	}
	redisClient := redis.NewClient(redisClientOptions)

	_, err := redisClient.Ping().Result()
	return redisClient, err
}

func ConnectKafkaProducer(config config.Configuration) (*kafka.Producer, error) {

	return kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": fmt.Sprintf("%s:%s", config.Kafka.Hostname, config.Kafka.Port),
		"sasl.mechanisms":   config.Kafka.AuthMethod,
		"security.protocol": "sasl_ssl",
		"sasl.username":     config.Kafka.User,
		"sasl.password":     config.Kafka.Password,
	})

}

func ConnectPrimaryDatabase(config config.Configuration) *gorm.DB {
	var sqlConnection *gorm.DB
	var err error

	if config.Db.Vendor == "mysql" {
		connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.Db.Username, config.Db.Password, config.Db.Host, config.Db.Port, config.Db.Database)
		sqlConnection, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	} else {
		connectionString := "host=" + config.Db.Host + " user=" + config.Db.Username +
			" password=" + config.Db.Password + " dbname=" + config.Db.Database +
			" search_path=" + config.Db.Schema + " port=" + config.Db.Port + " sslmode=disable"
		sqlConnection, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	}

	if err != nil {
		log.Fatal("Failed to Initiate SQL Primary Database")
		panic(err)
	}
	sqlPoolConnection, _ := sqlConnection.DB()
	sqlPoolConnection.SetMaxOpenConns(config.Db.MaxOpenConn)
	sqlPoolConnection.SetMaxIdleConns(config.Db.MaxIdleConn)
	sqlPoolConnection.SetConnMaxIdleTime(time.Hour)

	return sqlConnection
}
