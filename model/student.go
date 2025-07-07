package model

type Student struct {
	ID     uint64 `gorm:"primary"`
	Code   string
	Name   string
	Scores []Score
}
