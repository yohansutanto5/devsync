package system_test

import (
	"app/handler"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSystemHealth(t *testing.T) {

	handler.GetSystemHealth(ctx, database)
	assert.Equal(t, ctx.Writer.Status(), 200)
	// expectedJSON := `
	// 	{
	// 		"redis":              true,
	// 		"database_primary":   true,
	// 		"database_secondary": false,
	// 	}
	// `
	// // Get the actual result from the response body
	// var actual map[string]interface{}
	// err := json.Unmarshal(ctx.Writer.Body.Bytes(), &actual)
	// if err != nil {
	// 	t.Errorf("error parsing JSON response: %v", err)
	// }
	// assert.JSONEq(t, expectedJSON, actual)
}
