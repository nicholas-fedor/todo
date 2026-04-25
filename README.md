<!-- markdownlint-disable -->
<div align="center">

# Todo

<img src=".github/assets/todo.png" alt="todo logo" width="128"/>

A basic todo web app project for testing and breaking things.
Built with [Go](https://go.dev/), [BadgerDB](https://github.com/dgraph-io/badger), [Fiber](https://github.com/gofiber/fiber), [templ](https://github.com/a-h/templ), [templUI](https://github.com/templui/templui), [Tailwind CSS](https://tailwindcss.com/), and [HTMX](https://htmx.org/).

<!-- markdownlint-restore -->
</div>

> [!WARNING]
> This project is intended for testing everything from project architecture to release engineering.
> The application is not meant to be anything other than a lorem ipsum filler.
> If you wish to use this, then fork the project.

## Quickstart

### Docker CLI

Pull and Run the Docker Hub image:

```bash
docker run -p 3000:3000 -v todo-data:/data nickfedor/todo:latest
```

Or the GitHub Container Registry image:

```bash
docker run -p 3000:3000 -v todo-data:/data ghcr.io/nicholas-fedor/todo:latest
```

The app will be available at <http://localhost:3000>.

#### Notes

- The `-p 3000:3000` flag maps host port `3000` to container port `3000`.
- The `-v todo-data:/data` flag mounts a `todo-data` Docker volume to the `/data` mountpoint inside the container.

### Docker Compose

The following [`docker-compose.yml`](docker-compose.yml) configuration file is preferable for actual deployments:

```yaml
services:
  todo:
    image: nickfedor/todo:latest
    # Alternative: ghcr.io/nicholas-fedor/todo:latest
    container_name: todo
    ports:
      - "3000:3000"
    volumes:
      - todo-data:/data
    restart: unless-stopped

volumes:
  todo-data:
```

#### Start the service

```bash
docker compose up -d
```

- The app will be available at <http://localhost:3000>.

#### Stop the service

```bash
docker compose down
```

#### Remove the `todo-data` Docker volume

The preferred method is to use `docker compose down -v`, which stops the stack and removes all volumes regardless of project name:

```bash
docker compose down -v
```

Alternatively, you can remove the volume directly:

```bash
docker volume rm todo_todo-data
```

Note that this command assumes the default project name `todo` (derived from the directory name). If a different project name is used, the volume name will be prefixed accordingly (e.g., `${project_name}_todo-data`).

To ensure consistent volume naming, you can pin the project name by adding a `name` field to `docker-compose.yml`:

```yaml
name: todo
```

With this configuration, the volume will always be named `todo_todo-data` regardless of the directory name.

## Installation

### Docker

The application is primarily designed to be run as an application container (i.e. Docker).
Images are available via either  [Docker Hub](https://hub.docker.com/r/nickfedor/todo) or [GitHub Container Registry (GHCR)](https://github.com/nicholas-fedor/todo/packages).

Semantically-versioned images are available with `v{major}.{minor}.{patch}`-specific tags: `v0`, `v0.1`, `v0.1.0`, and `latest`.
Multi-platform images are available with platform-specific tags: `amd64-latest`, `arm64v8-latest`, `riscv64-latest`, etc.

### Networking

The application is currently hardcoded to use port `3000`.
Run the container with the desired port mapping to avoid conflicts, such as port `8080` instead:

```bash
docker run -p 8080:3000 -v todo-data:/data nickfedor/todo:latest
```

- The `-p 8080:3000` flag maps host port 8080 to container port 3000.
- The app will be available at <http://localhost:8080> in this configuration.

#### Storage

The application uses Badger DB for persistent storage, which is mounted to the container's `/data` volume mountpoint.

```bash
docker run  -p 3000:3000 -v todo-data:/data nickfedor/todo:latest
```

- The `-v todo-data:/data` flag mounts a Docker volume at `/data` inside the container, where BadgerDB stores its data.

If you wish to use a local bind mount, then you could use the following:

```bash
docker run  -p 3000:3000 -v ./data:/data nickfedor/todo:latest
```

- The `-v ./data:/data` flag mounts a local `data` directory relative to the current terminal's command context. If you are in `/root` and run the Docker CLI  command, then the `data` directory would be bind mounted to `/root/data`.

### GitHub Releases [![Latest](https://img.shields.io/badge/Latest-2ea44f?style=flat-square)](https://github.com/nicholas-fedor/todo/releases/latest)

Prebuilt binaries are available for the following platforms:

- Linux (amd64/386/arm/arm64/riscv64)
- macOS (amd64/arm64)
- Windows (amd64/386).

### From Source

Clone the repository:

```bash
git clone https://github.com/nicholas-fedor/todo.git
```

Change to the project directory:

```bash
cd todo
```

Full installation using Make (installs tools, generates assets, builds, and installs binary):

```bash
make install
```

Or using Task:

```bash
task install
```

The binary is installed to the invoking user's `$GOPATH/bin` (typically `~/go/bin/todo`). Using `sudo` would install to root's `$GOPATH` instead.

Alternatively, use `go install` for direct installation (requires Go toolchain):

```bash
go install github.com/nicholas-fedor/todo/cmd/server@latest
```

> [!NOTE]
> Building from source requires Tailwind CSS for the webpage styling.
> This project uses [bun](https://bun.com/docs/installation) for package management (already configured in `package.json`).
> Run `bun install` if building the frontend assets as part of a full development setup.

## Development

Refer to [CONTRIBUTING.md](CONTRIBUTING.md) for development-related documentation.

## License

AGPLv3 — see [LICENSE](LICENSE) for details.
