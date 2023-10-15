package cmd

import (
	"os"
	"github.com/gin-gonic/gin"
	"github.com/badalsura/goOtpAuth/azFunctions/register"
	"github.com/badalsura/goOtpAuth/azFunctions/verify"
)

func getPort() string{
	port := ":8080"
	if val,ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok{
		port = ":" + val
	}
	return port
}
func main(){
	router := gin.Default()

	router.POST("/api/register", register.RegistrationHandler)
	router.POST("/api/verify", verify.OTPVerificationHandler)

	
	port := getPort()
	router.Run(port)
}