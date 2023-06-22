package models

type IpAddres struct {
	ID   uint `gorm:"primaryKey"`
	Ip   string
	Port uint32
}
