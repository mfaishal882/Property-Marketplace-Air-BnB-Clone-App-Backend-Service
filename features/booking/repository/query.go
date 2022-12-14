package repository

import (
	"api-airbnb-alta/features/booking"
	"errors"

	"gorm.io/gorm"
)

type bookingRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) booking.RepositoryInterface {
	return &bookingRepository{
		db: db,
	}
}

// Create implements booking.RepositoryInterface
func (repo *bookingRepository) Create(input booking.Core) error {
	var properties Property

	tx1 := repo.db.First(&properties, input.PropertyID)
	if tx1.Error != nil {
		return tx1.Error
	}
	input.PricePerNight = properties.PricePerNight
	// input.GrossAmount = properties.PricePerNight
	input.BookingStatus = "Complete_payment"

	userGorm := fromCore(input)
	tx := repo.db.Create(&userGorm) // proses insert data
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}
	return nil
}

// GetAll implements booking.RepositoryInterface
func (repo *bookingRepository) GetAll() (data []booking.Core, err error) {
	var results []Booking

	tx := repo.db.Find(&results)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(results)
	return dataCore, nil
}

// GetById implements booking.RepositoryInterface
func (repo *bookingRepository) GetById(id int) (data booking.Core, err error) {
	var result Booking

	tx := repo.db.First(&result, id)

	if tx.Error != nil {
		return data, tx.Error
	}

	if tx.RowsAffected == 0 {
		return data, tx.Error
	}

	var dataCore = result.toCore()
	return dataCore, nil
}
