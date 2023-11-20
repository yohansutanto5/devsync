package system_test

import (
	"app/cmd/config"
	"app/db"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

// Create testing Contenxt
var ctx *gin.Context
var configuration config.Configuration
var database *db.DataStore
var route *gin.Engine

func TestMain(m *testing.M) {
	configuration = config.Load("test")
	var err error
	database = db.NewDatabase(configuration)
	if err != nil {
		log.Fatal("Can not initiate test")
	} else {
		// Create a mock HTTP request for testing
		req, _ := http.NewRequest("GET", "/sample", nil)
		w := httptest.NewRecorder()
		ctx, route = gin.CreateTestContext(w)
		ctx.Request = req
		// Run tests
		exitCode := m.Run()

		// Cleanup resources, close the database connection, etc.

		os.Exit(exitCode)
	}
}
