package response

import "tourism-monitoring/entities"

type VisitReportResponse struct {
	ID                       int     `json:"id"`
	WisatawanID              int     `json:"id_wisatawan"`
	ObjekWisataID            int     `json:"id_objek_wisata"`
	TanggalKunjungan         string  `json:"tanggal_kunjungan"`
	EstimasiEmisiKarbon      float64 `json:"estimasi_emisi_karbon"`
	CatatanSampahPerKilogram float64 `json:"catatan_sampah_per_kilogram"`
}

func FromVisitReportEntities(visitReports []entities.VisitReport) []VisitReportResponse {
	var responses []VisitReportResponse
	for _, visitReport := range visitReports {
		responses = append(responses, VisitReportResponse{
			ID:                       visitReport.ID,
			WisatawanID:              visitReport.WisatawanID,
			ObjekWisataID:            visitReport.ObjekWisataID,
			TanggalKunjungan:         visitReport.TanggalKunjungan.Format("2006-01-02"),
			EstimasiEmisiKarbon:      visitReport.EstimasiEmisiKarbon,
			CatatanSampahPerKilogram: visitReport.CatatanSampahPerKilogram,
		})
	}
	return responses
}

func FromVisitReportEntity(visitReport entities.VisitReport) VisitReportResponse {
	return VisitReportResponse{
		ID:                       visitReport.ID,
		WisatawanID:              visitReport.WisatawanID,
		ObjekWisataID:            visitReport.ObjekWisataID,
		TanggalKunjungan:         visitReport.TanggalKunjungan.Format("2006-01-02"),
		EstimasiEmisiKarbon:      visitReport.EstimasiEmisiKarbon,
		CatatanSampahPerKilogram: visitReport.CatatanSampahPerKilogram,
	}
}
