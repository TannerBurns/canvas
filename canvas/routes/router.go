package routes

import (
	"log"
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

func NewRouter() (*mux.Router, *controllers.Controller) {
	lc, err := models.NewConfig("./config/settings.conf")
	if err != nil {
		log.Fatal(err)
	}
	controller := &controllers.Controller{
		Name: "API.Controller",
	}
	controller.Logger = models.NewLogger(lc.Config["api"]["name"])
	controller.Session, _ = models.NewSession(lc)

	Routes := []Routes{
		StatusRoutes(controller),
	}

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
	return router, controller
}
