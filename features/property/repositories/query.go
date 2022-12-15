package repositories

import (
	property "api-airbnb-alta/features/property"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type propertyRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) property.RepositoryInterface {
	return &propertyRepository{
		db: db,
	}
}

// Create implements property.RepositoryInterface
func (repo *propertyRepository) Create(input property.Core) error {
	var user User
	userGorm := fromCore(input)
	tx := repo.db.Create(&userGorm) // proses insert data
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}

	tx2 := repo.db.Model(&user).Where("id = ?", input.UserID).Update("is_hosting", "Yes")
	if tx2.Error != nil {
		return tx.Error
	}
	return nil
}

// Delete implements property.RepositoryInterface
func (repo *propertyRepository) Delete(id int) error {
	var property Property
	tx := repo.db.Delete(&property, id) // proses delete
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("delete failed")
	}
	return nil
}

// Update implements property.RepositoryInterface
func (repo *propertyRepository) Update(input property.Core, id int) error {
	userGorm := fromCore(input)
	var property Property
	tx := repo.db.Model(&property).Where("ID = ?", id).Updates(&userGorm) // proses update
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("update failed")
	}
	return nil
}

// GetAll implements property.RepositoryInterface
func (repo *propertyRepository) GetAllWithSearch(queryPropertyName, queryCity, queryPropertyType string) (data []property.Core, err error) {
	var property []Property

	tx := repo.db.Where("property_name LIKE ?", "%"+queryPropertyName+"%").Where(&Property{City: queryCity, PropertyType: queryPropertyType}).Find(&property)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(property)

	// fmt.Println("\n\n isi repo getallwith search ", property)
	// fmt.Println("\n\n isi repo getallwith search datacore ", dataCore)
	return dataCore, nil
}

// GetAll implements property.RepositoryInterface
func (repo *propertyRepository) GetAll() (data []property.Core, err error) {
	var properties []Property

	tx := repo.db.Find(&properties)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(properties)

	// fmt.Println("\n\n isi repo getall ", properties)
	// fmt.Println("\n\n isi repo getall datacore ", dataCore)

	return dataCore, nil
}

// GetById implements property.RepositoryInterface
func (repo *propertyRepository) GetById(id int) (data property.Core, err error) {
	var property Property

	tx := repo.db.First(&property, id)

	if tx.Error != nil {
		return data, tx.Error
	}

	if tx.RowsAffected == 0 {
		return data, tx.Error
	}

	var dataCore = property.toCore()
	return dataCore, nil
}

// GetPropertyImages implements property.RepositoryInterface
func (repo *propertyRepository) GetPropertyImages(id int) (data []property.PropertyImage, err error) {
	var images []PropertyImage
	tx := repo.db.Where("property_id = ?", id).Find(&images)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toPropertyImagesList(images)
	return dataCore, nil
}

// GetPropertyComments implements property.RepositoryInterface
func (repo *propertyRepository) GetPropertyComments(id int) (data []property.Comment, err error) {
	var comments []Comment
	tx := repo.db.Where("property_id = ?", id).Find(&comments)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toPropertyCommentList(comments)
	return dataCore, nil
}

// GetAll implements property.RepositoryInterface
func (repo *propertyRepository) GetAvailability(propertyId uint, checkinDate time.Time, checkoutData time.Time) (result string, err error) {
	var properties []Property
	queryBuilder := fmt.Sprintf("SELECT * FROM bookings WHERE property_id = %d AND '%s' BETWEEN checkin_date AND checkout_date OR '%s' BETWEEN checkin_date AND checkout_date;", propertyId, checkinDate, checkoutData)
	// tx := repo.db.Raw(`
	// 	SELECT * FROM bookings WHERE property_id = @propertyID
	// 	AND @checkinDate BETWEEN checkin_date AND checkout_date
	// 	OR @checkoutDate BETWEEN checkin_date AND checkout_date;
	// `, sql.Named("propertyID", propertyId), sql.Named("checkinDate", checkinDate), sql.Named("checkoutDate", checkoutData)).Find(&properties)

	fmt.Println("\n\n query ", queryBuilder)

	tx := repo.db.Raw(queryBuilder).Find(&properties)

	if tx.Error != nil {
		return "Not Available", tx.Error
	}

	affectedRow := tx.RowsAffected
	fmt.Println("\n\nHasil check availbility, \n Checkin date = ", checkinDate, " \n Checkout date = ", checkoutData, " \n Hasil Row = ", affectedRow)

	if affectedRow == 0 {
		return "Available", nil
	}

	return "Not Available", nil
}
