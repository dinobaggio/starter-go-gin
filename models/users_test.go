package models

import (
	"fmt"
	"testing"
)

func TestGetUser(t *testing.T) {
	users, err := NewUser().GetUsersWithPagination()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(users)
}

func TestGetUserByID(t *testing.T) {
	user, err := NewUser().GetByID(1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(user)
}

func TestGetUserByUUID(t *testing.T) {
	user, err := NewUser().GetByUUID("3d6ebd8d-6a04-449d-b232-bd1ca3ead0bf")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(user)
}
