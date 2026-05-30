# Api-kit CLI

![Go Version](https://img.shields.io/badge/Go-1.25.5-blue.svg)
![License](https://img.shields.io/badge/License-MIT-green.svg)

## Project Name and Description

**Api-kit** is a powerful, user-friendly Command Line Interface (CLI) tool designed to generate predefined API templates across different languages, frameworks, and architectures. 

The primary goal is to increase productivity when studying, practicing, or starting personal, university, or experimental projects. By running a simple command, users can answer a few interactive questions or pass flags to instantly receive a functional project structure ready for development.

## Technology Stack

The CLI is built with simplicity, portability, and excellent UX in mind:

- **Language:** Go (Golang) `1.25.5`
- **CLI Framework:** [Cobra](https://github.com/spf13/cobra) for handling commands, subcommands, and flags.
- **Interactive Prompts:** [huh](https://github.com/charmbracelet/huh) (by Charm) for a highly polished, interactive user experience.
- **Template Engine:** Go's standard `text/template` combined with `embed.FS` to bundle templates directly into a single static binary.

## Project Architecture

The CLI is distributed as a single static binary, meaning users don't need any additional runtimes (like Node or JVM) to generate projects.

### Tool Architecture
- **Embedded Templates:** Templates are stored inside the binary using `embed.FS`, allowing the tool to work completely offline with zero runtime dependencies.
- **Hybrid UX:** Combines Cobra's flag parsing with Huh's interactive forms. If required flags are missing, the tool falls back to interactive prompts to ask only for the missing information.

### Generated Project Architecture
For the MVP, generated APIs follow a **Layered Architecture**:
`Controller -> Service -> Repository`
This separates responsibilities clearly without overengineering, making it easy to understand for developers coming from ecosystems like Java, NestJS, or Node.

## Getting Started

### Installation

**One command (recommended):**
```bash
curl -sSL https://github.com/JuD4Mo/api-kit/releases/latest/download/install.sh | sh
```

**With Go installed:**
```bash
go install github.com/Jud4Mo/Api-kit/cmd/api-kit@latest
```

**Manual:** Download the archive for your OS from [Releases](https://github.com/JuD4Mo/api-kit/releases), extract, and move `akit` to a directory in your `$PATH`.

### Usage

**Interactive Mode:**
Simply run the initialization command and follow the step-by-step prompts:
```bash
akit new
```
*This will ask for the language, framework, architecture, and project name.*

**Flag Mode (Non-interactive):**
Ideal for CI/CD, scripts, or quick generation:
```bash
akit new --lang=go --framework=gin --arch=layered --name=my-api
```

## Project Structure

The repository follows a clean, modular structure:

```text
├── cmd
│   └── api-kit/           # Main entry point of the CLI application
├── internal
│   ├── catalog/           # Available templates, languages, and frameworks definitions
│   ├── cli/               # Cobra commands definitions (root, new)
│   ├── generator/         # Template parsing, variable replacement, and file writing logic
│   ├── prompt/            # Interactive UI forms using 'huh'
│   └── templates/         # Embedded template files (.tmpl) grouped by lang/framework/arch
├── AGENTS.md              # Core AI/Developer guidelines, architectural decisions, and workflows
├── go.mod                 # Go module definition
└── README.md              # Project documentation
```

## Key Features

- **Interactive & Seamless UX:** Beautiful terminal UI for selecting options.
- **Flag-Driven Automation:** Fully scriptable through standard CLI flags.
- **Self-Contained:** Generates projects entirely offline.
- **Dynamic Variable Replacement:** Automatically configures module names, project names, and ports within the generated code.
- **Built-in Best Practices:** MVP templates include functional `Go + Gin + Layered Architecture` examples with auto-generated READMEs.

## Development Workflow

Our development philosophy emphasizes intentionality and simplicity:

1. **Think Before Coding:** Make no silent assumptions. Clarify ambiguities before implementing.
2. **Goal-Driven Execution:** Turn each task into verifiable steps. Define what "working" means before writing code.
3. **Surgical Changes:** Touch only what is necessary to solve the current problem. Avoid speculative flexibility.
4. **Teach While Building:** Code changes should be accompanied by brief, clear explanations of *why* decisions were made.

## Coding Standards

- **Simplicity First:** Write the minimum amount of code needed. Avoid over-abstractions and speculative features.
- **Idiomatic Go:** Follow standard Go formatting, naming conventions, and structure.
- **Clean Diffs:** Do not refactor unrelated code or adjust unrelated formatting. Keep pull requests focused.
- **Orphan Cleanup:** When refactoring, always remove unused imports, variables, and dead code left behind by your changes.

## Testing

Testing is an integral part of the development process. 
- Unit tests are located alongside their respective packages (e.g., `catalog_test.go`, `generator_test.go`).
- Run tests using the standard Go toolchain:
  ```bash
  go test ./...
  ```
- **Verification:** Every bug fix or new feature must include reproducible checks and verification steps. 

## Contributing

We welcome contributions! When contributing, please align with our goals:
1. Ensure your solution is simple and avoids overengineering.
2. Ensure every changed line is directly justified by the issue/feature request.
3. Write small, verifiable solutions and include tests where appropriate.
4. Keep the CLI experience hybrid: support both interactive (`huh`) and flag-based (`cobra`) usages.

## License

This project is licensed under the MIT License.
