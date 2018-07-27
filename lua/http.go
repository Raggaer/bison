package lua

import (
	glua "github.com/tul/gopher-lua"
	"github.com/valyala/fasthttp"
)

// HTTPModule module for all http actions
type HTTPModule struct {
	RequestContext *fasthttp.RequestCtx
	RequestParams  map[string]string
}

// NewHTTPModule returns a new http module
func NewHTTPModule(ctx *fasthttp.RequestCtx, params map[string]string) *Module {
	module := &HTTPModule{
		RequestContext: ctx,
		RequestParams:  params,
	}
	return &Module{
		Name: "http",
		Data: module,
		Funcs: map[string]glua.LGFunction{
			"redirect": module.Redirect,
			"param":    module.GetParam,
		},
	}
}

// GetParam retrieves a request param
func (h *HTTPModule) GetParam(state *glua.LState) int {
	v, ok := h.RequestParams[state.ToString(1)]
	if !ok {
		state.Push(glua.LNil)
		return 1
	}
	state.Push(glua.LString(v))
	return 1
}

// Redirect redirects the user to the given location
func (h *HTTPModule) Redirect(state *glua.LState) int {
	redirectURL := string(state.ToString(1))
	redirectCode := int(state.ToNumber(2))
	h.RequestContext.Redirect(redirectURL, redirectCode)
	return 0
}
