package model

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type Booking struct {
	ID           int64  `json:"id"`
	CustomerName string `json:"customer_name" validate:"required,min=2"`
	Date         string `json:"date" validate:"required,datetime=2006-01-02"`
	Status       string `json:"status" validate:"required,oneof=confirmed cancelled pending"`
}

func (b *Booking) Validate() error {
	validate := validator.New()
	_ = validate.RegisterValidation("datetime", func(fl validator.FieldLevel) bool {
		_, err := time.Parse(time.DateOnly, fl.Field().String())
		return err == nil
	})
	return validate.Struct(b)
}
