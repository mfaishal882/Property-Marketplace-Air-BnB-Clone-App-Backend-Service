package delivery

import (
	"api-airbnb-alta/features/booking"
	"time"
)

type InsertRequest struct {
	CheckinDate  string `json:"checkin_date" form:"checkin_date"`
	CheckoutDate string `json:"checkout_date" form:"checkout_date"`
	PropertyID   uint   `json:"property_id" form:"property_id"`
	UserID       uint   `json:"user_id" form:"user_id"`
}

var layout2 = "2006-01-02"

func toCore(bookingInput InsertRequest) booking.Core {
	checkin, _ := time.Parse(layout2, bookingInput.CheckinDate)
	checkout, _ := time.Parse(layout2, bookingInput.CheckoutDate)
	bookingCoreData := booking.Core{
		CheckinDate:  checkin,
		CheckoutDate: checkout,
		PropertyID:   bookingInput.PropertyID,
		UserID:       bookingInput.UserID,
	}
	return bookingCoreData
}
