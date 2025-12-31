package handlers

import (
	"github.com/kimihito-sandbox/pbgoframework/templates"
	"github.com/pocketbase/pocketbase/core"
)

func HomeHandler(e *core.RequestEvent) error {
	return templates.Home().Render(e.Request.Context(), e.Response)
}

func AboutHandler(e *core.RequestEvent) error {
	return templates.About().Render(e.Request.Context(), e.Response)
}
