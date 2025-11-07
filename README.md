# Shadow CLI

```
 ____  _   _    _    ____   _____        __   ____ _     ___ 
/ ___|| | | |  / \  |  _ \ / _ \ \      / /  / ___| |   |_ _|
\___ \| |_| | / _ \ | | | | | | \ \ /\ / /  | |   | |    | | 
 ___) |  _  |/ ___ \| |_| | |_| |\ V  V /   | |___| |___ | | 
|____/|_| |_/_/   \_|____/ \___/  \_/\_/     \____|_____|___|
```

A powerful command-line interface tool for managing Kubernetes cluster contexts with ease.

## 📋 Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
  - [Context Management](#context-management)
- [Project Structure](#project-structure)
- [Development](#development)
- [License](#license)

## 🎯 Overview

Shadow CLI is a lightweight, user-friendly tool designed to simplify Kubernetes cluster context management. Built with Go and powered by [Cobra](https://github.com/spf13/cobra), Shadow provides an intuitive interface for switching between multiple Kubernetes contexts without the hassle of remembering complex kubectl commands.

## ✨ Features

- **List Contexts**: View all available Kubernetes cluster contexts with visual highlighting of the active context
- **Switch Contexts**: Interactively switch between Kubernetes contexts using a numbered menu
- **Color-Coded Output**: Active context is highlighted in green for easy identification
- **Aliases Support**: Use short aliases for faster command execution (`ls`, `sw`)
- **kubectl Integration**: Seamlessly integrates with your existing kubectl configuration

## 📦 Prerequisites

Before using Shadow CLI, ensure you have:

- **Go 1.23.11** or higher installed
- **kubectl** installed and configured
- Access to one or more Kubernetes clusters with configured contexts

## 🚀 Installation

### From Source

1. Clone the repository:
```bash
git clone https://github.com/LiraGui/shadow.git
cd shadow
```

2. Build the binary:
```bash
go build -o shadow
```

3. (Optional) Move the binary to your PATH:
```bash
sudo mv shadow /usr/local/bin/
```

### Using Go Install

```bash
go install github.com/LiraGui/shadow@latest
```

## 📖 Usage

### Context Management

#### List All Contexts

Display all available Kubernetes cluster contexts. The currently active context is highlighted in green.

```bash
shadow context list
# or use the alias
shadow context ls
```

**Example Output:**
```
Listing Kubernetes cluster contexts:

1. dev-cluster
2. staging-cluster
3. production-cluster (highlighted in green if active)
```

#### Switch Context

Interactively switch to a different Kubernetes context by selecting from a numbered list.

```bash
shadow context switch
# or use the alias
shadow context sw
```

**Example Interaction:**
```
Select the cluster context:

1. dev-cluster
2. staging-cluster
3. production-cluster (highlighted in green if active)

Enter the number of the context to switch to: 2
Switched to context: staging-cluster
```

### Available Commands

```bash
shadow                    # Display help information
shadow context            # Context management commands
shadow context list       # List all contexts (alias: ls)
shadow context switch     # Switch context interactively (alias: sw)
shadow cluster            # Cluster management (coming soon)
```

## 📁 Project Structure

```
shadow/
├── cmd/                          # Command definitions
│   ├── root.go                   # Root command and CLI setup
│   ├── context.go                # Context command group
│   └── cluster.go                # Cluster command group (placeholder)
├── features/                     # Feature implementations
│   ├── context/                  # Context management features
│   │   ├── listClusterContext.go    # List contexts functionality
│   │   └── switchClusterContext.go  # Switch context functionality
│   └── cluster/                  # Cluster features (empty)
├── main.go                       # Application entry point
├── go.mod                        # Go module definition
├── go.sum                        # Go dependencies checksums
├── LICENSE                       # Apache License 2.0
└── README.md                     # This file
```

## 🛠️ Development

### Building from Source

```bash
# Clone the repository
git clone https://github.com/LiraGui/shadow.git
cd shadow

# Install dependencies
go mod download

# Build the project
go build -o shadow

# Run tests (when available)
go test ./...
```

### Dependencies

- [spf13/cobra](https://github.com/spf13/cobra) - CLI framework
- [spf13/pflag](https://github.com/spf13/pflag) - POSIX/GNU-style flags

### Adding New Features

The project follows a modular structure:

1. **Commands** are defined in the `cmd/` directory
2. **Feature implementations** reside in the `features/` directory
3. Each feature should have its own package for better organization

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## 👤 Author

**LiraGui**
- Email: liraguilherme252@gmail.com
- GitHub: [@LiraGui](https://github.com/LiraGui)

## 🙏 Acknowledgments

- Built with [Cobra](https://github.com/spf13/cobra) CLI framework
- Inspired by the need for simpler Kubernetes context management

---

**Note**: This project is under active development. More features and improvements are coming soon!