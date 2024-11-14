package entities

type Place struct {
	ID               int    `gorm:"primaryKey" json:"id_objek_wisata"`
	Lokasi           string `json:"lokasi"`
	KapasitasMaks    int    `json:"kapasitas_maks"`
	JumlahPengunjung int    `json:"jumlah_pengunjung"`
}
