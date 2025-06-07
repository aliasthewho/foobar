# Foobar

A RESTful API service built with Go and Gorilla Mux.

## Features

- RESTful API endpoints for CRUD operations
- Structured logging with Zap
- Graceful shutdown
- Health check endpoint
- GitHub Actions CI/CD pipeline

## Getting Started

### Prerequisites

- Go 1.21 or later
- Git

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/foobar.git
   cd foobar
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

### Running the Application

Start the server:
```bash
go run main.go
```

The server will start on `http://localhost:8080`

## API Endpoints

- `GET /health` - Health check
- `GET /examples` - List all examples
- `GET /examples/{id}` - Get a specific example
- `POST /examples` - Create a new example
- `PUT /examples/{id}` - Update an example
- `PATCH /examples/{id}` - Partially update an example
- `DELETE /examples/{id}` - Delete an example

## Development

### Running Tests

```bash
go test -v ./...
```

### Linting

```bash
golangci-lint run
```

## CI/CD

This project uses GitHub Actions for CI/CD. The workflow includes:

- Running tests
- Building the application
- Running linters

The workflow runs on every push to the `main` branch and on pull requests.

## License

[MIT](LICENSE)
