package factory

import (
	authDelivery "api-airbnb-alta/features/auth/delivery"
	authRepo "api-airbnb-alta/features/auth/repository"
	authService "api-airbnb-alta/features/auth/service"

	userDelivery "api-airbnb-alta/features/user/delivery"
	userRepo "api-airbnb-alta/features/user/repository"
	userService "api-airbnb-alta/features/user/service"

	propertyDelivery "api-airbnb-alta/features/property/delivery"
	propertyRepo "api-airbnb-alta/features/property/repositories"
	propertyService "api-airbnb-alta/features/property/service"

	propertyImageDelivery "api-airbnb-alta/features/propertyImage/delivery"
	propertyImageRepo "api-airbnb-alta/features/propertyImage/repository"
	propertyImageService "api-airbnb-alta/features/propertyImage/service"

	bookingDelivery "api-airbnb-alta/features/booking/delivery"
	bookingRepo "api-airbnb-alta/features/booking/repository"
	bookingService "api-airbnb-alta/features/booking/service"
  
	commentDelivery "api-airbnb-alta/features/comment/delivery"
	commentRepo "api-airbnb-alta/features/comment/repository"
	commentService "api-airbnb-alta/features/comment/service"

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

	propertyRepoFactory := propertyRepo.New(db)
	propertyServiceFactory := propertyService.New(propertyRepoFactory)
	propertyDelivery.New(propertyServiceFactory, e)

	propertyImageRepoFactory := propertyImageRepo.New(db)
	propertyImageServiceFactory := propertyImageService.New(propertyImageRepoFactory)
	propertyImageDelivery.New(propertyImageServiceFactory, e)

	bookingRepoFactory := bookingRepo.New(db)
	bookingServiceFactory := bookingService.New(bookingRepoFactory)
	bookingDelivery.New(bookingServiceFactory, e)

	commentRepoFactory := commentRepo.New(db)
	commentServiceFactory := commentService.New(commentRepoFactory)
	commentDelivery.New(commentServiceFactory, e)

}
