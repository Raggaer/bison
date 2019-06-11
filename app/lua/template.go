package lua

import (
	"html/template"

	"github.com/fasthttp-contrib/sessions"
	"github.com/valyala/fasthttp"
	glua "github.com/yuin/gopher-lua"
)

// TemplateModule defines a html template module
type TemplateModule struct {
	Tpl            *template.Template
	RequestContext *fasthttp.RequestCtx
	Session        sessions.Session
}

// NewTemplateModule returns a new template module
func NewTemplateModule(tpl *template.Template, ctx *fasthttp.RequestCtx, session sessions.Session) *Module {
	module := &TemplateModule{
		Tpl:            tpl,
		RequestContext: ctx,
		Session:        session,
	}
	return &Module{
		Name: "template",
		Data: module,
		Funcs: map[string]glua.LGFunction{
			"render": module.Render,
		},
	}
}

// Render renders the given template with the passed lua table
func (t *TemplateModule) Render(state *glua.LState) int {
	name := state.ToString(1)
	data := state.ToTable(2)
	tbl := TableToMap(data)

	// Add custom needed fields for the execute function
	tbl["_RequestContext"] = t.RequestContext
	tbl["_Session"] = t.Session

	if err := t.Tpl.ExecuteTemplate(t.RequestContext, name, tbl); err != nil {
		state.RaiseError("Unale to render template %s - %s", name, err)
	}
	return 0
}
