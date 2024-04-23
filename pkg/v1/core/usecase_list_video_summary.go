package core

import (
	"context"

	"github.com/dreammnck/plan_retirever/pkg/logger"
	"github.com/dreammnck/plan_retirever/pkg/v1/model"
	"github.com/sagikazarmark/slog-shim"
	"github.com/samber/lo"
)

func (p *planRetrieverSvc) ListVideoSummary(ctx context.Context, userID string) ([]model.VideoSummary, error) {
	logger.Logger.Info("list video summary", slog.String("tag", "usecase list video summary"))

	contents, err := p.videoSummaryFirestore.List(ctx, userID)
	if err != nil {
		return nil, err
	}

	videoSummary := lo.Map(contents, func(val model.VideoSummaryFireStore, _ int) model.VideoSummary {
		return model.VideoSummary{
			Content:         val.Content,
			CanGenerateTrip: val.CanGenerateTrip,
			UserID:          val.UserID,
			VIdeoID:         val.VideoID,
		}
	})

	logger.Logger.Info("list video summary done", slog.String("tag", "usecase list video summary"))
	return videoSummary, nil
}
