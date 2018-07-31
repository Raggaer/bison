package main

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/buaazp/fasthttprouter"
	"github.com/raggaer/bison/app/config"
	"github.com/raggaer/bison/app/controllers"
	"github.com/raggaer/bison/app/lua"
	"github.com/raggaer/bison/app/router"
	"github.com/raggaer/bison/app/template"

	"github.com/valyala/fasthttp"
)

func main() {
	// Load config file
	config, err := config.LoadConfig(filepath.Join("app", "config", "config.lua"))
	if err != nil {
		log.Fatal(err)
	}

	// Compile all lua files
	files, err := lua.CompileFiles(filepath.Join("app", "controllers"))
	if err != nil {
		log.Fatal(err)
	}

	// Load all templates
	tpl, err := template.LoadTemplates(&template.TemplateFuncData{
		Config: config,
		Files:  files,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create fasthttp router
	r := fasthttprouter.New()
	routes, err := router.LoadRoutes(filepath.Join("app", "router", "router.lua"))
	if err != nil {
		log.Fatal(err)
	}

	// Create fasthttp server
	handler := &controllers.Handler{
		Config: config,
		Routes: routes,
		Files:  files,
		Tpl:    tpl,
	}

	for _, rx := range routes {
		if rx.Method == http.MethodGet {
			r.GET(rx.Path, handler.MainRoute)
		}
		if rx.Method == http.MethodPost {
			r.POST(rx.Path, handler.MainRoute)
		}
	}
	if config.DevMode {
		log.Println("Running development mode - bison listening on address '" + config.Address + "'")
	} else {
		log.Println("bison listening on address '" + config.Address + "'")
	}
	fasthttp.ListenAndServe(config.Address, r.Handler)
}
