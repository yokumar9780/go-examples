package controller

import (
	"booking-service/model"
	"booking-service/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BookingController struct {
	bookingService service.BookingService
}

func NewBookingController(bookingService service.BookingService) *BookingController {
	return &BookingController{bookingService}
}

func (c *BookingController) RegisterRoutes(router *gin.Engine) {
	bookingRoutes := router.Group("/bookings")
	{
		bookingRoutes.GET("", c.GetAll)
		bookingRoutes.GET("/:id", c.GetByID)
		bookingRoutes.POST("", c.Create)
		bookingRoutes.PUT("/:id", c.Update)
		bookingRoutes.DELETE("/:id", c.Delete)
	}
}

// GetAll Bookings godoc
// @Summary List all bookings
// @Produce json
// @Success 200 {array} model.Booking
// @Router /bookings [get]
func (c *BookingController) GetAll(ctx *gin.Context) {
	bookings, err := c.bookingService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, bookings)
}

// GetByID Booking godoc
// @Summary Get a booking by ID
// @Produce json
// @Param id path int true "Booking ID"
// @Success 200 {object} model.Booking
// @Failure 404 {object} model.ErrorResponse
// @Router /bookings/{id} [get]
func (c *BookingController) GetByID(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	booking, err := c.bookingService.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, booking)
}

// Create Booking godoc
// @Summary Create a new booking
// @Accept json
// @Produce json
// @Param booking body model.Booking true "Booking JSON"
// @Success 201 {object} model.Booking
// @Failure 400 {object} model.ErrorResponse
// @Router /bookings [post]
func (c *BookingController) Create(ctx *gin.Context) {
	var booking model.Booking
	if err := ctx.ShouldBindJSON(&booking); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//validation
	if err := booking.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"validation_error": err.Error()})
		return
	}
	result, _ := c.bookingService.Create(booking)
	ctx.JSON(http.StatusCreated, result)
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
func (c *BookingController) Update(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	var booking model.Booking
	if err := ctx.ShouldBindJSON(&booking); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := booking.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"validation_error": err.Error()})
		return
	}
	updated, err := c.bookingService.Update(id, booking)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updated)
}

// Delete Booking godoc
// @Summary Delete a booking
// @Produce json
// @Param id path int true "Booking ID"
// @Success 200 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Router /bookings/{id} [delete]
func (c *BookingController) Delete(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	err := c.bookingService.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
