package actionresult

import (
	"context"
	"net/http"
)

type ActionContext struct {
	context.Context
	http.ResponseWriter
}

type ActionResult interface {
	Execute(*ActionContext) error
}
