# Go YouTube Video Fetcher API

## Overview
A Go API to fetch the latest videos from YouTube for a given search query, with data stored in a MySQL database.

## Requirements
- Docker
- Docker Compose

## Setup

1. Clone the repository:

   ```bash
   git clone <repo_url>
   cd <repo_name>
   ```

2. Create a `.env` file and set your environment variables:

   ```
   DB_HOST=mysql
   DB_PORT=3306
   DB_USER=root
   DB_PASSWORD=yourpassword
   DB_NAME=yourdbname
   ```

3. Start the application:

   ```bash
   docker-compose up --build
   ```

4. The API will be available at `http://localhost:8080`.


```plaintext
.
├── cmd
│   └── api
│       └── main.go                  # Entry point of the app, starts the server
├── internal
│   ├── controllers
│   │   └── videos_controller.go      # Handles HTTP requests for videos (GET)
│   ├── services
│   │   └── youtube_service.go        # Logic for fetching videos from YouTube
│   ├── database
│   │   ├── queries                   # SQL queries generated by sqlc (for CRUD operations)
│   │   └── migrations                # SQL migration files for setting up tables
│   ├── tasks
│   │   └── video_fetcher.go          # Background task to fetch videos asynchronously
│   ├── server
│   │   ├── routes.go                 # Define API routes (e.g., GET /videos)
│   │   └── server.go                 # Setup and configure HTTP server and middleware
│   ├── models
│   │   └── video.go                  # SQLC generated models (e.g., Video struct)
│   ├── merrors
│   │   ├── conflict_409.go
│   │   ├── constants.go
│   │   ├── downstream_550.go
│   │   ├── forbidden_403.go
│   │   ├── handle_service_errors.go
│   │   ├── internal_server_500.go
│   │   ├── not_found_404.go
│   │   ├── service_unavailable_503.go
│   │   ├── unauthorized_401.go
│   │   └── validation_422.go
│   ├── utils
│   │   └── base_response.go          # Helper functions to standardize API responses
│   ├── config
│   │   └── config.go                 # Configuration handling (API keys, DB connections)
│   └── middleware
│       └── logging_middleware.go     # Request logging middleware
├── Dockerfile                        # Dockerfile for building the app container
├── docker-compose.yml                # Docker setup for local development (DB setup)
├── Taskfile.yml                      # Taskfile for automating tasks (migrations, fetching videos, etc.)
├── go.mod                            # Go module dependencies
├── go.sum                            # Go module dependencies
├── .env                              # Environment variables (API keys, DB connection strings)
├── README.md                         # Instructions on how to run the project
└── sqlc.yaml                         # SQLC config for query generation
```

## API Endpoints
- **GET /videos**: Fetch paginated video data.

## Tasks
To run scheduled tasks, use:

```bash
task run
```

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for more details.