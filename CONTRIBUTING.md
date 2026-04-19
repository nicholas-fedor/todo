# Contributing

Complete development guide for the Todo application: setup, architecture, and workflows.

## Quick Start

```bash
go mod download
bun install
task dev  # http://localhost:3000
```

See the sections below for detailed tool installation and configuration.

---

## Project Structure

```text
├── main.go                  # Fiber app setup and routes
├── go.mod                   # Go dependencies
├── package.json             # Node dependencies (Tailwind CLI)
├── Taskfile.yml             # Development tasks (dev, tailwind, templ)
├── .air.toml                # Air configuration
├── .templui.json            # TemplUI CLI configuration
├── assets/                  # Embedded static assets
│   ├── css/
│   │   ├── input.css
│   │   └── output.css
│   └── js/                  # Component JavaScript (populated by templui)
├── components/              # Local TemplUI component overrides
├── internal/
│   ├── handler/            # HTTP handlers (home, create, toggle, delete)
│   ├── middleware/         # 404 catch-all
│   └── models/            # Todo data structures
├── pkg/
│   ├── pages/             # Top-level templ pages (home, notfound)
│   └── todo/              # Todo business logic
├── storage/               # BadgerDB store abstraction
│   ├── storage.go         # Interface
│   └── badger.go          # Implementation
└── data/                  # BadgerDB data directory (created at runtime)
```

---

## Development Tools & Configuration

### Air

Live-reloading for Go development via CTL+C restart and rebuild on file changes.

#### Installation

```bash
go install github.com/air-verse/air@latest
air init
```

#### Configured Behavior (`.air.toml`)

Air runs:

```bash
templ generate && go build -o ./tmp/main .
```

Template file changes (`*.templ`) trigger a rebuild. Generated files (`*_templ.go`) are excluded from watching.

#### Air Resources

- GitHub: <https://github.com/air-verse/air>
- Config reference: <https://github.com/air-verse/air/blob/master/air_example.toml>

---

### Fiber v3

Fast Go web framework used for routing, middleware, and static file serving.

Already included:

```bash
go get github.com/gofiber/fiber/v3
```

#### Application Structure (`main.go`)

```go
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

//go:embed assets
var assetsFS embed.FS

func main() {
    store, err := storage.NewBadgerStore(storage.BadgerOptions{Dir: "data"})
    if err != nil {
        log.Fatal(err)
    }
    defer store.Close()

    assetFS, _ := fs.Sub(assetsFS, "assets")
    faviconFS, _ := fs.Sub(assetsFS, "assets/images")

    app := fiber.New()
    app.Use(favicon.New(favicon.Config{File: "favicon.ico", FileSystem: faviconFS}))

    app.Get("/assets*", static.New("", static.Config{FS: assetFS}))
    app.Get("/", handler.HomeHandler(store))
    app.Post("/todos", handler.CreateHandler(store))
    app.Put("/todos/:id/toggle", handler.ToggleHandler(store))
    app.Delete("/todos/:id", handler.DeleteHandler(store))
    app.Use(middleware.NotFoundMiddleware)

    log.Fatal(app.Listen(":3000"))
}
```

#### Fiber Resources

- Docs: <https://docs.gofiber.io/>
- Air integration: <https://docs.gofiber.io/recipes/air/>

---

### templ

Type-safe HTML templating with Go code generation.

#### Install templ

```bash
go install github.com/a-h/templ/cmd/templ@latest
```

#### Generate Templates

```bash
templ generate
```

Compiles all `*.templ` files into `*_templ.go` Go source files. These are compiled directly into the binary — no runtime template parsing.

#### Hot Reload with Air

When a `.templ` file changes, Air's `cmd` config re-runs `templ generate` before rebuilding. No manual step needed.

#### Integration with Fiber

Render a templ component by calling it from `pkg/pages`:

```go
import "github.com/nicholas-fedor/todo/pkg/pages"

todos, _ := todo.ListTodos(store)
return Render(c, pages.Home(todos))
```

#### templ Resources

- Official guide: <https://templ.guide/>
- Fiber example: <https://github.com/a-h/templ/blob/main/examples/integration-gofiber/README.md>

---

### TemplUI

Pre-built component library for templ. Provides buttons, cards, inputs, dialogs, tables, etc.

#### Docs

<https://templui.io/>

#### Component Installation

