package delivery

import "api-airbnb-alta/features/auth"

type UserResponse struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Team     string `json:"team"`
	Role     string `json:"role"`
	Status   string `json:"status"`
	Token    string `json:"token"`
}

func FromCore(data auth.Core, token string) UserResponse {
	return UserResponse{
		ID:       data.ID,
		FullName: data.FullName,
		Email:    data.Email,
		Team:     data.Team,
		Role:     data.Role,
		Status:   data.Status,
		Token:    token,
	}
}
