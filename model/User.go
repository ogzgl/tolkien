package model

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const MinPasswordLength = 8
const PasswordKeyLength = 64

type User struct {
	Id                primitive.ObjectID `bson:"_id"`
	Username          string             `bson:"username"`
	Email             string             `bson:"email"`
	NameSurname       string             `bson:"name_surname"`
	PasswordHash      []byte             `bson:"password_hash"`
	PhoneNumber       string             `bson:"phone_number"`
	ProfilePictureUrl string             `bson:"profile_picture_url"`
	Bio               string             `bson:"bio"`
	City              string             `bson:"city"`
	Country           string             `bson:"country"`
	BirthDate         time.Time          `bson:"birth_date"`
	Job               string             `bson:"job"`
	RegistrationDate  time.Time          `bson:"registration_date"`
	IsVerified        bool               `bson:"is_verified"`
	TermsAccepted     bool               `bson:"terms_accepted"`
	Role              string             `bson:"role"`
}

type UserResponse struct {
	Id                primitive.ObjectID `json:"_id"`
	Username          string             `json:"username"`
	Email             string             `json:"email"`
	PhoneNumber       string             `json:"phone_number"`
	NameSurname       string             `json:"name_surname"`
	ProfilePictureUrl string             `json:"profile_picture_url"`
	Bio               string             `json:"bio"`
	City              string             `json:"city"`
	Country           string             `json:"country"`
	BirthDate         time.Time          `json:"birth_date"`
	Job               string             `json:"job"`
	RegistrationDate  time.Time          `json:"registration_date"`
	IsVerified        bool               `json:"is_verified"`
	TermsAccepted     bool               `json:"terms_accepted"`
	Role              string             `json:"role"`
	Token             string             `json:"token"`
}

func (u *User) Validate() error {
	if u.Username == "" {
		return NewInputError("username", "can't be blank")
	}
	if u.Email == "" {
		return NewInputError("email", "can't be blank")
	}
	if u.PhoneNumber == "" {
		return NewInputError("phoneNumber", "can't be blank")
	}
	if u.PasswordHash == nil || len(u.PasswordHash) != PasswordKeyLength {
		return NewInputError("password", "can't be blank")
	}
	return nil
}

func ValidatePassword(password string) error {
	if len(password) < MinPasswordLength {
		return NewInputError("password", fmt.Sprintf("must be at least %d characters in length", MinPasswordLength))
	}
	return nil
}
