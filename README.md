# Students API (Production Ready)

A lightweight, **production-ready REST API** built in **GoLang** to manage student records.  
Uses **SQLite** by default, but can easily switch to PostgreSQL/MySQL using a storage interface.  
Includes **CRUD operations**, **structured logging**, **input validation**, **graceful shutdown**, and **configurable environment settings**.

---

## Project Overview

This API is designed to manage student records in a clean and scalable way.  
It follows a modular structure so storage, handlers, and types are **separated for easy maintenance and upgrades**.

**Key Features:-

- Full **CRUD** support (Create, Read, Update, Delete)
- **Input validation** using `go-playground/validator`
- **Structured logging** using Go `slog`
- **Graceful server shutdown** on interrupts
- **Config-driven** using YAML (supports multiple environments)
- Supports **SQLite** by default; interface allows other DBs
- **Error handling** and structured responses for API clients
- **Production-ready setup** with environment configs and logs

---

## Technology Stack

- **Language:** Go 1.22
- **Database:** SQLite3 (default) – switchable to PostgreSQL/MySQL
- **Validation:** `github.com/go-playground/validator/v10`
- **Config Loader:** `github.com/ilyakaznacheev/cleanenv`
- **Logging:** Go `slog` for structured logging

---

## Folder Structure

students-api/
├── cmd/students-api/main.go # App entry point
├── internal/config/ # YAML config loader
├── internal/http/handlers/students/ # API endpoint handlers
├── internal/storage/ # Storage interface
├── internal/storage/sqlite/ # SQLite implementation
├── internal/types/ # Data models
├── internal/utils/responce/ # Standard API response helpers
├── config/ # YAML configuration files (dev, prod)
├── go.mod # Go module dependencies
├── go.sum # Checksums
└── README.md # Documentation


---

## Production Deployment Setup

### 1. Clone Repository

```bash
git clone https://github.com/Dhi390/students-api.git
cd students-api

2. Create Configuration for Production

config/prod.yaml example:

env: "production"
storage_path: "storage/storage.db"
http_server:
  address: ":8082"

    For production, you can use a different DB path or switch to PostgreSQL by implementing the storage interface.

3. Build & Run API Server

# Build binary
go build -o students-api cmd/students-api/main.go

# Run server
./students-api -config config/prod.yaml

Server runs at: http://0.0.0.0:8082
API Endpoints
Method	Endpoint	Description
POST	/api/students	Create new student
GET	/api/students/{id}	Get student by ID
GET	/api/students	Get all students
PUT	/api/students/{id}	Update student by ID
DELETE	/api/students/{id}	Delete student by ID
Postman Usage

To test the API endpoints, use Postman:

    Import Collection

        Create a new collection in Postman.

        Add requests for all endpoints (POST, GET, PUT, DELETE).

    Set Base URL

        Use: http://localhost:8082

    Request Example (Create Student)

POST /api/students
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@example.com",
  "age": 21
}

    Request Example (Get Student by ID)

GET /api/students/1

    Request Example (Update Student)

PUT /api/students/1
Content-Type: application/json

{
  "name": "John Smith",
  "email": "johnsmith@example.com",
  "age": 22
}

    Request Example (Delete Student)

DELETE /api/students/1

    Tip: Save the base URL as an environment variable in Postman for easier updates.

Validation Rules

    name – required

    email – required

    age – required, integer

Invalid input returns structured error:

{
  "status": "error",
  "error": "field Name is required field, field Email is required field"
}

Tools & Methods Used

    Go Modules (go.mod) – dependency management

    SQLite3 driver (github.com/mattn/go-sqlite3) – database connection

    Validator (github.com/go-playground/validator/v10) – request validation

    Cleanenv (github.com/ilyakaznacheev/cleanenv) – YAML config loader

    Slog – structured logging

    HTTP Server – http.Server with graceful shutdown

Production Considerations

    Use environment variables for secrets and DB credentials

    Use systemd service or Docker container for deployment

    Enable HTTPS with reverse proxy (nginx or Caddy)

    Log to files instead of console for production

    Apply DB migrations if switching to PostgreSQL/MySQL

    Implement JWT authentication for secure endpoints in future

Future Improvements

    JWT Authentication & Authorization

    Pagination & Search for student list

    Export data to CSV/PDF

    Rate limiting for endpoints

    Switch to PostgreSQL/MySQL for production-ready scaling

License

MIT License
