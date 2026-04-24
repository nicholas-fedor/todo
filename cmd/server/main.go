package main

import (
	"io/fs"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/gofiber/fiber/v3/middleware/favicon"
	"github.com/gofiber/fiber/v3/middleware/static"

	"github.com/nicholas-fedor/todo/internal/handlers"
	"github.com/nicholas-fedor/todo/internal/middleware"
	"github.com/nicholas-fedor/todo/internal/storage"
	"github.com/nicholas-fedor/todo/internal/web"
)

func main() {
	// ----- Setup Backend Storage -----

	// Provide storage using the "data" directory
	store, err := storage.NewBadgerStore(
		storage.BadgerOptions{
			// This uses current session's active directory.
			Dir: "data",
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	defer store.Close()

	// ----- Setup Frontend Storage -----

	// Create asset storage rooted at the assets directory to match URL paths
	assetFS, _ := fs.Sub(
		web.Assets,
		"assets",
	)

	// Create favicon storage pointing to `assets/images`
	faviconFS, _ := fs.Sub(
		web.Assets,
		"assets/images",
	)

	// ----- Setup Fiber App -----

	// Create new Fiber instance
	app := fiber.New()

	// Initialize default favicon config
	app.Use(favicon.New(
		favicon.Config{
			File:       "favicon.ico",
			FileSystem: faviconFS,
		},
	))

	// ----- Register Routes -----

	// Serve embedded assets
	app.Get("/assets*",
		static.New(
			"",
			static.Config{
				FS: assetFS,
			},
		))

	// Create new GET route on path "/" and pass store to handler
	app.Get(
		"/",
		handlers.HomeHandler(store),
	)

	// Create new POST route for adding todos
	app.Post(
		"/todos",
		handlers.CreateHandler(store),
	)

	// Create new PUT route for toggling todo completion
	app.Put(
		"/todos/:id/toggle",
		handlers.ToggleHandler(store),
	)

	// Create new DELETE route for removing todos
	app.Delete(
		"/todos/:id",
		handlers.DeleteHandler(store),
	)

	// ----- Add 404 catch-all -----
	// In Fiber, `app.Use()` middleware runs in registration order.
	// The order is: static assets → favicon → routes → 404 catch-all

	// Add 404 not found middleware
	app.Use(middleware.NotFoundMiddleware)

	// ----- Server Operations -----

	// Start server on http://localhost:3000
	log.Fatal(app.Listen(":3000"))
}
