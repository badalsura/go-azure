package models

import(
	"time"
)

type User struct {
	ID uint `json:"id" gorm:"primaryKey;unique"`
	Name string `json:"name" gorm:"not null" validate:"required"`
	Email string `json:"email" gorm:"not null; unique" validate:"required, email"`
	Password string `json:"password" gorm:"not null" validate:"required"`
	PhoneNumber string `json:"phone" gorm:"not null; unique" validate:"required"`
	EmailOtpSID string `json:"emailotpsid"`
	PhoneOtpSID string `json:"phoneotpsid"`
	EmailVerified bool `json:"emailverified"`
	PhoneVerified bool `json:"phoneverified"`
	CreatedAt time.Time
	UpdatedAt time.Time
}