package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/bookgetherorg/tolkien/model"
	"github.com/bookgetherorg/tolkien/service"
	"github.com/bookgetherorg/tolkien/util"
	"log"
	"time"
)

type UserRequest struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type UserResponse struct {
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Image       string    `json:"image"`
	Bio         string    `json:"bio"`
	Token       string    `json:"token"`
	PhoneNumber string    `json:"phone_number"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Job         string    `json:"job"`
	BirthDate   time.Time `json:"birth_date"`
}

func Handle(input events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	request := UserRequest{}
	err := json.Unmarshal([]byte(input.Body), &request)
	if err != nil {
		log.Print(err.Error())
		return util.NewErrorResponse(err)
	}

	err = model.ValidatePassword(request.Password)
	if err != nil {
		log.Print(err.Error())
		return util.NewErrorResponse(err)
	}

	passwordHash, err := model.Scrypt(request.Password)
	if err != nil {
		log.Print(err.Error())
		return util.NewErrorResponse(err)
	}
	log.Print(request)
	user := model.User{
		Username:     request.Username,
		Email:        request.Email,
		PasswordHash: passwordHash,
		PhoneNumber:  request.PhoneNumber,
	}

	err = service.CreateUser(user)
	if err != nil {
		log.Print(err.Error())
		return util.NewErrorResponse(err)
	}

	token, err := model.GenerateToken(user.Username)
	if err != nil {
		log.Print(err.Error())
		return util.NewErrorResponse(err)
	}

	response := UserResponse{
		Username: user.Username,
		Email:    user.Email,
		Image:    user.ProfilePictureUrl,
		Bio:      user.Bio,
		Token:    token,
	}

	return util.NewSuccessResponse(201, response)
}

func main() {
	lambda.Start(Handle)
}
