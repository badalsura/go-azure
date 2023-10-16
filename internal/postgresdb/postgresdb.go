package postgresdb

import (
	"fmt"
	"os"
	"log"

	"github.com/badalsura/goOtpAuth/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "github.com/lib/pq"

)

var DB *gorm.DB

const (
	dbHost     = "localhost"
	dbPort     = 5432
	dbUser     = "postgres"
	dbPassword = "Qwerty123"
	dbName   = "postgres"
)

func ConnectDB() (*gorm.DB, error) {
	connectionString := os.Getenv("POSTGRES_CONNECTION_STRING")
	if connectionString == "" {
		connectionString = fmt.Sprintf("postgreqsql://%v:%v@%v:%v//%v?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
	}

	DB, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Printf("Error opening database connection: %v\n", err)
		return nil, err
	}

	DB.AutoMigrate(&models.User{})
	return DB, nil
}
// func getConnection() (*sql.DB, error) {
// 	connectionString := os.Getenv("POSTGRES_CONNECTION_STRING")
// 	if connectionString == "" {
// 		connectionString = fmt.Sprintf("postgreqsql://%s:%s@%s:%s//%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
// 		// connectionString = "postgresql://username:Qwerty123@localhost:5432/yourdb?sslmode=disable"
// 	}
// 	db, err := sql.Open("postgres", connectionString)
// 	if err != nil {
// 		log.Printf("Error opening database connection: %v\n", err)
// 		return nil, err
// 	}
// 	return db, nil
// }
// func StoreUserData(email, phone, otpSecret string) error {
// 	db, err := getConnection()
// 	if err != nil {
// 		return err
// 	}
// 	defer db.Close()

// 	_, err = db.Exec("INSERT INTO users (email, phone, otp_secret) VALUES ($1, $2, $3)", email, phone, otpSecret)
// 	if err != nil {
// 		log.Printf("Error inserting: %v\n", err)
// 		return err
// 	}

// 	return nil
// }

// func GetUserTOTPSecret(emailOrPhone string) (string, error) {
// 	db, err := getConnection()
// 	if err != nil {
// 		return "", err
// 	}
// 	defer db.Close()

// 	var otpSecret string
// 	err = db.QueryRow("SELECT otp_secret FROM users WHERE email = $1 OR phone = $2", emailOrPhone, emailOrPhone).Scan(&otpSecret)
// 	if err != nil {
// 		log.Printf("Error retrieving TOTP secret: %v\n", err)
// 		return "", err
// 	}

// 	return otpSecret, nil

// }
