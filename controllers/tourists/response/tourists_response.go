package response

import "tourism-monitoring/entities"

type TouristResponse struct {
	Nama          string `json:"nama"`
	Usia          int    `json:"usia"`
	Asal          string `json:"asal"`
	JenisKelamin  string `json:"jenis_kelamin"`
	TipeWisatawan string `json:"tipe_wisatawan"`
}

func FromTouristEntities(users []entities.User) []TouristResponse {
	var responses []TouristResponse
	for _, user := range users {
		responses = append(responses, TouristResponse{
			Nama:          user.Nama,
			Usia:          user.Usia,
			Asal:          user.Asal,
			JenisKelamin:  user.JenisKelamin,
			TipeWisatawan: user.TipeWisatawan,
		})
	}
	return responses
}

func FromTouristEntity(user entities.User) TouristResponse {
    return TouristResponse{
        Nama:          user.Nama,
        Usia:          user.Usia,
        Asal:          user.Asal,
        JenisKelamin:  user.JenisKelamin,
        TipeWisatawan: user.TipeWisatawan,
    }
}
