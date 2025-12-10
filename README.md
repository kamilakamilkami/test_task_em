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
```

Services:

API: http://localhost:8080

Swagger: http://localhost:8080/swagger/index.html

Postgres: localhost:5432

## 📘 Swagger Documentation

The project automatically generates Swagger docs using swag.

To regenerate docs manually:

```bash
swag init -g cmd/app/main.go
```

## Access UI after starting:

➡ http://localhost:8080/swagger/index.html

## 🧩 API Endpoints

### 🔹 Subscriptions

| Method | Endpoint              | Description                   |
|--------|------------------------|-------------------------------|
| POST   | `/subscriptions/`      | Create a subscription         |
| GET    | `/subscriptions/`      | List all subscriptions        |
| GET    | `/subscriptions/{id}`  | Get subscription by ID        |
| PUT    | `/subscriptions/{id}`  | Update subscription           |
| DELETE | `/subscriptions/{id}`  | Delete subscription           |
| GET    | `/subscriptions/sum`   | Get total spending with filters |

## 🗄 Database Migration
# Migrations run automatically on service startup:
```
RunMigrations(db)
```
# Creates table:
```
subscriptions (
  id SERIAL PRIMARY KEY,
  service_name TEXT,
  price INTEGER,
  user_id UUID,
  start_date DATE,
  end_date DATE
)
```

## 🧪 Example Create Request
```
{
  "service_name": "Netflix",
  "price": 3500,
  "user_id": "1b4e28ba-2fa1-11d2-883f-0016d3cca427",
  "start_date": "01-2024",
  "end_date": "03-2024"
}
```

## ⚙️ Build Without Docker
```
go mod download
go run cmd/app/main.go
```
### 👩‍💻 Author
## Kamila N.
# Backend Developer (Go)

