# Ginstarter
![GolangVersion](https://img.shields.io/github/go-mod/go-version/go-nop/ginstarter)
![License](https://img.shields.io/github/license/go-nop/ginstarter)


Ginstarter is a modern and lightweight boilerplate for building robust and scalable RESTful APIs using the [Gin Gonic](https://github.com/gin-gonic/gin) framework. It is designed to help developers kickstart their web application projects with clean, modular, and maintainable code.

---

## Features
- **Modular Architecture**: Separate layers for controllers, services, repositories, and routes.
- **Environment Configuration**: Easy environment variable management with `.env` support.
- **Database Integration**: Pre-configured with GORM for seamless database operations.
- **Middleware Support**: Preconfigured middleware for logging, error handling, and CORS. (soon)
- **Error Handling**: Unified error response structure.
- **Dependency Injection**: Easy dependency management for scalability.
- **Testing Ready**: Simplified setup for unit and integration tests.

---

## Technologies
- [Gin Gonic](https://github.com/gin-gonic/gin): High-performance HTTP web framework.
- [GORM](https://gorm.io/): ORM library for database interaction.
- [Logrus](https://github.com/sirupsen/logrus): Logging library.
- [Zap](https://github.com/uber-go/zap): High-performance logging library.
- [Testify](https://github.com/stretchr/testify): Testing utilities.
- [Goose](https://github.com/pressly/goose): Database migration tool.

---

## Project Structure
```
ginstarter/
├── cmd/
│   ├── app.go                               # Application initialization
│   └── command.go                          # Command line utilities
├── config/
│   └── config.go                             # Configuration handling
├── internal/
│   ├── routes/                               # Routing logic
│   │   ├── route.go
│   │   ├── middleware.go
│   │   └── user_route.go
│   ├── handler/                              # Handler functions for routes
│   │   ├── v1/                              # API version
│   │   │   ├── users/
│   │   │   │   ├── constructor.go            # Dependency injection
│   │   │   │   ├── view_handler.go           # Request handling for view endpoints
│   │   │   │   └── list_handler.go           # Request handling for list endpoints
│   │   │   └── example/
│   │   │       ├── constructor.go
│   │   │       └── example_handler.go
│   ├── services/                             # Business logic
│   │   ├── service.go                        # Service interface
│   │   └── user_service.go                   # User service implementation
│   ├── repositories/                         # Database interaction
│   │   ├── repository.go                     # Repository interface
│   │   └── user_repository.go                # User repository implementation
│   ├── entity/                               # Data models
│   │   └── user.go
│   ├── utils/                                # Utility functions
│   │   └── response.go
│   └── db/
│       ├── migrations/                       # Database migration files
│       └── connection.go                     # Database connection logic
├── pkg/                                     # Reusable packages
│   └── log/
│       ├── driver/                           # Log driver implementation
│       └── log.go                            # Log configuration
├── docs/                                     # Documentation
│   └── api.md
├── test/                                     # Test cases
│   ├── user_test.go
│   └── post_test.go
├── main.go                                   # Application entry point
├── go.mod
└── go.sum
```

---

## Getting Started

### Prerequisites
- Go 1.23 or higher
- Docker (optional, for running database containers)

### Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/go-nop/ginstarter.git
    cd ginstarter
    ```

2. Install dependencies:
    ```bash
    go mod tidy
    ```

3. Set up your environment variables:
    Create a `.env` file in the root directory and configure the following:
    ```env
    APP_PORT=8080
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=yourusername
    DB_PASSWORD=yourpassword
    DB_NAME=yourdatabase
    ```

4. Run the database migrations:
    ```bash
    go run main.go db:migrate
    ```

5. Run the application:
    ```bash
    go run main.go
    ```

The API server will start at `http://localhost:8080`.

---

## Usage

### Define a New Endpoint
1. Create a new handler in the `handler/{version}` directory.
2. Define the logic for your endpoint in a service layer.
3. Add a new route in the `routes` directory.

### Example: User Management
#### View User
```http
GET /api/v1/users/:id
```
Response:
```json
{
  "status": 404,
  "message": "User not found",
  "data": null
}
```

---

## Customizing the Package Name

To update the package name and repository structure for your project, use the provided `setup.sh` script:

### Steps:

1. Run the script:

   ```bash
   ./setup.sh
   ```

2. Enter the new Go package name when prompted (e.g., `github.com/yourusername/yourproject`).

3. The script will:
   - Update the `go.mod` file.
   - Replace old imports in all `.go` files with the new package name.
   - Update the Git remote origin to the new repository URL.
   - Move the repository folder to match the new project name.

4. Verify the changes and start development!

---

---

## License
Ginstarter is open-source software licensed under the [MIT License](LICENSE).

---

## Contributing
Contributions are welcome! Please fork the repository and submit a pull request with your changes.

---

## Author
Developed and maintained by [go-nop](https://github.com/go-nop).
