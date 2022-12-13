package factory

import (
	authDelivery "api-airbnb-alta/features/auth/delivery"
	authRepo "api-airbnb-alta/features/auth/repository"
	authService "api-airbnb-alta/features/auth/service"

	userDelivery "api-airbnb-alta/features/user/delivery"
	userRepo "api-airbnb-alta/features/user/repository"
	userService "api-airbnb-alta/features/user/service"

	propertyImageDelivery "api-airbnb-alta/features/propertyImage/delivery"
	propertyImageRepo "api-airbnb-alta/features/propertyImage/repository"
	propertyImageService "api-airbnb-alta/features/propertyImage/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	userRepoFactory := userRepo.New(db)
	// userRepoFactory := userRepo.NewRaw(db)
	userServiceFactory := userService.New(userRepoFactory)
	userDelivery.New(userServiceFactory, e)

	authRepoFactory := authRepo.New(db)
	authServiceFactory := authService.New(authRepoFactory)
	authDelivery.New(authServiceFactory, e)

	propertyImageRepoFactory := propertyImageRepo.New(db)
	propertyImageServiceFactory := propertyImageService.New(propertyImageRepoFactory)
	propertyImageDelivery.New(propertyImageServiceFactory, e)

}
