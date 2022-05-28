package sessions

import (
	"context"
	gorilla "github.com/gorilla/sessions"
	"platform/config"
	"platform/pipeline"
	"time"
)

type SessionComponent struct {
	store *gorilla.CookieStore
	config.Configuration
}

func (sc *SessionComponent) Init() {
	cookieKey, found := sc.Configuration.GetString("sessions:key")
	if !found {
		panic("Session key not found in configuration")
	}
	if sc.GetBoolDefault("sessions:cyclekey", true) {
		cookieKey += time.Now().String()
	}
	sc.store = gorilla.NewCookieStore([]byte(cookieKey))
}

func (sc *SessionComponent) ProcessRequest(ctx *pipeline.ComponentContext, next func(*pipeline.ComponentContext)) {
	session, _ := sc.store.Get(ctx.Request, SESSION_CONTEXT_KEY)
	c := context.WithValue(ctx.Request.Context(), SESSION_CONTEXT_KEY, session)
	ctx.Request = ctx.Request.WithContext(c)
	next(ctx)
	session.Save(ctx.Request, ctx.ResponseWriter)
}
