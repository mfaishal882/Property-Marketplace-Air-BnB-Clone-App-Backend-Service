package service

import (
	"api-airbnb-alta/features/propertyImage"
	"api-airbnb-alta/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	repo := new(mocks.PropertyImageRepo)
	returnData := []propertyImage.Core{{ID: 1, Title: "alta", ImageUrl: "ahdskjdhkask.jpg", PropertyID: 1}}
	t.Run("Success Get All data", func(t *testing.T) {
		repo.On("GetAll").Return(returnData, nil).Once()
		srv := New(repo)
		response, err := srv.GetAll("query")
		assert.Nil(t, err)
		assert.Equal(t, returnData[0].Title, response[0].Title)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Get All data", func(t *testing.T) {
		repo.On("GetAll").Return(nil, errors.New("failed to get data")).Once()
		srv := New(repo)
		response, err := srv.GetAll("query")
		assert.NotNil(t, err)
		assert.Nil(t, response)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Get data, no data found", func(t *testing.T) {
		repo.On("GetAll").Return(nil, nil).Once()
		srv := New(repo)
		response, err := srv.GetAll("query")
		assert.NotNil(t, err)
		assert.Nil(t, response)
		repo.AssertExpectations(t)
	})
}

// func TestCreate(t *testing.T) {
// 	repo := new(mocks.PropertyImageRepo)
// 	t.Run("Success Add Property Image", func(t *testing.T) {
// 		inputData := propertyImage.Core{Title: "alta", ImageUrl: "ahdskjdhkask.jpg", PropertyID: 1}
// 		c := echo.Context
// 		repo.On("Create", inputData).Return(1, nil).Once()
// 		srv := New(repo)
// 		err := srv.Create(inputData, c)
// 		assert.NoError(t, err)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("Failed Create user, duplicate entry", func(t *testing.T) {
// 		inputRepo := propertyImage.Core{Name: "alta", Email: "alta@mail.id", Password: "qwerty", Phone: "081234", Address: "Jakarta", Role: "user"}
// 		inputData := propertyImage.Core{Name: "alta", Email: "alta@mail.id", Password: "qwerty", Phone: "081234", Address: "Jakarta"}
// 		repo.On("Create", inputRepo).Return(0, errors.New("failed to insert data, error query")).Once()
// 		srv := New(repo)
// 		err := srv.Create(inputData)
// 		assert.NotNil(t, err)
// 		assert.Equal(t, "failed insert data, error query", err.Error())
// 		repo.AssertExpectations(t)
// 	})

/*
		dalam testcase ini kita akan melakukan test saat kondisi validationnya error,
		sehingga pengkondisian (if) validation akan terpenuhi dan akan menjalankan perintah return.
// 		jadi fungsi Create yang ada di repository tidak akan dijalankan dalam test case ini. sooo kita tidak perlu memanggil mock "Create" repo.On
// 	*/
// 	t.Run("Failed Create user, name empty", func(t *testing.T) {
// 		// inputRepo := user.Core{Email: "alta@mail.id", Password: "qwerty", Phone: "081234", Address: "Jakarta", Role: "user"}
// 		inputData := propertyImage.Core{Email: "alta@mail.id", Password: "qwerty", Phone: "081234", Address: "Jakarta"}
// 		// repo.On("Create", inputRepo).Return(0, errors.New("failed to insert data, error query")).Once()
// 		srv := New(repo)
// 		err := srv.Create(inputData)
// 		assert.NotNil(t, err)
// 		repo.AssertExpectations(t)
// 	})
// }

func TestGetById(t *testing.T) {
	repo := new(mocks.PropertyImageRepo)
	returnData := propertyImage.Core{ID: 1, Title: "alta", ImageUrl: "ahdskjdhkask.jpg", PropertyID: 1}
	t.Run("Success Get property image by id ", func(t *testing.T) {
		repo.On("GetById", 1).Return(returnData, nil).Once()
		srv := New(repo)
		response, err := srv.GetById(1)
		assert.NoError(t, err)
		assert.Equal(t, returnData.Title, response.Title)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Get property image by id", func(t *testing.T) {
		PropertyImage := propertyImage.Core{}
		repo.On("GetById", 1).Return(PropertyImage, errors.New("failed to get data by id")).Once()
		srv := New(repo)
		response, err := srv.GetById(1)
		assert.NotNil(t, err)
		assert.Equal(t, PropertyImage.Title, response.Title)
		repo.AssertExpectations(t)
	})
	t.Run("Success Get property image by id, no data found ", func(t *testing.T) {
		PropertyImage := propertyImage.Core{}
		repo.On("GetById", 1).Return(PropertyImage, nil).Once()
		srv := New(repo)
		_, err := srv.GetById(1)
		assert.NotNil(t, err)
		assert.Nil(t, nil)
		repo.AssertExpectations(t)
	})
}

