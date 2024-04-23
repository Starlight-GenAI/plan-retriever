package repo

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/dreammnck/plan_retirever/pkg/v1/model"
	"google.golang.org/api/iterator"
)

type videoSummaryFirestoreRepo struct {
	collection *firestore.CollectionRef
}

func NewVideoSummaryFirestoreRepo(client *firestore.Client, collectionName string) *videoSummaryFirestoreRepo {
	collection := client.Collection(collectionName)
	return &videoSummaryFirestoreRepo{collection: collection}
}

var _ model.IVideoSummaryFirestoreRepo = new(videoSummaryFirestoreRepo)

func (v *videoSummaryFirestoreRepo) Get(ctx context.Context, id string) (*model.VideoSummaryFireStore, error) {
	iter := v.collection.Where("queue_id", "==", id).Limit(1).Documents(ctx)
	result := new(model.VideoSummaryFireStore)

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

func (v *videoSummaryFirestoreRepo) List(ctx context.Context, userID string) ([]model.VideoSummaryFireStore, error) {
	iter := v.collection.Where("user_id", "==", userID).Documents(ctx)
	videoSummaryList := []model.VideoSummaryFireStore{}

	for {
		videoSummary := new(model.VideoSummaryFireStore)
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Println(err)
			return nil, err
		}

		if err := doc.DataTo(videoSummary); err != nil {
			return nil, err
		}

		videoSummaryList = append(videoSummaryList, *videoSummary)
	}

	return videoSummaryList, nil
}
