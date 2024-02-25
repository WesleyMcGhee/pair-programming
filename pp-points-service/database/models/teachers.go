package models

type Teachers struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Group []Groups
}