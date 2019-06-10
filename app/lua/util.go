package lua

import (
	"bufio"
	"fmt"
	"math"
	"os"

	glua "github.com/yuin/gopher-lua"
	"github.com/yuin/gopher-lua/parse"
)

// MapToTable converts a Go map into a lua table
func MapToTable(dst map[string]interface{}) *glua.LTable {
	src := &glua.LTable{}
	for k, e := range dst {
		v := GoValueToLua(e)
		src.RawSetString(k, v)
	}
	return src
}

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

// GoValueToLua converts a go value into a lua value
func GoValueToLua(val interface{}) glua.LValue {
	switch v := val.(type) {
	case string:
		return glua.LString(v)
	case bool:
		return glua.LBool(v)
	case int:
		return glua.LNumber(v)
	case int32:
		return glua.LNumber(v)
	case int64:
		return glua.LNumber(v)
	case float64:
		return glua.LNumber(v)
	default:
		return glua.LNil
	}
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
		// gopher-lua converts numbers as float64
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

// CompileLua reads the passed lua file from disk and compiles it.
func CompileLua(filePath string) (*glua.FunctionProto, error) {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(file)
	chunk, err := parse.Parse(reader, filePath)
	if err != nil {
		return nil, err
	}
	proto, err := glua.Compile(chunk, filePath)
	if err != nil {
		return nil, err
	}
	return proto, nil
}

// DoCompiledFile takes a FunctionProto, as returned by CompileLua, and runs it in the LState. It is equivalent
// to calling DoFile on the LState with the original source file
func DoCompiledFile(state *glua.LState, proto *glua.FunctionProto) error {
	lfunc := state.NewFunctionFromProto(proto)
	state.Push(lfunc)
	return state.PCall(0, glua.MultRet, nil)
}
