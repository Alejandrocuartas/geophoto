package user

import (
	"context"

	"github.com/Alejandrocuartas/geophoto/database/collections"
	"github.com/Alejandrocuartas/geophoto/graph/model"
	"github.com/Alejandrocuartas/geophoto/helpers"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func SaveUser(ctx context.Context, username string, password string) (model.User, error) {
	salt, e := helpers.EncryptPassword(password)
	if e != nil {
		return model.User{}, e
	}
	newUser := model.User{
		Username: username,
		Password: salt,
	}
	var User *mongo.Collection = collections.UserCollection()
	newU, err := User.InsertOne(ctx, newUser)
	id := newU.InsertedID.(primitive.ObjectID).Hex()
	newUser.ID = id
	if err != nil {
		return model.User{}, err
	}
	return newUser, nil
}
