package request

import (
	"fmt"
	"time"
	"tourism-monitoring/entities"
)

type TrashReportRequest struct {
	ID             int     `json:"id_laporan_sampah"`
	ObjekWisataID  int     `json:"objek_wisata_id"`
	JumlahSampah   float64 `json:"jumlah_sampah"`
	TipeSampah     string  `json:"tipe_sampah"`
	TanggalLaporan string  `json:"tanggal_laporan"`
}

func (trashReportRequest TrashReportRequest) ToEntities() (entities.TrashReport, error) {
	parsedDate, err := time.Parse("2006-01-02", trashReportRequest.TanggalLaporan)
	if err != nil {
		return entities.TrashReport{}, fmt.Errorf("invalid date format for tanggal_kunjungan: %v", err)
	}
	return entities.TrashReport{
		ID:             trashReportRequest.ID,
		ObjekWisataID:  trashReportRequest.ObjekWisataID,
		JumlahSampah:   trashReportRequest.JumlahSampah,
		TipeSampah:     trashReportRequest.TipeSampah,
		TanggalLaporan: parsedDate,
	}, nil
}
