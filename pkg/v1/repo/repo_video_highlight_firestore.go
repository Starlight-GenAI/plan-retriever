package repo

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/dreammnck/plan_retirever/pkg/v1/model"
	"google.golang.org/api/iterator"
)

type videoHightlightFirebaseRepo struct {
	collection *firestore.CollectionRef
}

func NewVideoHighlightFirebaseRepo(client *firestore.Client, collectionName string) *videoHightlightFirebaseRepo {
	collection := client.Collection(collectionName)
	return &videoHightlightFirebaseRepo{collection: collection}
}

var _ model.IVideoHighlightFirestoreRepo = new(videoHightlightFirebaseRepo)

func (v *videoHightlightFirebaseRepo) Get(ctx context.Context, id string) (*model.VideoHighlightFirestore, error) {
	iter := v.collection.Where("queue_id", "==", id).Limit(1).Documents(ctx)
	result := new(model.VideoHighlightFirestore)

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
