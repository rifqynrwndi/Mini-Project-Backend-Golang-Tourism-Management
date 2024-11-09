package entities

import "time"

type LaporanKunjungan struct {
    ID                      int       `gorm:"primaryKey" json:"id_laporan"`
    WisatawanID             int       `json:"id_wisatawan"`
    ObjekWisataID           int       `json:"id_objek_wisata"`
    TanggalKunjungan        time.Time `json:"tanggal_kunjungan"`
    EstimasiEmisiKarbon     float64   `json:"estimasi_emisi_karbon"`
    CatatanSampahTiapKunjungan string `json:"catatan_sampah_tiap_kunjungan"`
}
