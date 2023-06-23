package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Gethering struct {
	ID                      primitive.ObjectID `bson:"_id"`
	Title                   string             `bson:"title"`
	Description             string             `bson:"description"`
	CreatorUserId           primitive.ObjectID `bson:"creator_user_id"`
	MeetingConnectionString string             `bson:"meeting_connection_string"`
	IsActive                bool               `bson:"is_active"`
	DateTime                time.Time          `bson:"date_time"`
	DurationInMinutes       int64              `bson:"duration_in_minutes"`
	MaximumAttendeeCount    int64              `bson:"maximum_attendee_count"`
	GoogleBooksId           string             `bson:"google_books_id"`
	TagList                 []string           `bson:"tag_list"`
	CreatedAt               time.Time          `bson:"created_at"`
	UpdatedAt               time.Time          `bson:"updated_at"`
}

func (gethering *Gethering) Validate() error {
	if gethering.Title == "" {
		return NewInputError("title", "can't be blank")
	}
	if gethering.Description == "" {
		return NewInputError("description", "can't be blank")
	}
	if gethering.GoogleBooksId == "" {
		return NewInputError("book", "Google Books Id can not be blank")
	}
	if gethering.DateTime.Before(time.Now()) {
		return NewInputError("time", "Event can not be before now")
	}
	return nil
}

type GetheringAttendees struct {
	GetheringId string `bson:"gethering_id"`
	UserId      string `bson:"user_id"`
}

type GetheringFavoriters struct {
	GetheringId string `bson:"gethering_id"`
	UserId      string `bson:"user_id"`
}

type GetheringModerators struct {
	GetheringId string `bson:"gethering_id"`
	UserId      string `bson:"user_id"`
}
