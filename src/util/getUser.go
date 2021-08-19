package util

import (
	"context"

	"github.com/the-kaustubh/task1/database"
	"github.com/the-kaustubh/task1/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUser(uname, pwd string) (models.User, error) {
	var user models.User
	collection := database.GetCollection()

	filter := bson.M{
		"username": uname,
		"password": pwd,
	}
	e := collection.FindOne(context.TODO(), filter).Decode(&user)
	if e != nil {
		return models.User{}, e
	}

	return user, nil
}
