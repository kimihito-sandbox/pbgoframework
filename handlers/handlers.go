package handlers

import (
	"bytes"
	"html/template"
	"io/fs"
	"net/http"
	"os"

	"github.com/kimihito-sandbox/pbgoframework/templates"
	"github.com/olivere/vite"
	"github.com/pocketbase/pocketbase/core"
)

var (
	isDev     bool
	viteEntry = "src/main.js"
	distFS    fs.FS // Set by main.go for production
)

func init() {
	isDev = os.Getenv("DEV") == "1"
}

// SetDistFS sets the embedded dist filesystem for production
func SetDistFS(fsys fs.FS) {
	distFS = fsys
}

// IsDev returns whether the app is running in development mode
func IsDev() bool {
	return isDev
}

func getViteTags() (template.HTML, error) {
	var fsys fs.FS
	if isDev {
		fsys = os.DirFS("frontend")
	} else {
		fsys = distFS
	}

	fragment, err := vite.HTMLFragment(vite.Config{
		FS:        fsys,
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

	var buf bytes.Buffer
	if err := templates.Home(viteTags).Render(e.Request.Context(), &buf); err != nil {
		return err
	}

	return e.HTML(http.StatusOK, buf.String())
}

func AboutHandler(e *core.RequestEvent) error {
	viteTags, err := getViteTags()
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := templates.About(viteTags).Render(e.Request.Context(), &buf); err != nil {
		return err
	}

	return e.HTML(http.StatusOK, buf.String())
}
