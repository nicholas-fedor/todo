# Todo

Basic todo web app for testing. Built with Go, Fiber, templ, Tailwind CSS, BadgerDB, and HTMX.

## Quick Start

```bash
go mod download
bun install
task dev  # http://localhost:3000
```

## Production

```bash
templ generate && go build -o todo .
./todo  # Listens on :3000
```

Single binary with all static assets (CSS/JS) embedded via `//go:embed`.

Full documentation (installation, architecture, storage, HTMX patterns, troubleshooting): [CONTRIBUTING.md](CONTRIBUTING.md).

## License

AGPLv3 — see [LICENSE](LICENSE) for details.
