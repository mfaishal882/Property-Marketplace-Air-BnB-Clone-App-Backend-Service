package delivery

import (
	"api-airbnb-alta/features/booking"
	"time"
)

type DataResponse struct {
	ID            uint      `json:"id"`
	CheckinDate   time.Time `json:"checkin_date" form:"checkin_date"`
	CheckoutDate  time.Time `json:"checkout_date" form:"checkout_date"`
	PricePerNight float64   `json:"price_per_night" form:"price_per_night"`
	GrossAmount   float64   `json:"gross_amount" form:"gross_amount"`
	BookingStatus string    `json:"booking_status" form:"booking_status"`
	UserName      string    `json:"user_name" form:"user_name"`
	PropertyName  string    `json:"property_name" form:"property_name"`
	CreatedAt     time.Time `json:"created_at" form:"created_at"`
}

func fromCore(dataCore booking.Core) DataResponse {
	return DataResponse{
		ID:            dataCore.ID,
		CheckinDate:   dataCore.CheckinDate,
		CheckoutDate:  dataCore.CheckoutDate,
		PricePerNight: dataCore.PricePerNight,
		GrossAmount:   dataCore.GrossAmount,
		BookingStatus: dataCore.BookingStatus,
		UserName:      dataCore.User.FullName,
		PropertyName:  dataCore.Property.PropertyName,
	}
}

func fromCoreList(dataCore []booking.Core) []DataResponse {
	var dataResponse []DataResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
