package entities

import (
	"time"
)

type TrashReport struct {
	ID               	int       	`gorm:"primaryKey" json:"id_laporan_sampah"`
    ObjekWisataID    	int       	`json:"objek_wisata_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ObjekWisata			Place		`json:"objek_wisata"`
    JumlahSampah     	float64   	`json:"jumlah_sampah"`   
    TipeSampah       	string    	`json:"tipe_sampah"`   
    TanggalLaporan   	time.Time 	`json:"tanggal_laporan,omitempty"`
}