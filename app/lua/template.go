package lua

import (
	"html/template"

	glua "github.com/tul/gopher-lua"
	"github.com/valyala/fasthttp"
)

// TemplateModule defines a html template module
type TemplateModule struct {
	Tpl            *template.Template
	RequestContext *fasthttp.RequestCtx
}

// NewTemplateModule returns a new template module
func NewTemplateModule(tpl *template.Template, ctx *fasthttp.RequestCtx) *Module {
	module := &TemplateModule{
		Tpl:            tpl,
		RequestContext: ctx,
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
	if err := t.Tpl.ExecuteTemplate(t.RequestContext, name, TableToMap(data)); err != nil {
		state.RaiseError("Unale to render template %s - %s", name, err)
	}
	return 0
}
