package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/kimihito-sandbox/pbgoframework/handlers"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

//go:embed all:frontend/dist
var distFS embed.FS

func main() {
	app := pocketbase.New()

	// Set up embedded dist for production
	if !handlers.IsDev() {
		dist, err := fs.Sub(distFS, "frontend/dist")
		if err != nil {
			log.Fatal(err)
		}
		handlers.SetDistFS(dist)
	}

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		// Serve static assets
		if !handlers.IsDev() {
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
		se.Router.GET("/", handlers.HomeHandler)
		se.Router.GET("/about", handlers.AboutHandler)

		return se.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
