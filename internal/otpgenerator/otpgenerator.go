// package otpgenerator

// import (
// 	"time"
// 	"github.com/pquerna/otp/totp"
// )

// // GenerateSecret generates a new TOTP secret.
// func GenerateSecret() string {
// 	// You can use a library like "github.com/pquerna/otp" to generate a secret.
// 	// Example:
// 	// key, _ := totp.Generate(totp.GenerateOpts{
// 	//     Issuer:      "MyApp",
// 	//     AccountName: "User",
// 	// })

// 	// Return the secret
// 	// return key.Secret()
// 	return "your-secret-key" // Replace with a generated secret
// }

// // GenerateTOTP generates a Time-based One-Time Password (TOTP) using the provided secret.
// func GenerateTOTP(secret string) string {
// 	key, _ := totp.Generate(totp.GenerateOpts{
// 		Issuer:      "MyApp",
// 		AccountName: "User",
// 		Secret:      []byte(secret),
// 		// Other options like Period and Digits
// 	})

// 	// Generate the TOTP code
// 	otp, _ := totp.GenerateCode(key.Secret(), time.Now())

// 	return otp
// }

// // ValidateOTP validates the provided OTP against the given secret.
// func ValidateOTP(otp, secret string) bool {
// 	key, _ := totp.Generate(totp.GenerateOpts{
// 		Secret: []byte(secret),
// 	})

// 	return totp.Validate(otp, key)
// }

// // IsOTPExpired checks if the provided OTP has expired.
// func IsOTPExpired(otp, secret string) bool {
// 	key, _ := totp.Generate(totp.GenerateOpts{
// 		Secret: []byte(secret),
// 	})

// 	return !totp.Validate(otp, key)
// }
