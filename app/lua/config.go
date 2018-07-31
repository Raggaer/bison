package lua

import glua "github.com/tul/gopher-lua"

// ConfigModule provides access to config values
type ConfigModule struct {
	Config map[string]interface{}
}

// NewConfigModule returns a new config module
func NewConfigModule(data map[string]interface{}) *Module {
	module := &ConfigModule{
		Config: data,
	}
	return &Module{
		Name: "config",
		Data: module,
		Funcs: map[string]glua.LGFunction{
			"get": module.Get,
		},
	}
}

// Get retrieves a config value and pushes into stack
func (c *ConfigModule) Get(state *glua.LState) int {
	v, ok := c.Config[state.ToString(1)]
	if !ok {
		state.Push(glua.LNil)
		return 1
	}
	state.Push(GoValueToLua(v))
	return 1
}
