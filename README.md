# Pingo

Pingo is a lightweight self-hosted Linux server monitoring dashboard built with Go, HTML, CSS, and JavaScript.

The goal of the project is to provide a simple and clean web interface for monitoring important Linux system information without relying on a large external monitoring stack.

## Features

Planned and current features include:

* CPU usage
* Memory usage
* Disk usage
* System information
* Network information
* systemd services
* System logs
* Responsive web dashboard
* Lightweight self-hosted setup

## Tech Stack

* Go
* HTML
* CSS
* JavaScript
* Linux
* systemd

## Project Structure

```text
pingo/
├── cmd/
│   └── pingo/
│       └── main.go
├── internal/
│   └── router/
│       └── router.go
├── web/
│   ├── templates/
│   ├── static/
│   │   ├── css/
│   │   └── js/
│   └── partials/
├── scripts/
├── systemd/
├── go.mod
└── README.md
```

## Architecture

```text
Linux System
     │
     ▼
Go Backend
     │
     ▼
HTTP Server
     │
     ▼
HTML / CSS / JavaScript
     │
     ▼
Pingo Dashboard
```

## Run Locally

```bash
go run ./cmd/pingo
```

Open:

```text
http://localhost:8080
```

