package lua

import (
	"github.com/fasthttp-contrib/sessions"
	glua "github.com/yuin/gopher-lua"
)

// SessionModule defines a request session module
type SessionModule struct {
	Session sessions.Session
}

// NewSessionModule returns a new session module
func NewSessionModule(sess sessions.Session) *Module {
	module := &SessionModule{
		Session: sess,
	}
	return &Module{
		Name: "session",
		Data: module,
		Funcs: map[string]glua.LGFunction{
			"set": module.Set,
			"get": module.Get,
		},
	}
}

// Set sets the given session value
func (s *SessionModule) Set(state *glua.LState) int {
	s.Session.Set(state.ToString(1), LuaValueToGo(state.Get(2)))
	return 0
}

// Get retrieves the given session value
func (s *SessionModule) Get(state *glua.LState) int {
	v := GoValueToLua(s.Session.Get(state.ToString(1)))
	state.Push(v)
	return 1
}
