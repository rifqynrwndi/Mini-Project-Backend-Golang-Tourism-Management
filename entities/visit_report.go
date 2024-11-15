package entities

import (
	"time"
)

type VisitReport struct {
    ID                       int       `gorm:"primaryKey" json:"id_laporan"`
    WisatawanID              int       `json:"id_wisatawan" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
    Wisatawan                User      `json:"wisatawan"`
    ObjekWisataID            int       `json:"id_objek_wisata" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
    ObjekWisata              Place     `json:"objek_wisata"`
    TanggalKunjungan         time.Time `json:"tanggal_kunjungan,omitempty"`
    EstimasiEmisiKarbon      float64   `json:"estimasi_emisi_karbon"`
    CatatanSampahPerKilogram float64   `json:"catatan_sampah_per_kilogram"`
}
