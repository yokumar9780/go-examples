package controller

import (
	"booking-service/model"
	"booking-service/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type BookingController struct {
	bookingService service.BookingService
}

func NewBookingController(bookingService service.BookingService) *BookingController {
	return &BookingController{bookingService}
}

func (c *BookingController) RegisterRoutes(router *echo.Echo) {
	bookingRoutes := router.Group("/bookings")
	bookingRoutes.GET("", c.GetAll)
	bookingRoutes.GET("/:id", c.GetByID)
	bookingRoutes.POST("", c.Create)
	bookingRoutes.PUT("/:id", c.Update)
	bookingRoutes.DELETE("/:id", c.Delete)

}

// GetAll Bookings godoc
// @Summary List all bookings
// @Produce json
// @Success 200 {array} model.Booking
// @Router /bookings [get]
func (c *BookingController) GetAll(ctx echo.Context) error {
	bookings, err := c.bookingService.GetAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})

	}
	return ctx.JSON(http.StatusOK, bookings)
}

// GetByID Booking godoc
// @Summary Get a booking by ID
// @Produce json
// @Param id path int true "Booking ID"
// @Success 200 {object} model.Booking
// @Failure 404 {object} model.ErrorResponse
// @Router /bookings/{id} [get]
func (c *BookingController) GetByID(ctx echo.Context) error {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	booking, err := c.bookingService.GetByID(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})

	}
	return ctx.JSON(http.StatusOK, booking)
}

// Create Booking godoc
// @Summary Create a new booking
// @Accept json
// @Produce json
// @Param booking body model.Booking true "Booking JSON"
// @Success 201 {object} model.Booking
// @Failure 400 {object} model.ErrorResponse
// @Router /bookings [post]
func (c *BookingController) Create(ctx echo.Context) error {
	var booking model.Booking
	if err := ctx.Bind(&booking); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})

	}
	//validation
	if err := booking.Validate(); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"validation_error": err.Error()})
	}
	result, _ := c.bookingService.Create(booking)
	return ctx.JSON(http.StatusCreated, result)
}

// Update Booking godoc
// @Summary Update a booking
// @Accept json
// @Produce json
// @Param id path int true "Booking ID"
// @Param booking body model.Booking true "Booking JSON"
// @Success 200 {object} model.Booking
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Router /bookings/{id} [put]
func (c *BookingController) Update(ctx echo.Context) error {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	var booking model.Booking
	if err := ctx.Bind(&booking); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	if err := booking.Validate(); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"validation_error": err.Error()})
	}
	updated, err := c.bookingService.Update(id, booking)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, updated)
}

// Delete Booking godoc
// @Summary Delete a booking
// @Produce json
// @Param id path int true "Booking ID"
// @Success 200 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Router /bookings/{id} [delete]
func (c *BookingController) Delete(ctx echo.Context) error {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	err := c.bookingService.Delete(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})

	}
	return ctx.JSON(http.StatusOK, echo.Map{"message": "deleted"})
}
