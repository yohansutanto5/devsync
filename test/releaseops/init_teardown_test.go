package release_ops_test

import (
	"app/Integration"
	"app/cmd/config"
	"app/db"
	"app/model"
	"app/service"
	"fmt"
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
var ReleaseOPSService service.ReleaseOPSService

func TestMain(m *testing.M) {
	configuration = config.Load("test")
	var err error
	database = db.NewDatabase(configuration)
	external := Integration.NewExternalService(&configuration)
	ReleaseOPSService = service.NewReleaseOPSService(database, external)
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

		if err := database.Db.Delete(&model.ReleaseTicket{}).Error; err != nil {
			panic(fmt.Sprintf("Error truncating table: %s", err))
		}
		// if err := database.Db.Exec("SELECT setval(pg_get_serial_sequence('release_tickets', 'id'), coalesce(max(id), 1), false) FROM release_tickets").Error; err != nil {
		// 	panic(fmt.Sprintf("Error resetting sequence: %s", err))
		// }

		os.Exit(exitCode)
	}
}
