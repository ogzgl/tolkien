package main

import (
	"bytes"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/bookgetherorg/tolkien/model"
	"github.com/bookgetherorg/tolkien/service"
	"github.com/bookgetherorg/tolkien/util"
	"log"
)

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Handle(input events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	request := UserRequest{}
	err := json.Unmarshal([]byte(input.Body), &request)
	if err != nil {
		log.Print(err.Error())
		return util.NewErrorResponse(err)
	}

	user, err := service.GetUserByEmail(request.Email)
	if err != nil {
		log.Print(err.Error())
		return util.NewErrorResponse(err)
	}

	passwordHash, err := model.Scrypt(request.Password)
	if err != nil {
		log.Print(err.Error())
		return util.NewErrorResponse(err)
	}

	if !bytes.Equal(passwordHash, user.PasswordHash) {
		return util.NewErrorResponse(model.NewInputError("password", "wrong password"))
	}

	token, err := model.GenerateToken(user.Username)
	if err != nil {
		log.Print(err.Error())
		return util.NewErrorResponse(err)
	}

	response := model.UserResponse{
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
