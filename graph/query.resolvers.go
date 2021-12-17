package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"kanmaa-backend/graph/generated"
	"kanmaa-backend/graph/model"
)

func (r *queryResolver) Search(ctx context.Context, text string) ([]model.SearchResult, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) School(ctx context.Context, email string) (*model.School, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) StartTimeKeeper(ctx context.Context, school string, clearPasses bool) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetQr(ctx context.Context, email string, password string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, email string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context, first *int, after *string) (*model.UsersConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
