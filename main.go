package main

import (
	"log"

	"github.com/kimihito-sandbox/pbgoframework/handlers"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		// SSR routes
		se.Router.GET("/", handlers.HomeHandler)
		se.Router.GET("/about", handlers.AboutHandler)

		return se.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
