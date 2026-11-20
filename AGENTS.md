# Repository Guidelines

## Project Scope & Current State

Bookin is a planned Go backend for a cabin-rental property. Its core concepts are cabin types, individual cabins, availability, and bookings. The repository is currently empty: no Go module, application, tests, or Git history exists. Treat the structure below as guidance for new code, not existing infrastructure.

## Project Structure & Module Organization

- Place the server entry point in `cmd/api/main.go`.
- Keep application packages under `internal/`, organized by responsibility as the code grows.
- Place tests beside their implementation using `*_test.go` files.
- Store database migrations in `migrations/` when database support is introduced.
- Document setup and configuration in `README.md`; provide configuration examples in `.env.example` without secrets.

Start with one backend application. Add directories and abstractions only when a feature needs them.

## Build, Test, and Development Commands

Initialize `go.mod` with `go mod init <module-path>` using the intended repository path. Once the module and entry point exist, run commands from the repository root:

- `go run ./cmd/api` — start the API locally.
- `go build ./...` — compile all packages.
- `go test ./...` — run all tests.
- `go test -cover ./...` — inspect test coverage.
- `go vet ./...` — check for common code mistakes.
- `gofmt -w .` — format Go source files.

## Coding Style & Naming Conventions

Use standard Go formatting, including tabs supplied by `gofmt`. Use short, lowercase package names, exported `PascalCase` identifiers, and unexported `camelCase` identifiers. Handle errors explicitly and add useful context when wrapping them. Prefer the standard library unless a dependency provides clear value.

## Testing Guidelines

Use Go's built-in `testing` package and `net/http/httptest` for HTTP handlers. Name tests `TestBehavior` and use table-driven cases where helpful. No coverage threshold is established. Prioritize date validation, capacity limits, permissions, cancellation, and overlapping bookings. When persistence is added, test concurrent booking attempts against the database constraints.

## Commit & Pull Request Guidelines

There is no existing commit convention. Use concise, imperative messages such as `Add cabin listing endpoint`. Keep changes focused. Pull requests should describe behavior, link relevant issues, list verification commands, and explain configuration or migration changes. Include request/response examples for API changes.

## Security & Domain Rules

Never commit credentials or guest personal data. Distinguish cabin types from bookable cabin instances. Define booking statuses, timezone handling, and check-in/check-out boundaries before implementing availability rules.
