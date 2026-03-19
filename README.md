# REST API (Go + Gin)

A simple REST API built with Go and [Gin](https://github.com/gin-gonic/gin) for managing events in memory.

## Features

- `GET /events` to list all events
- `POST /events` to create a new event
- Request validation for required fields (`name`, `description`, `location`)
- In-memory event storage (no database yet)

## Tech Stack

- Go
- Gin (`github.com/gin-gonic/gin`)

## Project Structure

```text
.
|-- main.go
|-- models/
|   `-- event.go
|-- go.mod
`-- go.sum
```

## Prerequisites

- Go installed (version in `go.mod`)

## Getting Started

1. Install dependencies:

   ```bash
   go mod tidy
   ```

2. Run the server:

   ```bash
   go run main.go
   ```

3. Server starts on:

   ```text
   http://localhost:8080
   ```

## API Endpoints

### 1) Get all events

- **Method:** `GET`
- **Path:** `/events`

Example:

```bash
curl http://localhost:8080/events | jq
```

### 2) Create an event

- **Method:** `POST`
- **Path:** `/events`
- **Content-Type:** `application/json`

Example request:

```bash
curl -X POST http://localhost:8080/events \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Go Meetup",
    "description": "Learning REST APIs with Gin",
    "location": "Bengaluru"
  }' | jq
```

Example response (`201 Created`):

```json
{
  "id": 1,
  "name": "Go Meetup",
  "description": "Learning REST APIs with Gin",
  "location": "Bengaluru",
  "dateTime": "2026-03-19T12:34:56.789Z",
  "userId": 1
}
```

## Validation Rules

For `POST /events`, these fields are required:

- `name`
- `description`
- `location`

If validation fails, the API returns `400 Bad Request` with an error message.

## Notes

- Data is stored in memory (`models/events` slice), so restarting the app clears all events.
- `id`, `dateTime`, and `userId` are set by the server in `Save()`.

## Next Improvements

- Add `GET /events/:id`
- Add update and delete endpoints
- Persist data with a database (PostgreSQL/MySQL/SQLite)
- Add tests for handlers and model logic
