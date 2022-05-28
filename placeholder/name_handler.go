package placeholder

import (
	"fmt"
	"platform/http/actionresult"
	"platform/http/handling"
	"platform/logging"
	"platform/validation"
)

var names = []string{"Alice", "Bob", "Charlie", "Dora"}

type NameHandler struct {
	logging.Logger
	handling.URLGenerator
	validation.Validator
}

func (n NameHandler) GetName(i int) actionresult.ActionResult {
	n.Logger.Debugf("GetName method invoked with argument: %v", i)
	var response string
	if i < len(names) {
		response = fmt.Sprintf("Name #%v: %v", i, names[i])
	} else {
		response = fmt.Sprintf("Index out of bounds")
	}
	return actionresult.NewTemplateAction("simple_message.html", response)
}
func (n NameHandler) GetNames() actionresult.ActionResult {
	n.Logger.Debug("GetNames method invoked")
	return actionresult.NewTemplateAction("simple_message.html", names)
}

type NewName struct {
	Name          string `validation:"required,min:3"`
	InsertAtStart bool
}

func (n NameHandler) GetForm() actionresult.ActionResult {
	postUrl, _ := n.URLGenerator.GenerateUrl(NameHandler.PostName)
	return actionresult.NewTemplateAction("name_form.html", postUrl)
}

func (n NameHandler) PostName(new NewName) actionresult.ActionResult {
	n.Logger.Debugf("PostName method invoked with argument %v", new)
	if ok, errs := n.Validator.Validate(&new); !ok {
		return actionresult.NewTemplateAction("validation_errors.html", errs)
	}
	if new.InsertAtStart {
		names = append([]string{new.Name}, names...)
	} else {
		names = append(names, new.Name)
	}
	return n.redirectOrError(NameHandler.GetNames)
}

func (n NameHandler) GetRedirect() actionresult.ActionResult {
	return n.redirectOrError(NameHandler.GetNames)
}

func (n NameHandler) GetJsonData() actionresult.ActionResult {
	return actionresult.NewJsonAction(names)
}

func (n NameHandler) redirectOrError(handler any, data ...any) actionresult.ActionResult {
	url, err := n.GenerateUrl(handler)
	if err == nil {
		return actionresult.NewRedirectAction(url)
	} else {
		return actionresult.NewErrorAction(err)
	}
}
