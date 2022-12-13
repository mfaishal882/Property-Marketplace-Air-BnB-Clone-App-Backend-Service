package delivery

import "api-airbnb-alta/features/comment"

type DataResponse struct {
	ID         uint    `json:"id"`
	Title      string  `json:"title"`
	Comment    string  `json:"comment"`
	Rating     float64 `json:"rating"`
	UserID     uint    `json:"user_id"`
	PropertyID uint    `json:"property_id"`
}

func fromCore(dataCore comment.Core) DataResponse {
	return DataResponse{
		ID:         dataCore.ID,
		Title:      dataCore.Title,
		Comment:    dataCore.Comment,
		Rating:     dataCore.Rating,
		UserID:     dataCore.UserID,
		PropertyID: dataCore.PropertyID,
	}
}

func fromCoreList(dataCore []comment.Core) []DataResponse {
	var dataResponse []DataResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
