#!/bin/bash

# Usage: ./create-go-spring-boot-style.sh booking-service

set -e

PROJECT_NAME=$1

if [ -z "$PROJECT_NAME" ]; then
  echo "❌ Project name required!"
  echo "Usage: $0 my-app"
  exit 1
fi

echo "📁 Creating project: $PROJECT_NAME"

mkdir -p $PROJECT_NAME/{cmd,config,controller,service,repository,model,dto,middleware,routes,utils}

cd $PROJECT_NAME

# Initialize go module
go mod init github.com/yokumar9780/$PROJECT_NAME

# Create main.go
cat <<EOF > cmd/main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("🚀 $PROJECT_NAME started")
}
EOF

# Create config placeholder
cat <<EOF > config/config.go
package config

import "fmt"

func InitConfig() {
	fmt.Println("🔧 Config initialized")
}
EOF

# Sample controller
cat <<EOF > controller/booking_controller.go
package controller

import "fmt"

func GetBookings() {
	fmt.Println("📡 Get bookings")
}
EOF

# Sample service
cat <<EOF > service/booking_service.go
package service

import "fmt"

func ListBookings() {
	fmt.Println("💼 Listing bookings")
}
EOF

# Sample repository
cat <<EOF > repository/booking_repository.go
package repository

import "fmt"

func FindAll() {
	fmt.Println("💾 Fetching all bookings")
}
EOF

# Sample model
cat <<EOF > model/booking.go
package model

type Booking struct {
	ID    int
	Title string
}
EOF

# Sample DTO
cat <<EOF > dto/booking_dto.go
package dto

type BookingRequest struct {
	Title string \`json:"title"\`
}
EOF

# Sample middleware
cat <<EOF > middleware/logger.go
package middleware

import "fmt"

func Logger() {
	fmt.Println("📘 Logger middleware")
}
EOF

# Sample routes
cat <<EOF > routes/routes.go
package routes

import "fmt"

func InitRoutes() {
	fmt.Println("📍 Routes initialized")
}
EOF

# Utils
cat <<EOF > utils/logger.go
package utils

import "log"

func Info(msg string) {
	log.Println("[INFO]", msg)
}
EOF

# .env file
cat <<EOF > .env
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=admin
DB_PASSWORD=secret
DB_NAME=bookingdb
EOF

# .gitignore
cat <<EOF > .gitignore
# Go
bin/
*.exe
*.out
*.test

# Vendor
vendor/

# Env
.env

# IDE
.vscode/
.idea/
EOF

# README
cat <<EOF >  README.md
# $PROJECT_NAME

A Go web application with layered architecture inspired by Spring Boot.

## 📁 Structure

- \`cmd\` - App entry point
- \`config\` - Configuration setup
- \`controller\` - Request handlers
- \`service\` - Business logic
- \`repository\` - Data access
- \`model\` - Domain models
- \`dto\` - Request/response structures
- \`middleware\` - Middlewares (auth, logging)
- \`routes\` - Router initialization
- \`utils\` - Common utilities

## 🚀 Getting Started

``$(bash
go run cmd/main.go
)``
EOF