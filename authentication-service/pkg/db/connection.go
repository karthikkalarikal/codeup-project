package db

import (
	"authentication/pkg/config"
	"authentication/pkg/domain"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg *config.Config) (*gorm.DB, error) {
	// psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)

	psqlInfo := os.Getenv("DSN")
	fmt.Println("psqlInfo", psqlInfo)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return nil, err
	}
	// user table
	if err = db.AutoMigrate(
		&domain.User{},
	); err != nil {
		return nil, err
	}

	// verify emails table
	if err = db.AutoMigrate(
		&domain.VerifyEmails{},
	); err != nil {
		return nil, err
	}

	return db, nil

}
