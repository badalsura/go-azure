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
superAdmin
creater 
reveiwer
admin
publisher
directory

4 products
subscriptions
register
signing
company and userAuth
superAdminverifcation
razorpay
