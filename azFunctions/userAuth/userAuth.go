package userAuth

import (
	"net/http"
	"strconv"
	"time"

	"github.com/badalsura/goOtpAuth/internal/auth"
	"github.com/badalsura/goOtpAuth/internal/models"
	"github.com/badalsura/goOtpAuth/internal/postgresdb"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func UserAuthHandler(c *gin.Context) {
	authType := c.Param("action")
	if authType == "login" && c.Request.Method == http.MethodPost {
		loginUserHandler(c)
	} else if authType == "logout" && c.Request.Method == http.MethodGet {
		logoutUserHandler(c)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"Error" : "Path not found!"})
	}
}
func loginUserHandler(c *gin.Context){
	type userData struct {
		EmailOrPhone string `json:"user" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	var user userData
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error" : "Invalid User Data Format"})
		return
	}

	var checkUser models.User
	db := postgresdb.DB
	userExists := db.First(&checkUser, "email LIKE ? OR phone_number LIKE ?", user.EmailOrPhone, user.EmailOrPhone)
	if userExists.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error" : "User doesn't exists!"})
		return
	}
	//TODO: Check if User has already verified Email and Phone no or not and redirect them to verification page
	
	if !checkUser.PhoneVerified {
		c.JSON(http.StatusUnauthorized, gin.H{"Error" : "Phone not verified yet!"})
		return
	}
	if !checkUser.EmailVerified {
		c.JSON(http.StatusUnauthorized, gin.H{"Error" : "Email not verified yet!"})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(checkUser.Password), []byte(user.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error" : "Incorrect Password"})
		return
	}

	userID := strconv.Itoa(int(checkUser.ID))
	tokenString := auth.GenerateToken(userID)
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("authToken", tokenString, int(time.Hour) * 24 * 30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{"Message" : "Logged in Successfully!"})

}

func logoutUserHandler(c *gin.Context){
	c.SetCookie("authToken", "", -1, "", "", false, false)
	c.JSON(http.StatusOK, gin.H{"Message" : "Logged Out Successfully!"})
}