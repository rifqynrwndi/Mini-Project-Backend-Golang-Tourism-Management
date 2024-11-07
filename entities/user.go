package entities

type User struct {
    ID           int   `gorm:"primaryKey"`
    Nama         string
    Usia         int
    Asal         string
    JenisKelamin string
    TipeWisatawan string
    Email        string `gorm:"unique"`
    Password     string
    Token        string
    Role         string
}
