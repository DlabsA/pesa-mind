package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port               string
	Env                string
	DBHost             string
	DBPort             string
	DBUser             string
	DBPassword         string
	DBName             string
	DBSSLMode          string
	DatabaseURL        string
	JWTSecret          string
	JWTExpiry          string
	RefreshTokenExpiry string
	RedisURL           string
	LogLevel           string
	CORSOrigins        string
}

func LoadConfig() *Config {
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("ENV", "development")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "5432")
	viper.SetDefault("DB_USER", "pesamind")
	viper.SetDefault("DB_PASSWORD", "pesamind123")
	viper.SetDefault("DB_NAME", "pesamind")
	viper.SetDefault("DB_SSLMODE", "disable")
	viper.SetDefault("DATABASE_URL", "")
	viper.SetDefault("JWT_SECRET", "super-secret-change-in-production")
	viper.SetDefault("JWT_EXPIRY", "15m")
	viper.SetDefault("REFRESH_TOKEN_EXPIRY", "30d")
	viper.SetDefault("REDIS_URL", "redis://localhost:6379")
	viper.SetDefault("LOG_LEVEL", "info")
	viper.SetDefault("CORS_ORIGINS", "http://localhost:3000,https://pesamind.app")

	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("No .env file found: %v", err)
	}
	return &Config{
		Port:               viper.GetString("PORT"),
		Env:                viper.GetString("ENV"),
		DBHost:             viper.GetString("DB_HOST"),
		DBPort:             viper.GetString("DB_PORT"),
		DBUser:             viper.GetString("DB_USER"),
		DBPassword:         viper.GetString("DB_PASSWORD"),
		DBName:             viper.GetString("DB_NAME"),
		DBSSLMode:          viper.GetString("DB_SSLMODE"),
		JWTSecret:          viper.GetString("JWT_SECRET"),
		JWTExpiry:          viper.GetString("JWT_EXPIRY"),
		RefreshTokenExpiry: viper.GetString("REFRESH_TOKEN_EXPIRY"),
		RedisURL:           viper.GetString("REDIS_URL"),
		LogLevel:           viper.GetString("LOG_LEVEL"),
		CORSOrigins:        viper.GetString("CORS_ORIGINS"),
	}
}
