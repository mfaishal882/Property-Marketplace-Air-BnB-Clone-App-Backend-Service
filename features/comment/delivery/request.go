package delivery

import "api-airbnb-alta/features/comment"

type InsertRequest struct {
	Title      string  `json:"title" form:"title"`
	Comment    string  `json:"comment" form:"comment"`
	Rating     float64 `json:"rating" form:"rating"`
	PropertyID uint    `json:"property_id" form:"property_id"`
}

type UpdateRequest struct {
	ID         uint    `json:"id" form:"id"`
	Title      string  `json:"title" form:"title"`
	Comment    string  `json:"comment" form:"comment"`
	Rating     float64 `json:"rating" form:"rating"`
	PropertyID uint    `json:"property_id" form:"property_id"`
}

func toCore(i interface{}, userId uint) comment.Core {
	switch i.(type) {
	case InsertRequest:
		cnv := i.(InsertRequest)
		return comment.Core{
			Title:      cnv.Title,
			Comment:    cnv.Comment,
			Rating:     cnv.Rating,
			UserID:     userId,
			PropertyID: cnv.PropertyID,
		}

	case UpdateRequest:
		cnv := i.(UpdateRequest)
		return comment.Core{
			Title:      cnv.Title,
			Comment:    cnv.Comment,
			Rating:     cnv.Rating,
			UserID:     userId,
			PropertyID: cnv.PropertyID,
		}
	}

	return comment.Core{}
}
