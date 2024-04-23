package core

import (
	"context"
	"fmt"

	"log/slog"

	"github.com/dreammnck/plan_retirever/pkg/logger"
	"github.com/dreammnck/plan_retirever/pkg/v1/model"
)

func (p *planRetrieverSvc) VideoSummary(ctx context.Context, id string) (*model.VideoSummary, error) {
	logger.Logger.Info("get video summary", slog.String("tag", "usecase video summary"))

	content, err := p.videoSummaryFirestore.Get(ctx, id)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("get video summary fail with %s", err.Error()), slog.String("tag", "usecase video summary"))
		return nil, err
	}

	uniqueLocation := map[string]string{}
	filterContents := []model.VideoSummaryContent{}

	for _, val := range content.Content {
		key := fmt.Sprintf("%v%v", val.Lat, val.Lng)
		if _, ok := uniqueLocation[key]; !ok {
			uniqueLocation[key] = val.LocationName
			filterContents = append(filterContents, val)
		}
	}

	logger.Logger.Info("successfully get video summary", slog.String("tag", "usecase video summary"))
	return &model.VideoSummary{
		Content:         filterContents,
		CanGenerateTrip: content.CanGenerateTrip,
		UserID:          content.UserID,
		VIdeoID:         content.VideoID,
	}, nil
}
