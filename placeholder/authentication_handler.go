package placeholder

import (
	"fmt"
	"platform/authorization/identity"
	"platform/http/actionresult"
)

type AuthenticationHandler struct {
	identity.User
	identity.SignInManager
	identity.UserStore
}

func (h AuthenticationHandler) GetSignIn() actionresult.ActionResult {
	return actionresult.NewTemplateAction("signin.html", fmt.Sprintf("Signed in as: %v", h.User.GetDisplayName()))
}

type Credentials struct {
	Username string
	Password string
}

func (h AuthenticationHandler) PostSignIn(creds Credentials) actionresult.ActionResult {
	if creds.Password == "mysecret" {
		user, ok := h.UserStore.GetUserByName(creds.Username)
		if ok {
			h.SignInManager.SignIn(user)
			return actionresult.NewTemplateAction("signin.html", fmt.Sprintf("Signed in as: %v", h.User.GetDisplayName()))
		}
	}
	return actionresult.NewTemplateAction("signin.html", "Access Denied")
}

func (h AuthenticationHandler) PostSignOut() actionresult.ActionResult {
	h.SignInManager.SignOut(h.User)
	return actionresult.NewTemplateAction("signin.html", "Signed out")
}
