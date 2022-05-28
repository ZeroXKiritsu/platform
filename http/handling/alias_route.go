package handling

import (
	"fmt"
	"net/http"
	"platform/http/actionresult"
	"platform/services"
	"reflect"
	"regexp"
)

func (rc *RouterComponent) AddMethodAlias(srcUrl string, method any, data ...any) *RouterComponent {
	var urlGen URLGenerator
	services.GetService(&urlGen)
	url, err := urlGen.GenerateUrl(method, data...)
	if err == nil {
		return rc.AddUrlAlias(srcUrl, url)
	} else {
		panic(err)
	}
}

func (rc *RouterComponent) AddUrlAlias(srcUrl string, targetUrl string) *RouterComponent {
	aliasFunc := func(any) actionresult.ActionResult {
		return actionresult.NewRedirectAction(targetUrl)
	}
	alias := Route{
		httpMethod:  http.MethodGet,
		handlerName: "Alias",
		actionName:  "Redirect",
		expression:  *regexp.MustCompile(fmt.Sprintf("^%v[/]?$", srcUrl)),
		handlerMethod: reflect.Method{
			Type: reflect.TypeOf(aliasFunc),
			Func: reflect.ValueOf(aliasFunc),
		},
	}
	rc.routes = append([]Route{alias}, rc.routes...)
	return rc
}
