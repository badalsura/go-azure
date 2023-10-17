package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)


func GenerateToken(id string) string {
	var JwtSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid" : id,
		"exp" : time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString(JwtSecretKey)
	if err != nil {
		panic(err)
	}
	return tokenString
}