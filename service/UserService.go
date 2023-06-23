package service

import (
	"context"
	"github.com/bookgetherorg/tolkien/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func CreateUser(user model.User) error {
	err := user.Validate()
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
	collection := client.Database(DB).Collection(UserCollection)
	//Perform InsertOne operation & validate against the error.
	user.Id = primitive.NewObjectID()
	_, err = collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Print(err.Error())
		return err
	}
	//Return success without any error.
	return nil

}

//func UpdateUser(oldUser model.User, newUser model.User) error {
//	err := newUser.Validate()
//	if err != nil {
//		return err
//	}
//
//	transactItems := make([]*dynamodb.TransactWriteItem, 0, 3)
//
//	if oldUser.Email != newUser.Email {
//		newUserIdentifier := model.UserIdentifier{
//			Email:       newUser.Email,
//			Username:    newUser.Username,
//			PhoneNumber: newUser.PhoneNumber,
//		}
//
//		newEmailUserItem, err := dynamodbattribute.MarshalMap(newUserIdentifier)
//		if err != nil {
//			return err
//		}
//
//		// Link user with the new email
//		transactItems = append(transactItems, &dynamodb.TransactWriteItem{
//			Put: &dynamodb.Put{
//				TableName:           aws.String(EmailUserTableName),
//				Item:                newEmailUserItem,
//				ConditionExpression: aws.String("attribute_not_exists(Email)"),
//			},
//		})
//
//		// Unlink user from the old email
//		transactItems = append(transactItems, &dynamodb.TransactWriteItem{
//			Delete: &dynamodb.Delete{
//				TableName:           aws.String(EmailUserTableName),
//				Key:                 StringKey("Email", oldUser.Email),
//				ConditionExpression: aws.String("attribute_exists(Email)"),
//			},
//		})
//	}
//
//	newUserItem, err := dynamodbattribute.MarshalMap(newUser)
//	if err != nil {
//		return err
//	}
//
//	// Update user info
//	transactItems = append(transactItems, &dynamodb.TransactWriteItem{
//		Put: &dynamodb.Put{
//			TableName:                 aws.String(UserTableName),
//			Item:                      newUserItem,
//			ConditionExpression:       aws.String("Email = :email"),
//			ExpressionAttributeValues: StringKey(":email", oldUser.Email),
//		},
//	})
//
//	_, err = DynamoDB().TransactWriteItems(&dynamodb.TransactWriteItemsInput{
//		TransactItems: transactItems,
//	})
//	if err != nil {
//		return err
//	}
//
//	return nil
//}

func GetUserByEmail(email string) (model.User, error) {
	if email == "" {
		return model.User{}, model.NewInputError("email", "can't be blank")
	}

	result := model.User{}
	filter := bson.D{primitive.E{Key: "email", Value: email}}
	client, err := GetMongoClient()
	if err != nil {
		log.Print(err.Error())
		return result, err
	}
	collection := client.Database(DB).Collection(UserCollection)
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Print(err.Error())
		return result, err
	}
	return result, nil
}

func GetUserByUsername(username string) (model.User, error) {
	if username == "" {
		return model.User{}, model.NewInputError("username", "can't be blank")
	}

	result := model.User{}
	filter := bson.D{primitive.E{Key: "username", Value: username}}
	client, err := GetMongoClient()
	if err != nil {
		log.Print(err.Error())
		return result, err
	}
	collection := client.Database(DB).Collection(UserCollection)
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Print(err.Error())
		return result, err
	}
	return result, nil
}

func GetCurrentUser(auth string) (*model.User, string, error) {
	username, token, err := model.VerifyAuthorization(auth)
	if err != nil {
		log.Print(err.Error())
		return nil, "", err
	}

	user, err := GetUserByUsername(username)
	if err != nil {
		log.Print(err.Error())
		return nil, "", err
	}

	return &user, token, nil
}
