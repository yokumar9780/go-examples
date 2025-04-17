# Booking Service (Go Spring Boot-Style)

This project demonstrates a Go project structure inspired by Java Spring Boot, using
the [Gin](https://github.com/gin-gonic/gin) web framework and Swagger for API documentation.

---

#### Features

* Clean modular structure inspired by Spring Boot
* RESTful CRUD endpoints using Gin
* Integrated Swagger documentation
* Easily extensible and production-friendly

#### Project Structure

```bash
booking-service/
├── cmd/
│   └── main.go
├── controller/
├── model/
├── router/
├── service/
├── docs/
├── go.mod
├── go.sum
└── README.md
```

## ▶️ Run the Service

```bash
go run cmd/main.go
```

Then access the API at:
http://localhost:8080/bookings

#### Swagger

Generate Swagger documentation:

```bash
swag init --generalInfo cmd/main.go --output ./docs
```

Open the Swagger UI at:  http://localhost:8080/swagger/index.html

#### Rebuild & Regenerate

Clean, tidy modules, regenerate Swagger docs, and restart the app:

```bash
go clean
go mod tidy
swag init --generalInfo cmd/main.go --output ./docs
go run cmd/main.go
```

