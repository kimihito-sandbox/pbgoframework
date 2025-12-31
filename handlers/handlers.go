package handlers

import (
	"html/template"
	"os"

	"github.com/kimihito-sandbox/pbgoframework/templates"
	"github.com/olivere/vite"
	"github.com/pocketbase/pocketbase/core"
)

var (
	isDev     bool
	viteEntry = "src/main.js"
)

func init() {
	isDev = os.Getenv("DEV") == "1"
}

func getViteTags() (template.HTML, error) {
	var fs = os.DirFS("frontend/dist")
	if isDev {
		fs = os.DirFS("frontend")
	}

	fragment, err := vite.HTMLFragment(vite.Config{
		FS:        fs,
		IsDev:     isDev,
		ViteURL:   "http://localhost:5173",
		ViteEntry: viteEntry,
	})
	if err != nil {
		return "", err
	}
	return fragment.Tags, nil
}

func HomeHandler(e *core.RequestEvent) error {
	viteTags, err := getViteTags()
	if err != nil {
		return err
	}
	return templates.Home(viteTags).Render(e.Request.Context(), e.Response)
}

func AboutHandler(e *core.RequestEvent) error {
	viteTags, err := getViteTags()
	if err != nil {
		return err
	}
	return templates.About(viteTags).Render(e.Request.Context(), e.Response)
}
