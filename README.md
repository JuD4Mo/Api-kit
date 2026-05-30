# Api-kit CLI

![Go Version](https://img.shields.io/badge/Go-1.25.5-blue.svg)
![License](https://img.shields.io/badge/License-MIT-green.svg)

**Api-kit** is a CLI tool that scaffolds ready-to-run API projects across different languages, frameworks, and architectures.

## Installation

**One command (recommended):**
```bash
curl -sSL https://github.com/JuD4Mo/api-kit/releases/latest/download/install.sh | sh
```

**With Go installed:**
```bash
go install github.com/JuD4Mo/api-kit/cmd/api-kit@latest
```

**Manual:** Download the archive for your OS from [Releases](https://github.com/JuD4Mo/api-kit/releases), extract, and move `akit` to a directory in your `$PATH`.

## Usage

**Interactive:**
```bash
akit new
```

**Non-interactive (flags):**
```bash
akit new --lang=go --framework=gin --arch=layered --name=my-api
```

**List available stacks:**
```bash
akit list-catalog
```

## Supported Stacks

| Language | Framework | Type | Architecture |
|----------|-----------|------|-------------|
| Go       | Gin       | Monolith | Layered |
| Java     | Spring Boot | Monolith | Layered |

## Project Structure

```
├── cmd/api-kit/           # Main entry point
├── internal
│   ├── catalog/           # Available stacks and validation
│   ├── cli/               # Cobra commands
│   ├── generator/         # Template rendering and file generation
│   ├── prompt/            # Interactive forms (huh) and styling (lipgloss)
│   └── templates/         # Embedded .tmpl files grouped by stack
├── .goreleaser.yaml       # Automated release configuration
├── install.sh             # One-command install script
└── AGENTS.md              # Agent and development guidelines
```

## Development

This project uses a `development` → `main` git flow. Releases are cut from `main` via tags, which trigger GoReleaser to build binaries for Linux, macOS, and Windows.

```bash
# Tests
go test ./...

# Build locally
go build -o akit ./cmd/api-kit
```

## License

MIT
