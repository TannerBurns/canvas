package routes

import (
	"net/http"

	"../controllers"
	"../models"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() (*mux.Router, *models.Logger) {
	controller := &controllers.Controller{Name: "API.Controller"}
	controller.Logger = models.NewLogger()
	controller.Session = models.NewSession()
	api_route := "/canvas/api/v1"

	StatusRoutes := Routes{
		Route{
			"Status",
			"GET",
			api_route + "/status",
			controller.Status,
		},
	}

	Routes := []Routes{StatusRoutes}

	router := mux.NewRouter().StrictSlash(true)
	for _, routes := range Routes {
		for _, route := range routes {
			var handler http.Handler
			handler = route.HandlerFunc

			router.
				Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(handler)
		}
	}
	return router, controller.Logger
}
