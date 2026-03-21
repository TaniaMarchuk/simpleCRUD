# Simple Go CRUD API

A simple RESTful API server written in Golang, demonstrating basic CRUD operations, database integration with PostgreSQL, and common development patterns. This project was created as a learning exercise to explore Go, REST APIs, JSON handling, and database interactions.

## Features

- **RESTful API**: Built using the `gorilla/mux` router.
- **Database Integration**: Uses PostgreSQL for data storage with `lib/pq` driver.
- **Logging**: Implements structured logging with `logrus`.
- **Configuration**: Managed via TOML files using `BurntSushi/toml`.
- **Testing**: Includes unit tests and integration tests.
- **Migrations**: Database schema management using SQL migrations.

## Tech Stack

- **Language**: Go (Golang)
- **Router**: [gorilla/mux](https://github.com/gorilla/mux)
- **Logger**: [logrus](https://github.com/sirupsen/logrus)
- **Database**: PostgreSQL
- **Config**: TOML

## Getting Started

### Prerequisites

- Go 1.16 or higher
- PostgreSQL

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/TaniaMarchuk/simpleCRUD.git
   cd simpleCRUD
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

### Configuration

The application uses a TOML configuration file. A sample configuration can be found in `configs/apiserver.toml`. Ensure your database credentials are correctly set in the configuration file.

### Building and Running

You can use the provided `Makefile` for common tasks:

- **Build the project**:
  ```bash
  make
  ```

- **Run tests**:
  ```bash
  make test
  ```

- **Start the API server**:
  ```bash
  ./apiserver
  ```

Once the server is running, you can access the health check endpoint:
[http://localhost:8080/hello](http://localhost:8080/hello)

## Project Structure

- `cmd/apiserver`: Main entry point for the application.
- `internal/app/apiserver`: Core API server logic and routing.
- `internal/app/model`: Data models and validations.
- `internal/app/store`: Database interaction and repositories.
- `migrations`: SQL migration files for database schema.
- `configs`: Configuration files.

## License

This project is open-source and available under the MIT License.
