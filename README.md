# GitHub Container Registry (GHCR) CLI

[![Go CI/CD](https://github.com/chethanm99/GitHub-Actions-Container-Registry-CLI/actions/workflows/go-ci.yml/badge.svg)](https://github.com/chethanm99/GitHub-Actions-Container-Registry-CLI/actions/workflows/go-ci.yml)

[![Go Version](https://img.shields.io/badge/Go-1.24.3-blue.svg)](https://go.dev/) [![Cobra](https://img.shields.io/badge/CLI-Cobra-blue)](https://cobra.dev/) [![GitHub Actions](https://img.shields.io/badge/CI/CD-GitHub%20Actions-brightgreen.svg)](https://github.com/features/actions)

A simple and efficient command-line interface (CLI) for interacting with the GitHub Container Registry (ghcr.io).

This tool allows you to quickly list container image tags for any public package hosted on GHCR directly from your terminal.

## Key Features

- **List Image Tags:** Easily view all available tags for a specified container image.
- **Simple Authentication:** Uses a standard GitHub Personal Access Token (PAT) for authentication.
- **Environment Variable Support:** Reads credentials from environment variables (`GITHUB_USER`, `GITHUB_TOKEN`) for secure, scriptable use.
- **Built with Go & Cobra:** A robust and modern CLI application built with industry-standard libraries.

---

## Tech Stack

This project is built with a focus on simplicity and performance using the following core technologies:

- **Language:** [Go](https://go.dev/)
- **CLI Framework:** [Cobra](https://cobra.dev/)
- **CI/CD:** [GitHub Actions](https://github.com/features/actions)

  ---

## Prerequisites

Before using this tool, you will need:

1.  **A GitHub Account:** You must have a GitHub account.
2.  **A GitHub Personal Access Token (Classic):** The tool requires a **Classic** PAT with the `read:packages` scope to authenticate with the GHCR API.
    - You can create a Classic PAT [here](https://github.com/settings/tokens/new?scopes=read:packages&description=GHCR-CLI-Token).

---

## Installation

Currently, the primary way to install and run this tool is from the source code.

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/chethanm99/GitHub-Actions-Container-Registry-CLI.git
    cd GitHub-Actions-Container-Registry-CLI
    ```

2.  **Build the binary:**
    This command compiles the source code and creates an executable file named `ghcr-cli` (or `ghcr-cli.exe` on Windows).
    ```bash
    go build -o ghcr-cli
    ```

3.  **Move the binary to your PATH (Optional):**
    For easy access from anywhere in your terminal, move the compiled binary to a directory in your system's `PATH`.
    ```bash
    # For Linux/macOS
    sudo mv ghcr-cli /usr/local/bin/
    ```

---

## Usage

The tool requires your GitHub username and your Classic Personal Access Token to function. It is **highly recommended** to set these as environment variables for security and convenience.

### 1. Set Environment Variables

```bash
export GITHUB_USER="your-github-username"
export GITHUB_TOKEN="ghp_YourClassicPersonalAccessToken"
```

### 2. Run the list-tags command

Once the environment variables are set, you can run the command by providing the repository name.

```bash
ghcr-cli ghcr list-tags --repo "actions/actions-runner"
```

The expected output will be like: 

```bash
Tags for repository actions/actions-runner:
2.308.0
2.309.0
2.310.0
...
latest
```

Alternatively, you can provide your credentials directly as flags:

```bash
ghcr-cli ghcr list-tags \
  --repo "actions/actions-runner" \
  --user "your-github-username" \
  --token "ghp_YourClassicPersonalAccessToken"
```

---

## CI/CD Pipeline

This project uses GitHub Actions for its Continuous Integration pipeline. The workflow, defined in .github/workflows/go-ci.yml, automatically performs the following on every push or pull request to the main branch:

✅ Builds the application to ensure it compiles correctly.

✅ Tests the codebase to verify functionality.

The current build status can be seen from the badge at the top of this README.


***Thank your for exploring this project !***


