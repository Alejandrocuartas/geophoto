package photo

import (
	"context"

	"github.com/Alejandrocuartas/geophoto/database/collections"
	"github.com/Alejandrocuartas/geophoto/graph/model"
	"github.com/Alejandrocuartas/geophoto/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetPhotos(ctx context.Context, lat string, long string) ([]*model.Photo, error) {
	var coll *mongo.Collection = collections.PhotoCollection()
	var photos []*model.Photo
	lo, e := helpers.ParseStringToFloat(long)
	if e != nil {
		return []*model.Photo{}, &MyError{message: "Error parsing long into float."}
	}
	la, e := helpers.ParseStringToFloat(lat)
	if e != nil {
		return []*model.Photo{}, &MyError{message: "Error parsing lat into float."}
	}
	filter := bson.D{
		{Key: "location",
			Value: bson.D{
				{Key: "$near", Value: bson.D{
					{Key: "$geometry", Value: &model.Location{
						Type:        "Point",
						Coordinates: []float64{lo, la},
					}},
					{Key: "$maxDistance", Value: 70000},
				}},
			}},
	}

	cur, err := coll.Find(ctx, filter)
	for cur.Next(ctx) {
		var elem *model.Photo
		err := cur.Decode(&elem)
		if err != nil {
			return []*model.Photo{}, err
		}
		photos = append(photos, elem)
	}
	if err != nil {
		return []*model.Photo{}, err
	}
	return photos, nil
}
