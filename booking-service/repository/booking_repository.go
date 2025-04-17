package repository

import (
	"booking-service/model"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type BookingRepository interface {
	Create(booking model.Booking) (model.Booking, error)
	FindAll() ([]model.Booking, error)
	FindByID(id int64) (model.Booking, error)
	Update(id int64, booking model.Booking) (model.Booking, error)
	Delete(id int64) error
}

type bookingRepositoryImpl struct {
	db *gorm.DB
}

// NewBookingRepository same as constructor in java
func NewBookingRepository(db *gorm.DB) BookingRepository {
	return &bookingRepositoryImpl{db: db}
}
func (repo *bookingRepositoryImpl) Create(booking model.Booking) (model.Booking, error) {
	booking.ID = int64(uuid.New().ID())
	err := repo.db.Create(&booking).Error
	log.Info("A booking appears", booking)
	return booking, err
}

func (repo *bookingRepositoryImpl) FindAll() ([]model.Booking, error) {
	var bookings []model.Booking
	err := repo.db.Find(&bookings).Error
	return bookings, err
}

func (repo *bookingRepositoryImpl) FindByID(id int64) (model.Booking, error) {
	var booking model.Booking
	err := repo.db.First(&booking, id).Error
	return booking, err
}

func (repo *bookingRepositoryImpl) Update(id int64, booking model.Booking) (model.Booking, error) {
	booking.ID = id
	err := repo.db.Save(&booking).Error
	return booking, err
}

func (repo *bookingRepositoryImpl) Delete(id int64) error {
	return repo.db.Delete(&model.Booking{}, id).Error
}
