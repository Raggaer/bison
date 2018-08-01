package lua

import (
	"time"

	"github.com/patrickmn/go-cache"
	glua "github.com/tul/gopher-lua"
)

// CacheModule provices access to the memory cache object
type CacheModule struct {
	Cache *cache.Cache
}

// NewCacheModule returns a new cache module
func NewCacheModule(c *cache.Cache) *Module {
	module := &CacheModule{
		Cache: c,
	}
	return &Module{
		Name: "cache",
		Data: module,
		Funcs: map[string]glua.LGFunction{
			"get": module.Get,
			"set": module.Set,
		},
	}
}

// Get retrieves a cache value
func (c *CacheModule) Get(state *glua.LState) int {
	key := state.ToString(1)
	v, ok := c.Cache.Get(key)
	if !ok {
		state.Push(glua.LNil)
		return 1
	}
	state.Push(GoValueToLua(v))
	return 1
}

// Set sets a cache value
func (c *CacheModule) Set(state *glua.LState) int {
	key := state.ToString(1)
	val := LuaValueToGo(state.Get(2))
	dur := state.Get(3)
	if dur.Type() == glua.LTString {
		d, err := time.ParseDuration(state.ToString(3))
		if err != nil {
			state.RaiseError("Unable to parse cache duration - %v", err)
			return 0
		}
		c.Cache.Set(key, val, d)
		return 0
	}
	c.Cache.Set(key, val, time.Minute*5)
	return 0
}
