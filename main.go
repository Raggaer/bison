package main

import (
	"log"
	"net/http"

	"github.com/buaazp/fasthttprouter"

	"github.com/valyala/fasthttp"
)

func main() {
	// Load config file
	config, err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Create fasthttp router
	router := fasthttprouter.New()
	routes, err := loadRoutes()
	if err != nil {
		log.Fatal(err)
	}

	// Create fasthttp server
	handler := &Handler{
		Config: config,
		Routes: routes,
	}

	for _, r := range routes {
		if r.Method == http.MethodGet {
			router.GET(r.Path, handler.MainRoute)
		}
		if r.Method == http.MethodPost {
			router.POST(r.Path, handler.MainRoute)
		}
	}
	fasthttp.ListenAndServe(config.Address, router.Handler)
}
