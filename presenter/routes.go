package presenter

import (
	"github.com/TomasCruz/go-service-template/config/routes"
)

var rts routes.Routes

// SetRoutes sets the routes from main
func SetRoutes(r routes.Routes) {
	rts = r
}
