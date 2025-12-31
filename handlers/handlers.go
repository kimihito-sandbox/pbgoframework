package handlers

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/kimihito-sandbox/pbgoframework/templates"
	"github.com/pocketbase/pocketbase/core"
)

// ViteTagsFunc is a function that returns Vite HTML tags
type ViteTagsFunc func() (template.HTML, error)

// New creates handlers with the given ViteTags function
func New(getViteTags ViteTagsFunc) *Handlers {
	return &Handlers{getViteTags: getViteTags}
}

// Handlers holds the dependencies for HTTP handlers
type Handlers struct {
	getViteTags ViteTagsFunc
}

func (h *Handlers) HomeHandler(e *core.RequestEvent) error {
	viteTags, err := h.getViteTags()
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := templates.Home(viteTags).Render(e.Request.Context(), &buf); err != nil {
		return err
	}

	return e.HTML(http.StatusOK, buf.String())
}

func (h *Handlers) AboutHandler(e *core.RequestEvent) error {
	viteTags, err := h.getViteTags()
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := templates.About(viteTags).Render(e.Request.Context(), &buf); err != nil {
		return err
	}

	return e.HTML(http.StatusOK, buf.String())
}
