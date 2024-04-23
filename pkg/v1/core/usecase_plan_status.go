package core

import (
	"context"
	"fmt"

	"github.com/dreammnck/plan_retirever/pkg/logger"
	"github.com/samber/lo"

	"github.com/sagikazarmark/slog-shim"
)

const (
	notInQueue = "not_in_queue"
)

func (p *planRetrieverSvc) PlanStatus(ctx context.Context, id string) (*string, error) {
	logger.Logger.Info("get plan status", slog.String("tag", "usecase plan status"))

	queueHistory, err := p.queueHistoryFirestore.Get(ctx, id)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("get plan status fail with %s", err.Error()), slog.String("tag", "usecase plan status"))
		return nil, err
	}

	if queueHistory.Status == "" {
		return lo.ToPtr(notInQueue), nil
	}

	logger.Logger.Info("successfully get plan status", slog.String("tag", "usecase plan status"))
	return &queueHistory.Status, nil
}
