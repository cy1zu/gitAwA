package main

import (
	"backend/app/api"
	"backend/app/db/postgres"
	"backend/config"
	"backend/logger"
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// init config
	err := config.Init()
	if err != nil {
		panic(err)
	}
	// init logger
	err = logger.Init()
	if err != nil {
		panic(err)
	}
	defer zap.L().Sync()
	// init database
	err = postgres.Init("release")
	if err != nil {
		panic(err)
	}
	// init router
	r := api.Init()
	//start server
	srv := &http.Server{
		Addr:    fmt.Sprintf("127.0.0.1:%s", config.Conf.Port),
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown", zap.Error(err))
	}
	zap.L().Info("Server exiting")

}
