package actionresult

func NewErrorAction(err error) ActionResult {
	return &ErrorActionResult{err}
}

type ErrorActionResult struct {
	err error
}

func (action *ErrorActionResult) Execute(ctx *ActionContext) error {
	return action.err
}
