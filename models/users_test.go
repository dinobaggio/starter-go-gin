package models

import (
	"fmt"
	"testing"
)

func TestGetUser(t *testing.T) {
	users, _ := NewUser().GetUsersWithPagination()
	fmt.Println(users)
}
