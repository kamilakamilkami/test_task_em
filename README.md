# 📦 Subscription Service (Go + Gin + PostgreSQL + Swagger)

A production-ready REST API service for managing user subscriptions.  
Built using **Go**, **Gin**, **PostgreSQL**, **Docker**, and **Swagger** documentation.

---

## 🚀 Features

- Create, update, delete user subscriptions  
- Filter subscriptions and calculate total spending  
- PostgreSQL migrations on startup  
- Fully containerized (Docker + docker-compose)  
- Auto-generated Swagger documentation  
- Clean architecture (Handlers → Service → Repository)

---

## 📁 Project Structure
/cmd/app → application entry point
/internal/config → configuration
/internal/db → migrations and DB initialization
/internal/subscription
dto.go → request DTO
handler.go → HTTP handlers
model.go → database/model structs
repository.go → database operations
service.go → business logic
/docs → Swagger files

---

## 🛠 Technologies

| Tech | Description |
|------|-------------|
| Go 1.22 | Main language |
| Gin | HTTP framework |
| PostgreSQL 16 | Database |
| Docker | Containerization |
| Swagger / Swaggo | API documentation |
| Clean Architecture | Separation of layers |

---

## 🔧 Environment Variables

Create `.env` file:
DB_DSN=postgres://postgres:0000@db:5432/subscriptions?sslmode=disable

---

## 🐳 Running with Docker

Make sure Docker is installed, then run:

```bash
docker compose up --build

Services:

API: http://localhost:8080

Swagger: http://localhost:8080/swagger/index.html

Postgres: localhost:5432
