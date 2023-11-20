package handler

import (
	"app/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSystemHealth(c *gin.Context, ds *db.DataStore) {
	// Variable
	var redis bool = true
	var database_primary bool = true
	var database_secondary bool = true
	var err error

	_, err = ds.Redis.Ping().Result()
	if err != nil {
		redis = false
	}

	sqlDB, err := ds.Db.DB()
	if err != nil {
		database_secondary = false
	} else {
		err = sqlDB.Ping()
		if err != nil {
			database_secondary = false
		}
	}

	// sqlDBView, err := ds.DbView.DB()
	// if err != nil {
	// 	database_primary = false
	// } else {
	// 	err = sqlDBView.Ping()
	// 	if err != nil {
	// 		database_primary = false
	// 	}
	// }

	result := map[string]bool{
		"redis":              redis,
		"database_primary":   database_primary,
		"database_secondary": database_secondary,
	}

	c.JSON(http.StatusOK, result)

}
