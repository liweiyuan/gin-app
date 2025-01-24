# Gin App

Gin App is a web application built using the Gin framework, a fast and lightweight HTTP web framework for Go. This application demonstrates a simple setup with logging, configuration management, and error handling.

## Features

- **Fast HTTP Routing**: Utilizes Gin's high-performance HTTP router.
- **Structured Logging**: Uses Logrus for structured logging with support for log rotation.
- **Configuration Management**: Employs Viper for flexible configuration management.
- **Error Handling**: Includes middleware for centralized error logging and handling.

## Getting Started

### Prerequisites

- Go 1.18 or later
- Git

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/gin-app.git
   cd gin-app
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Run the application:

   ```bash
   go run main.go
   ```

### Configuration

The application uses a `config.yaml` file for configuration. You can customize the application settings such as logging level, format, and output.
