package handlers

import (
	"html/template"

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

func (h *Handlers) html(e *core.RequestEvent) {
	e.Response.Header().Set("Content-Type", "text/html; charset=utf-8")
}

func (h *Handlers) HomeHandler(e *core.RequestEvent) error {
	viteTags, err := h.getViteTags()
	if err != nil {
		return err
	}
	h.html(e)
	return templates.Home(viteTags).Render(e.Request.Context(), e.Response)
}

func (h *Handlers) AboutHandler(e *core.RequestEvent) error {
	viteTags, err := h.getViteTags()
	if err != nil {
		return err
	}
	h.html(e)
	return templates.About(viteTags).Render(e.Request.Context(), e.Response)
}

// htmx handlers

func (h *Handlers) GreetingHandler(e *core.RequestEvent) error {
	h.html(e)
	return templates.Greeting("こんにちは！サーバーからの挨拶です 👋").Render(e.Request.Context(), e.Response)
}

func (h *Handlers) CounterIncrementHandler(e *core.RequestEvent) error {
	h.counter++
	h.html(e)
	return templates.Counter(h.counter).Render(e.Request.Context(), e.Response)
}

func (h *Handlers) CounterDecrementHandler(e *core.RequestEvent) error {
	h.counter--
	h.html(e)
	return templates.Counter(h.counter).Render(e.Request.Context(), e.Response)
}
