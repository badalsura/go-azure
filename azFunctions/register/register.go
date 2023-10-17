package register

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/badalsura/goOtpAuth/internal/models"
	"github.com/badalsura/goOtpAuth/internal/postgresdb"
	"github.com/badalsura/goOtpAuth/internal/twilioapi"
)

type userData struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
	PhoneNumber string `json:"phone" binding:"required"`
	PhoneVerified bool 
	EmailVerified bool
}


func RegistrationHandler(c *gin.Context) {

	var Data userData

	if err := c.ShouldBindJSON(&Data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid JSON Data Format!"})
		return
	}

	var tempUser models.User
	//generate bcrypt hash of the password (ideally i would do this on client side)
	passHash, err := bcrypt.GenerateFromPassword([]byte(Data.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Password hashing failed!"})
		return
	}

	db := postgresdb.DB
	isOldUser := db.First(&tempUser, "email LIKE ?", Data.Email).Error
	if isOldUser != nil {

		phoneOtpSid, err := twilioapi.SendPhoneOTP(Data.PhoneNumber)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Error sending mobile otp"})
			return
		}

		emailOtpSid, err := "123", nil
		//TODO: Implement email OTP system
		// emailOtpSid, err := twilioapi.SendEmailOTP(Data.Email)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Error sending email otp"})
			return
		}

		user := models.User{
			Name:        Data.Name,
			Email:       Data.Email,
			Password:    string(passHash),
			PhoneNumber: Data.PhoneNumber,
			EmailOtpSID: emailOtpSid,
			PhoneOtpSID: phoneOtpSid,
			PhoneVerified: Data.PhoneVerified,
			EmailVerified: Data.EmailVerified,
		}

		newUser := db.Create(&user)
		if newUser.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error creating new user!"})
		} else {
			c.JSON(http.StatusAccepted, gin.H{"Message": "Registration Successful! \n Go to /api/verify/email and /api/verify/phone to verify emailID and PhoneNo respectively"})
		}
	} else {
		c.JSON(http.StatusConflict, gin.H{"Error": "User already Exists!"})
		return
	}
}

func SendOTP() {

}
