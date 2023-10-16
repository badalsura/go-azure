package twilioapi

import (
	"errors"
	"log"
	"os"

	"github.com/twilio/twilio-go"
	twilioVerify "github.com/twilio/twilio-go/rest/verify/v2"
)

var twilioSID = os.Getenv("TWILIO_SERVICES_ID")

var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
    Username: os.Getenv("TWILIO_ACCOUNT_SID"),
    Password: os.Getenv("TWILIO_AUTHTOKEN"),
})

func SendPhoneOTP(phone string) (string, error){
	params := &twilioVerify.CreateVerificationParams{}
	params.SetTo(phone)
	params.SetChannel("sms")

	resp, err := client.VerifyV2.CreateVerification(twilioSID, params)
	if err != nil {
		log.Printf("Error Sending OTP: %v", err)
		return "", err
	}

	log.Printf("Sent verification '%s'\n", *resp.Sid)

	return *resp.Sid, nil
}

func SendEmailOTP(email string) (string, error){
	params := &twilioVerify.CreateVerificationParams{}
	params.SetTo(email)
	params.SetChannel("sms")

	resp, err := client.VerifyV2.CreateVerification(twilioSID, params)
	if err != nil {
		log.Printf("Error Sending OTP: %v", err)
		return "", err
	}

	log.Printf("Sent verification '%s'\n", *resp.Sid)

	return *resp.Sid, nil
}

func VerifyPhone(phone string, otp string) error {
	params := &twilioVerify.CreateVerificationCheckParams{}
	params.SetTo(phone)
	params.SetCode(otp)

	resp, err := client.VerifyV2.CreateVerificationCheck(twilioSID, params)
	if err != nil {
		log.Printf("Error Verifying Phone: %v", err)
		return err
	} else if *resp.Status == "approved" {
		return nil
	}

	return errors.New("invalid otp")
}

func VerifyEmail(email string, otp string) error {
	params := &twilioVerify.CreateVerificationCheckParams{}
	params.SetTo(email)
	params.SetCode(otp)

	resp, err := client.VerifyV2.CreateVerificationCheck(twilioSID, params)
	if err != nil {
		log.Printf("Error Verifying Email: %v", err)
		return err
	} else if *resp.Status == "approved" {
		return nil
	}

	return errors.New("invalid otp")
}