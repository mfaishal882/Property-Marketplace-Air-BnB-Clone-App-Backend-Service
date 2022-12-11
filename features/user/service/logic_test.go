package service

import (
	"api-airbnb-alta/features/user"
	"api-airbnb-alta/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	repo := new(mocks.UserRepository)
	t.Run("Success get all user", func(t *testing.T) {
		inputRepo := []user.Core{{ID: 1, FullName: "alta", Email: "alta", Password: "alta", Phone: "089213912", Gender: "Male", ProfileImageUrl: "https://preview.keenthemes.com/metronic-v4/theme/assets/pages/media/profile/profile_user.jpg"}}
		repo.On("GetAll").Return(inputRepo, nil).Once()
		srv := New(repo)
		response, err := srv.GetAll("")
		assert.NoError(t, err)
		assert.Equal(t, inputRepo[0].FullName, response[0].FullName)
		repo.AssertExpectations(t)
	})

	t.Run("Failed get all", func(t *testing.T) {
		repo.On("GetAll").Return(nil, errors.New("failed")).Once()
		srv := New(repo)
		response, err := srv.GetAll("")
		assert.NotNil(t, err)
		assert.Nil(t, response)
		repo.AssertExpectations(t)
	})

	t.Run("Success get all with search ", func(t *testing.T) {
		inputRepo := []user.Core{{ID: 1, FullName: "alta", Email: "alta", Password: "alta", Phone: "089213912", Gender: "Male", ProfileImageUrl: "https://preview.keenthemes.com/metronic-v4/theme/assets/pages/media/profile/profile_user.jpg"}}
		repo.On("GetAllWithSearch", "query").Return(inputRepo, nil).Once()
		srv := New(repo)
		response, err := srv.GetAll("query")
		assert.NoError(t, err)
		assert.Equal(t, inputRepo[0].FullName, response[0].FullName)
		repo.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	repo := new(mocks.UserRepository)
	t.Run("Success get by id user", func(t *testing.T) {
		returnData := user.Core{ID: 1, FullName: "alta", Email: "alta", Password: "alta", Phone: "089213912", Gender: "Male", ProfileImageUrl: "https://preview.keenthemes.com/metronic-v4/theme/assets/pages/media/profile/profile_user.jpg"}
		repo.On("GetById", 1).Return(returnData, nil).Once()
		srv := New(repo)
		response, err := srv.GetById(1)
		assert.NoError(t, err)
		assert.Equal(t, returnData.FullName, response.FullName)
		repo.AssertExpectations(t)
	})

	t.Run("Failed get by id user", func(t *testing.T) {
		inputRepo := user.Core{ID: 1, FullName: "alta", Email: "alta", Password: "alta", Phone: "089213912", Gender: "Male", ProfileImageUrl: "https://preview.keenthemes.com/metronic-v4/theme/assets/pages/media/profile/profile_user.jpg"}
		repo.On("GetById", 1).Return(user.Core{}, errors.New("failed")).Once()
		srv := New(repo)
		response, err := srv.GetById(1)
		assert.NotNil(t, err)
		assert.NotEqual(t, inputRepo.FullName, response.FullName)
		repo.AssertExpectations(t)
	})
}

// func TestCreate(t *testing.T) {
// 	repo := new(mocks.UserRepository)
// 	t.Run("Success Create ", func(t *testing.T) {
// 		inputRepo := user.Core{FullName: "alta", Email: "alta", Status: "alta", Password: "alta", Team: "alta", Role: "alta"}
// 		repo.On("Create", inputRepo).Return(nil).Once()
// 		srv := New(repo)
// 		err := srv.Create(inputRepo)
// 		assert.NoError(t, err)
// 		repo.AssertExpectations(t)
// 	})

// 	// t.Run("Failed Create ", func(t *testing.T) {
// 	// 	inputRepo := user.Core{FullName: "alta", Email: "alta", Status: "alta", Password: "alta", Team: "alta", Role: "alta"}
// 	// 	repo.On("Create", inputRepo).Return(errors.New("failed")).Once()
// 	// 	srv := New(repo)
// 	// 	err := srv.Create(inputRepo)
// 	// 	assert.NotNil(t, err)
// 	// 	// assert.Equal(t, "failed", err.Error())
// 	// 	repo.AssertExpectations(t)
// 	// })
// }

// func TestUpdate(t *testing.T) {
// 	repo := new(mocks.UserRepository)
// 	t.Run("Success update ", func(t *testing.T) {
// 		inputRepo := user.Core{FullName: "alta", Email: "alta", Password: "alta", Team: "alta", Role: "alta", Status: "alta"}
// 		inputRepo.Password = ""
// 		repo.On("Update", inputRepo, 1).Return(nil).Once()
// 		srv := New(repo)
// 		err := srv.Update(inputRepo, 1)
// 		assert.NoError(t, err)
// 		repo.AssertExpectations(t)
// 	})

// t.Run("Failed Update ", func(t *testing.T) {
// 	inputRepo := user.Core{FullName: "alta", Email: "alta", Password: "alta", Team: "alta", Role: "alta", Status: "alta"}
// 	repo.On("Update", inputRepo, 1).Return(errors.New("failed")).Once()
// 	srv := New(repo)
// 	err := srv.Update(inputRepo, 1)
// 	assert.NotNil(t, err)
// 	// assert.Equal(t, "failed", err.Error())
// 	repo.AssertExpectations(t)
// })
// }

// func TestDeleteLog(t *testing.T) {
// 	repo := new(mocks.LogRepository)
// 	t.Run("Success delete log", func(t *testing.T) {
// 		repo.On("DeleteLog", 1).Return(nil).Once()
// 		srv := New(repo)
// 		err := srv.DeleteLog(1)
// 		assert.NoError(t, err)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("Failed delete log", func(t *testing.T) {
// 		repo.On("DeleteLog", 1).Return(errors.New("failed")).Once()
// 		srv := New(repo)
// 		err := srv.DeleteLog(1)
// 		assert.NotNil(t, err)
// 		// assert.Equal(t, "failed", err.Error())
// 		repo.AssertExpectations(t)
// 	})
// }

// func TestGetAllLog(t *testing.T) {
// 	repo := new(mocks.LogRepository)
// 	t.Run("Success get all log", func(t *testing.T) {
// 		inputRepo := []log.Core{{ID: 1, Title: "alta", Feedback: "alta", Status: "alta", UserID: 1, MenteeID: 1}}
// 		repo.On("GetAllLog").Return(inputRepo, nil).Once()
// 		srv := New(repo)
// 		response, err := srv.GetAllLog("")
// 		assert.NoError(t, err)
// 		assert.Equal(t, inputRepo[0].Title, response[0].Title)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("Failed get all log", func(t *testing.T) {
// 		repo.On("GetAllLog").Return(nil, errors.New("failed")).Once()
// 		srv := New(repo)
// 		response, err := srv.GetAllLog("")
// 		assert.NotNil(t, err)
// 		assert.Nil(t, response)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("Success get all with search log", func(t *testing.T) {
// 		inputRepo := []log.Core{{ID: 1, Title: "alta", Feedback: "alta", Status: "alta", UserID: 1, MenteeID: 1}}
// 		repo.On("GetAllWithSearchLog", "query").Return(inputRepo, nil).Once()
// 		srv := New(repo)
// 		response, err := srv.GetAllLog("query")
// 		assert.NoError(t, err)
// 		assert.Equal(t, inputRepo[0].Title, response[0].Title)
// 		repo.AssertExpectations(t)
// 	})
// }
