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

First, set the volume permissions to UID 1000:GID 1000 (matching the container user):

```bash
docker run --rm -v todo.data:/data alpine chown -R 1000:1000 /data
```

Pull and Run the `todo` Docker Hub image:

```bash
docker run -p 3000:3000 -v todo.data:/data nickfedor/todo:latest
```

Or the `todo` GitHub Container Registry image:

```bash
docker run -p 3000:3000 -v todo.data:/data ghcr.io/nicholas-fedor/todo:latest
```

The app will be available at <http://localhost:3000>.

#### Notes

- The `-p 3000:3000` flag maps host port `3000` to container port `3000`.
- The `-v todo.data:/data` flag mounts a `todo.data` Docker volume to the `/data` mountpoint inside the container.
- The initial volume setup step is necessary due to the `todo` container running nonroot `1000:1000` by default.

### Docker Compose

The following [`docker-compose.yml`](docker-compose.yml) configuration file is preferable for actual deployments:

```yaml
services:
  #  one-time sidecar to set volume permissions to 1000:100
  volume-manager:
    command: chown -R 1000:1000 /data
    container_name: volume-manager
    image: alpine
    user: root:root # runs as root to match Docker daemon for volume management
    volumes:
      - todo.data:/data

  todo:
    container_name: todo
    depends_on:
      volume-manager:
        condition: service_completed_successfully
    image: nickfedor/todo:latest-amd64
    # image: ghcr.io/nicholas-fedor/todo:latest
    volumes:
      - todo.data:/data
    ports:
      - "3000:3000"
    restart: unless-stopped
    security_opt:
      - "no-new-privileges=true"

volumes:
  todo.data:
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

#### Remove the `todo.data` Docker volume

The preferred method is to use `docker compose down -v`, which stops the stack and removes all volumes regardless of project name:

```bash
docker compose down -v
```

Alternatively, you can remove the volume directly:

```bash
docker volume rm todo_todo.data
```

> [!NOTE]
> This command assumes the default project name `todo` (derived from the directory name). If a different project name is used, the volume name will be prefixed accordingly (e.g., `${project_name}_todo.data`).
>
> To ensure consistent volume naming, you can pin the project name by adding a `name` field to `docker-compose.yml`:
>
> ```yaml
> name: todo
> ```
>
> With this configuration, the volume will always be named `todo_todo.data` regardless of the directory name.

## Installation

### Docker

The application is primarily designed to be run as an application container (i.e. Docker).
Images are available via either  [Docker Hub](https://hub.docker.com/r/nickfedor/todo) or [GitHub Container Registry (GHCR)](https://github.com/nicholas-fedor/todo/packages).

Semantically-versioned images are available with `v{major}.{minor}.{patch}`-specific tags: `v0`, `v0.1`, `v0.1.0`, and `latest`.
Multi-platform images are available with platform-specific tags: `amd64-latest`, `arm64v8-latest`, `riscv64-latest`, etc.

#### Networking

The application is currently hardcoded to use port `3000`.
Run the container with the desired port mapping to avoid conflicts, such as port `8080` instead:

```bash
docker run -p 8080:3000 -v todo.data:/data nickfedor/todo:latest
```

- The `-p 8080:3000` flag maps host port 8080 to container port 3000.
- The app will be available at <http://localhost:8080> in this configuration.

#### Storage

The application uses Badger DB for persistent storage, which is mounted to the container's `/data` volume mountpoint.
If running via Docker, the image uses a default nonroot user of 1000:1000, so your storage location will need to use the corresponding permissions.

- First, set the volume permissions to UID:GID of 1000:1000

```bash
docker run --rm -v todo.data:/data alpine chown -R 1000:1000 /data
```

- Then run the container:

```bash
docker run  -p 3000:3000 -v todo.data:/data nickfedor/todo:latest
```

- The `-v todo.data:/data` flag mounts a Docker volume at `/data` inside the container, where BadgerDB stores its data.

If you wish to use a local bind mount, then you could use the following:

```bash
docker run  -p 3000:3000 -v ./data:/data nickfedor/todo:latest
```

- The `-v ./data:/data` flag mounts a local `data` directory relative to the current terminal's command context. If you are in `/root` and run the Docker CLI  command, then the `data` directory would be bind mounted to `/root/data`.

> [!NOTE]
> You can also opt to run the container using a different `user:group`; however, this will still require ensuring correct filesystem permissions.
> The use of a `1000:1000` default is an attempt at balancing security and convenience -- given the Docker default is `root` level access.

### GitHub Releases 

[![Latest](https://img.shields.io/badge/Latest-2ea44f?style=flat-square)](https://github.com/nicholas-fedor/todo/releases/latest)

Prebuilt binaries are available for the following platforms:

- Linux (amd64/386/arm/arm64/riscv64)
- macOS (amd64/arm64)
- Windows (amd64/386).

### From Source

> [!NOTE]
> Building from source requires Tailwind CSS (managed via [bun](https://bun.com/docs/installation)).

Clone the repository:

```bash
git clone https://github.com/nicholas-fedor/todo.git
```

Change to the project directory:

```bash
cd todo
```

**Prerequisite:** Install frontend dependencies and build Tailwind CSS assets:

```bash
bun install
make tailwind-build   # or: task tailwind-build
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

## Development

Refer to [CONTRIBUTING.md](CONTRIBUTING.md) for development-related documentation.

## License

AGPLv3 — see [LICENSE](LICENSE) for details.
