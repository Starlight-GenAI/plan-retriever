package core

import (
	"context"

	"github.com/dreammnck/plan_retirever/pkg/logger"
	"github.com/dreammnck/plan_retirever/pkg/v1/model"
	"github.com/sagikazarmark/slog-shim"
	"github.com/samber/lo"
)

func (p *planRetrieverSvc) ListQueueHistory(ctx context.Context, userID string) ([]model.QueueHistory, error) {
	logger.Logger.Info("list queue history", slog.String("tag", "usecase list queue history"))

	contents, err := p.queueHistoryFirestore.List(ctx, userID)
	if err != nil {
		return nil, err
	}

	queueHistories := lo.Map(contents, func(val model.QueueHistoryFirestore, _ int) model.QueueHistory {
		return model.QueueHistory{
			ID:            val.ID,
			VideoUrl:      val.VideoUrl,
			VideoID:       val.VideoID,
			Status:        val.Status,
			Title:         val.Title,
			Description:   val.Description,
			Thumbnails:    val.Thumbnails,
			CreatedAt:     val.CreatedAt,
			UpdatedAt:     val.UpdatedAt,
			ChannelName:   val.ChannelName,
			IsUseSubTitle: val.IsUseSubTitle,
		}
	})
	logger.Logger.Info("list queue history done", slog.String("tag", "usecase list queue history"))
	return queueHistories, nil
}
