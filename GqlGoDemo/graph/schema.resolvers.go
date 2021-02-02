package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"GqlGoDemo/graph/generated"
	"GqlGoDemo/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) CreateNewUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	//fmt.Println(input)
	u := &model.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}
	r.user = append(r.user, u)
	return u, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUser) (*model.User, error) {

	for _, i := range r.user {

		if input.OldFirstname == i.FirstName {
			i.FirstName = input.NewFirstname
			i.LastName = input.NewLastname
			fmt.Println("", i)
			return i, nil
		}
	}

	return &model.User{}, nil
}

func (r *mutationResolver) UpdateFirstName(ctx context.Context, input *model.Fname) (*model.User, error) {
	for _, i := range r.user {
		if input.OldFirstname == i.FirstName {
			i.FirstName = input.FirstName
			return i, nil
		}

	}
	return &model.User{}, nil
}

func (r *queryResolver) User(ctx context.Context) ([]*model.User, error) {
	return r.user, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
