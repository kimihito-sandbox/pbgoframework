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
	counter     int
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

// API handlers for htmx

func (h *Handlers) GreetingHandler(e *core.RequestEvent) error {
	var buf bytes.Buffer
	if err := templates.Greeting("こんにちは！サーバーからの挨拶です 👋").Render(e.Request.Context(), &buf); err != nil {
		return err
	}
	return e.HTML(http.StatusOK, buf.String())
}

func (h *Handlers) CounterIncrementHandler(e *core.RequestEvent) error {
	h.counter++
	return h.renderCounter(e)
}

func (h *Handlers) CounterDecrementHandler(e *core.RequestEvent) error {
	h.counter--
	return h.renderCounter(e)
}

func (h *Handlers) renderCounter(e *core.RequestEvent) error {
	var buf bytes.Buffer
	if err := templates.Counter(h.counter).Render(e.Request.Context(), &buf); err != nil {
		return err
	}
	return e.HTML(http.StatusOK, buf.String())
}
