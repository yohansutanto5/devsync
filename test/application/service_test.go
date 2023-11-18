package application_test

import (
	"app/model"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetList_Normal_1(t *testing.T) {
	apps, err := applicationService.GetList()
	assert.ErrorIs(t, err, nil)
	fmt.Println(len(apps))
	assert.Greater(t, len(apps), 2)
}
func TestInsert_Normal_1(t *testing.T) {
	apps := []*model.Application{
		{
			Name:        "App1",
			ID:          "APP1",
			Category:    1,
			OwnerID:     4,
			LeadID:      5,
			Description: "Description for App1",
		},
		{
			Name:        "App2",
			ID:          "APP2",
			Category:    2,
			OwnerID:     6,
			LeadID:      7,
			Description: "Description for App2",
		},
		{
			Name: "App3",
			ID:   "APP3",

			Category:    1,
			OwnerID:     7,
			LeadID:      8,
			Description: "Description for App3",
		},
		{
			Name: "App4",
			ID:   "APP4",

			Category:    3,
			OwnerID:     5,
			LeadID:      7,
			Description: "Description for App4",
		},
		{
			Name: "App5",
			ID:   "APP5",

			Category:    2,
			OwnerID:     4,
			LeadID:      7,
			Description: "Description for App5",
		},
	}

	for _, app := range apps {
		err := applicationService.Insert(app)
		assert.ErrorIs(t, err, nil)
	}

}
