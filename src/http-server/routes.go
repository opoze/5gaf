package httpserver

import (
	"net/http"

	"github.com/opoze/5gaf/http-server/controllers"
)

// Define routes here
func RegisterRoutes(httpRouter *http.ServeMux) {
	RegisterHTTPRoute("/notify", httpRouter, HttpJsonPostHandlerFactory(controllers.HandleNefNotification))
}
