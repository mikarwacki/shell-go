# shell-go

A lightweight command-line utility written in Go to execute shell commands directly from your terminal. This project serves as a practical implementation of Go's `os/exec` package for system-level command execution.

## Overview

`shell-go` takes a command and its arguments, executes them in the underlying shell, and prints the output to standard out. It's a simple yet powerful tool for anyone looking to understand how Go can interact with the operating system or for use in basic automation scripts.

## Features

*   **Direct Command Execution:** Run any shell command that your system's PATH can resolve.
*   **Standard I/O:** Captures and displays both standard output and standard error from the executed command.
*   **Cross-Platform:** Written in Go, it can be compiled to run on Windows, macOS, and Linux.
*   **Self-Contained:** Compiles to a single, dependency-free binary.

## Getting Started

### Prerequisites

*   You need to have [Go](https://go.dev/doc/install) (version 1.18 or newer) installed on your system.

### Installation & Build

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/mikarwacki/shell-go.git
    cd shell-go
    ```

2.  **Build the executable:**
    ```bash
    go build
    ```
    This command will create a `shell-go` (or `shell-go.exe` on Windows) executable in the project directory.

### Usage

To run a command, execute `shell-go` followed by the command and its arguments.

**Basic Examples:**

```bash
# List files in the current directory
./shell-go ls -la

# Print a message to the console
./shell-go echo "Hello from shell-go!"

# Check the version of git
./shell-go git --version
```

**Running Complex Commands:**

For commands involving pipes (`|`), redirection (`>`), or other shell-specific syntax, it's best to wrap the command in `sh -c` or `bash -c`:

```bash
# List Go files and count them
./shell-go bash -c "ls -l | grep .go"
```

## How It Works

The program takes all command-line arguments provided to it, uses the first argument as the command to execute, and the rest as its parameters. It leverages Go's standard `os/exec` package to create a new process and run the command, capturing its output.
