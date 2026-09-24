# 📚 Book Shop Backend (Go + Fiber v3 + GORM + PostgreSQL)

A high-performance, robust, and clean-architecture RESTful API for managing a Book Shop backend built using **Go (Golang)**, **Fiber v3**, **GORM ORM**, and **PostgreSQL (Neon DB)**.

---

## 📑 Table of Contents

- [📚 Book Shop Backend (Go + Fiber v3 + GORM + PostgreSQL)](#-book-shop-backend-go--fiber-v3--gorm--postgresql)
  - [📑 Table of Contents](#-table-of-contents)
  - [🚀 Features](#-features)
  - [🏗 Architecture \& Folder Structure](#-architecture--folder-structure)
  - [🧠 Code Logic \& Request Flow](#-code-logic--request-flow)

---

## 🚀 Features

- **Clean Layered Architecture:** Handlers ➡️ Services ➡️ Repositories ➡️ Database Models.
- **Ultra-Fast HTTP Framework:** Built on top of [Fiber v3](https://github.com/gofiber/fiber).
- **ORM & Auto Migration:** Automated schema creation and synchronization with [GORM](https://gorm.io/).
- **UUID Primary Keys:** Secure UUID v4 primary keys with automatic generation via `BeforeCreate` GORM hook.
- **Robust Input Validation:** Strict field validation (Name, Author, Non-negative Price/Pages, UUID checks).
- **Standard RESTful Responses:** Appropriate HTTP status codes (`200 OK`, `201 Created`, `400 Bad Request`, `404 Not Found`, `500 Internal Server Error`).
- **Environment Configuration:** Secure config loading using `.env` via `godotenv`.

---
[ HTTP Client / Postman ]
│ (1) HTTP Request (e.g. POST /api/v1/books)
▼
[ Router (internal/router/router.go) ]
│ (2) Matches route and calls handler
▼
[ Handler (internal/handlers/book_handlers.go) ]
│ (3) Binds request body / validates UUID format
▼
[ Service (internal/service/book_service.go) ]
│ (4) Validates business rules (Price >= 0, Name not empty)
▼
[ Repository (internal/repository/book_repository.go) ]
│ (5) Executes database query via GORM
▼
[ PostgreSQL / Neon DB ]
│ (6) Executes SQL query & returns data
▼
[ Client Response ] (7) Returns standard JSON with HTTP status code

## 🏗 Architecture & Folder Structure

book-shop-backend-go-fiber/
├── cmd/
│ └── main.go # Application entry point: initializes app, registers routes & starts server
├── config/
│ └── config.go # Environment loader: parses .env into Config struct
├── internal/
│ ├── app/
│ │ └── app.go # Dependency injection container & bootstrap logic
│ ├── database/
│ │ └── db.go # PostgreSQL database connection & GORM AutoMigrate
│ ├── handlers/
│ │ └── book_handlers.go # HTTP handlers: parses requests, binds JSON, formats HTTP responses
│ ├── models/
│ │ └── book_model.go # GORM models, database struct definitions & hooks
│ ├── repository/
│ │ └── book_repository.go # Data access layer: executes GORM SQL queries
│ ├── router/
│ │ └── router.go # Route definitions & Fiber v3 endpoint mappings
│ └── service/
│ └── book_service.go # Business logic layer: field validations & service rules
├── .env.example # Example environment variables template
├── .gitignore # Git ignore rules (.env, binaries)
├── API.md # Detailed API endpoint reference & testing payloads
├── go.mod # Go module dependencies
├── go.sum # Go checksums
└── README.md # Project documentation

---

## 🧠 Code Logic & Request Flow

Every incoming HTTP request follows a strict single-responsibility lifecycle:

[ HTTP Client / Postman ]
│ (1) HTTP Request (e.g. POST /api/v1/books)
▼
[ Router (internal/router/router.go) ]
│ (2) Matches route and calls handler
▼
[ Handler (internal/handlers/book_handlers.go) ]
│ (3) Binds request body / validates UUID format
▼
[ Service (internal/service/book_service.go) ]
│ (4) Validates business rules (Price >= 0, Name not empty)
▼
[ Repository (internal/repository/book_repository.go) ]
│ (5) Executes database query via GORM
▼
[ PostgreSQL / Neon DB ]
│ (6) Executes SQL query & returns data
▼
[ Client Response ] (7) Returns standard JSON with HTTP status code
