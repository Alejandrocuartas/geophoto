package photo

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

func SavePhoto(ctx context.Context, input model.NewPhoto) (model.Photo, error) {
	var Photo *mongo.Collection = collections.PhotoCollection()
	var User *mongo.Collection = collections.UserCollection()
	//validate user exists
	objectID, er := primitive.ObjectIDFromHex(input.UserID)
	if er != nil {
		return model.Photo{}, &MyError{message: "Error parsing id into hex for Mongo. Probably no Mongo Id."}
	}
	filter := bson.M{"_id": objectID}
	userExist := User.FindOne(ctx, filter)
	var existing model.User
	err := userExist.Decode(&existing)
	if err == mongo.ErrNoDocuments {
		return model.Photo{}, &MyError{message: "User does not exist."}
	}
	//save user
	lo, e := helpers.ParseStringToFloat(input.Long)
	if e != nil {
		return model.Photo{}, &MyError{message: "Error parsing long into float."}
	}
	la, e := helpers.ParseStringToFloat(input.Lat)
	if e != nil {
		return model.Photo{}, &MyError{message: "Error parsing lat into float."}
	}
	coords := []float64{lo, la}
	newLocation := &model.Location{
		Type:        "Point",
		Coordinates: coords,
	}
	newPhoto := model.Photo{
		Location: newLocation,
		URL:      input.URL,
		User:     &existing,
	}
	newP, err := Photo.InsertOne(ctx, newPhoto)
	if err != nil {
		return model.Photo{}, err
	}
	id := newP.InsertedID.(primitive.ObjectID).Hex()
	newPhoto.ID = id
	return newPhoto, nil
}
