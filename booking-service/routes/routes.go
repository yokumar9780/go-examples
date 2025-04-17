package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yokumar9780/booking-service/controller"
)

func SetupRouter(controller *controller.BookingController) *gin.Engine {
	r := gin.Default()
	controller.RegisterRoutes(r)
	return r
}
