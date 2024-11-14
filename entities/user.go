package entities

type User struct {
    ID           int    `gorm:"primaryKey" json:"id_wisatawan"`
    Nama         string `json:"nama"`
    Usia         int    `json:"usia"`
    Asal         string `json:"asal"`
    JenisKelamin string `json:"jenis_kelamin"`
    TipeWisatawan string `json:"tipe_wisatawan"`
    Email        string `json:"email" gorm:"unique"`
    Password     string `json:"password"`
    Token        string `json:"token"`
    Role         string `json:"role"`
}
