package handler

import (
	"fmt"
	"net/http"

	"github.com/dreammnck/plan_retirever/pkg/v1/model"
	"github.com/dreammnck/plan_retirever/pkg/v1/serializer"
	"github.com/labstack/echo/v4"
)

func (h *planRetrieverHandler) TripSummaryHandler(c echo.Context) error {
	body := new(serializer.TripSummaryRequest)
	if err := c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, "cannot bind request body")
	}

	if err := body.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("validate request fail with %s", err.Error()))
	}

	content, err := h.planRetrieverSvc.TripSummary(c.Request().Context(), body.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	contentResp := []serializer.TripSummaryContent{}
	for _, item := range content.Content {
		locationWithSummaryList := []serializer.LocationWithSummary{}
		countDining := 0
		c := serializer.TripSummaryContent{
			Day: item.Day,
		}

		for _, loc := range item.LocationWithSummary {
			locationWithSummary := serializer.LocationWithSummary{
				LocationName:             loc.LocationName,
				Summary:                  loc.Summary,
				PlaceID:                  loc.PlaceID,
				Lat:                      loc.Lat,
				Lng:                      loc.Lng,
				Category:                 loc.Category,
				Rating:                   loc.Rating,
				HasRecommendedRestaurant: loc.HasRecommendedRestaurant,
				RecommendedRestaurant: serializer.RestaurantDetail{
					Name:    loc.RecommendedRestaurant.Name,
					Summary: loc.RecommendedRestaurant.Summary,
					Rating:  loc.RecommendedRestaurant.Rating,
					Lat:     loc.RecommendedRestaurant.Lat,
					Lng:     loc.RecommendedRestaurant.Lng,
				},
			}

			if len(loc.RecommendedRestaurant.Photos) > 0 {
				locationWithSummary.RecommendedRestaurant.Photo = loc.RecommendedRestaurant.Photos[0].Reference
			}

			if len(loc.Photos) > 0 {
				locationWithSummary.Photo = loc.Photos[0].Reference
			}

			if loc.Category == model.DINING {
				countDining += 1
			}

			locationWithSummaryList = append(locationWithSummaryList, locationWithSummary)
		}
		c.CountDining = countDining
		c.LocationWithSummary = locationWithSummaryList
		contentResp = append(contentResp, c)
	}

	return c.JSON(http.StatusOK, serializer.TripSummaryResponse{Content: contentResp, UserID: content.UserID, VideoID: content.VideoID})
}
