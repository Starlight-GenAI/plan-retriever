package core

import (
	"context"
	"fmt"

	"log/slog"

	"github.com/dreammnck/plan_retirever/pkg/logger"

	"github.com/dreammnck/plan_retirever/pkg/v1/model"
)

func (p *planRetrieverSvc) TripSummary(ctx context.Context, id string) (*model.TripSummary, error) {
	logger.Logger.Info("get trip summary", slog.String("tag", "usecase trip summary"))

	content, err := p.tripSummaryFirestore.Get(ctx, id)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("get trip summary fail with %s", err.Error()), slog.String("tag", "usecase trip summary"))
		return nil, err
	}

	uniqueLocation := map[string]string{}
	uniqueTripContents := []model.TripSummaryContent{}

	for _, val := range content.Content {
		locationWithSummary := []model.LocationWithSummary{}
		for _, intVal := range val.LocationWithSummary {
			key := fmt.Sprintf("%v%v", intVal.Lat, intVal.Lng)
			if _, ok := uniqueLocation[key]; !ok {
				locationWithSummary = append(locationWithSummary, intVal)
				uniqueLocation[key] = intVal.LocationName
			}
		}
		uniqueTripContents = append(uniqueTripContents, model.TripSummaryContent{Day: val.Day, LocationWithSummary: locationWithSummary})
	}

	logger.Logger.Info("successfully get trip summary", slog.String("tag", "usecase trip summary"))
	return &model.TripSummary{Content: uniqueTripContents, UserID: content.UserID, VideoID: content.VideoID}, nil
}
