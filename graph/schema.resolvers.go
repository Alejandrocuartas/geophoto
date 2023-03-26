package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"

	"github.com/Alejandrocuartas/geophoto/graph/model"
	"github.com/Alejandrocuartas/geophoto/helpers"
	"github.com/Alejandrocuartas/geophoto/services/user"
)

// NewUser is the resolver for the newUser field.
func (r *mutationResolver) NewUser(ctx context.Context, password string, username string) (*model.UserRegistration, error) {
	user, err := user.SaveUser(ctx, username, password)
	if err != nil {
		return nil, err
	}
	jwt, err := helpers.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}
	return &model.UserRegistration{
		ID:       user.ID,
		Username: user.Username,
		Jwt:      jwt,
	}, nil
}

// NewPhoto is the resolver for the newPhoto field.
func (r *mutationResolver) NewPhoto(ctx context.Context, input model.NewPhoto) (*model.Photo, error) {
	//create a new photo with Lon, Lat, User, URL and ID keys and return it
	photo := model.Photo{
		ID:   "id",
		URL:  "url",
		User: &model.User{ID: "user"},
		Long: "lon",
		Lat:  "lat",
	}
	return &photo, nil
}

// Login is the resolver for the login field.
func (r *queryResolver) Login(ctx context.Context, password string, username string) (*model.UserRegistration, error) {
	u, e := user.Login(ctx, username, password)
	if e != nil {
		return nil, e
	}
	jwt, err := helpers.GenerateToken(u.ID)
	if err != nil {
		return nil, err
	}
	return &model.UserRegistration{
		ID:       u.ID,
		Username: u.Username,
		Jwt:      jwt,
	}, nil
}

// Photos is the resolver for the photos field.
func (r *queryResolver) Photos(ctx context.Context, lat string, long string) ([]*model.Photo, error) {
	// Create a list of photos with lat, long, url, id and User and return them
	photos := []*model.Photo{
		{
			ID:   "photo1",
			URL:  "http://www.google.com",
			User: &model.User{ID: "user1"},
			Lat:  lat,
			Long: long,
		},
		{
			ID:   "photo2",
			URL:  "http://www.google.com",
			User: &model.User{ID: "user2"},
			Lat:  lat,
			Long: long,
		},
	}
	return photos, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
