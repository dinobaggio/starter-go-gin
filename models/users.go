package models

import (
	"database/sql"
	"fmt"
	"reflect"
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

func scanUserRow(result *User, row interface{}) error {
	var err error
	switch reflect.TypeOf(row).String() {
	case "*sql.Rows":
		err = row.(*sql.Rows).Scan(&result.ID,
			&result.UUID,
			&result.Name,
			&result.Email,
			&result.Password,
			&result.CreatedAt,
			&result.UpdatedAt, &result.DeletedAt)
	case "*sql.Row":
		err = row.(*sql.Row).Scan(&result.ID,
			&result.UUID,
			&result.Name,
			&result.Email,
			&result.Password,
			&result.CreatedAt,
			&result.UpdatedAt, &result.DeletedAt)
	default:
		return fmt.Errorf("incorrect type")
	}

	if err != nil {
		return err
	}

	return nil
}

func (*User) GetUsersWithPagination() ([]User, error) {
	conn := config.SQLDBConn()
	var users []User
	rows, err := conn.Query("select id, uuid, name, email, password, created_at, updated_at, deleted_at from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user = User{}
		err := scanUserRow(&user, rows)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (*User) GetByID(id int64) (*User, error) {
	var result User

	conn := config.SQLDBConn()
	row := conn.QueryRow("select id, uuid, name, email, password, created_at, updated_at, deleted_at from users where id = ?", id)

	err := scanUserRow(&result, row)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (*User) GetByUUID(uuid string) (*User, error) {
	var result User

	conn := config.SQLDBConn()
	row := conn.QueryRow("select id, uuid, name, email, password, created_at, updated_at, deleted_at from users where uuid = ?", uuid)

	err := scanUserRow(&result, row)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (*User) GetByEmail(email string) (*User, error) {
	var result User

	conn := config.SQLDBConn()
	row := conn.QueryRow("select id, uuid, name, email, password, created_at, updated_at, deleted_at from users where email = ?", email)

	err := scanUserRow(&result, row)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
