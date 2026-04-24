<!-- markdownlint-disable -->
<div align="center">

# Todo

<img src=".github/assets/todo.png" alt="todo logo" width="128"/>

A basic todo web app for testing.
Built with Go, Fiber, templ, Tailwind CSS, BadgerDB, and HTMX.

<br/><br/>
<!-- markdownlint-restore -->
</div>

## Quick Start

```bash
# Install Go dependencies
go mod download

# Install Tailwind CSS via npm/bun (dev dependency)
# The project uses bun for package management
bun install

# Or using npm:
# npm install

# Install development tools (optional, for hot-reload)
go install github.com/air-verse/air@latest
go install github.com/a-h/templ/cmd/templ@latest

# Start development server
task dev  # http://localhost:3000
```

## Production

```bash
# Generate templates and build
templ generate && go build -o todo ./cmd/server

# Run
./todo  # Listens on :3000
```

Single binary with all static assets (CSS/JS) embedded via `//go:embed`.

Full documentation (installation, architecture, storage, HTMX patterns, troubleshooting): [CONTRIBUTING.md](CONTRIBUTING.md).

## License

AGPLv3 — see [LICENSE](LICENSE) for details.
