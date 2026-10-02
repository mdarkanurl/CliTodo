# CliTodo

## Description

A simple command-line todo manager built with Go. CliTodo lets users add, list, edit, toggle, and delete todos directly from the terminal, with persistent JSON file storage.

Key capabilities:
- Add, edit, delete, and toggle todos via CLI flags.
- Tabular todo listing with completion status and timestamps.
- Persistent storage in a local JSON file with generic storage helper.
- Zero-config setup with standard library flag parsing.

## Features

- Add Todo — Create a new todo with `-add "title"`.
- List Todos — Display all todos in a formatted table with `-list`.
- Toggle Complete — Mark a todo as complete/incomplete with `-toggle <index>`.
- Edit Todo — Update a todo title with `-edit <id:new_title>`.
- Delete Todo — Remove a todo by index with `-del <index>`.
- Persistent Storage — Todos are saved to `todos.json` and loaded on every run.
- Validation — Index bounds checking for edit, toggle, and delete operations.

## Tech Stack

- Language: Go 1.27
- CLI Parsing: Standard library `flag`
- Storage: `encoding/json` + `os` file I/O
- UI: github.com/aquasecurity/table for tabular output
- Dev Tooling: `go build`, `go run`, `go mod`

## Architecture

The application is a small modular CLI app with separation by responsibility. Each file owns a single concern with no external dependencies except table rendering.

Core modules:
- `main.go` — Entry point, wires storage loading, flag execution, and saving.
- `command.go` — CLI flag definitions (`add`, `edit`, `del`, `toggle`, `list`) and command dispatch.
- `todo.go` — `Todo` / `Todos` types with add, delete, toggle, edit, validate, and print logic.
- `storage.go` — Generic `Storage[T]` helper for JSON `Load` / `Save` to file.
- `todos.json` — Local JSON file used as the database.

## Setup Instructions

### Prerequisites
- Go 1.27+

### Steps

1. Clone the repo:
   ```bash
   git clone <repo-url>
   cd CliTodo
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Build the binary:
   ```bash
   go build -o CliTodo .
   ```

4. Run commands:
   ```bash
   ./CliTodo -list
   ```

## Usage

- `./CliTodo -add "Buy groceries"` — Add a new todo.
- `./CliTodo -list` — List all todos.
- `./CliTodo -toggle 0` — Toggle complete status for todo at index 0.
- `./CliTodo -edit 0:New_title` — Edit title of todo at index 0.
- `./CliTodo -del 0` — Delete todo at index 0.

## Contributing

Contributions, issues, and feature requests are welcome.

## License

MIT License.
