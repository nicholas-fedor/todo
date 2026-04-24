package main

import (
	"io/fs"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/gofiber/fiber/v3/middleware/favicon"
	"github.com/gofiber/fiber/v3/middleware/static"

	"github.com/nicholas-fedor/todo/internal/handlers"
	"github.com/nicholas-fedor/todo/internal/middleware"
	"github.com/nicholas-fedor/todo/internal/storage"
	"github.com/nicholas-fedor/todo/internal/web"
)

const shutdownTimeout = 10 * time.Second

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

	// ----- Register Fiber shutdown hooks -----
	// Add pre-shutdown hook to close database
	app.Hooks().OnPreShutdown(func() error {
		log.Info("Pre-shutdown: closing Badger store...")

		err := store.Close()
		if err != nil {
			log.Error(err)
			// Logging instead of returning error to ensure shutdown continues.
		}

		return nil
	})

	// Add post-shutdown hook for final logging
	app.Hooks().OnPostShutdown(func(err error) error {
		if err != nil {
			log.Errorf("Shutdown completed with error: %v", err)
		} else {
			log.Info("Server shutdown completed successfully")
		}

		return nil
	})

	// ----- Server Operations -----

	// Start the server in a goroutine (required for hooks to fire)
	go func() {
		err := app.Listen(":3000")
		if err != nil {
			log.Error(err)
		}
	}()

	// Wait for termination signal (Ctrl+C, Docker/K8s stop, etc.)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Info("Received shutdown signal — gracefully shutting down...")

	// Graceful shutdown with 10 second timeout
	err = app.ShutdownWithTimeout(shutdownTimeout)
	if err != nil {
		log.Error(err)
	}
}
