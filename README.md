
# Fampay Backend Assignment

This is a backend assignment for Fampay. The assignment is to create a REST API that fetches videos from Youtube and stores them in a database. The API should also provide a way to search the stored videos.

### Prerequisites

- Go (version 1.16 or higher)
- MySQL server
- [Taskfile](https://taskfile.dev/) (for task automation)
- Docker (optional, for containerization)

### Environment Variables

Create a `.env` file in the root directory of the project based on the `.example.env` file provided. Ensure to fill in the necessary values for your environment.


## Taskfile Commands

The following commands can be run using the Taskfile:

- **migrate**: Run database migrations.
  ```bash
  task migrate
  ```

- **run**: Run the Go application.
  ```bash
  task run
  ```

- **fetch_videos**: Fetch videos from YouTube.
  ```bash
  task fetch_videos
  ```

- **build**: Build the Go application.
  ```bash
  task build
  ```

- **clean**: Clean up built binaries.
  ```bash
  task clean
  ```

- **start**: Run migrations and start the server.
  ```bash
  task start
  ```

- **docker_build**: Build the Docker image.
  ```bash
  task docker_build
  ```

- **docker_run**: Run the Docker container.
  ```bash
  task docker_run
  ```

- **docker_stop**: Stop the Docker container.
  ```bash
  task docker_stop
  ```

- **docker_compose_up**: Start services using Docker Compose.
  ```bash
  task docker_compose_up
  ```

- **docker_compose_down**: Stop services using Docker Compose.
  ```bash
  task docker_compose_down
  ```

## API Endpoints

### Fetch Videos

- **GET** `/v1/videos`
  - Retrieves stored video data in a paginated response sorted in descending order of published date-time.

- **POST** `/v1/videos/fetch`
  - Initiates the fetching of the latest videos from YouTube and stores them in the database.

## Error Handling

Errors are handled using the utility functions defined in the `utils.Basetemplate` and the error helpers in the `merrors` folder. When handling errors, provide the context and error message without redefining them constantly.

## Database Setup

This project uses XAMPP for the MySQL database. Ensure that your database is running and configured correctly.

## Instructions

1. Clone the repository and navigate to the project directory.
2. Create the `.env` file from the `.example.env` template.
3. Set up your MySQL database and ensure the connection details in the `.env` file are correct.
4. Run the initial migrations:
   ```bash
   task migrate
   ```
5. Start the server:
   ```bash
   task start
   ```

## Bonus Features

- Support for multiple YouTube API keys to handle quota exhaustion.
- Optional dashboard for viewing stored videos with filtering and sorting options.

## Reference

- [YouTube Data API v3](https://developers.google.com/youtube/v3/getting-started)
- [YouTube Search API Reference](https://developers.google.com/youtube/v3/docs/search/list)



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


## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for more details.
