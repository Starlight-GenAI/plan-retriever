package handler

import (
	"fmt"
	"net/http"

	"github.com/dreammnck/plan_retirever/pkg/v1/model"
	"github.com/dreammnck/plan_retirever/pkg/v1/serializer"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
)

func (h *planRetrieverHandler) VideoHighlightHandler(c echo.Context) error {
	body := new(serializer.VideoHighlightRequest)
	if err := c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, "cannot bind request body")
	}

	if err := body.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("validate request fail with %s", err.Error()))
	}

	content, err := h.planRetrieverSvc.VideoHighlight(c.Request().Context(), body.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	contentResp := lo.Map[model.VideoHighlightContent](content.Content, func(val model.VideoHighlightContent, _ int) serializer.VideoHighlightContent {
		return serializer.VideoHighlightContent{
			HighlightName:   val.HighlightName,
			HighlightDetail: val.HighlightDetail,
		}
	})
	return c.JSON(http.StatusOK, serializer.VideoHightlightResponse{
		Content: contentResp,
		UserID:  content.UserID,
		QueueID: content.QueueID,
		VideoID: content.VideoID,
	})
}
