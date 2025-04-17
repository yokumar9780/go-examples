package model

type Booking struct {
	ID           int64  `json:"id"`
	CustomerName string `json:"customer_name"`
	Date         string `json:"date"`
	Status       string `json:"status"`
}
