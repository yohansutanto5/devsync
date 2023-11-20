package db

import (
	"app/cmd/config"
	"app/constanta"
	"app/pkg/log"
	"app/pkg/util"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DataStore struct {
	Db     *gorm.DB
	DbView *gorm.DB
	Redis  *redis.Client
	// add elasticsearch here
	// add slave connection here
}

func NewDatabase(config config.Configuration) *DataStore {
	// Initiate Primary SQL Database
	connectionString := "host=" + config.Db.Host + " user=" + config.Db.Username +
		" password=" + config.Db.Password + " dbname=" + config.Db.Database +
		" search_path=" + config.Db.Schema + " port=" + config.Db.Port + " sslmode=disable"
	sqlConnection, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to Initiate SQL Primary Database")
		panic(err)
	}
	sqlPoolConnection, _ := sqlConnection.DB()
	sqlPoolConnection.SetMaxOpenConns(config.Db.MaxOpenConn)
	sqlPoolConnection.SetMaxIdleConns(config.Db.MaxIdleConn)
	sqlPoolConnection.SetConnMaxIdleTime(time.Hour)

	// Initiate Redis Connection
	redisClientOptions := &redis.Options{
		Addr:     config.Redis.Address,
		Password: config.Redis.Password,
		DB:       config.Redis.DB,
	}
	redisClient := redis.NewClient(redisClientOptions)

	_, err = redisClient.Ping().Result()
	if err != nil {
		// log.Fatal("Failed to Initiate Redis Connection")
		// panic(err)
	}

	// defer func() {
	// 	redisClient.Close()
	// 	sqlPoolConnection.Close()
	// }()

	return &DataStore{
		Db:    sqlConnection,
		Redis: redisClient,
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

func ConnectDB(username, password, host, dbName string) (*sqlx.DB, error) {
	connectionString := "user=" + username + " password=" + password + " host=" + host + " dbname=" + dbName + " sslmode=disable" + " search_path=app"
	db, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	// Set the maximum number of open and idle connections
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
