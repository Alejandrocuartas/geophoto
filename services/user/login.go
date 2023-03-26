package user

import (
	"context"
	"fmt"

	"github.com/Alejandrocuartas/geophoto/database/collections"
	"github.com/Alejandrocuartas/geophoto/graph/model"
	"github.com/Alejandrocuartas/geophoto/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type myerror struct {
	message string
}

func (e *myerror) Error() string {
	return fmt.Sprintf("my error occurred: %s", e.message)
}

func Login(ctx context.Context, username string, password string) (model.User, error) {
	var User *mongo.Collection = collections.UserCollection()
	//validate username exists
	filter := bson.M{"username": username}
	userExist := User.FindOne(ctx, filter)
	var existing model.User
	err := userExist.Decode(&existing)
	if err == mongo.ErrNoDocuments {
		return model.User{}, &myerror{message: "Username does not exist."}
	}
	//validate pass
	correctPass, e := helpers.ValidatePassword(password, existing.Password)
	if e != nil {
		return model.User{}, e
	}
	if !correctPass {
		return model.User{}, &myerror{message: "Incorrect password."}
	}
	return existing, nil
}
