package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/Raggaer/bison/app/config"
	"github.com/Raggaer/bison/app/controllers"
	"github.com/Raggaer/bison/app/lua"
	"github.com/Raggaer/bison/app/router"
	"github.com/Raggaer/bison/app/template"
	"github.com/buaazp/fasthttprouter"
	cache "github.com/patrickmn/go-cache"

	"github.com/valyala/fasthttp"
)

func main() {
	var controllersPath string
	flag.StringVar(&controllersPath, "controllers", "", "Filepath for your controllers folder. Default 'app/controllers'")
	var configPath string
	flag.StringVar(&configPath, "config", "", "Filepath for your config file. Default 'app/config/config.lua'")
	var viewsPath string
	flag.StringVar(&viewsPath, "views", "", "Filepath for your views folder. Default 'app/views'")
	var routerPath string
	flag.StringVar(&routerPath, "router", "", "Filepath for your router file. Default 'app/router/router.lua'")

	flag.Parse()

	// Load config file
	if configPath == "" {
		configPath = filepath.Join("app", "config", "config.lua")
	}
	config, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	// Compile all lua files
	if controllersPath == "" {
		controllersPath = filepath.Join("app", "controllers")
	}
	files, err := lua.CompileFiles(controllersPath)
	if err != nil {
		log.Fatal(err)
	}

	// Load all templates
	if viewsPath == "" {
		viewsPath = filepath.Join("app", "views")
	}
	tpl, err := template.LoadTemplates(viewsPath, &template.TemplateFuncData{
		Config: config,
		Files:  files,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create fasthttp router
	if routerPath == "" {
		routerPath = filepath.Join("app", "router", "router.lua")
	}
	r := fasthttprouter.New()
	routes, err := router.LoadRoutes(routerPath)
	if err != nil {
		log.Fatal(err)
	}

	// Create fasthttp server
	handler := &controllers.Handler{
		Config: config,
		Routes: routes,
		Files:  files,
		Tpl:    tpl,
		Cache:  cache.New(time.Minute*5, time.Minute*10),
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
		fmt.Println("Running development mode - bison listening on address '" + config.Address + "'")
	} else {
		fmt.Println("bison listening on address '" + config.Address + "'")
	}
	fasthttp.ListenAndServe(config.Address, r.Handler)
}
