package models

type Groups struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Students []Students
	Teacher  Teachers
}