package repository

import (
	"booking-service/model"
	"errors"
)

type BookingRepository interface {
	Create(booking model.Booking) (model.Booking, error)
	FindAll() ([]model.Booking, error)
	FindByID(id int64) (model.Booking, error)
	Update(id int64, booking model.Booking) (model.Booking, error)
	Delete(id int64) error
}

type bookingRepositoryImpl struct {
	data            map[int64]model.Booking
	autoIncrementID int64
}

// NewBookingRepository same as constructor in java
func NewBookingRepository() BookingRepository {
	return &bookingRepositoryImpl{
		data:            make(map[int64]model.Booking),
		autoIncrementID: 1,
	}
}
func (repo *bookingRepositoryImpl) Create(booking model.Booking) (model.Booking, error) {
	booking.ID = repo.autoIncrementID
	repo.data[repo.autoIncrementID] = booking
	repo.autoIncrementID++
	return booking, nil
}

func (repo *bookingRepositoryImpl) FindAll() ([]model.Booking, error) {
	var bookings []model.Booking
	for _, b := range repo.data {
		bookings = append(bookings, b)
	}
	return bookings, nil
}

func (repo *bookingRepositoryImpl) FindByID(id int64) (model.Booking, error) {
	b, ok := repo.data[id]
	if !ok {
		return model.Booking{}, errors.New("booking not found")
	}
	return b, nil
}

func (repo *bookingRepositoryImpl) Update(id int64, booking model.Booking) (model.Booking, error) {
	if _, ok := repo.data[id]; !ok {
		return model.Booking{}, errors.New("booking not found")
	}
	booking.ID = id
	repo.data[id] = booking
	return booking, nil
}

func (repo *bookingRepositoryImpl) Delete(id int64) error {
	if _, ok := repo.data[id]; !ok {
		return errors.New("booking not found")
	}
	delete(repo.data, id)
	return nil
}
