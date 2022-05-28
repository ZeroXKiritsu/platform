package templates

import "io"

type TemplateExecutor interface {
	ExecTemplate(writer io.Writer, name string, data any) (err error)
	ExecTemplateWithFunc(writer io.Writer, name string, data any, handlerFunc InvokeHandlerFunc) (err error)
}

type InvokeHandlerFunc func(handlerName string, methodName string, arg ...any) any
