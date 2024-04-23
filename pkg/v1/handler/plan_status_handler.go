package handler

import (
	"fmt"
	"net/http"

	"github.com/dreammnck/plan_retirever/pkg/v1/serializer"
	"github.com/labstack/echo/v4"
)

func (h *planRetrieverHandler) PlanStatusHandler(c echo.Context) error {

	body := new(serializer.PlanStatusRequest)
	if err := c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, "cannot bind request body")
	}

	if err := body.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("validate request fail with %s", err.Error()))
	}

	status, err := h.planRetrieverSvc.PlanStatus(c.Request().Context(), body.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, serializer.PlanStatusResponse{Status: *status})
}
