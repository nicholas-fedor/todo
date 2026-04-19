package main

import (
	"embed"
	"io/fs"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/gofiber/fiber/v3/middleware/favicon"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/nicholas-fedor/todo/internal/handler"
	"github.com/nicholas-fedor/todo/internal/middleware"
	"github.com/nicholas-fedor/todo/storage"
)

// Embed assets from the assets directory
//
//go:embed assets
var assetsFS embed.FS

func main() {
	// ----- Setup Backend Storage -----

	// Provide storage using the "data" directory
	store, err := storage.NewBadgerStore(
		storage.BadgerOptions{
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
		assetsFS,
		"assets",
	)

	// Create favicon storage pointing to `assets/images`
	faviconFS, _ := fs.Sub(
		assetsFS,
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
		handler.HomeHandler(store),
	)

	// Create new POST route for adding todos
	app.Post(
		"/todos",
		handler.CreateHandler(store),
	)

	// Create new PUT route for toggling todo completion
	app.Put(
		"/todos/:id/toggle",
		handler.ToggleHandler(store),
	)

	// Create new DELETE route for removing todos
	app.Delete(
		"/todos/:id",
		handler.DeleteHandler(store),
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
