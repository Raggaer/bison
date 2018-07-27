package lua

import glua "github.com/yuin/gopher-lua"

// LoaderValue sets values on a module loader func
type LoaderValue struct {
	Name  string
	Value interface{}
}

// MakeLoader returns a module loader func
func MakeLoader(state *glua.LState, values []LoaderValue) func(*glua.LState) int {
	mod := state.SetFuncs(state.NewTable(), nil)
	for _, v := range values {
		data := state.NewUserData()
		data.Value = v.Value
		state.SetField(mod, v.Name, data)
	}
	return func(s *glua.LState) int {
		s.Push(mod)
		return 1
	}
}
