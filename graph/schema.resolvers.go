package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"main/Structures"
	"main/db"
	"main/graph/generated"
	"main/graph/model"
)

func (r *mutationResolver) CreateRecord(ctx context.Context, input model.NewRecord) (*model.Record, error) {
	newRecord := Structures.Record{
		ID:        input.ID,
		Title:     input.Title,
		Content:   input.Content,
		Timestamp: input.Timestamp,
		Views:     input.Views,
	}

	db.AddRecord(newRecord)

	record := &model.Record{
		ID:        input.ID,
		Title:     input.Title,
		Content:   input.Content,
		Timestamp: input.Timestamp,
		Views:     input.Views,
	}

	return record, nil
}

func (r *queryResolver) Records(ctx context.Context) ([]*model.Record, error) {
	return nil, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
