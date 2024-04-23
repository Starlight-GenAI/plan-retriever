package handler

import (
	"fmt"
	"net/http"

	"github.com/dreammnck/plan_retirever/pkg/v1/serializer"
	"github.com/labstack/echo/v4"
)

func (h *planRetrieverHandler) ListVideoSummary(c echo.Context) error {
	body := new(serializer.ListVideoSummaryRequest)
	if err := c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, "cannot bind request body")
	}
	if err := body.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("validate request fail with %s", err.Error()))
	}

	contents, err := h.planRetrieverSvc.ListVideoSummary(c.Request().Context(), body.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	resp := []serializer.VideoSummaryResponse{}
	for _, val := range contents {
		videoSummaryContent := []serializer.VideoSummaryContent{}
		content := serializer.VideoSummaryResponse{
			CanGenerateTrip: val.CanGenerateTrip,
			UserID:          val.UserID,
			VideoID:         val.VIdeoID,
		}
		for _, intVal := range val.Content {
			c := serializer.VideoSummaryContent{
				LocationName: intVal.LocationName,
				StartTime:    intVal.StartTime,
				EndTime:      intVal.EndTime,
				Summary:      intVal.Summary,
				PlaceID:      intVal.PlaceID,
				Lat:          intVal.Lat,
				Lng:          intVal.Lng,
			}
			if len(intVal.Photos) > 0 {
				c.Photo = intVal.Photos[0].Reference
			}

			videoSummaryContent = append(videoSummaryContent, c)
		}
		content.Content = videoSummaryContent

		resp = append(resp, content)
	}

	// resp := lo.Map(contents, func(val model.VideoSummary, _ int) serializer.VideoSummaryResponse {
	// 	return serializer.VideoSummaryResponse{
	// 		CanGenerateTrip: val.CanGenerateTrip,
	// 		UserID:          val.UserID,
	// 		VideoID:         val.VIdeoID,
	// 		Content: lo.Map(val.Content, func(intVal model.VideoSummaryContent, _ int) serializer.VideoSummaryContent {
	// 			return serializer.VideoSummaryContent{
	// 				LocationName: intVal.LocationName,
	// 				StartTime:    intVal.StartTime,
	// 				EndTime:      intVal.EndTime,
	// 				Summary:      intVal.Summary,
	// 				PlaceID:      intVal.PlaceID,
	// 				Lat:          intVal.Lat,
	// 				Lng:          intVal.Lng,
	// 				Photo:        intVal.Photos[0].Reference,
	// 			}
	// 		}),
	// 	}
	// })

	return c.JSON(http.StatusOK, resp)
}
