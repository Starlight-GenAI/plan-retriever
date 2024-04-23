package handler

import (
	"fmt"
	"net/http"
	"sort"

	"github.com/dreammnck/plan_retirever/pkg/v1/model"
	"github.com/dreammnck/plan_retirever/pkg/v1/serializer"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
)

func (h *planRetrieverHandler) ListQueueIDHandler(c echo.Context) error {

	body := new(serializer.ListQueueHistoryRequest)
	if err := c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, "cannot bind request body")
	}

	if err := body.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("validate request fail with %s", err.Error()))
	}

	contents, err := h.planRetrieverSvc.ListQueueHistory(c.Request().Context(), body.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	resp := lo.Map(contents, func(val model.QueueHistory, _ int) serializer.QueueHistory {
		return serializer.QueueHistory{
			ID:            val.ID,
			VideoUrl:      val.VideoUrl,
			VideoID:       val.VideoID,
			Status:        val.Status,
			Title:         val.Title,
			Description:   val.Description,
			Thumbnails:    val.Thumbnails,
			CreatedAt:     val.CreatedAt,
			UpdatedAt:     val.UpdatedAt,
			ChannelName:   val.ChannelName,
			IsUseSubTitle: val.IsUseSubTitle,
		}
	})

	sort.Slice(resp, func(i, j int) bool {
		return resp[i].CreatedAt.After(resp[j].CreatedAt)
	})

	return c.JSON(http.StatusOK, serializer.ListQueueHistoryResponse{Data: resp})
}
