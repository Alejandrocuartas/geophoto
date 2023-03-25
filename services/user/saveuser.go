package user

import (
	"context"
	"fmt"

	"github.com/Alejandrocuartas/geophoto/database/collections"
	"github.com/Alejandrocuartas/geophoto/graph/model"
	"github.com/Alejandrocuartas/geophoto/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MyError struct {
	message string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("my error occurred: %s", e.message)
}

func SaveUser(ctx context.Context, username string, password string) (model.User, error) {
	var User *mongo.Collection = collections.UserCollection()
	//validate username is unique
	filter := bson.M{"username": username}
	userExist := User.FindOne(ctx, filter)
	var existing model.User
	err := userExist.Decode(&existing)
	if err != mongo.ErrNoDocuments {
		return model.User{}, err
	}
	//save user
	salt, e := helpers.EncryptPassword(password)
	if e != nil {
		return model.User{}, e
	}
	newUser := model.User{
		Username: username,
		Password: salt,
	}
	newU, err := User.InsertOne(ctx, newUser)
	if err != nil {
		return model.User{}, err
	}
	id := newU.InsertedID.(primitive.ObjectID).Hex()
	newUser.ID = id
	return newUser, nil
}
