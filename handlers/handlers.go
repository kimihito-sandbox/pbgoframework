package handlers

import (
	"html/template"

	"github.com/a-h/templ"
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

// renderPage renders a full page with Vite tags
func (h *Handlers) renderPage(e *core.RequestEvent, page func(template.HTML) templ.Component) error {
	viteTags, err := h.getViteTags()
	if err != nil {
		return err
	}
	h.html(e)
	return page(viteTags).Render(e.Request.Context(), e.Response)
}

// renderFragment renders an HTML fragment (for htmx)
func (h *Handlers) renderFragment(e *core.RequestEvent, component templ.Component) error {
	h.html(e)
	return component.Render(e.Request.Context(), e.Response)
}

// Page handlers

func (h *Handlers) HomeHandler(e *core.RequestEvent) error {
	return h.renderPage(e, templates.Home)
}

func (h *Handlers) AboutHandler(e *core.RequestEvent) error {
	return h.renderPage(e, templates.About)
}

// htmx handlers

func (h *Handlers) GreetingHandler(e *core.RequestEvent) error {
	return h.renderFragment(e, templates.Greeting("こんにちは！サーバーからの挨拶です 👋"))
}

func (h *Handlers) CounterIncrementHandler(e *core.RequestEvent) error {
	h.counter++
	return h.renderFragment(e, templates.Counter(h.counter))
}

func (h *Handlers) CounterDecrementHandler(e *core.RequestEvent) error {
	h.counter--
	return h.renderFragment(e, templates.Counter(h.counter))
}
