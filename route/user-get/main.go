package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/bookgetherorg/tolkien/service"
	"github.com/bookgetherorg/tolkien/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

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

func Handle(input events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	user, token, err := service.GetCurrentUser(input.Headers["Authorization"])
	if err != nil {
		log.Print(err.Error())
		return util.NewUnauthorizedResponse()
	}

	response := UserResponse{
		Id:                user.Id,
		Username:          user.Username,
		Email:             user.Email,
		PhoneNumber:       user.PhoneNumber,
		NameSurname:       user.NameSurname,
		ProfilePictureUrl: user.ProfilePictureUrl,
		Bio:               user.Bio,
		City:              user.City,
		Country:           user.Country,
		BirthDate:         user.BirthDate,
		Job:               user.Job,
		RegistrationDate:  user.RegistrationDate,
		IsVerified:        user.IsVerified,
		TermsAccepted:     user.TermsAccepted,
		Role:              user.Role,
		Token:             token,
	}

	return util.NewSuccessResponse(200, response)
}

func main() {
	lambda.Start(Handle)
}
