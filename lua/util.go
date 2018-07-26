package lua

import (
	"fmt"
	"math"

	glua "github.com/yuin/gopher-lua"
)

// TableToMap converts a lua table into a go map[string]interface
func TableToMap(src *glua.LTable) map[string]interface{} {
	if src == nil {
		return map[string]interface{}{}
	}
	dst := map[string]interface{}{}
	src.ForEach(func(i glua.LValue, e glua.LValue) {
		dst[fmt.Sprint(LuaValueToGo(i))] = LuaValueToGo(e)
	})
	return dst
}

// LuaValueToGo converts a lua value into a go value
func LuaValueToGo(dst glua.LValue) interface{} {
	switch v := dst.(type) {
	case *glua.LNilType:
		return nil
	case glua.LBool:
		return bool(v)
	case glua.LString:
		return string(v)
	case glua.LNumber:
		// gopher-lua converts numbers are float64
		// we check if the number needs to be a float64
		if isNumberFloat(v) {
			return float64(v)
		}
		return int64(v)
	case *glua.LTable:
		// tables can be map-like or arrays {1, 2, ...}
		// map-like are converted to a map[string]interface
		// arrays are converted to []interface
		if v.MaxN() == 0 {
			// Table is map-like
			ret := map[string]interface{}{}
			v.ForEach(func(i glua.LValue, e glua.LValue) {
				ret[fmt.Sprint(LuaValueToGo(i))] = LuaValueToGo(e)
			})
			return ret
		}

		ret := make([]interface{}, 0, v.MaxN())
		for i := 0; i <= v.MaxN(); i++ {
			ret = append(ret, v.RawGetInt(i))
		}
		return ret
	default:
		return nil
	}
}

func isNumberFloat(v glua.LValue) bool {
	vf := float64(v.(glua.LNumber))
	return vf != math.Trunc(vf)
}

// IsValueTable checks if the lua value is a table
func IsValueTable(v glua.LValue) bool {
	return v.Type() == glua.LTTable
}
