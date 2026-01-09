# 🦅 HawkLens

<p align="center">
  <img src="docs/logo.png" alt="HawkLens Logo" width="300px">
</p>

<p align="center">
  <strong>Advanced Multi-Platform OSINT & Real-time Analytics Framework</strong>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Language-Go-00ADD8?style=for-the-badge&logo=go" alt="Go">
  <img src="https://img.shields.io/badge/Platform-Docker-2496ED?style=for-the-badge&logo=docker" alt="Docker">
  <img src="https://img.shields.io/badge/Architecture-Modular-orange?style=for-the-badge" alt="Modular">
  <img src="https://img.shields.io/badge/License-MIT-green?style=for-the-badge" alt="License">
</p>

---

## 🌟 Overview

HawkLens is a high-performance OSINT (Open Source Intelligence) framework built in Golang. It is designed to concurrently scrape, analyze, and visualize data from major social platforms in real-time. By leveraging a custom worker pool architecture and Server-Sent Events (SSE), HawkLens delivers insights with unprecedented speed and efficiency.

## 🚀 Key Features

- **⚡ High-Concurrency Engine**: Utilizes a managed Worker Pool (Dispatcher) for efficient multi-platform scanning.
- **📡 Real-time Streaming**: Instant data delivery to the dashboard via Server-Sent Events (SSE).
- **🔌 Multi-Platform Support**: Built-in connectors for **Twitter (X)**, **YouTube**, **Reddit**, **Instagram**, and **TikTok**.
- **📊 Advanced Analytics**: Real-time sentiment analysis and topic modeling using a custom NLP pipeline.
- **🐳 Docker Orchestration**: One-click deployment with a full stack including PostgreSQL, Redis, and Elasticsearch.
- **🖥️ Premium Dashboard**: Modern glassmorphism UI with live-streaming data visualizations.
- **📦 Data Export**: Export findings in structured JSON or CSV formats.

## 🛠️ Technology Stack

- **Core**: Go (Golang)
- **Infrastructure**: Docker, Docker Compose
- **Databases**: PostgreSQL (Persistence), Redis (Caching/Rate-limiting), Elasticsearch (Search)
- **Frontend**: Vanilla HTML/JS/CSS, Chart.js
- **CLI**: Cobra, Fatih/Color, TableWriter, ProgressBar

## 🏁 Getting Started

### Quick Start with Docker (Recommended)

```bash
# Clone the repository
git clone https://github.com/ismailtsdln/HawkLens.git
cd HawkLens

# Start the entire stack
docker-compose up --build
```
The dashboard will be available at `http://localhost:8080`.

### Manual Installation

```bash
# Build the binary
go build -o hawklens cmd/hawklens/main.go

# Run a scan from CLI
./hawklens scan "cybersecurity"
```

## 📂 Project Structure

- `cmd/`: Application entry points (CLI & Serve).
- `internal/`: Core business logic.
  - `analytics/`: NLP and data export engines.
  - `api/`: SSE-enabled REST API.
  - `engine/`: Managed worker pool dispatcher.
  - `plugins/`: Platform-specific OSINT collectors.
- `pkg/`: Shared libraries and registry logic.
- `dashboard/`: Premium web frontend assets.
- `docs/`: Project assets and documentation.

## 🤝 Contributing

We welcome contributions! Please check our [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## 📄 License

HawkLens is licensed under the MIT License. See [LICENSE](LICENSE) for details.

---
<p align="center">Built for intelligence. Designed for speed. 🦅</p>
