package delivery

import "api-airbnb-alta/features/propertyImage"

type DataResponse struct {
	ID         uint   `json:"id"`
	Title      string `json:"title" form:"title"`
	ImageUrl   string `json:"image_url" form:"image_url"`
	PropertyID uint   `json:"property_id" form:"property_id"`
}

func fromCore(dataCore propertyImage.Core) DataResponse {
	return DataResponse{
		ID:         dataCore.ID,
		Title:      dataCore.Title,
		ImageUrl:   dataCore.ImageUrl,
		PropertyID: dataCore.PropertyID,
	}
}

func fromCoreList(dataCore []propertyImage.Core) []DataResponse {
	var dataResponse []DataResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
