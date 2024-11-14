package request

import "tourism-monitoring/entities"

type RegisterRequest struct {
	Nama          string `json:"nama"`
	Usia          int    `json:"usia"`
	Asal          string `json:"asal"`
	JenisKelamin  string `json:"jenis_kelamin"`
	TipeWisatawan string `json:"tipe_wisatawan"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	Role		  string `json:"role,omitempty"`
}

func (registerRequest RegisterRequest) ToEntities() entities.User {
	return entities.User{
		Nama:         registerRequest.Nama,
		Usia:         registerRequest.Usia,
		Asal:         registerRequest.Asal,
		JenisKelamin: registerRequest.JenisKelamin,
		TipeWisatawan: registerRequest.TipeWisatawan,
		Email:        registerRequest.Email,
		Password:     registerRequest.Password,
		Role:		  registerRequest.Role,
	}
}
