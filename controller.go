package main

import (
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

// Handler main fasthttp handler
type Handler struct {
	Config *Config
	Routes []*Route
}

// MainRoute handles all http requests
func (h *Handler) MainRoute(ctx *fasthttp.RequestCtx) {
	// Retrieve current route
	params := map[string]string{}
	ctx.VisitUserValues(func(b []byte, i interface{}) {
		params[string(b)] = fmt.Sprint(i)
	})
	route := retrieveCurrentRoute(params, string(ctx.Method()), string(ctx.Path()), h.Routes)
	log.Println(route)
}
