package routes

import (
	"booking-service/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter(controller *controller.BookingController) *gin.Engine {
	//r := gin.Default()
	r := gin.New()
	controller.RegisterRoutes(r)
	return r
}
