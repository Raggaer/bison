package lua

import (
	"fmt"

	glua "github.com/tul/gopher-lua"
	"github.com/valyala/fasthttp"
)

// HTTPModule module for all http actions
type HTTPModule struct {
	RequestContext *fasthttp.RequestCtx
	RequestParams  map[string]string
}

// NewHTTPModule returns a new http module
func NewHTTPModule(ctx *fasthttp.RequestCtx, params map[string]string, values ...ModuleValue) *Module {
	module := &HTTPModule{
		RequestContext: ctx,
		RequestParams:  params,
	}
	return &Module{
		Name:   "http",
		Data:   module,
		Values: values,
		Funcs: map[string]glua.LGFunction{
			"write":         module.Write,
			"method":        module.GetRequestMethod,
			"uri":           module.GetRelativeURL,
			"remoteAddress": module.GetRemoteAddress,
			"getCookie":     module.GetCookie,
			"setCookie":     module.SetCookie,
			"redirect":      module.Redirect,
			"param":         module.GetParam,
		},
	}
}

// Write writes data to the request response writer
func (h *HTTPModule) Write(state *glua.LState) int {
	if _, err := h.RequestContext.Write([]byte(state.ToString(1))); err != nil {
		state.RaiseError("Unable to write data into request response - %s", err)
	}
	return 0
}

// GetRequestMethod retrieves the request method
func (h *HTTPModule) GetRequestMethod(state *glua.LState) int {
	state.Push(glua.LString(string(h.RequestContext.Method())))
	return 1
}

// GetRelativeURL retrieves the request uri
func (h *HTTPModule) GetRelativeURL(state *glua.LState) int {
	state.Push(glua.LString(string(h.RequestContext.RequestURI())))
	return 1
}

// GetRemoteAddress retrieves the request remote address
func (h *HTTPModule) GetRemoteAddress(state *glua.LState) int {
	state.Push(glua.LString(h.RequestContext.RemoteIP().String()))
	return 1
}

// SetCookie sets a HTTP cookie
func (h *HTTPModule) SetCookie(state *glua.LState) int {
	var cookie fasthttp.Cookie
	cookie.SetKey(state.ToString(1))
	cookie.SetValue(fmt.Sprint(LuaValueToGo(state.Get(2))))
	cookie.SetHTTPOnly(true)
	h.RequestContext.Response.Header.SetCookie(&cookie)
	return 0
}

// GetCookie retrieves a HTTP cookie
func (h *HTTPModule) GetCookie(state *glua.LState) int {
	n := state.ToString(1)
	f := false
	h.RequestContext.Request.Header.VisitAllCookie(func(k, v []byte) {
		if f {
			return
		}
		if string(k) == n {
			state.Push(glua.LString(string(v)))
			f = true
		}
	})
	if f {
		return 1
	}
	return 0
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
