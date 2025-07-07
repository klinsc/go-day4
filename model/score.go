package model

import "time"

type Score struct {
	ID           uint64
	ExerciseDate time.Time
	Score        uint64
	StudentID    uint64
}
