package repo

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/dreammnck/plan_retirever/pkg/v1/model"
	"google.golang.org/api/iterator"
)

type tripSummaryFirestoreRepo struct {
	collection *firestore.CollectionRef
}

func NewTripSummaryFirestoreRepo(client *firestore.Client, collectionName string) *tripSummaryFirestoreRepo {
	collection := client.Collection(collectionName)
	return &tripSummaryFirestoreRepo{collection: collection}
}

var _ model.ITripSummaryFirestoreRepo = new(tripSummaryFirestoreRepo)

func (t *tripSummaryFirestoreRepo) Get(ctx context.Context, id string) (*model.TripSummaryFirestore, error) {

	iter := t.collection.Where("queue_id", "==", id).Limit(1).Documents(ctx)
	result := new(model.TripSummaryFirestore)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Println(err)
			return nil, err
		}

		if err := doc.DataTo(result); err != nil {
			return nil, err
		}

	}
	return result, nil
}

func (t *tripSummaryFirestoreRepo) List(ctx context.Context, userID string) ([]model.TripSummaryFirestore, error) {

	iter := t.collection.Where("user_id", "==", userID).Documents(ctx)
	tripSummaryList := []model.TripSummaryFirestore{}

	for {
		tripSummary := new(model.TripSummaryFirestore)
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Println(err)
			return nil, err
		}

		if err := doc.DataTo(tripSummary); err != nil {
			return nil, err
		}

		tripSummaryList = append(tripSummaryList, *tripSummary)
	}

	return tripSummaryList, nil
}
