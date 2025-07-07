package model

import "time"

type Scores struct {
	ID           uint `gorm:"primaryKey"`
	ExerciseDate time.Time
	Score        int
	StudentID    uint
}
