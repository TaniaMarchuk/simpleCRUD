# Simple Go CRUD API

A production-ready boilerplate for a RESTful API server written in Golang. This project demonstrates best practices for building scalable Go applications, including structured logging, robust configuration, database migrations, and comprehensive testing.

## Table of Contents
- [Features](#features)
- [Architecture](#architecture)
- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Configuration](#configuration)
  - [Installation](#installation)
  - [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
- [Database Schema](#database-schema)
- [Development](#development)
  - [Testing](#testing)
  - [Migrations](#migrations)
- [License](#license)

## Features
- **Clean Architecture**: Decoupled store, model, and server layers.
- **RESTful Design**: Standard HTTP methods and status codes.
- **PostgreSQL Integration**: High-performance database operations using `lib/pq`.
- **Structured Logging**: Context-aware logging with `logrus`.
- **Schema Management**: SQL-based migrations for versioned database changes.
- **Validation**: Input validation using `ozzo-validation`.
- **Security**: Password hashing with `bcrypt`.

## Architecture
The project follows a modular structure where concerns are separated:
- **Server**: Handles HTTP routing and middleware.
- **Store**: Abstracted database layer with repository pattern.
- **Models**: Defines data structures and business rules (e.g., validation).

## Tech Stack
- **Language**: Go 1.16+
- **Router**: [gorilla/mux](https://github.com/gorilla/mux)
- **Database**: PostgreSQL
- **Logging**: [logrus](https://github.com/sirupsen/logrus)
- **Validation**: [ozzo-validation](https://github.com/go-ozzo/ozzo-validation)
- **Config**: TOML ([BurntSushi/toml](https://github.com/BurntSushi/toml))

## Project Structure
```text
.
├── cmd/apiserver          # Entry point (main.go)
├── configs/               # TOML configuration files
├── internal/app/
│   ├── apiserver/         # API logic & routing
│   ├── model/             # Domain entities & validation
│   └── store/             # Repository layer & DB connection
├── migrations/            # SQL migration files
├── Makefile               # Build and test shortcuts
└── README.md
```

## Getting Started

### Prerequisites
- Go 1.16 or higher installed.
- PostgreSQL server instance.
- `golang-migrate` (optional, for manual migration management).

### Configuration
Update `configs/apiserver.toml` with your environment settings:
```toml
bind_addr = ":8080"
log_level = "debug"

[store]
database_url = "host=localhost user=postgres password=password dbname=restapi_dev sslmode=disable"
```

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/TaniaMarchuk/simpleCRUD.git
   cd simpleCRUD
   ```
2. Download dependencies:
   ```bash
   go mod download
   ```

### Running the Application
Use the `Makefile` to build and run:
```bash
make          # Compiles the binary
./apiserver   # Starts the server
```

## API Endpoints
| Method | Endpoint | Description | Status |
|--------|----------|-------------|--------|
| GET    | `/hello` | Health check / Greeting | Done |
| POST   | `/users` | Create a new user | Planned |

## Database Schema
The primary table is `users`, created via migrations:
- `id`: serial (Primary Key)
- `email`: varchar (Required, Unique)
- `encrypted_password`: varchar (Required)

## Development

### Testing
We use `testing` and `testify/assert` for validation.
```bash
make test
```

### Migrations
Migrations are stored in `/migrations`. They follow the `up` and `down` pattern.
To apply migrations (example using `migrate` CLI):
```bash
migrate -path migrations -database "postgres://localhost/restapi_dev?sslmode=disable" up
```

## License
MIT License. Feel free to use this for your own projects!
