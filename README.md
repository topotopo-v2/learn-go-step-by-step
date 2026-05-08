# 🦫 Go Learning Journey (Commit-by-Commit)

This repository documents my journey learning Go by building a simple API step by step.
Each commit represents a small, intentional improvement — with explanations captured here.

---

## 🐈 Goals

* Learn Go fundamentals through hands-on practice
* Understand backend concepts (HTTP, routing, testing)
* Build clean, scalable, and testable code
* Document thinking process, not just final code

---

## 🐈 Project Overview

A simple REST API for managing albums:

* `GET /albums` → list all albums
* `POST /albums` → create a new album
* `GET /albums/:id` → get album by ID
---
Here’s your service layer commit entry in the same format:

---

## 🐈 Learning Log (by Commit)

### 💙 Commit 04: Introduce Service Layer and Business Logic Orchestration

**Message:** `add service layer to separate business logic from handler and repository`

**Summary:**

Refactored the application to include a service layer that handles business logic and orchestration, improving separation of concerns, testability, and scalability of the codebase.

**What I did:**

* Introduced `service.go` to act as an intermediate layer between handler and repository
* Refactored handler to call service instead of repository directly
* Injected dependencies via constructor (repository → service → handler)
* Created `ServiceI` interface for better abstraction and testability
* Updated unit tests:
  * handler tests now mock service instead of repository
  * added service tests with mocked repository using `testify/mock`

**Key Concepts:**

* Service Layer - A layer responsible for handling business logic and coordinating operations between handler and repository
* Separation of Concerns - Dividing responsibilities across layers (HTTP, business logic, data access)
* Orchestration - Coordinating multiple steps (validation, checks, DB calls) to complete a business operation
* Dependency Injection - Passing dependencies through constructors for better modularity and testability
* Interface Abstraction - Using interfaces (`ServiceI`) to decouple implementation and enable mocking
* Layered Architecture - Structuring code into handler → service → repository for scalability and maintainability

---

## 🐈 Learning Log (by Commit)

### 💙 Commit 03: Add Structured Logging and Request Observability

**Message:** `add a logging system with slog and a middlewear-based logging`

**Summary:**

Introduced a logging system with structured JSON logs, middleware-based request logging to improve debugging, and traceability.

**What I did:**
* Added centralized logger using `slog` with JSON output
* Replaced `fmt.Println` and `log.Fatal` with structured logging
* Implemented request logging middleware:
  * logs method, path, status, latency, client IP
* Added request ID middleware:
  * generates unique request ID per request
  * attaches ID to logs and response headers
* Injected logger via dependency injection (main → router → middleware → db)
* Updated database connector to use structured logging instead of standard logging

**Key Concepts:**
* Structured Logging  - Logging in JSON format for machine-readable and searchable logs
* Middleware Logging - Automatically logging request lifecycle without manual logging in handlers
* Request ID / Correlation ID - Tracking a request across layers for debugging and tracing

### 💙 Commit 02: Add PostgreSQL Integration, Unit Tests, and Clean Project Structure

**Message:** `add postgreSQL integration and update handler, repository and project structure`

**Summary:**

Refactored the Go REST API to introduce PostgreSQL integration, implement unit tests for handler and repository layers,
and restructure the project into a scalable architecture.

**What I did:**

* Integrated PostgreSQL using `database/sql` and `lib/pq`
* Added .env support for managing sensitive configuration, currently only used to store database url
* Created a dedicated database layer (`internal/storage/postgres/connector.go`)
* Refactored repository to use SQL queries for CRUD operations
* Implemented repository pattern with interface for testability
* Updated handler unit tests using Gin + httptest
* Added repository unit tests using `sqlmock`
* Introduced mock-based testing with `testify/mock`
* Restructured project into clean architecture folders:
    * `cmd/` for application entry point
    * `internal/` for domain logic (album, router, db)
* Improved dependency injection (handler → repository → db)

**Key Concepts:**

* **Environment Configuration**
    * Using .env files for managing sensitive values like DB_URL, separating config from code
* **Dependency Injection**
    * Passing repository interfaces into handlers for testability
* **Repository Pattern**
    * Separating database logic from business logic
* **Unit Testing in Go**
    * Handler testing using `httptest`
    * Repository testing using `sqlmock`
* **Mocking**
    * Using `testify/mock` for simulating dependencies
    * Using `sqlmock` for simulating database queries
* **Table-Driven Tests**
    * Structuring multiple test cases cleanly and scalably
* **Clean Architecture**
    * Separation of concerns across layers:
        * HTTP layer (handler)
        * Business/data access layer (repository)
        * Infrastructure layer (DB)

**Project Structure Improvement:**

```text id="structure-final"
cmd/api                → entry point (main.go)
internal/album         → handler, repository, model, tests
internal/router        → route definitions
internal/storage       → database connection setup
```

**References:**

* https://dev.to/unkletayo/setting-up-postgresql-on-macos-a-fresh-start-guide-198j

---

### 💙 Commit 01: Build Basic Album API with Tests

**Message:** `build album API with handlers, tests`

**What I did:**

* Built a simple REST API using Gin:
    * `GET /albums` → list all albums
    * `POST /albums` → create a new album
    * `GET /albums/:id` → get album by ID
* Implemented in-memory data storage using a slice
* Added HTTP handler tests using `httptest`
* Covered both **success and error scenarios** (e.g. album not found)
* Used **table-driven tests** for better scalability
* Improved error handling for invalid JSON requests

**Key Concepts:**

* HTTP routing and handlers (Gin)
* JSON binding and response handling
* Testing HTTP handlers with `httptest`
* Table-driven tests in Go
* Proper HTTP status codes and error handling
* Test isolation (resetting shared state)

**References:**

- https://go.dev/doc/tutorial/web-service-gin

**Notes:**

* Current implementation uses global state (`albums`), which is not ideal for scalability
* Future improvement: refactor into layered architecture (handler → service → store)
* Tests simulate real HTTP requests without running a server

---

## 🐈 How to Run

```bash
go run ./cmd/api
```

---

## 🐈 How to Test

```bash
go test ./...
```

---

## 🐈 Why This Repo Exists

This is not just about building an API —
it’s about **learning in public, step by step**, and documenting the thinking behind each change.

---

## 🐈 Learning Sources

- Tutorial: Developing a RESTful API with Go and Gin: https://go.dev/doc/tutorial/web-service-gin
- Setting Up PostgreSQL on macOS: A Fresh Start
  Guide: https://dev.to/unkletayo/setting-up-postgresql-on-macos-a-fresh-start-guide-198j