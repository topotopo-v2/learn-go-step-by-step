# 🐹 Go Learning Journey (Commit-by-Commit)

This repository documents my journey learning Go by building a simple API step by step.
Each commit represents a small, intentional improvement — with explanations captured here.

---

## 📌 Goals

* Learn Go fundamentals through hands-on practice
* Understand backend concepts (HTTP, routing, testing)
* Build clean, scalable, and testable code
* Document thinking process, not just final code

---

## 🧱 Project Overview

A simple REST API for managing albums:

* `GET /albums` → list all albums
* `POST /albums` → create a new album
* `GET /albums/:id` → get album by ID

---

## 📚 Learning Log (by Commit)

### 🟢 Commit — Build Basic Album API with Tests

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

---

**Key Concepts:**

* HTTP routing and handlers (Gin)
* JSON binding and response handling
* Testing HTTP handlers with `httptest`
* Table-driven tests in Go
* Proper HTTP status codes and error handling
* Test isolation (resetting shared state)

**References:**
- https://go.dev/doc/tutorial/web-service-gin
---

**Notes:**

* Current implementation uses global state (`albums`), which is not ideal for scalability
* Future improvement: refactor into layered architecture (handler → service → store)
* Tests simulate real HTTP requests without running a server

---

**Takeaway:**
Built a complete, testable API while learning core Go backend concepts, with a focus on simplicity, correctness, and incremental improvement.


## 🚀 How to Run

```bash
go run main.go
```

---

## 🧪 How to Test

```bash
go test ./...
```

---

## 📈 Future Improvements

* Add persistence (database)
* Add middleware (logging, auth)
* Standardize API responses
* Improve project structure (handler/service/store)

---

## ✍️ Why This Repo Exists

This is not just about building an API —
it’s about **learning in public, step by step**, and documenting the thinking behind each change.

---

## 📚 Learning Sources
- Tutorial: Developing a RESTful API with Go and Gin: https://go.dev/doc/tutorial/web-service-gin
