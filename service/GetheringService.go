package service

import (
	"context"
	"github.com/bookgetherorg/tolkien/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func CreateGethering(gethering model.Gethering) error {
	err := gethering.Validate()
	if err != nil {
		log.Print(err.Error())
		return err
	}

	client, err := GetMongoClient()
	if err != nil {
		log.Print(err.Error())
		return err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database(DB).Collection(GetheringCollection)
	//Perform InsertOne operation & validate against the error.
	_, err = collection.InsertOne(context.TODO(), gethering)
	if err != nil {
		log.Print(err.Error())
		return err
	}
	//Return success without any error.
	return nil
}

func GetGetheringsPaginated() ([]model.Gethering, error) {
	client, err := GetMongoClient()
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	collection := client.Database(DB).Collection(GetheringCollection)
	var getherings []model.Gethering

	filter := bson.D{{}} //bson.D{{}} specifies 'all documents'

	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return getherings, findError
	}
	//Map result to slice
	for cur.Next(context.TODO()) {
		t := model.Gethering{}
		err := cur.Decode(&t)
		if err != nil {
			return getherings, err
		}
		getherings = append(getherings, t)
	}

	if len(getherings) == 0 {
		return getherings, mongo.ErrNoDocuments
	}
	return getherings, nil

}
