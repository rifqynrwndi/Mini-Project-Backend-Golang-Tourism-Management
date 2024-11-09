package request

import (
	"fmt"
	"time"
	"tourism-monitoring/entities"
)

type VisitReportRequest struct {
	WisatawanID              int     `json:"id_wisatawan"`
	ObjekWisataID            int     `json:"id_objek_wisata"`
	TanggalKunjungan         string  `json:"tanggal_kunjungan"`
	EstimasiEmisiKarbon      float64 `json:"estimasi_emisi_karbon"`
	CatatanSampahPerKilogram float64 `json:"catatan_sampah_per_kilogram"`
	TransportMode            string  `json:"transport_mode"`
	DistanceKM               float64 `json:"distance_km"`
}

func (visitReportRequest VisitReportRequest) ToEntities() (entities.VisitReport, error) {
	parsedDate, err := time.Parse("2006-01-02", visitReportRequest.TanggalKunjungan)
	if err != nil {
		return entities.VisitReport{}, fmt.Errorf("invalid date format for tanggal_kunjungan: %v", err)
	}
	return entities.VisitReport{
		WisatawanID:              visitReportRequest.WisatawanID,
		ObjekWisataID:            visitReportRequest.ObjekWisataID,
		TanggalKunjungan:         parsedDate,
		EstimasiEmisiKarbon:      visitReportRequest.EstimasiEmisiKarbon,
		CatatanSampahPerKilogram: visitReportRequest.CatatanSampahPerKilogram,
	} , nil
}
