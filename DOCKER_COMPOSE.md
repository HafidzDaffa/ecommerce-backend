# Docker Compose Notes

## Command Changes

Starting from Docker Compose v2, the command has changed from `docker-compose` to `docker compose` (without hyphen).

### Old (Docker Compose v1)
```bash
docker-compose up
docker-compose build
docker-compose down
```

### New (Docker Compose v2)
```bash
docker compose up
docker compose build
docker compose down
```

## Installation

If you don't have Docker Compose v2 installed:

```bash
# On Ubuntu/Debian
sudo apt-get update
sudo apt-get install docker-compose-plugin

# On macOS with Homebrew
brew install docker

# On Windows
# Docker Desktop includes docker compose plugin
```

## Verify Installation

```bash
docker compose version
# Should show: Docker Compose version v2.x.x
```

## Makefile

The `Makefile` has been updated to use `docker compose` commands. All make commands should work correctly:
- `make build`
- `make run`
- `make stop`
- `make logs`
- etc.

If you still need to use `docker-compose` (v1), you can modify the Makefile or run the commands directly:

```bash
docker-compose build
docker-compose up -d
```
