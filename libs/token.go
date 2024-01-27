package libs

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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

func ValidToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := ExtractToken(c)

		_, err := ParseToken(tokenString)

		if err != nil {
			ResponseUnauthorize(c)
			return
		}

		c.Next()
	}
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}

	bearerToken := c.Request.Header.Get("Authorization")

	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}

	return ""
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(EnvVariable("TOKEN_SECRET")), nil
	})
}

type TokenPayload struct {
	UserID int64
	Email  string
}

func ExtractTokenPayload(c *gin.Context) (*TokenPayload, error) {
	var result TokenPayload
	tokenString := ExtractToken(c)
	token, err := ParseToken(tokenString)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		result.UserID = int64(claims["user_id"].(float64))
		result.Email = claims["email"].(string)
		return &result, nil
	}

	return nil, fmt.Errorf("invalid format token")
}
