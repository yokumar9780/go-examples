package main

import (
	"booking-service/config"
	"booking-service/controller"
	_ "booking-service/docs" // This registers the Swagger docs
	"booking-service/middleware"
	"booking-service/repository"
	"booking-service/routes"
	"booking-service/service"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Booking API
// @version 1.0
// @description API for managing bookings
// @host localhost:8080
// @BasePath /
func main() {
	config.InitConfig()
	middleware.SetupJWKS()
	repo := repository.NewBookingRepository(config.DB)
	svc := service.NewBookingService(repo)
	ctrl := controller.NewBookingController(svc)

	r := routes.SetupRouter(ctrl)
	r.GET("/swagger/*", echoSwagger.WrapHandler)

	r.Logger.Fatal(r.Start(":8080"))

}
