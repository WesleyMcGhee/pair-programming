package models

type Users struct {
	ID            int `gorm:"primaryKey"`
	Username      string
	Password_hash string
	Email         string
}