package core

import "github.com/dreammnck/plan_retirever/pkg/v1/model"

type planRetrieverSvc struct {
	videoSummaryFirestore   model.IVideoSummaryFirestoreRepo
	tripSummaryFirestore    model.ITripSummaryFirestoreRepo
	videoHighlightFirestore model.IVideoHighlightFirestoreRepo
	queueHistoryFirestore   model.IQueueHistoryFirestoreRepo
}

type NewPlanRetriverSvcCfgs struct {
	Cache                   model.ICacheRepo
	VideoSummaryFirestore   model.IVideoSummaryFirestoreRepo
	TripSummaryFirestore    model.ITripSummaryFirestoreRepo
	VideoHighlightFirestore model.IVideoHighlightFirestoreRepo
	QueueHistoryFirestore   model.IQueueHistoryFirestoreRepo
}

func NewPlanRetrieverSvc(cfgs NewPlanRetriverSvcCfgs) *planRetrieverSvc {
	return &planRetrieverSvc{
		videoSummaryFirestore:   cfgs.VideoSummaryFirestore,
		tripSummaryFirestore:    cfgs.TripSummaryFirestore,
		videoHighlightFirestore: cfgs.VideoHighlightFirestore,
		queueHistoryFirestore:   cfgs.QueueHistoryFirestore,
	}
}

var _ model.IPlanRetireverSvc = new(planRetrieverSvc)
