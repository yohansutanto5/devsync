package user_test

import (
	"app/model"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetList_Normal_1(t *testing.T) {
	users, _ := userService.GetList()
	fmt.Println(len(users))
	assert.Greater(t, len(users), 2)
}
func TestInsert_Normal_1(t *testing.T) {
	users := []*model.User{
		{
			Username:  "q",
			FirstName: "w",
			LastName:  "r",
			Email:     "e",
			ProfileID: 1,
			Active:    true,
		},
		{
			Username:  "user2",
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john@example.com",
			ProfileID: 1,
			Active:    false,
		},
		{
			Username:  "user3",
			FirstName: "Alice",
			LastName:  "Smith",
			Email:     "alice@example.com",
			ProfileID: 1,
			Active:    true,
		},
		{
			Username:  "user4",
			FirstName: "Bob",
			LastName:  "Johnson",
			Email:     "bob@example.com",
			ProfileID: 1,
			Active:    false,
		},
		{
			Username:  "user5",
			FirstName: "Eva",
			LastName:  "Davis",
			Email:     "eva@example.com",
			ProfileID: 1,
			Active:    true,
		},
	}

	for _, user := range users {
		err := userService.Insert(user)
		assert.ErrorIs(t, err, nil)
	}

}
