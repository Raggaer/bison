package lua

import glua "github.com/tul/gopher-lua"
import "net/url"

// URLModule module for some url manipulation actions
type URLModule struct {
}

// NewURLModule returns a new url module
func NewURLModule() *Module {
	module := &URLModule{}
	return &Module{
		Name: "url",
		Data: module,
		Funcs: map[string]glua.LGFunction{
			"pathEscape":    module.PathEscape,
			"pathUnescape":  module.PathUnescape,
			"queryEscape":   module.QueryEscape,
			"queryUnescape": module.QueryUnescape,
		},
	}
}

// QueryEscape escapes the string so it can be safely placed inside a URL path segment
func (u *URLModule) QueryEscape(state *glua.LState) int {
	state.Push(glua.LString(url.QueryEscape(state.ToString(1))))
	return 1
}

// QueryUnescape inverse process of query escape
func (u *URLModule) QueryUnescape(state *glua.LState) int {
	p, err := url.QueryUnescape(state.ToString(1))
	if err != nil {
		state.RaiseError("Unable to unescape query %s - %v", state.ToString(1), err)
	}
	state.Push(glua.LString(p))
	return 1
}

// PathEscape escapes the string so it can be safely placed inside a URL path segment
func (u *URLModule) PathEscape(state *glua.LState) int {
	state.Push(glua.LString(url.PathEscape(state.ToString(1))))
	return 1
}

// PathUnescape inverse process of path escape
func (u *URLModule) PathUnescape(state *glua.LState) int {
	p, err := url.PathUnescape(state.ToString(1))
	if err != nil {
		state.RaiseError("Unable to unescape path %s - %v", state.ToString(1), err)
	}
	state.Push(glua.LString(p))
	return 1
}
