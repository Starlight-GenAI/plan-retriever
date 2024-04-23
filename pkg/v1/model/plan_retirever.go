package model

import (
	"context"

	"cloud.google.com/go/pubsub"
)

const (
	Success = "success"
	Pending = "pending"
)

type IPlanRetireverSvc interface {
	PlanStatus(ctx context.Context, id string) (*string, error)
	TripSummary(ctx context.Context, id string) (*TripSummary, error)
	VideoSummary(ctx context.Context, id string) (*VideoSummary, error)
	VideoHighlight(ctx context.Context, id string) (*VideoHightlight, error)
	ListTripSummary(ctx context.Context, userID string) ([]TripSummary, error)
	ListVideoSummary(ctx context.Context, userID string) ([]VideoSummary, error)
	ListQueueHistory(ctx context.Context, userID string) ([]QueueHistory, error)
	GetVideoSummaryByCategory(ctx context.Context, id string, category string) (*VideoSummary, error)
}

type ICacheRepo interface {
	Get(ctx context.Context, id string) (*string, error)
	Set(ctx context.Context, id, status string) error
}

type ITripSummaryFirestoreRepo interface {
	Get(ctx context.Context, id string) (*TripSummaryFirestore, error)
	List(ctx context.Context, userID string) ([]TripSummaryFirestore, error)
}

type IVideoSummaryFirestoreRepo interface {
	Get(ctx context.Context, id string) (*VideoSummaryFireStore, error)
	List(ctx context.Context, userID string) ([]VideoSummaryFireStore, error)
}

type IVideoHighlightFirestoreRepo interface {
	Get(ctx context.Context, id string) (*VideoHighlightFirestore, error)
}

type IQueueHistoryFirestoreRepo interface {
	Get(ctx context.Context, id string) (*QueueHistoryFirestore, error)
	List(ctx context.Context, userID string) ([]QueueHistoryFirestore, error)
	Update(ctx context.Context, id string, status string) error
}

type IPubSubAdapter interface {
	Subscribe(ctx context.Context, msg *pubsub.Message)
}
