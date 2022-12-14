package repository

import (
	"api-airbnb-alta/features/comment"
	"errors"

	"gorm.io/gorm"
)

type commentRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) comment.RepositoryInterface {
	return &commentRepository{
		db: db,
	}
}

// Create implements user.Repository
func (repo *commentRepository) Create(input comment.Core) error {
	var property Property
	userGorm := fromCore(input)
	tx := repo.db.Create(&userGorm) // proses insert data
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}

	tx1 := repo.db.First(&property, input.PropertyID)
	if tx1.Error != nil {
		return tx1.Error
	}
	avg := (property.RatingAverage + input.Rating) / 2
	tx2 := repo.db.Model(&property).Where("id = ?", input.PropertyID).Update("rating_average", avg)
	if tx2.Error != nil {
		return tx.Error
	}
	return nil
}

// GetAll implements user.Repository
func (repo *commentRepository) GetAll() (data []comment.Core, err error) {
	var results []Comment

	tx := repo.db.Find(&results)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(results)
	return dataCore, nil
}

// GetAll with search by name implements user.Repository
func (repo *commentRepository) GetAllWithSearch(query string) (data []comment.Core, err error) {
	var results []Comment

	tx := repo.db.Where("title LIKE ?", "%"+query+"%").Find(&results)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(results)
	return dataCore, nil
}

// GetById implements user.RepositoryInterface
func (repo *commentRepository) GetById(id int) (data comment.Core, err error) {
	var result Comment

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
func (repo *commentRepository) Update(input comment.Core, id int) error {
	resultGorm := fromCore(input)
	var result Comment
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
func (repo *commentRepository) Delete(id int) error {
	var result Comment
	tx := repo.db.Delete(&result, id) // proses delete
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("delete failed")
	}
	return nil
}

func (repo *commentRepository) FindUser(email string) (result comment.Core, err error) {
	var data Comment
	tx := repo.db.Where("email = ?", email).First(&data)
	if tx.Error != nil {
		return comment.Core{}, tx.Error
	}

	result = data.toCore()

	return result, nil
}
