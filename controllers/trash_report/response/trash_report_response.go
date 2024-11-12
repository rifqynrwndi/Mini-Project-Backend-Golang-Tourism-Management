package response

import (
	"tourism-monitoring/entities"
)

type TrashReportResponse struct {
	ID               int       `json:"id_laporan_sampah"`
    ObjekWisataID    int       `json:"objek_wisata_id"`
    JumlahSampah     float64   `json:"jumlah_sampah"`   
    TipeSampah       string    `json:"tipe_sampah"`   
    TanggalLaporan   string	   `json:"tanggal_laporan"`
}

func FromTrashReportEntities(trashReports []entities.TrashReport) []TrashReportResponse {
	responses := make([]TrashReportResponse, len(trashReports))
	for i, trashReport := range trashReports {
		responses[i] = TrashReportResponse{
			ID:               trashReport.ID,
			ObjekWisataID:    trashReport.ObjekWisataID,
			JumlahSampah:     trashReport.JumlahSampah,
			TipeSampah:       trashReport.TipeSampah,
			TanggalLaporan:   trashReport.TanggalLaporan.Format("2006-01-02"),
		}
	}
	return responses
}

func FromTrashReportEntity(trashReport entities.TrashReport) TrashReportResponse {
	return TrashReportResponse{
		ID:               trashReport.ID,
		ObjekWisataID:    trashReport.ObjekWisataID,
		JumlahSampah:     trashReport.JumlahSampah,
		TipeSampah:       trashReport.TipeSampah,
		TanggalLaporan:   trashReport.TanggalLaporan.Format("2006-01-02"),
	}
}
