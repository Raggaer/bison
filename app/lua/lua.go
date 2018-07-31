package lua

import (
	"os"
	"path/filepath"
	"strings"

	glua "github.com/tul/gopher-lua"
)

// Module represents a lua module
type Module struct {
	Name   string
	Funcs  map[string]glua.LGFunction
	Values []ModuleValue
	Data   interface{}
}

// ModuleValue data used for module values
type ModuleValue struct {
	Name  string
	Value glua.LValue
}

// LoaderValue sets values on a module loader func
type LoaderValue struct {
	Name  string
	Value interface{}
}

func makeLoader(state *glua.LState, funcs map[string]glua.LGFunction, moduleData interface{}, values []ModuleValue) func(*glua.LState) int {
	mod := state.SetFuncs(state.NewTable(), funcs)
	data := state.NewUserData()
	data.Value = moduleData
	state.SetField(mod, "__data", data)
	for _, v := range values {
		state.SetField(mod, v.Name, v.Value)
	}
	return func(s *glua.LState) int {
		s.Push(mod)
		return 1
	}
}

// CompileFiles compiles all lua files into function protos
func CompileFiles(dir string) (map[string]*glua.FunctionProto, error) {
	files := map[string]*glua.FunctionProto{}
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(info.Name(), ".lua") {
			// Compile lua file
			proto, err := CompileLua(path)
			if err != nil {
				return err
			}
			files[path] = proto
		}
		return nil
	})
	return files, err
}

// NewState creates a new lua state with bison modules
func NewState(modules []*Module) *glua.LState {
	state := glua.NewState()
	for _, module := range modules {
		state.PreloadModule(module.Name, makeLoader(state, module.Funcs, module.Data, module.Values))
	}
	return state
}
