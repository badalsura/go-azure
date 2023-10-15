package postgresdb

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func getConnection() (*sql.DB, error) {
	connectionString := os.Getenv("POSTGRES_CONNECTION_STRING")
	if connectionString == "" {
		connectionString = "postgresql://username:Qwerty123@localhost:5432/yourdb?sslmode=disable"
	}
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Printf("Error opening database connection: %v\n", err)
		return nil, err
	}
	return db, nil
}
func StoreUserData(email, phone, otpSecret string) error {
	db, err := getConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO users (email, phone, otp_secret) VALUES ($1, $2, $3)", email, phone, otpSecret)
	if err != nil {
		log.Printf("Error inserting: %v\n", err)
		return err
	}

	return nil
}

func GetUserTOTPSecret(emailOrPhone string) (string, error) {
	db, err := getConnection()
	if err != nil {
		return "", err
	}
	defer db.Close()

	var otpSecret string
	err = db.QueryRow("SELECT otp_secret FROM users WHERE email = $1 OR phone = $2", emailOrPhone, emailOrPhone).Scan(&otpSecret)
	if err != nil {
		log.Printf("Error retrieving TOTP secret: %v\n", err)
		return "", err
	}

	return otpSecret, nil

}
