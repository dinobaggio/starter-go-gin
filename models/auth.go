package models

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func LoginCheck(email, password string) (*User, error) {
	user := NewUser()
	user, err := user.GetByEmail(email)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("not found user")
	}

	err = VerifyPassword(password, *user.Password)
	if err != nil {
		return user, fmt.Errorf("wrong email and password")
	}

	return user, nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GeneratePassword(password string) (*string, error) {
	hashByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	hashString := string(hashByte)

	return &hashString, nil
}
