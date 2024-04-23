package adapter

import (
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/pubsub"
	"github.com/dreammnck/plan_retirever/pkg/logger"
	"github.com/dreammnck/plan_retirever/pkg/v1/model"
	"github.com/sagikazarmark/slog-shim"
)

type pubSubAdapter struct {
	queueHistoryRepo model.IQueueHistoryFirestoreRepo
}

func NewPubSubAdapter(queueHistoryRepo model.IQueueHistoryFirestoreRepo) *pubSubAdapter {
	return &pubSubAdapter{queueHistoryRepo: queueHistoryRepo}
}

var _ model.IPubSubAdapter = new(pubSubAdapter)

func (p *pubSubAdapter) Subscribe(ctx context.Context, msg *pubsub.Message) {
	logger.Logger.Info("start consume data", slog.String("tag", "pubsub consumer"))

	notification := new(model.NotificationEventMessage)
	if err := json.Unmarshal(msg.Data, notification); err != nil {
		logger.Logger.Error(fmt.Sprintf("fail to unmarshal data with %s", err.Error()), slog.String("tag", "pubsub consumer"))
		return
	}

	err := p.queueHistoryRepo.Update(ctx, notification.QueueID, notification.Status)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("set status into cache fail with %s", err.Error()), slog.String("tag", "pubsub consumer"))
		return
	}

	msg.Ack()
	logger.Logger.Info("successfully consume data", slog.String("tag", "pubsub consumer"))
}
