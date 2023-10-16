package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func UserAuth(c *gin.Context){

	tokenString, err := c.Cookie("authToken")

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"Message" : "Unauthorized request, User logout",
		})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	//validate Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface {}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error" : "Error! invalid token format"})
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("userid", claims[""])
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}