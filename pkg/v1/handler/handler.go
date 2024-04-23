package handler

import "github.com/dreammnck/plan_retirever/pkg/v1/model"

type planRetrieverHandler struct {
	planRetrieverSvc model.IPlanRetireverSvc
}

func NewPlanRetrieverHandler(planRetrieverSvc model.IPlanRetireverSvc) *planRetrieverHandler {
	return &planRetrieverHandler{planRetrieverSvc: planRetrieverSvc}
}
