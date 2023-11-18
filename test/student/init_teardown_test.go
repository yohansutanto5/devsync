package student_test

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
var studenService service.StudentService

func TestMain(m *testing.M) {
	configuration = config.Load("dev")
	var err error
	database = db.NewDatabase(configuration)
	studenService = service.NewStudentService(database)
	if err != nil {
		log.Fatal("asd")
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
