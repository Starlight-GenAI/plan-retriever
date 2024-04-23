package handler

import (
	"fmt"
	"net/http"

	"github.com/dreammnck/plan_retirever/pkg/v1/serializer"
	"github.com/labstack/echo/v4"
)

func (h *planRetrieverHandler) VideoSummaryHandler(c echo.Context) error {
	body := new(serializer.VideoSummaryRequest)
	if err := c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, "cannot bind request body")
	}

	if err := body.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("validate request fail with %s", err.Error()))
	}

	content, err := h.planRetrieverSvc.VideoSummary(c.Request().Context(), body.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	contentResp := []serializer.VideoSummaryContent{}
	for _, item := range content.Content {
		content := serializer.VideoSummaryContent{
			LocationName: item.LocationName,
			StartTime:    item.StartTime,
			EndTime:      item.EndTime,
			Summary:      item.Summary,
			Lat:          item.Lat,
			PlaceID:      item.PlaceID,
			Lng:          item.Lng,
			Category:     item.Category,
		}
		if len(item.Photos) > 0 {
			content.Photo = item.Photos[0].Reference
		}

		contentResp = append(contentResp, content)
	}

	return c.JSON(http.StatusOK, serializer.VideoSummaryResponse{Content: contentResp, CanGenerateTrip: content.CanGenerateTrip, UserID: content.UserID, VideoID: content.VIdeoID})
}
