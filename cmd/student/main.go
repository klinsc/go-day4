package main

import (
	"fmt"
	"log"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/klinsc/config"
	"github.com/klinsc/model"
	student "github.com/klinsc/students"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// main.go

func main() {
	// Load configuration
	var conf config.Config

	if err := env.Parse(&conf); err != nil {
		log.Fatal("Error when decoding configuration:", err)
	}

	// Database connection
	db, err := gorm.Open(postgres.Open(conf.Database.URL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Error when connect to database:", err)
	}
	// Ensure close connection when terminated
	defer func() {
		sqldb, _ := db.DB()
		sqldb.Close()
	}()

	repo := student.NewRepository(db)
	insertValue(repo)
}

func insertValue(repo student.IRepository) {
	student := &model.Students{
		Code: "0001",
		Name: "Java",
		Scores: []model.Scores{
			{
				ExerciseDate: time.Date(2025, time.June, 20, 20, 30, 00, 00, time.Now().Location()),
				Score:        20,
			},
			{
				ExerciseDate: time.Date(2025, time.June, 25, 20, 30, 00, 00, time.Now().Location()),
				Score:        50,
			},
		},
	}
	err := repo.Create(student)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Student: %+v", student)
}
