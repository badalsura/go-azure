package verify

import (
	"errors"
	"net/http"

	"github.com/badalsura/goOtpAuth/internal/models"
	"github.com/badalsura/goOtpAuth/internal/postgresdb"
	"github.com/badalsura/goOtpAuth/internal/twilioapi"
	"github.com/gin-gonic/gin"
)

func OTPVerificationHandler(c *gin.Context) {
	verificationType := c.Param("type")
	switch verificationType {
	case "email" :
		verifyEmail(c)
	case "phone" :
		verifyPhone(c)
	default : 
		c.JSON(http.StatusNotFound, gin.H{"Error" : "Path not found!"})
	}
}

func verifyEmail(c *gin.Context) {
	type userData struct {
		Email string `json:"email" binding:"required,email"`
		OTP   string `json:"otp" binding:"required"`
	}

	var user userData
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Request Format!"})
		return
	}

	err := twilioapi.VerifyEmail(user.Email, user.OTP)

	if err != nil {
		if err == errors.New("invalid otp") {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid OTP"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err})
		}
		return
	}

	var db = postgresdb.DB
	var userTable models.User
	db.Model(&userTable).Where("email LIKE ?", user.Email).Update("email_verified", true)
	c.JSON(http.StatusOK, gin.H{"data": "Email verified Successfuly!"})
}

func verifyPhone(c *gin.Context) {
	type userData struct {
		PhoneNumber string `json:"phone" binding:"required"`
		OTP         string `json:"otp" binding:"required"`
	}

	var user userData
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Request Format!"})
		return
	}

	err := twilioapi.VerifyPhone(user.PhoneNumber, user.OTP)
	if err != nil {
		if err.Error() == "invalid otp" {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid OTP"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		}
		return
	}

	var db = postgresdb.DB
	var userTable models.User
	db.Model(&userTable).Where("phone_number LIKE ?", user.PhoneNumber).Update("phone_verified", true)
	c.JSON(http.StatusOK, gin.H{"data": "Phone verified Successfuly!"})
}
