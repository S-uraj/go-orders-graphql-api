package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.27

import (
	"context"
	"fmt"

	"github.com/S-uraj/go-orders-graphql-api/graph/model"
)

// CreatePost is the resolver for the CreatePost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
	panic(fmt.Errorf("not implemented: CreatePost - CreatePost"))
}

// UpdatePost is the resolver for the UpdatePost field.
func (r *mutationResolver) UpdatePost(ctx context.Context, postID int, input *model.NewPost) (*model.Post, error) {
	panic(fmt.Errorf("not implemented: UpdatePost - UpdatePost"))
}

// GetAllPosts is the resolver for the GetAllPosts field.
func (r *queryResolver) GetAllPosts(ctx context.Context) ([]*model.Post, error) {
	panic(fmt.Errorf("not implemented: GetAllPosts - GetAllPosts"))
}

// GetOnePost is the resolver for the GetOnePost field.
func (r *queryResolver) GetOnePost(ctx context.Context, id int) (*model.Post, error) {
	panic(fmt.Errorf("not implemented: GetOnePost - GetOnePost"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
