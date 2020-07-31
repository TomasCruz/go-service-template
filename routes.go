package main

import (
	"net/http"

	"github.com/TomasCruz/go-service-template/config/routes"

	"github.com/TomasCruz/go-service-template/presenter"
)

func newRoutes() routes.Routes {
	return routes.Routes{
		HealthRoute: "/health/",
		HelloRoute:  "/hello/",
	}
}

func bindRoutes(routesStruct routes.Routes) {
	// health
	healthHandlerFunc := presenter.HealthHandler
	http.HandleFunc(routesStruct.HealthRoute, healthHandlerFunc)

	// hello
	helloHandlerFunc := presenter.HelloHandler
	http.HandleFunc(routesStruct.HelloRoute, helloHandlerFunc)
}
