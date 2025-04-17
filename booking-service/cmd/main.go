package main

import (
	"booking-service/config"
	"booking-service/controller"
	"booking-service/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	_ "booking-service/docs" // This registers the Swagger docs
	"booking-service/repository"
	"booking-service/routes"
	"booking-service/service"
)

// @title Booking API
// @version 1.0
// @description API for managing bookings
// @host localhost:8080
// @BasePath /
func main() {
	config.InitConfig()
	repo := repository.NewBookingRepository(config.DB)
	svc := service.NewBookingService(repo)
	ctrl := controller.NewBookingController(svc)

	r := routes.SetupRouter(ctrl)

	// Register logging middleware
	r.Use(gin.Recovery()) // Panic recovery
	//Request ID: adds a unique ID per request for traceability
	r.Use(middleware.RequestIDMiddleware())
	r.Use(middleware.LoggingMiddleware()) // Logging: logs all requests with method, path, status, duration, and request ID
	//CORS: enables cross-origin requests
	r.Use(middleware.CORSMiddleware())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
