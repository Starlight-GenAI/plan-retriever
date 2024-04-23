package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/pubsub"
	"github.com/dreammnck/plan_retirever/config"
	routes "github.com/dreammnck/plan_retirever/pkg"
	"github.com/dreammnck/plan_retirever/pkg/logger"
	"github.com/labstack/echo/v4"
	"github.com/sagikazarmark/slog-shim"
	"go.uber.org/zap"
	"google.golang.org/api/option"
)

func main() {

	logger.Init()
	configs := config.InitConfig()
	ctx := context.Background()

	fireStoreClient, err := firestore.NewClientWithDatabase(ctx, configs.Firestore.ProjectID, configs.Firestore.Database, option.WithCredentialsFile(configs.Firestore.CredentialFilePath))
	if err != nil {
		panic(err)
	}
	defer fireStoreClient.Close()

	pubsubClient, err := pubsub.NewClient(ctx, configs.PubSub.ProjectID, option.WithCredentialsFile(configs.PubSub.CredentialFilePath))
	if err != nil {
		panic(err)
	}

	defer pubsubClient.Close()

	apiRouter := routes.NewRouter(configs, fireStoreClient, pubsubClient).RegisterRouter()
	go RunServer(apiRouter, configs.Server.Port)

	Shutdown(apiRouter)
}

func RunServer(router *echo.Echo, port int) {
	startPort := fmt.Sprintf(":%d", port)
	router.Logger.Fatal(router.Start(startPort))
}

func Shutdown(router *echo.Echo) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	if err := router.Shutdown(context.Background()); err != nil {
		slog.Error(err.Error(), zap.String("tag", "shutdown Server"))
	}
}
