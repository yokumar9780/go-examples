package routes

import (
	"booking-service/controller"
	"booking-service/middleware"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func SetupRouter(controller *controller.BookingController) *echo.Echo {
	e := echo.New()
	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.CORS())
	e.Use(middleware.RequestIDMiddleware)
	e.Use(middleware.LoggingMiddleware)
	e.Use(middleware.JWTAuthMiddleware)
	controller.RegisterRoutes(e)
	return e
}
