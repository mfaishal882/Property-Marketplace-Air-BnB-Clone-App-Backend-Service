package repository

import (
	"api-airbnb-alta/features/propertyImage"
	"errors"

	"gorm.io/gorm"
)

type propertyImageRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) propertyImage.RepositoryInterface {
	return &propertyImageRepository{
		db: db,
	}
}

// Create implements user.Repository
func (repo *propertyImageRepository) Create(input propertyImage.Core) error {
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

// GetAll implements user.Repository
func (repo *propertyImageRepository) GetAll() (data []propertyImage.Core, err error) {
	var results []PropertyImage

	tx := repo.db.Find(&results)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(results)
	return dataCore, nil
}

// GetAll with search by name implements user.Repository
func (repo *propertyImageRepository) GetAllWithSearch(query string) (data []propertyImage.Core, err error) {
	var results []PropertyImage

	tx := repo.db.Where("title LIKE ?", "%"+query+"%").Find(&results)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(results)
	return dataCore, nil
}

// GetById implements user.RepositoryInterface
func (repo *propertyImageRepository) GetById(id int) (data propertyImage.Core, err error) {
	var result PropertyImage

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

// Update implements user.Repository
func (repo *propertyImageRepository) Update(input propertyImage.Core, id int) error {
	resultGorm := fromCore(input)
	var result PropertyImage
	tx := repo.db.Model(&result).Where("ID = ?", id).Updates(&resultGorm) // proses update
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("update failed")
	}
	return nil
}

// Delete implements user.Repository
func (repo *propertyImageRepository) Delete(id int) error {
	var result PropertyImage
	tx := repo.db.Delete(&result, id) // proses delete
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("delete failed")
	}
	return nil
}

func (repo *propertyImageRepository) FindUser(email string) (result propertyImage.Core, err error) {
	var data PropertyImage
	tx := repo.db.Where("email = ?", email).First(&data)
	if tx.Error != nil {
		return propertyImage.Core{}, tx.Error
	}

	result = data.toCore()

	return result, nil
}
