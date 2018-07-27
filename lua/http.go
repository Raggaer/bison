package lua

import (
	glua "github.com/tul/gopher-lua"
	"github.com/valyala/fasthttp"
)

// HTTPModule module for all http actions
type HTTPModule struct {
	RequestContext *fasthttp.RequestCtx
}

// NewHTTPModule returns a new http module
func NewHTTPModule(ctx *fasthttp.RequestCtx) *Module {
	module := &HTTPModule{
		RequestContext: ctx,
	}
	return &Module{
		Name: "http",
		Data: module,
		Funcs: map[string]glua.LGFunction{
			"redirect": module.Redirect,
		},
	}
}

// Redirect redirects the user to the given location
func (h *HTTPModule) Redirect(state *glua.LState) int {
	redirectURL := string(state.ToString(1))
	redirectCode := int(state.ToNumber(2))
	h.RequestContext.Redirect(redirectURL, redirectCode)
	return 0
}
