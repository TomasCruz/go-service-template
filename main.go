package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/TomasCruz/go-service-template/presenter"

	"github.com/TomasCruz/go-service-template/config"

	"github.com/TomasCruz/go-service-template/environment"
	"github.com/TomasCruz/go-service-template/service"
)

func setupFromEnvVars(conf *config.Config) {
	conf.Port = environment.ReadAndCheckEnvVar("TEMPLATE_SVC_PORT")
}

func main() {
	var conf config.Config

	// populate configuration
	setupFromEnvVars(&conf)

	// graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, os.Kill)

	// create and bindRoutes
	routes := newRoutes()
	bindRoutes(routes)

	// not necessary in this case
	/*
		// pass config
		presenter.SetConfig(config)
		service.SetConfig(config)
	*/

	// pass routes
	presenter.SetRoutes(routes)

	// fire up the server
	var httpServer *http.Server
	go func() {
		httpServer = &http.Server{Addr: fmt.Sprintf(":%s", conf.Port)}

		log.Printf("starting web server")
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to start HTTP server: %s", err)
		}
	}()

	<-stop
	gracefulShutdown(httpServer)
}

func gracefulShutdown(httpServer *http.Server) {
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	httpServer.Shutdown(shutdownCtx)
	service.Shutdown()
}
