package model

type Students struct {
	ID     uint `gorm:"primaryKey"`
	Code   string
	Name   string
	Scores []Scores `gorm:"foreignKey:StudentID"`
}
