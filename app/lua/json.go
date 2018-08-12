package lua

import (
	"github.com/clbanning/mxj"
	glua "github.com/tul/gopher-lua"
)

// JSONModule provides access to json methods
type JSONModule struct {
}

// NewJSONModule returns a new JSON module
func NewJSONModule() *Module {
	module := &JSONModule{}
	return &Module{
		Name: "json",
		Data: module,
		Funcs: map[string]glua.LGFunction{
			"marshal": module.MarshalJSON,
		},
	}
}

// MarshalJSON marshals the given lua table into a JSON string
func (j *JSONModule) MarshalJSON(state *glua.LState) int {
	tbl := state.ToTable(1)
	r := mxj.Map(TableToMap(tbl))
	buff, err := r.Json(true)
	if err != nil {
		state.RaiseError("Unable to marshal lua table - %s", err)
		return 0
	}
	state.Push(glua.LString(string(buff)))
	return 1
}
