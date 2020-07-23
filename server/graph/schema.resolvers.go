package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/CKjapan/go-graphql-tutorial/server/graph/generated"
	"github.com/CKjapan/go-graphql-tutorial/server/graph/model"
	"github.com/CKjapan/go-graphql-tutorial/server/models"
	"github.com/CKjapan/go-graphql-tutorial/server/util"
)

func (r *linkResolver) User(ctx context.Context, obj *models.Link) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*models.Link, error) {

	link := &models.Link{
		ID:      util.CreateUniqueID(),
		Title:   input.Title,
		Address: input.Address,
	}
	res := r.DB.Create(link)
	if err := res.Error; err != nil {
		return nil, err
	}
	return link, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Links(ctx context.Context) ([]*models.Link, error) {
	panic(fmt.Errorf("not implemented"))
}

// Link returns generated.LinkResolver implementation.
func (r *Resolver) Link() generated.LinkResolver { return &linkResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type linkResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
