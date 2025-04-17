package main

import (
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"github.com/yokumar9780/booking-service/controller"

	_ "github.com/volvogroup/booking-service/docs" // This registers the Swagger docs
	"github.com/yokumar9780/booking-service/repository"
	"github.com/yokumar9780/booking-service/routes"
	"github.com/yokumar9780/booking-service/service"
)

// @title Booking API
// @version 1.0
// @description API for managing bookings
// @host localhost:8080
// @BasePath /
func main() {

	repo := repository.NewBookingRepository()
	svc := service.NewBookingService(repo)
	ctrl := controller.NewBookingController(svc)

	r := routes.SetupRouter(ctrl)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
