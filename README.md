# PinGo

PinGo is a monitoring web application built with Go.

The project is currently in its early development stage. The first goal is to build a clean and secure foundation for a monitoring platform that can later include authentication, dashboards, service checks, status tracking and alerting.

## Current Features

- Go HTTP server
- Responsive landing page
- Static HTML and CSS
- Basic routing with `http.ServeMux`
- Local-only server binding on `127.0.0.1:8080`
- Configured HTTP timeouts

## Project Structure

```text
pingo/
├── go.mod
├── main.go
└── static/
    ├── index.html
    └── css/
        └── style.css
```

## Run the Application

Start the development server with:

```bash
go run .
```

The application will listen on:

```text
http://127.0.0.1:8080
```

You can test it with:

```bash
curl http://127.0.0.1:8080
```

## Development Goals

Planned development includes:

- Secure login system
- User authentication
- Session management
- Protected dashboard
- Service monitoring
- Availability checks
- Status history
- Alerting
- Production-ready deployment
- Security hardening

## Security

Security is a core requirement of PinGo.

The application is being developed with the goal of using secure defaults and production-oriented practices, including:

- Explicit HTTP server configuration
- Restricted network exposure
- Request timeouts
- Secure authentication
- Password hashing
- Secure session cookies
- Input validation
- CSRF protection
- TLS termination
- Principle of least privilege

## Status

PinGo is currently under active development.

The current focus is the application foundation and routing. Authentication and monitoring functionality will be implemented incrementally after the core concepts are fully understood.
