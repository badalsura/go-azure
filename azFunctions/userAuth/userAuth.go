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

func LoginUserHandler(c *gin.Context){
	type userData struct {
		EmailOrPhone string
		Password string
	}
	var user userData
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error" : "Invalid User Data Format"})
		return
	}

	var checkUser models.User
	db := postgresdb.DB
	result := db.First(&checkUser, "email LIKE ? OR phone LIKE ?", user.EmailOrPhone)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error" : "User doesn't exists!"})
		return
	}
	//TODO: Check if User has already verified Email and Phone no or not
	
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

func LogoutUserHandler(c *gin.Context){
	c.SetCookie("authToken", "", -1, "", "", false, false)
	c.JSON(http.StatusOK, gin.H{"Message" : "Logged Out Successfully!"})
}