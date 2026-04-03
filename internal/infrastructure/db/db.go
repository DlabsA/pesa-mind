package db

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() error {
	// Initialize Viper to load environment variables
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: Could not read .env file: %v", err)
	}

	dsn := viper.GetString("DATABASE_URL")
	if dsn == "" {
		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			viper.GetString("DB_HOST"),
			viper.GetString("DB_PORT"),
			viper.GetString("DB_USER"),
			viper.GetString("DB_PASSWORD"),
			viper.GetString("DB_NAME"),
			viper.GetString("DB_SSLMODE"),
		)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db
	return nil
}
