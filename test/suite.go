package test

import (
	"io"
	"log"
	"net"
	"net/http"
	"path/filepath"
	"testing"
	"time"

	"github.com/buaazp/fasthttprouter"
	"github.com/patrickmn/go-cache"
	"github.com/Raggaer/bison/app/config"
	"github.com/Raggaer/bison/app/controllers"
	"github.com/Raggaer/bison/app/lua"
	"github.com/Raggaer/bison/app/router"
	"github.com/Raggaer/bison/app/template"
	"github.com/valyala/fasthttp"
)

func createTestServer(p chan<- int, t *testing.T) io.Closer {
	t.Parallel()
	// Load config file
	config, err := config.LoadConfig(filepath.Join("config", "config.lua"))
	if err != nil {
		log.Fatal(err)
	}
	config.TestMode = true

	// Compile all lua files
	files, err := lua.CompileFiles(filepath.Join("controllers"))
	if err != nil {
		log.Fatal(err)
	}

	// Load all templates
	tpl, err := template.LoadTemplates(filepath.Join("views"), &template.TemplateFuncData{
		Config: config,
		Files:  files,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create fasthttp router
	r := fasthttprouter.New()
	routes, err := router.LoadRoutes(filepath.Join("router", "router.lua"))
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

	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("cannot start tcp server on port %d: %s", 0, err)
	}
	p <- ln.Addr().(*net.TCPAddr).Port
	go fasthttp.Serve(ln, r.Handler)
	return ln
}
