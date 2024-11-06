package response

import "tourism-monitoring/entities"

type AuthResponse struct {
	ID    int    `json:"id"`
	Nama  string `json:"nama"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func FromEntities(user entities.User) AuthResponse {
	return AuthResponse{
		ID:    user.ID,
		Nama:  user.Nama,
		Email: user.Email,
		Token: user.Token,
	}
}
