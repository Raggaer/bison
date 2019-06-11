package template

import (
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/Raggaer/bison/app/config"
	"github.com/Raggaer/bison/app/lua"
	"github.com/fasthttp-contrib/sessions"
	cache "github.com/patrickmn/go-cache"
	"github.com/valyala/fasthttp"
	glua "github.com/yuin/gopher-lua"
)

// TemplateFuncData data needed for template functions
type TemplateFuncData struct {
	Cache           *cache.Cache
	Config          *config.Config
	ControllersPath string
	Files           map[string]*glua.FunctionProto
}

// LoadTemplates load the given view directory
func LoadTemplates(dir string, data *TemplateFuncData) (*template.Template, error) {
	tpl := template.New("bison")
	tpl.Funcs(templateFuncMap(data))
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(info.Name(), ".html") {
			if _, err := tpl.ParseFiles(path); err != nil {
				return err
			}
		}
		return nil
	})
	return tpl, err
}

func templateFuncMap(h *TemplateFuncData) template.FuncMap {
	return map[string]interface{}{
		"execute": func(file string, data map[string]interface{}) template.HTML {
			// Load values from the map
			_, ok := data["_RequestContext"].(*fasthttp.RequestCtx)
			if !ok {
				return ""
			}
			session, ok := data["_Session"].(sessions.Session)
			if !ok {
				return ""
			}

			proto, ok := h.Files[filepath.Join(h.ControllersPath, file)]
			if !ok {
				return ""
			}

			// Create state with basic bison modules
			state := lua.NewState([]*lua.Module{
				lua.NewConfigModule(h.Config.Custom),
				lua.NewURLModule(),
				lua.NewCacheModule(h.Cache),
				lua.NewSessionModule(session),
				lua.NewJSONModule(),
				lua.NewEnvironmentModule(),
			})
			defer state.Close()

			// Execute compiled state and return top value as html text
			if err := lua.DoCompiledFile(state, proto); err != nil {
				return ""
			}
			executeData := state.Get(-1)
			if executeData.Type() == glua.LTString {
				return template.HTML(string(executeData.(glua.LString)))
			}
			return ""
		},
	}
}
