package verify

import (
	"errors"
	"net/http"

	"github.com/badalsura/goOtpAuth/internal/postgresdb"
	"github.com/badalsura/goOtpAuth/internal/twilioapi"
	"github.com/gin-gonic/gin"
)


func OTPVerificationHandler(c *gin.RouterGroup) (){
	c.POST("/verify/email", verifyEmail)
	c.POST("/verify/phone", verifyPhone)
}


func verifyEmail(c *gin.Context) (){
	type userData struct{
		Email string
		OTP string
	}

	var user userData
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error" : "Invalid Request Format!"})
		return
	}

	err := twilioapi.VerifyEmail(user.Email, user.OTP)

	if err != nil {
		if err == errors.New("invalid otp") {
			c.JSON(http.StatusBadRequest, gin.H{"Error" : "Invalid OTP"})
		}else {
			c.JSON(http.StatusBadRequest, gin.H{"Error" : err})
		}
		return
	}
	

	db := postgresdb.DB
	db.Where("phone LIKE ?", user.Email).Update("emailverified", true)
	c.JSON(http.StatusOK, gin.H{"data": "Email verified Successfuly!"})
}

func verifyPhone(c *gin.Context) (){
	type userData struct{
		PhoneNumber string
		OTP string
	}

	var user userData
	if c.Bind(&user) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error" : "Invalid Request Format!"})
		return
	}

	err := twilioapi.VerifyPhone(user.PhoneNumber, user.OTP)
	if err != nil {
		if err == errors.New("invalid otp") {
			c.JSON(http.StatusBadRequest, gin.H{"Error" : "Invalid OTP"})
		}else {
			c.JSON(http.StatusBadRequest, gin.H{"Error" : err})
		}
		return
	}

	db := postgresdb.DB
	db.Where("phone LIKE ?", user.PhoneNumber).Update("phoneverified", true)
	c.JSON(http.StatusOK, gin.H{"data": "Phone verified Successfuly!"})
}