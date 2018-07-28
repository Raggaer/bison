package main

import (
	"fmt"
	"html/template"
	"log"
	"path/filepath"

	"github.com/raggaer/bison/lua"
	glua "github.com/tul/gopher-lua"
	"github.com/valyala/fasthttp"
)

// Handler main fasthttp handler
type Handler struct {
	Config *Config
	Routes []*Route
	Files  map[string]*glua.FunctionProto
	Tpl    *template.Template
}

// MainRoute handles all http requests
func (h *Handler) MainRoute(ctx *fasthttp.RequestCtx) {
	// If we are running under development mode reload stuff
	if h.Config.DevMode {
		tpl, err := loadTemplates()
		if err != nil {
			ctx.Error("Unable to reload templates", 500)
			return
		}
		h.Tpl = tpl
		routes, err := loadRoutes()
		if err != nil {
			ctx.Error("Unable to reload routes", 500)
			return
		}
		h.Routes = routes
		luaFiles, err := lua.CompileFiles("controllers")
		if err != nil {
			ctx.Error("Unable to reload controllers", 500)
			return
		}
		h.Files = luaFiles
	}

	// Retrieve current route
	params := map[string]string{}
	ctx.VisitUserValues(func(b []byte, i interface{}) {
		params[string(b)] = fmt.Sprint(i)
	})
	route := retrieveCurrentRoute(params, string(ctx.Method()), string(ctx.Path()), h.Routes)

	// Retrieve compiled file for this route
	proto, ok := h.Files[filepath.Join("controllers", route.File)]
	if !ok {
		ctx.NotFound()
		return
	}

	// Create state with bison modules
	state := lua.NewState([]*lua.Module{
		lua.NewHTTPModule(ctx, params, lua.ModuleValue{
			Name:  "method",
			Value: glua.LString(string(ctx.Method())),
		}),
		lua.NewConfigModule(h.Config.Custom),
		lua.NewTemplateModule(h.Tpl, ctx),
	})
	defer state.Close()

	// Execute compiled state
	if err := lua.DoCompiledFile(state, proto); err != nil {
		log.Println(err)
		ctx.Error("Unable to execute "+route.Path, 500)
		return
	}
}
