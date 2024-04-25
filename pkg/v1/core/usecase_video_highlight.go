package core

import (
	"context"
	"fmt"

	"github.com/dreammnck/plan_retirever/pkg/logger"
	"github.com/dreammnck/plan_retirever/pkg/v1/model"
	"github.com/sagikazarmark/slog-shim"
)

func (s *planRetrieverSvc) VideoHighlight(ctx context.Context, id string) (*model.VideoHightlight, error) {
	logger.Logger.Info("get video highligh", slog.String("tag", "usecase video highlight"))

	content, err := s.videoHighlightFirestore.Get(ctx, id)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("get video highlight fail with %s", err.Error()), slog.String("tag", "usecase video highlight"))
		return nil, err
	}

	logger.Logger.Info("get video highligh done", slog.String("tag", "usecase video highlight"))

	return &model.VideoHightlight{Content: content.Content, UserID: content.UserID, QueueID: content.QueueID, VideoID: content.VideoID, ContentSumary: content.ContentSumary}, nil
}
