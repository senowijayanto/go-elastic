## Overview
The `app` folder contains the main application source code, including models, handlers, services, and configurations.

## Structure
```
├── config             # Configuration files (e.g., database connection)
├── handlers           # HTTP handlers
├── migrations         # Database migrations
├── models             # Database models
├── repositories       # Data access layer
├── services           # Business logic layer
├── go.mod             # Go module file
├── go.sum             # Go module dependencies
└── main.go            # Application entry point
```

## Running the Application
1. Install dependencies:
   ```bash
   go mod tidy
   ```
2. Run the application:
   ```bash
   go run main.go
   ```
