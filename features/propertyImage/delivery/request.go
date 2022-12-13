package delivery

import "api-airbnb-alta/features/propertyImage"

type InsertRequest struct {
	Title      string `json:"title" form:"title"`
	ImageUrl   string `json:"image_url" form:"image_url"`
	PropertyID uint   `json:"property_id" form:"property_id"`
}

type UpdateRequest struct {
	ID         uint   `json:"id" form:"id"`
	Title      string `json:"title" form:"title"`
	ImageUrl   string `json:"image_url" form:"image_url"`
	PropertyID uint   `json:"property_id" form:"property_id"`
}

func toCore(i interface{}) propertyImage.Core {
	switch i.(type) {
	case InsertRequest:
		cnv := i.(InsertRequest)
		return propertyImage.Core{
			Title:      cnv.Title,
			ImageUrl:   cnv.ImageUrl,
			PropertyID: cnv.PropertyID,
		}

	case UpdateRequest:
		cnv := i.(UpdateRequest)
		return propertyImage.Core{
			ID:         cnv.ID,
			Title:      cnv.Title,
			ImageUrl:   cnv.ImageUrl,
			PropertyID: cnv.PropertyID,
		}
	}

	return propertyImage.Core{}
}
