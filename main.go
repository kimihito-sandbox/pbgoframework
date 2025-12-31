package main

import (
	"embed"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/kimihito-sandbox/pbgoframework/handlers"
	"github.com/olivere/vite"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

//go:embed all:frontend/dist
var distFS embed.FS

var (
	isDev     = os.Getenv("DEV") == "1"
	viteEntry = "src/main.js"
)

func getDistFS() fs.FS {
	fsys, err := fs.Sub(distFS, "frontend/dist")
	if err != nil {
		log.Fatal(err)
	}
	return fsys
}

func getViteTags() (template.HTML, error) {
	var fsys fs.FS
	if isDev {
		fsys = os.DirFS("frontend")
	} else {
		fsys = getDistFS()
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

func main() {
	app := pocketbase.New()

	h := handlers.New(getViteTags)

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		// Serve static assets in production
		if !isDev {
			assets, err := fs.Sub(distFS, "frontend/dist/assets")
			if err != nil {
				log.Fatal(err)
			}
			se.Router.GET("/assets/{path...}", func(e *core.RequestEvent) error {
				http.StripPrefix("/assets/", http.FileServerFS(assets)).ServeHTTP(e.Response, e.Request)
				return nil
			})
		}

		// SSR routes
		se.Router.GET("/", h.HomeHandler)
		se.Router.GET("/about", h.AboutHandler)

		return se.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
