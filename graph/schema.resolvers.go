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

func (r *mutationResolver) DeleteRecord(ctx context.Context, input model.DeleteRecord) (*model.DeletedRecord, error) {
	db.DeleteRecord(input.ID)
	deletedrecord := &model.DeletedRecord{
		ID: input.ID,
	}
	return deletedrecord, nil
}

func (r *queryResolver) Records(ctx context.Context) ([]*model.Record, error) {
	//r.records = r.records[:0] // clear the array
	var tarRecords []*model.Record

	records := db.GetRecords() // Get all raws from table
	for _, origRecord := range records {
		//var destRecord model.Record
		destRecord := model.Record{ // convert all structures to
			ID:        origRecord.ID,
			Title:     origRecord.Title,
			Content:   origRecord.Content,
			Views:     origRecord.Views,
			Timestamp: origRecord.Timestamp,
		}
		tarRecords = append(tarRecords, &destRecord)
	}
	return tarRecords, nil
}

func (r *queryResolver) Record(ctx context.Context, input model.GetRecord) (*model.Record, error) {
	origRecord := db.GetRecord(input.ID)
	var foundRecord model.Record
	foundRecord = model.Record{
		ID:        origRecord.ID,
		Title:     origRecord.Title,
		Content:   origRecord.Content,
		Views:     origRecord.Views,
		Timestamp: origRecord.Timestamp,
	}

	return &foundRecord, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
