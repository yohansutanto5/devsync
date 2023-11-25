package application_test

import (
	"app/cmd/config"
	"app/db"
	"app/service"
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
var applicationService service.ApplicationService

func TestMain(m *testing.M) {
	configuration = config.Load("dev")
	var err error
	database = db.NewDatabase(configuration)
	applicationService = service.NewApplicationService(database)
	if err != nil {
		log.Fatal("Can not initiate test")
	} else {
		// Create a mock HTTP request for testing
		req, _ := http.NewRequest("GET", "/sample", nil)
		w := httptest.NewRecorder()
		ctx, _ = gin.CreateTestContext(w)
		ctx.Request = req
		// Run tests
		exitCode := m.Run()

		// Cleanup resources, close the database connection, etc.

		os.Exit(exitCode)
	}
}
