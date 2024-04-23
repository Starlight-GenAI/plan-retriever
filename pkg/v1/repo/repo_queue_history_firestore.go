package repo

import (
	"context"
	"errors"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/dreammnck/plan_retirever/pkg/v1/model"
	"google.golang.org/api/iterator"
)

type queueHistoryFirestoreRepo struct {
	collection *firestore.CollectionRef
}

func NewQueueHistoryFirestoreRepo(client *firestore.Client, collectionName string) *queueHistoryFirestoreRepo {
	collection := client.Collection(collectionName)
	return &queueHistoryFirestoreRepo{collection: collection}
}

var _ model.IQueueHistoryFirestoreRepo = new(queueHistoryFirestoreRepo)

func (q *queueHistoryFirestoreRepo) Update(ctx context.Context, id string, status string) error {

	var ref *firestore.DocumentRef
	iter := q.collection.Where("queue_id", "==", id).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return err
		}

		ref = doc.Ref
	}

	if ref == nil {
		return errors.New("ref is nil")
	}

	if _, err := q.collection.Doc(ref.ID).Update(ctx, []firestore.Update{
		{
			Path:  "status",
			Value: status,
		},
		{
			Path:  "updated_at",
			Value: time.Now(),
		},
	}); err != nil {
		return err
	}

	return nil
}

func (q *queueHistoryFirestoreRepo) List(ctx context.Context, userID string) ([]model.QueueHistoryFirestore, error) {

	queueHistories := []model.QueueHistoryFirestore{}
	iter := q.collection.Where("user_id", "==", userID).Documents(ctx)
	for {
		queueHistory := new(model.QueueHistoryFirestore)
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Println(err)
			return nil, err
		}

		if err := doc.DataTo(queueHistory); err != nil {
			return nil, err
		}

		queueHistories = append(queueHistories, *queueHistory)
	}

	return queueHistories, nil
}

func (q *queueHistoryFirestoreRepo) Get(ctx context.Context, id string) (*model.QueueHistoryFirestore, error) {

	result := new(model.QueueHistoryFirestore)
	iter := q.collection.Where("queue_id", "==", id).Documents(ctx)
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
