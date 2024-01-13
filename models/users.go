package models

import (
	"context"
	"starter-go-gin/config"
	"time"
)

type User struct {
	ID        int64      `json:"id"`
	UUID      string     `json:"uuid"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  *string    `json:"-"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func NewUser() *User {
	return &User{}
}

func (User) GetUsersWithPagination(ctx context.Context) ([]User, error) {
	conn := config.SQLDBConn()
	var users []User
	rows, err := conn.Query("select id, uuid, name, email, password, created_at, updated_at, deleted_at from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user = User{}
		err := rows.Scan(&user.ID,
			&user.UUID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt, &user.DeletedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
