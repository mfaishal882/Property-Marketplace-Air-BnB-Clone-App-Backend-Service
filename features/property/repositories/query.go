package repositories

import (
	property "api-airbnb-alta/features/property"
	"errors"

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

	tx2 := repo.db.Model(&user).Where("id = ?", input.UserID).Update("isHosting", "Yes")
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
