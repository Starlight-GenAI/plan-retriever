package core

import (
	"context"
	"fmt"

	"github.com/dreammnck/plan_retirever/pkg/logger"
	"github.com/dreammnck/plan_retirever/pkg/v1/model"
	"github.com/sagikazarmark/slog-shim"
	"github.com/samber/lo"
)

func (p *planRetrieverSvc) ListTripSummary(ctx context.Context, userID string) ([]model.TripSummary, error) {
	logger.Logger.Info("list trip summary", slog.String("tag", "usecase list trip summary"))

	contents, err := p.tripSummaryFirestore.List(ctx, userID)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("list trip summary fail with %s", err.Error()), slog.String("tag", "usecase list trip summary"))
		return nil, err
	}

	tripsSummary := lo.Map(contents, func(val model.TripSummaryFirestore, _ int) model.TripSummary {
		return model.TripSummary{
			Content: val.Content,
			UserID:  val.UserID,
			VideoID: val.VideoID,
		}
	})

	logger.Logger.Info("list trip summary done", slog.String("tag", "usecase list trip summary"))
	return tripsSummary, nil
}
