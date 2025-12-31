package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kimihito-sandbox/pbgoframework/handlers"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()

	isDev := os.Getenv("DEV") == "1"

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		// Serve Vite assets
		if isDev {
			// In dev mode, serve src/assets for images etc.
			se.Router.GET("/src/assets/{path...}", func(e *core.RequestEvent) error {
				http.StripPrefix("/src/assets/", http.FileServerFS(os.DirFS("frontend/src/assets"))).ServeHTTP(e.Response, e.Request)
				return nil
			})
		} else {
			// In prod mode, serve built assets
			se.Router.GET("/assets/{path...}", func(e *core.RequestEvent) error {
				http.StripPrefix("/assets/", http.FileServerFS(os.DirFS("frontend/dist/assets"))).ServeHTTP(e.Response, e.Request)
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
