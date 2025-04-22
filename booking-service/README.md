# Booking Service (Go Spring Boot-Style)

This project demonstrates a Go project structure inspired by Java Spring Boot, using:

* [Echo](https://github.com/labstack/echo): High-performance HTTP web framework for building RESTful APIs in Go.
* **Clean Folder Structure**: Modular, layered project layout inspired by Spring Boot for better maintainability and
  separation of concerns.
* [GORM](https://gorm.io/): Powerful ORM for Go that simplifies database interactions with models and migrations.
* [Swagger](https://github.com/swaggo/echo-swagger): Auto-generates interactive API documentation from code annotations.
* [Validator](https://github.com/go-playground/validator): Robust data validation library for struct-based input
  checking.
* **PostgreSQL**: Reliable open-source relational database used for persisting application data.
*  Dockerfile and Makefile setup
   
---

#### Features

* Clean modular structure inspired by Spring Boot
* RESTful CRUD endpoints using Gin
* Integrated Swagger documentation
* Modular layers (controller, service, repository)
* RESTful CRUD routes
* Manual dependency injection
* PostgreSQL DB storage
* Booking model includes field validation using go-playground/validator.
* JWT Auth

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

#### Use Makefile:

```bash
make build     # to build the Go app
make run       # to run the app
make swag      # to generate Swagger docs
```
