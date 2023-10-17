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

