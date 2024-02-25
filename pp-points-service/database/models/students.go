package models

type Students struct {
	ID     int `gorm:"primaryKey"`
	Name   string
	Group  Groups
	Points int
}