#### Option A — Import from upstream (simpler, references remote module)

```go
import "github.com/templui/templui/components/button"

@button.Button() {
    Click me
}
```

#### Option B — Vendor locally (downloads source into project)

```bash
templui add button
```

Then import from your project path:

```go
import "github.com/nicholas-fedor/todo/components/button"
```

Option B lets you customize components locally.

#### Interactive Components Need Scripts

Components with JavaScript behavior (checkbox, datepicker, input, etc.) require:

```templ
<head>
    @checkbox.Script()
    @datepicker.Script()
</head>
```

The `Script()` call generates the appropriate `<script src="...">` tag pointing to `assets/js/` (populated by `templui add`). For debugging, set `utils.UseUnminifiedScripts = true` in `main.go`.

#### Key Files

- `.templui.json` — module name, directories
- `components/` — vendored component overrides
- `assets/js/` — component JavaScript files

#### TemplUI Resources

- Getting started: <https://templui.io/docs/how-to-use>
- CLI: <https://templui.io/docs/cli>

---

### Tailwind CSS v4

Utility-first CSS framework. Project uses Tailwind CLI (no config file).

#### Install Dependencies

```bash
bun init -y
bun add -D tailwindcss
```

#### Build CSS

Input: `assets/css/input.css` (source with `@import` directives and custom rules)
Output: `assets/css/output.css` (compiled, referenced in templates)

```bash
bunx @tailwindcss/cli -i assets/css/input.css -o assets/css/output.css
```

#### Watch During Development

```bash
bunx @tailwindcss/cli -i assets/css/input.css -o assets/css/output.css --watch
```

Or via Taskfile: `task tailwind`

#### Link in Templates

```templ
<link rel="stylesheet" href="/assets/css/output.css">
```

Tailwind automatically scans your `.templ` files for class names. The `sources.generated.css` file (auto-generated by the Tailwind task) lists all template paths so Tailwind can purge unused classes.

---

### BadgerDB

Embedded key-value store used to persist todos.

#### Add to Project

```bash
go get github.com/dgraph-io/badger/v4
```

#### Store Interface & Implementation

Interface: `internal/storage/storage.go`

```go
type Store interface {
    Set(key string, value []byte) error
    Get(key string) ([]byte, error)
    Delete(key string) error
    Close() error
}
```

Implementation: `internal/storage/badger.go` — wraps BadgerDB operations, stores serialized `models.TodoView` JSON under keys prefixed with `todo:`.

#### Initialize in `main.go`

```go
store, err := storage.NewBadgerStore(storage.BadgerOptions{Dir: "data"})
if err != nil { log.Fatal(err) }
defer store.Close()
```

#### Data Directory

`data/` created automatically on first run. Contains BadgerDB SST files and logs.

#### Badger CLI (optional)

```bash
go install github.com/dgraph-io/badger/v4/badger@latest
badger --help
```

#### BadgerDB Resources

- Quick start: <https://dgraph-io.github.io/badger/quickstart.html>
- GoDoc: <https://pkg.go.dev/github.com/dgraph-io/badger/v4>

---

### Task

Task runner to automate dev workflows.

#### Install

```bash
go install github.com/go-task/task/v3/cmd/task@latest
```

#### Defined Tasks (`Taskfile.yml`)

| Task | Description |
| ------ | ------------ |
| `task dev` | Parallel: Tailwind watch + Air live reload |
| `task tailwind` | Build CSS (and watch with `--watch`) |
| `task templ` | Generate templates |

---

### bun

JavaScript runtime used to invoke Tailwind CLI (`bunx`). Already present as `devDependency` in `package.json`.

Install system-wide (if missing):

```bash
curl -fsSL https://bun.sh/install | bash
```

---

### UUID

Google UUID library generates unique keys for todos.

Already added:

```bash
go get github.com/google/uuid
```

Used in handler: `id := uuid.New().String()`

Package: <https://pkg.go.dev/github.com/google/uuid>

---

## Tech Stack

| Layer      | Tech                   |
|------------|------------------------|
| Backend    | Go 1.26+ + Fiber v3    |
| Templating | templ + TemplUI        |
| Styling    | Tailwind CSS v4        |
| Database   | BadgerDB v4 (embedded) |
| Dynamic UI | HTMX                   |
| Dev Tools  | Air, Task, bun         |

---

## API Reference

