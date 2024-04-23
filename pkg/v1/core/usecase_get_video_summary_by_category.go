package core

import (
	"context"

	"github.com/dreammnck/plan_retirever/pkg/v1/model"
	"github.com/samber/lo"
)

func (s *planRetrieverSvc) GetVideoSummaryByCategory(ctx context.Context, id string, category string) (*model.VideoSummary, error) {

	videoSummary, err := s.videoSummaryFirestore.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	filteredContent := lo.Filter(videoSummary.Content, func(val model.VideoSummaryContent, _ int) bool {
		return val.Category == category
	})

	return &model.VideoSummary{Content: filteredContent, CanGenerateTrip: videoSummary.CanGenerateTrip, UserID: videoSummary.UserID, VIdeoID: videoSummary.VideoID}, nil
}
