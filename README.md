# 🦅 HawkLens

**HawkLens** is a modern, high-performance, and extensible multi-platform social OSINT and digital data analysis framework built with **Go**.

## 🔥 Features

- **Multi-Platform Support**: Twitter (X), YouTube, Instagram (Planned), Reddit (Planned).
- **Concurrent Scanning**: High-speed data collection using Go routines.
- **Modular Plugin Architecture**: Easily extendable with new platform collectors.
- **CLI & API**: Interaction via professional CLI and robust REST API (FastAPI-inspired Gin implementation).
- **Scalable Storage**: Integration with PostgreSQL, Elasticsearch, and Redis.

## 🚀 Getting Started

### Prerequisites

- Go 1.21+
- PostgreSQL
- Redis
- Elasticsearch

### Installation

```bash
git clone https://github.com/ismailtsdln/HawkLens.git
cd HawkLens
go mod download
```

### Usage

#### CLI

Start a concurrent scan across all platforms:

```bash
go run cmd/hawklens/main.go scan "intelligence"
```

Scan a specific platform:

```bash
go run cmd/hawklens/main.go twitter "OSINT"
go run cmd/hawklens/main.go youtube "DeepMind"
```

#### API

Start the API server:

```bash
# Configuration required in internal/api/server.go
```

## 🏗️ Architecture

HawkLens uses a modular, interface-based architecture:

- `cmd/`: Application entry points.
- `internal/`: Core logic, API, CLI, and database handlers.
- `pkg/`: Publicly exportable packages like the plugin interface.
- `plugins/`: Independent platform collectors.

## 🤝 Contributing

Contributions are welcome! Please read the contribution guidelines for more information.

## 📜 License

This project is licensed under the MIT License.
