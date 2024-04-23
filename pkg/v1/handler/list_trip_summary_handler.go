package handler

import (
	"fmt"
	"net/http"

	"github.com/dreammnck/plan_retirever/pkg/v1/serializer"
	"github.com/labstack/echo/v4"
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

				locationWithSummaryList = append(locationWithSummaryList, locationWithSummary)
			}
			c.LocationWithSummary = locationWithSummaryList

			tripSummaryContent = append(tripSummaryContent, c)
		}
		tripSummary.Content = tripSummaryContent
		resp = append(resp, tripSummary)
	}

	// resp := lo.Map(contents, func(val model.TripSummary, _ int) serializer.TripSummaryResponse {
	// 	return serializer.TripSummaryResponse{
	// 		Content: lo.Map(val.Content, func(intVal model.TripSummaryContent, _ int) serializer.TripSummaryContent {
	// 			return serializer.TripSummaryContent{
	// 				Day: intVal.Day,
	// 				LocationWithSummary: lo.Map(intVal.LocationWithSummary, func(loc model.LocationWithSummary, _ int) serializer.LocationWithSummary {
	// 					return serializer.LocationWithSummary{
	// 						LocationName:             loc.LocationName,
	// 						Summary:                  loc.Summary,
	// 						PlaceID:                  loc.PlaceID,
	// 						Lat:                      loc.Lat,
	// 						Lng:                      loc.Lng,
	// 						Category:                 loc.Category,
	// 						Rating:                   loc.Rating,
	// 						HasRecommendedRestaurant: loc.HasRecommendedRestaurant,
	// 						RecommendedRestaurant: serializer.RestaurantDetail{
	// 							Name:    loc.RecommendedRestaurant.Name,
	// 							Summary: loc.RecommendedRestaurant.Summary,
	// 							Rating:  loc.RecommendedRestaurant.Rating,
	// 							Lat:     loc.RecommendedRestaurant.Lat,
	// 							Lng:     loc.RecommendedRestaurant.Lng,
	// 							Photos: lo.Map(loc.RecommendedRestaurant.Photos, func(photo model.Photo, _ int) serializer.Photo {
	// 								return serializer.Photo{
	// 									Reference: photo.Reference,
	// 									MaxWidth:  photo.MaxWidth,
	// 									MaxHeight: photo.MaxHeight,
	// 								}
	// 							}),
	// 						},
	// 						Photos: lo.Map(loc.Photos, func(photo model.Photo, _ int) serializer.Photo {
	// 							return serializer.Photo{
	// 								Reference: photo.Reference,
	// 								MaxWidth:  photo.MaxWidth,
	// 								MaxHeight: photo.MaxHeight,
	// 							}
	// 						}),
	// 					}
	// 				}),
	// 			}
	// 		}),
	// 		UserID:  val.UserID,
	// 		VideoID: val.VideoID,
	// 	}
	// })

	return c.JSON(http.StatusOK, resp)
}
