package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"agricoladb/graph/generated"
	"agricoladb/graph/model"
	"context"
)

// Cards is the resolver for the cards field.
func (r *queryResolver) Cards(ctx context.Context) ([]*model.Card, error) {
	return r.cards, nil
}

// FindCardByRevisionAndLiteralID is the resolver for the findCardByRevisionAndLiteralId field.
func (r *queryResolver) FindCardByRevisionAndLiteralID(ctx context.Context, input model.RevisionAndLiteralID) (*model.Card, error) {
	return r.cards[0], nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
