package handler

import (
	"fmt"
	"net/http"

	"github.com/dreammnck/plan_retirever/pkg/v1/model"
	"github.com/dreammnck/plan_retirever/pkg/v1/serializer"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
)

func (h *planRetrieverHandler) GetVideoSummaryByCategoryHandler(c echo.Context) error {
	body := new(serializer.GetVideoSummaryByCategory)
	if err := c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, "cannot bind request body")
	}

	if err := body.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("validate request fail with %s", err.Error()))
	}

	content, err := h.planRetrieverSvc.GetVideoSummaryByCategory(c.Request().Context(), body.ID, body.Category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	contentResp := []serializer.VideoSummaryContent{}
	for _, item := range content.Content {
		c := serializer.VideoSummaryContent{
			LocationName: item.LocationName,
			StartTime:    item.StartTime,
			EndTime:      item.EndTime,
			Summary:      item.Summary,
			Lat:          item.Lat,
			PlaceID:      item.PlaceID,
			Lng:          item.Lng,
			Category:     item.Category,
			Photos: lo.Map(item.Photos, func(photo model.Photo, _ int) string {
				return photo.Reference
			}),
		}

		contentResp = append(contentResp, c)
	}

	return c.JSON(http.StatusOK, serializer.VideoSummaryResponse{Content: contentResp, CanGenerateTrip: content.CanGenerateTrip, UserID: content.UserID, VideoID: content.VIdeoID})
}
