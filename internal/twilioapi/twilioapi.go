package twilioapi

import(
	"os"
	"github.com/twilio/twilio-go"
)

func SendOtp(phone string) (string, error){
	to := "+91" + phone
	return to,nil
}