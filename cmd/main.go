package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/badalsura/goOtpAuth/azFunctions/register"
	"github.com/badalsura/goOtpAuth/azFunctions/userAuth"
	"github.com/badalsura/goOtpAuth/azFunctions/verify"
	"github.com/badalsura/goOtpAuth/internal/initializer"
	"github.com/badalsura/goOtpAuth/internal/middleware"
	"github.com/badalsura/goOtpAuth/internal/postgresdb"
)

var DB *gorm.DB

func getPort() string{
	port := ":8080"
	if val,ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok{
		port = ":" + val
	}
	return port
}

func init() {
	initializer.LoadEnv()
	var err error
	postgresdb.DB, err = postgresdb.ConnectDB()
	if err != nil {
		log.Printf("Database Connection Failed!")
		panic(err)
	}
}

func main(){
	router := gin.Default()

	api := router.Group("/api")
	api.POST("/register", register.RegistrationHandler)
	api.GET("/auth/:action", middleware.AuthValidator, userAuth.UserAuthHandler)
	api.POST("/auth/:action", userAuth.UserAuthHandler)
	api.POST("/verify/:type", verify.OTPVerificationHandler)

	port := getPort()
	router.Run(port)
}