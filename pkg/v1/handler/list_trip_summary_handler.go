package handler

import (
	"fmt"
	"net/http"

	"github.com/dreammnck/plan_retirever/pkg/v1/model"
	"github.com/dreammnck/plan_retirever/pkg/v1/serializer"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
)

func (h *planRetrieverHandler) ListTripSummary(c echo.Context) error {
	body := new(serializer.ListTripSummaryRequest)
	if err := c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, "cannot bind request body")
	}

	if err := body.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("validate request fail with %s", err.Error()))
	}

	contents, err := h.planRetrieverSvc.ListTripSummary(c.Request().Context(), body.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	resp := []serializer.TripSummaryResponse{}
	for _, val := range contents {
		tripSummary := serializer.TripSummaryResponse{
			UserID:  val.UserID,
			VideoID: val.VideoID,
		}

		tripSummaryContent := []serializer.TripSummaryContent{}
		for _, intVal := range val.Content {
			c := serializer.TripSummaryContent{
				Day: intVal.Day,
			}
			locationWithSummaryList := []serializer.LocationWithSummary{}
			for _, loc := range intVal.LocationWithSummary {

				locationWithSummary := serializer.LocationWithSummary{
					LocationName: loc.LocationName,
					Summary:      loc.Summary,
					PlaceID:      loc.PlaceID,
					Lat:          loc.Lat,
					Lng:          loc.Lng,
					Category:     loc.Category,
					Rating:       loc.Rating,
					Photos: lo.Map(loc.Photos, func(photo model.Photo, _ int) string {
						return photo.Reference
					}),
				}
				locationWithSummaryList = append(locationWithSummaryList, locationWithSummary)

				if loc.HasRecommendedRestaurant {
					recommendedRestaurant := serializer.LocationWithSummary{
						LocationName: loc.RecommendedRestaurant.Name,
						Summary:      loc.RecommendedRestaurant.Summary,
						PlaceID:      loc.RecommendedRestaurant.PlaceID,
						Lat:          loc.RecommendedRestaurant.Lat,
						Lng:          loc.RecommendedRestaurant.Lng,
						Category:     model.RECOMMENDED_DINING,
						Rating:       loc.RecommendedRestaurant.Rating,
						Photos: lo.Map(loc.RecommendedRestaurant.Photos, func(photo model.Photo, _ int) string {
							return photo.Reference
						}),
					}

					c.LocationWithSummary = append(c.LocationWithSummary, recommendedRestaurant)
				}
			}
			c.LocationWithSummary = locationWithSummaryList

			tripSummaryContent = append(tripSummaryContent, c)
		}
		tripSummary.Content = tripSummaryContent
		resp = append(resp, tripSummary)
	}

	return c.JSON(http.StatusOK, resp)
}
