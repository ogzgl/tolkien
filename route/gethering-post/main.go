package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/bookgetherorg/tolkien/model"
	"github.com/bookgetherorg/tolkien/service"
	"github.com/bookgetherorg/tolkien/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

type GetheringRequest struct {
	Title                string    `json:"title"`
	Description          string    `json:"description"`
	CreatorUserId        string    `json:"creator_user_id"`
	DateTime             string    `json:"date_time"`
	DurationInMinutes    int64     `json:"duration_in_minutes"`
	MaximumAttendeeCount int64     `json:"maximum_attendee_count"`
	GoogleBooksId        string    `json:"google_books_id"`
	TagList              []string  `json:"tag_list"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}

func Handle(input events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	user, _, err := service.GetCurrentUser(input.Headers["Authorization"])
	if err != nil {
		log.Print(err.Error())
		return util.NewUnauthorizedResponse()
	}
	request := GetheringRequest{}

	err = json.Unmarshal([]byte(input.Body), &request)
	if err != nil {
		log.Print(err.Error())
		return util.NewErrorResponse(err)
	}

	date, err := time.Parse("02-01-2006 15:04", request.DateTime)
	if err != nil {
		log.Print(err.Error())
		return util.NewErrorResponse(err)
	}

	if date.Before(time.Now()) {
		log.Print("Invalid date time")
		return util.NewErrorResponse(model.NewInputError("date", "Event date can not be before now!"))
	}

	log.Print(request)
	gethering := model.Gethering{
		ID:                   primitive.NewObjectID(),
		Title:                request.Title,
		Description:          request.Description,
		CreatorUserId:        user.Id,
		DateTime:             date,
		DurationInMinutes:    request.DurationInMinutes,
		MaximumAttendeeCount: request.MaximumAttendeeCount,
		GoogleBooksId:        request.GoogleBooksId,
		TagList:              request.TagList,
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
	}

	err = service.CreateGethering(gethering)
	if err != nil {
		log.Print(err.Error())
		return util.NewErrorResponse(err)
	}

	return util.NewSuccessResponse(200, nil)
}

func main() {
	lambda.Start(Handle)
}
