# 🚀 Customer Service API

<div align="center">

[![Go Report Card](https://goreportcard.com/badge/github.com/pusrenk/customer-service)](https://goreportcard.com/report/github.com/pusrenk/customer-service)
[![GoDoc](https://godoc.org/github.com/pusrenk/customer-service?status.svg)](https://godoc.org/github.com/pusrenk/customer-service)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

</div>

A lightning-fast, production-ready customer service API built with Go. Designed for scalability and maintainability, this service provides a robust foundation for managing customer interactions and data.

## ✨ Features

- ⚡ **High Performance**: Built with Echo framework for blazing-fast response times
- 🔒 **Secure**: Built-in security best practices and middleware
- 📊 **Structured Logging**: Comprehensive logging with Zap
- 🔄 **Database Ready**: PostgreSQL with GORM for efficient data management
- ⚙️ **Configurable**: Flexible configuration management with Viper
- 📝 **API Documentation**: Auto-generated Swagger documentation with swaggo
- 🧪 **Test Coverage**: Comprehensive test suite with mockery for mocking
- 🔍 **API Testing**: Easy API testing with Swagger UI

## 🛠️ Tech Stack

| Component | Technology | Description |
|-----------|------------|-------------|
| Framework | [Echo](https://echo.labstack.com/) | High performance, minimalist Go web framework |
| Database | PostgreSQL | Advanced open source database |
| ORM | [GORM](https://gorm.io/) | The fantastic ORM library for Golang |
| Logger | [Zap](https://github.com/uber-go/zap) | Blazing fast, structured, leveled logging |
| Config | [Viper](https://github.com/spf13/viper) | Complete configuration solution |
| Swagger | [swaggo/swag](https://github.com/swaggo/swag) | Auto Swagger documentation generator |
| Testing | [mockery](https://github.com/vektra/mockery) | Mock generation for Go interfaces |

## 📋 Prerequisites

- Go 1.24 or higher
- PostgreSQL
- Make (for using Makefile commands)
- golangci-lint (for linting)
- swag (for Swagger documentation)
- mockery (for generating mocks)

## 🚀 Quick Start

```bash
# Clone the repository
git clone https://github.com/pusrenk/customer-service.git
cd customer-service

# Install dependencies
make deps

# Generate Swagger documentation
make swagger

# Generate mocks
make mocks

# Build and run
make build
make run
```

## 📁 Project Structure

```
.
├── cmd/            # Application entry points
├── configs/        # Configuration files
├── database/       # Database related code
├── log/           # Logging related code
├── docs/          # Swagger documentation
└── mocks/         # Generated mocks for testing
```

## 🛠️ Available Commands

| Command | Description |
|---------|-------------|
| `make build` | Build the project |
| `make run` | Run the application |
| `make test` | Run tests |
| `make coverage` | Generate test coverage report |
| `make lint` | Run linter |
| `make fmt` | Format code |
| `make deps` | Download dependencies |
| `make tidy` | Tidy dependencies |
| `make swagger` | Generate Swagger documentation |
| `make mocks` | Generate mocks for testing |
| `make help` | Show all available commands |

## 📚 API Documentation

The API documentation is automatically generated using swaggo/swag. To view the documentation:

1. Generate the Swagger documentation:
   ```bash
   make swagger
   ```

2. Start the application:
   ```bash
   make run
   ```

3. Access the Swagger UI at: `http://localhost:8080/swagger/index.html`

The documentation includes:
- Detailed API endpoints
- Request/Response schemas
- Authentication requirements
- Example requests and responses

## 🧪 Testing

The project uses mockery for generating mocks, making it easy to write unit tests. To generate mocks:

```bash
make mocks
```

This will create mock implementations of your interfaces in the `mocks` directory.

## ⚙️ Configuration

The application uses Viper for configuration management. Configuration files should be placed in the `configs` directory. The application supports multiple configuration formats (YAML, JSON, TOML, etc.).

## 📝 Logging

The application uses Zap for structured logging. Logs are written to the `log` directory by default.

## 💾 Database

The application uses PostgreSQL with GORM as the ORM. Make sure to set up your database connection details in the configuration files.

## 🤝 Contributing

We welcome contributions! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the GNU General Public License v3.0 - see the [LICENSE](LICENSE) file for details.

---

<div align="center">
Made with ❤️ by Pusrenk Team
</div>