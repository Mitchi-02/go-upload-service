# Upload Service

A file upload service built with Go, featuring JWT authentication and MinIO storage.

## Quick Start

### With Docker

```bash
cd build
docker-compose up -d
```

### Without Docker

```bash
# Start PostgreSQL and MinIO manually
# Set environment variables (see Configuration section)
go mod download
go run main.go
```

## Application URLs

| Service           | URL                   | Description                        |
| ----------------- | --------------------- | ---------------------------------- |
| **API**           | http://localhost:8080 | Main application                   |
| **Swagger UI**    | http://localhost:8081 | Interactive API documentation      |
| **MinIO Console** | http://localhost:9090 | File storage admin (root/password) |
| **PostgreSQL**    | localhost:5432        | Database (upload-service/password) |

## Project Structure

```
upload-service/
├── build/
│   ├── docker-compose.yml     # Docker services configuration
│   └── Dockerfile            # Application container
├── configs/
│   ├── app.go               # Application configuration
│   ├── database.go          # Database configuration
│   └── storage.go           # MinIO storage configuration
├── docs/
│   ├── swagger.yaml         # OpenAPI specification
│   └── README.md           # API documentation
├── pkg/
│   ├── api/
│   │   ├── auth/           # Authentication endpoints
│   │   ├── documents/      # Document management endpoints
│   │   └── users/          # User data models
│   ├── common/             # Shared utilities and models
│   ├── database/           # Database connection and migrations
│   └── middlewares/        # HTTP middlewares (auth, CORS, JSON)
├── main.go                 # Application entry point
```

## Configuration

Set these environment variables (or use defaults):

```bash
# Application
APP_PORT=8080
JWT_SECRET=your-secret-key

# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=upload-service
DB_PASSWORD=password
DB_NAME=main

# MinIO Storage
MINIO_ENDPOINT=localhost:9000
MINIO_ACCESS_KEY=minioadmin
MINIO_SECRET_KEY=minioadmin
MINIO_BUCKET_NAME=upload-service
MINIO_FOLDER_NAME=documents
MINIO_BASE_URL=http://localhost:9000
```
