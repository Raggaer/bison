package lua

import (
	"os"

	glua "github.com/yuin/gopher-lua"
)

// EnvironmentModule module for set and get env variables
type EnvironmentModule struct {
}

// NewEnvironmentModule returns a new env module
func NewEnvironmentModule() *Module {
	module := &EnvironmentModule{}
	return &Module{
		Name: "env",
		Data: module,
		Funcs: map[string]glua.LGFunction{
			"set": module.Set,
			"get": module.Get,
		},
	}
}

// Set defines a env variable
func (e *EnvironmentModule) Set(state *glua.LState) int {
	key := state.ToString(1)
	val := state.ToString(2)
	if err := os.Setenv(key, val); err != nil {
		state.RaiseError("Unable to set environment variable %s with value %s", key, val)
	}
	return 0
}

// Get retrieves a env variable
func (e *EnvironmentModule) Get(state *glua.LState) int {
	key := state.ToString(1)
	val := os.Getenv(key)
	state.Push(glua.LString(val))
	return 1
}
