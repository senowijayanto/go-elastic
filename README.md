## Project Overview
This project is a web application developed in Golang, integrated with the ELK (Elasticsearch, Logstash, Kibana) stack for monitoring and logging. The application includes functionalities for managing products and utilizes PostgreSQL as the primary database.

## Directory Structure
```
├── README.md          # Project overview
├── app                # Application source code
│   ├── README.md      # Details about the app folder
│   ├── config         # Configuration files (e.g., database connection)
│   ├── handlers       # HTTP handlers
│   ├── migrations     # Database migrations
│   ├── models         # Database models
│   ├── repositories   # Data access layer
│   ├── services       # Business logic layer
│   ├── go.mod         # Go module file
│   ├── go.sum         # Go module dependencies
│   └── main.go        # Application entry point
└── elk                # ELK stack configuration
    ├── README.md      # Details about the elk folder
    ├── docker-compose.yml  # Docker setup for ELK stack
    ├── elasticsearch  # Elasticsearch configurations
    ├── kibana         # Kibana configurations
    └── logstash       # Logstash configurations
```

## Requirements
- Golang 1.20+
- Docker & Docker Compose
- PostgreSQL

## Running the Application
1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd <repository-name>
   ```
2. Start the ELK stack:
   ```bash
   cd elk
   docker-compose up -d
   ```
3. Run the application:
   ```bash
   cd app
   go run main.go
   ```
4. Access Kibana at [http://localhost:5601](http://localhost:5601).