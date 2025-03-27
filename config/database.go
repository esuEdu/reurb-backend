package config

import (
	"fmt"
	"log"

	"github.com/esuEdu/reurb-backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {

	env := LoadEnv()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		env.DBHost,
		env.DBUser,
		env.DBPass,
		env.DBName,
		env.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	db.AutoMigrate(&models.User{}) // initial migration

	return db

}
