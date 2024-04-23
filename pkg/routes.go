package routes

import (
	"context"
	"fmt"
	"net/http"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/pubsub"
	"github.com/dreammnck/plan_retirever/config"
	"github.com/dreammnck/plan_retirever/pkg/logger"
	"github.com/dreammnck/plan_retirever/pkg/v1/adapter"
	"github.com/dreammnck/plan_retirever/pkg/v1/core"
	"github.com/dreammnck/plan_retirever/pkg/v1/handler"
	"github.com/dreammnck/plan_retirever/pkg/v1/repo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sagikazarmark/slog-shim"
)

type apiRouter struct {
	config          *config.Config
	firestoreClient *firestore.Client
	pubsubClient    *pubsub.Client
}

func NewRouter(config *config.Config, firestoreClient *firestore.Client, pubsubClient *pubsub.Client) *apiRouter {
	return &apiRouter{
		config:          config,
		firestoreClient: firestoreClient,
		pubsubClient:    pubsubClient,
	}
}

func (a *apiRouter) RegisterRouter() *echo.Echo {
	router := newEcho()
	ctx := context.Background()
	videoSummaryFirestre := repo.NewVideoSummaryFirestoreRepo(a.firestoreClient, a.config.Firestore.VideoSummaryCollection)
	tripSummaryFirestore := repo.NewTripSummaryFirestoreRepo(a.firestoreClient, a.config.Firestore.TripSummaryCollection)
	videoHighlightFirestore := repo.NewVideoHighlightFirebaseRepo(a.firestoreClient, a.config.Firestore.VideoHighlightCollection)
	queueHistoryFirestore := repo.NewQueueHistoryFirestoreRepo(a.firestoreClient, a.config.Firestore.QueueHistoryCollection)
	planRetrieverSvc := core.NewPlanRetrieverSvc(core.NewPlanRetriverSvcCfgs{
		VideoSummaryFirestore:   videoSummaryFirestre,
		TripSummaryFirestore:    tripSummaryFirestore,
		VideoHighlightFirestore: videoHighlightFirestore,
		QueueHistoryFirestore:   queueHistoryFirestore,
	})
	h := handler.NewPlanRetrieverHandler(planRetrieverSvc)
	pubsubAdapter := adapter.NewPubSubAdapter(queueHistoryFirestore)

	router.POST("/plan-status", h.PlanStatusHandler)
	router.POST("/get-video-summary", h.VideoSummaryHandler)
	router.POST("/get-trip-summary", h.TripSummaryHandler)
	router.POST("/get-video-summary-by-category", h.GetVideoSummaryByCategoryHandler)
	router.POST("/list-video-summary", h.ListVideoSummary)
	router.POST("/list-trip-summary", h.ListTripSummary)
	router.POST("/get-video-highlight", h.VideoHighlightHandler)
	router.POST("/list-queue-history", h.ListQueueIDHandler)
	router.GET("/health", func(e echo.Context) error {
		return e.String(http.StatusOK, "OK")
	})

	// pub/sub
	sub := a.pubsubClient.Subscription(a.config.PubSub.SubscriptionID)
	go func() {
		err := sub.Receive(ctx, pubsubAdapter.Subscribe)
		if err != nil {
			logger.Logger.Error(fmt.Sprintf("consume error with %s", err.Error()), slog.String("tag", "consume event"))
		}
	}()

	return router
}

func newEcho() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Use(middleware.Secure())
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		HSTSMaxAge:            3600,
		ContentSecurityPolicy: "default-src 'self'",
		HSTSExcludeSubdomains: true,
	}))
	e.Use(middleware.CORS())

	return e
}
