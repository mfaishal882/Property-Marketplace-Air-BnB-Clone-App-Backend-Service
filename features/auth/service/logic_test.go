package service

import (
	"api-airbnb-alta/features/auth"
	"api-airbnb-alta/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	repo := new(mocks.AuthRepository)

	t.Run("Success Login", func(t *testing.T) {
		inputData := auth.Core{Email: "Budi@gmail.com", Password: "13123123"}
		returnData := auth.Core{ID: 1, FullName: "Budi", Email: "Budi@gmail.com", Password: "13123123", ProfileImageUrl: ".jpg", Phone: "0899", Gender: "Male", IsHosting: "No"}

		repo.On("FindUser", inputData.Email).Return(returnData, nil).Once()
		srv := New(repo)
		result, token, errLogin := srv.Login(inputData)

		assert.NoError(t, errLogin, "Error is null")
		assert.NotEmpty(t, token, "Token generated.")
		assert.NotEmpty(t, result, "Result data.")
		repo.AssertExpectations(t)
	})

	t.Run("Failed Login - Empty Data", func(t *testing.T) {
		inputData1 := auth.Core{Email: "Budi@gmail.com", Password: ""}
		inputData2 := auth.Core{Email: "Budi@gmail.com", Password: ""}
		inputData3 := auth.Core{Email: "", Password: ""}
		returnData := auth.Core{}

		repo.On("FindUser").Return(returnData, errors.New("Failed process query")).Once()
		srv := New(repo)
		_, _, errLogin1 := srv.Login(inputData1)
		_, _, errLogin2 := srv.Login(inputData2)
		_, _, errLogin3 := srv.Login(inputData3)

		assert.Contains(t, errLogin1, "validate input", "Failed login. Empty Email.")
		assert.Contains(t, errLogin2, "validate input", "Failed login. Empty Password.")
		assert.Contains(t, errLogin3, "validate input", "Failed login. Empty Email and Password.")
		repo.AssertExpectations(t)
	})
}
