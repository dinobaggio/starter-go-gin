package libs

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(userID int64, email string) (*string, error) {
	expToken, err := strconv.Atoi(EnvVariable("TOKEN_HOUR_LIFESPAN"))
	if err != nil {
		return nil, err
	}
	secret := EnvVariable("TOKEN_SECRET")

	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(expToken)).Unix()

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	return &token, nil
}