All endpoints return HTML. Static assets at `/assets*`. Non-asset handlers can return full pages or HTMX fragments.

| Method | Route | Handler | Purpose |
| -------- | ------- | ---------- | ---------- |
| `GET` | `/` | `handler.HomeHandler` | Render home page with all todos |
| `POST` | `/todos` | `handler.CreateHandler` | Create todo from `title` form field |
| `PUT` | `/todos/:id/toggle` | `handler.ToggleHandler` | Flip `completed` boolean |
| `DELETE` | `/todos/:id` | `handler.DeleteHandler` | Remove todo by key |

Request/response content-type: HTML (or fragments). The `CreateHandler`, `ToggleHandler`, and `DeleteHandler` delegate to `HomeHandler` to re-render the list after mutation.

---

## HTMX Integration

Dynamic UI updates powered by HTMX attributes — no custom frontend JavaScript needed.

### How It Works

1. User interaction (submit form, click, change) triggers an HTMX request
2. Server returns an HTML fragment (partial template)
3. HTMX swaps the fragment into the DOM based on `hx-target` and `hx-swap`

### Attributes in This Project

- `hx-post="/todos"` — POST to create a todo
- `hx-put="/todos/{{.Key}}/toggle"` — PUT to toggle completion
- `hx-delete="/todos/{{.Key}}"` — DELETE to remove
- `hx-target="#todo-list"` — target container to update
- `hx-swap="outerHTML"` — replace entire target element

### Templates

- `pkg/pages/home_templ.go` — compiled from `home.templ`
- Home handler passes `[]models.TodoView` to `pages.Home(todos)`
- List partial: `templ` component within `home.templ` that renders each todo item with HTMX attributes

### HTMX Resources

- HTMX docs: <https://htmx.org/docs/>
- templ + HTMX patterns: <https://templ.guide/server-side-rendering/htmx>

---

## Storage Deep Dive

### Key Schema

- Keys: `todo:{uuid}` (e.g. `todo:550e8400-e29b-41d4-a716-446655440000`)
- Value: JSON-serialized `models.TodoView`

### models.TodoView

```go
type TodoView struct {
    Key       string    `json:"key"`
    Title     string    `json:"title"`
    Completed bool      `json:"completed"`
    CreatedAt time.Time `json:"created_at"`
}
```

### CRUD Operations

- **Create** — `store.Set(key, json.Marshal(todo))` in `CreateHandler`
- **Read (all)** — `todo.ListTodos(store)` scans with `store.Iterator()` and decodes each JSON value
- **Update** — `store.Set(key, updatedJSON)` after toggling `completed`
- **Delete** — `store.Delete(key)`

### Badger Options

`storage.BadgerOptions` wraps `badger.Options`. Defaults used: `Dir: "data"`, no encryption, auto-remove corruption.

---

## Manual Commands

If not using Air/Task:

```bash
# One-off: generate templates and run
templ generate && go run .

# Generate templates only (produces *_templ.go)
templ generate

# Build CSS (no watch)
bunx @tailwindcss/cli -i assets/css/input.css -o assets/css/output.css

# Build binary
templ generate && go build -o todo .
```

---

## Production Notes

- Single binary: `//go:embed` bundles `assets/` at compile time
- No external files needed at runtime
- Static middleware serves `/assets*` from embedded FS via `fs.Sub(assetsFS, "assets")`
- Server listens on `:3000` (no HTTPS, no reverse-proxy config — add as needed)
- Ensure `data/` directory is writable by the process

---

## Troubleshooting

### Templates not updating?

- Confirm Air is running; check `.air.toml` includes `templ generate && go build` in `cmd`
- Verify `exclude_regex` prevents watching `*_templ.go`

### CSS changes not appearing?

- Tailwind watch must be running (`task dev` or `task tailwind`)
- Check `assets/css/input.css` imports the template sources via `@source` directives (auto-generated by Task)

### Assets 404?

- Static route registered: `app.Get("/assets*", static.New(...))`
- Embedded FS sub-rooted correctly: `fs.Sub(assetsFS, "assets")`

### Port already in use?

- Server hardcoded to `:3000` in `main.go:107`. Change or use `pkill -f todo`.

---

## Additional Resources

- Fiber middlewares: <https://docs.gofiber.io/middleware>
- templ patterns: <https://templ.guide/patterns/>
- HTMX examples: <https://htmx.org/examples/>
