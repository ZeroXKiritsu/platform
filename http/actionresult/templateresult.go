package actionresult

import "platform/templates"

func NewTemplateAction(name string, data any) ActionResult {
	return &TemplateActionResult{templateName: name, data: data}
}

type TemplateActionResult struct {
	templateName string
	data         any
	templates.TemplateExecutor
	templates.InvokeHandlerFunc
}

func (action *TemplateActionResult) Execute(ctx *ActionContext) error {
	return action.TemplateExecutor.ExecTemplateWithFunc(ctx.ResponseWriter, action.templateName, action.data, action.InvokeHandlerFunc)
}
