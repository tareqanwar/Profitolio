package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/tareqanwar/profitolio/graph/generated"
	"github.com/tareqanwar/profitolio/graph/model"
)

func (r *mutationResolver) AddAssetToPortfolio(ctx context.Context, input model.AddAssetToPortfolioInput) (*model.AddAssetToPortfolioPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddPortfolioTransaction(ctx context.Context, input model.AddPortfolioTransactionInput) (*model.AddPortfolioTransactionPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreatePortfolio(ctx context.Context, input model.CreatePortfolioInput) (*model.CreatePortfolioPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) FetchPortfolio(ctx context.Context, id string) (*model.Portfolio, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) FetchPortfolios(ctx context.Context) ([]*model.Portfolio, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
