package photo

import (
	"context"
	"testing"

	"github.com/Alejandrocuartas/geophoto/graph/model"
)

const expectedPhoto = &model.Photo{
	ID:  "123",
	URL: "http://",
	Location: &model.Location{
		Type: "Point",
		Coordinates: []float64{
			-74.08175,
			4.60971,
		},
	},
	User: &model.User{
		ID:       "123",
		Username: "ale31jo",								
		Password: "jwt",
	},
},

func TestGetPhotos(t *testing.T) {
	tables := struct {
		lat            string
		long           string
		mockFunc       func()
		expectedPhotos []model.Photo
	}{
		lat:  "4.60971",
		long: "-74.08175",
		mockFunc: func() {
			GetPhotos = func(ctx *context.Context, lat string, long string) ([]*model.Photo, error) {
				return []*model.Photo{
					expectedPhoto,
				}, nil
			}
		},
		expectedPhotos: []*model.Photo{
			expectedPhoto,
		},
	}
	
	originalGetPhotos := GetPhotos

	tables.mockFunc()

	photosList := GetPhotos(context.Background(), "lat", "lon")
	if photosList[0].ID != "123" {
		t.Errorf("Error getting id")
	}
	GetPhotos = originalGetPhotos
}