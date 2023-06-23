package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Comment struct {
	ID          primitive.ObjectID `bson:"_id"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
	Body        string             `bson:"body"`
	UserId      string             `bson:"user_id"`
	GetheringId string             `bson:"gethering_id"`
}

func (comment *Comment) Validate() error {
	if comment.Body == "" {
		return NewInputError("body", "can't be blank")
	}
	return nil
}
