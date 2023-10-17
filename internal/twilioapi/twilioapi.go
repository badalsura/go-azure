package twilioapi

import (
	"errors"
	"log"
	"os"

	"github.com/badalsura/goOtpAuth/internal/models"
	"github.com/badalsura/goOtpAuth/internal/postgresdb"
	"github.com/twilio/twilio-go"
	twilioVerify "github.com/twilio/twilio-go/rest/verify/v2"
)

//TODO: Figure out why .env file was not working and remove this hardcoded username and password as they are a security risk
var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
	Username: "AC4a9bcb4a8c697170f25abd80369e8f99",
	Password: "6a0036c0ed3931efae9a954811e8bbb0",
})

func SendPhoneOTP(phone string) (string, error) {
	params := &twilioVerify.CreateVerificationParams{}
	if phone[0] != '+'{
		phone = "+" + phone
	}
	params.SetTo(phone)
	params.SetChannel("sms")

	resp, err := client.VerifyV2.CreateVerification(os.Getenv("TWILIO_SERVICES_ID"), params)
	if err != nil {
		log.Printf("Error Sending Mobile OTP: %v", err)
		return "", err
	}

	log.Printf("Sent verification '%s'\n", *resp.Sid)

	return *resp.Sid, nil
}

func SendEmailOTP(email string) (string, error) {
	params := &twilioVerify.CreateVerificationParams{}
	params.SetTo(email)
	params.SetChannel("email")

	resp, err := client.VerifyV2.CreateVerification(os.Getenv("TWILIO_SERVICES_ID"), params)
	if err != nil {
		log.Printf("Error Sending Email OTP: %v", err)
		return "", err
	}

	log.Printf("Sent verification '%s'\n", *resp.Sid)

	return *resp.Sid, nil
}

func VerifyPhone(phone string, otp string) error {
	params := &twilioVerify.CreateVerificationCheckParams{}
	if phone[0] != '+'{
		phone = "+" + phone
	}
	params.SetTo(phone)
	params.SetCode(otp)

	resp, err := client.VerifyV2.CreateVerificationCheck(os.Getenv("TWILIO_SERVICES_ID"), params)
	if err != nil {
		log.Printf("Error Verifying Phone: %v", err)
		return err
	} else if resp.Status != nil && *resp.Status == "approved" {
		return nil
	}

	return errors.New("invalid otp")
}

func VerifyEmail(email string, otp string) error {
	params := &twilioVerify.CreateVerificationCheckParams{}
	params.SetTo(email)
	params.SetCode(otp)

	db := postgresdb.DB
	var checkUser models.User
	userExists := db.First(&checkUser, "email LIKE ?", email)
	if userExists.Error != nil {
		return errors.New("email doesn't exists or otp not sent yet")
	}

	resp, err := client.VerifyV2.CreateVerificationCheck(checkUser.EmailOtpSID, params)
	if err != nil {
		log.Printf("Error Verifying Email: %v", err)
		return err
	} else if resp.Status != nil && *resp.Status == "approved" {
		return nil
	}

	return errors.New("invalid otp")
}
