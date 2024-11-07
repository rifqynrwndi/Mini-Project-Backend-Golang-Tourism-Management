package response

import "tourism-monitoring/entities"


type PlaceResponse struct {
	ID          int    `json:"id"`
	Lokasi      string `json:"lokasi"`
	KapasitasMaks int `json:"kapasitas_maks"`
	JumlahPengunjung int `json:"jumlah_pengunjung"`
}

func FromPlaceEntities(places []entities.Place) []PlaceResponse {
	var responses []PlaceResponse
	for _, place := range places {
		responses = append(responses, PlaceResponse{
			ID:          place.ID,
			Lokasi:      place.Lokasi,
			KapasitasMaks: place.KapasitasMaks,
			JumlahPengunjung: place.JumlahPengunjung,
		})
	}
	return responses
}

func FromPlaceEntity(place entities.Place) PlaceResponse {
	return PlaceResponse{
		ID:          place.ID,
		Lokasi:      place.Lokasi,
		KapasitasMaks: place.KapasitasMaks,
		JumlahPengunjung: place.JumlahPengunjung,
	}
}