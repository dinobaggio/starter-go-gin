package models

import (
	"context"
	"fmt"
	"testing"
)

func TestGetUser(t *testing.T) {
	ctx := context.Background()
	users, _ := NewUser().GetUsersWithPagination(ctx)
	fmt.Println(users)
}
