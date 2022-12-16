package service

import (
	"api-airbnb-alta/features/comment"
	"api-airbnb-alta/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	repo := new(mocks.CommentRepo)
	returnData := []comment.Core{{ID: 1, Title: "alta", Comment: "bagus", Rating: 5, UserID: 1, PropertyID: 1}}
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

func TestGetById(t *testing.T) {
	repo := new(mocks.CommentRepo)
	returnData := comment.Core{ID: 1, Title: "alta", Comment: "bagus", Rating: 5, UserID: 1, PropertyID: 1}
	t.Run("Success Get property image by id ", func(t *testing.T) {
		repo.On("GetById", 1).Return(returnData, nil).Once()
		srv := New(repo)
		response, err := srv.GetById(1)
		assert.NoError(t, err)
		assert.Equal(t, returnData.Title, response.Title)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Get property image by id", func(t *testing.T) {
		PropertyImage := comment.Core{}
		repo.On("GetById", 1).Return(PropertyImage, errors.New("failed to get data by id")).Once()
		srv := New(repo)
		response, err := srv.GetById(1)
		assert.NotNil(t, err)
		assert.Equal(t, PropertyImage.Title, response.Title)
		repo.AssertExpectations(t)
	})
	t.Run("Success Get property image by id, no data found ", func(t *testing.T) {
		PropertyImage := comment.Core{}
		repo.On("GetById", 1).Return(PropertyImage, nil).Once()
		srv := New(repo)
		_, err := srv.GetById(1)
		assert.NotNil(t, err)
		assert.Nil(t, nil)
		repo.AssertExpectations(t)
	})
}
