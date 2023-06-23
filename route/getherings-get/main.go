package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/bookgetherorg/tolkien/model"
	"github.com/bookgetherorg/tolkien/service"
	"github.com/bookgetherorg/tolkien/util"
	"log"
)

type GetheringsResponse struct {
	Getherings []model.Gethering `json:"getherings"`
}

func Handle(input events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	_, _, err := service.GetCurrentUser(input.Headers["Authorization"])
	if err != nil {
		log.Print(err.Error())
		return util.NewUnauthorizedResponse()
	}

	getherings, err := service.GetGetheringsPaginated()

	response := GetheringsResponse{
		Getherings: getherings,
	}

	return util.NewSuccessResponse(200, response)
}

func main() {
	lambda.Start(Handle)
}