// func TestDelete(t *testing.T) {
// 	repo := new(mocks.PropertyImageRepo)
// 	// returnData := propertyImage.Core{ID: 1, Title: "alta", ImageUrl: "ahdskjdhkask.jpg", PropertyID: 1}
// 	t.Run("Success Get property image by id ", func(t *testing.T) {
// 		repo.On("Delete", 1).Return(nil).Once()
// 		srv := New(repo)
// 		err := srv.Delete(1)
// 		assert.NoError(t, err)
// 		// assert.Equal(t, returnData.Title, response.Title)
// 		repo.AssertExpectations(t)
// 	})

// 	// t.Run("Failed Get property image by id", func(t *testing.T) {
// 	// 	User := propertyImage.Core{}
// 	// 	repo.On("Delete", 1).Return(errors.New("failed to get data by id")).Once()
// 	// 	srv := New(repo)
// 	// 	response, err := srv.GetById(1)
// 	// 	assert.NotNil(t, err)
// 	// 	assert.Equal(t, User.Title, response.Title)
// 	// 	repo.AssertExpectations(t)
// 	// })
// 	// t.Run("Success Delete Property Image", func(t *testing.T) {
// 	// 	repo.On("Delete", 1).Return(nil).Once()
// 	// 	srv := New(repo)
// 	// 	err := srv.Delete(1)
// 	// 	assert.NoError(t, err)
// 	// 	repo.AssertExpectations(t)
// 	// })

// 	// t.Run("Failed Delete Property Image", func(t *testing.T) {
// 	// 	repo.On("Delete", 1).Return(errors.New("failed delete property image, error query")).Once()
// 	// 	srv := New(repo)
// 	// 	err := srv.Delete(0)
// 	// 	assert.NotNil(t, err)
// 	// 	assert.Equal(t, "failed delete property image, error query", err.Error())
// 	// 	repo.AssertExpectations(t)
// 	// })
// }

// func TestUpdate(t *testing.T) {
// 	repo := new(mocks.PropertyImageRepo)
// 	inputData := propertyImage.Core{Title: "alta", ImageUrl: "ahdskjdhkask.jpg", PropertyID: 1}
// 	t.Run("Success update property image", func(t *testing.T) {
// 		repo.On("Update", inputData, 1).Return(nil).Once()
// 		srv := New(repo)
// 		err := srv.Update(inputData, 1)
// 		assert.NoError(t, err)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("Failed Update Property Image", func(t *testing.T) {
// 		repo.On("Update", inputData, 1).Return(errors.New("failed update data, error query")).Once()
// 		srv := New(repo)
// 		err := srv.Update(inputData, 1)
// 		assert.NotNil(t, err)
// 		assert.Equal(t, "failed update data, error query", err.Error())
// 		repo.AssertExpectations(t)
// 	})
// 	// t.Run("Failed Create user, name empty", func(t *testing.T) {
// 	// 	// inputRepo := user.Core{Email: "alta@mail.id", Password: "qwerty", Phone: "081234", Address: "Jakarta", Role: "user"}
// 	// 	inputData := user.Core{Email: "alta@mail.id", Password: "qwerty", Phone: "081234", Address: "Jakarta"}
// 	// 	// repo.On("Create", inputRepo).Return(0, errors.New("failed to insert data, error query")).Once()
// 	// 	srv := New(repo)
// 	// 	err := srv.Create(inputData)
// 	// 	assert.NotNil(t, err)
// 	// 	repo.AssertExpectations(t)
// 	// })
// }

func TestPropertyById(t *testing.T) {
	repo := new(mocks.PropertyImageRepo)
	returnData := propertyImage.Property{ID: 1, PropertyName: "alta", PricePerNight: 1200, Description: "bagus", Address: "disana", City: "tokyo", ContactNumber: "123", Facilities: "gada", PropertyType: "Hotel", RatingAverage: 5, ImageThumbnailUrl: ".mp4", UserID: 1}
	t.Run("Success Get property image by id ", func(t *testing.T) {
		repo.On("GetPropertyById", 1).Return(returnData, nil).Once()
		srv := New(repo)
		response, err := srv.GetPropertyById(1)
		assert.NoError(t, err)
		assert.Equal(t, returnData.PropertyName, response.PropertyName)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Get property by id", func(t *testing.T) {
		Property := propertyImage.Property{}
		repo.On("GetPropertyById", 1).Return(Property, errors.New("failed to get data by id")).Once()
		srv := New(repo)
		response, err := srv.GetPropertyById(1)
		assert.NotNil(t, err)
		assert.Equal(t, Property.PropertyName, response.PropertyName)
		repo.AssertExpectations(t)
	})
	t.Run("Success Get property by id, no data found ", func(t *testing.T) {
		Property := propertyImage.Property{}
		repo.On("GetPropertyById", 1).Return(Property, nil).Once()
		srv := New(repo)
		_, err := srv.GetPropertyById(1)
		assert.NotNil(t, err)
		assert.Nil(t, nil)
		repo.AssertExpectations(t)
	})
}
