package service

import (
	"booking-service/model"
	"booking-service/repository"
)

type BookingService interface {
	GetAll() ([]model.Booking, error)
	GetByID(id int64) (model.Booking, error)
	Create(booking model.Booking) (model.Booking, error)
	Update(id int64, booking model.Booking) (model.Booking, error)
	Delete(id int64) error
}

type bookingServiceImpl struct {
	repo repository.BookingRepository
}

func NewBookingService(repo repository.BookingRepository) BookingService {
	return &bookingServiceImpl{repo}
}

func (service *bookingServiceImpl) GetAll() ([]model.Booking, error) {
	return service.repo.FindAll()
}

func (service *bookingServiceImpl) GetByID(id int64) (model.Booking, error) {
	return service.repo.FindByID(id)
}

func (service *bookingServiceImpl) Create(booking model.Booking) (model.Booking, error) {
	return service.repo.Create(booking)
}

func (service *bookingServiceImpl) Update(id int64, booking model.Booking) (model.Booking, error) {
	return service.repo.Update(id, booking)
}

func (service *bookingServiceImpl) Delete(id int64) error {
	return service.repo.Delete(id)
}
