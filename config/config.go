package config

import (
	"fmt"
	"taxi-service/utils"
)

type Config struct {
	DatabaseDSN string
}

func LoadConfig() Config {
	dbHost := utils.GetEnv("DB_HOST", "localhost")
	dbUser := utils.GetEnv("DB_USER", "postgres")
	dbPassword := utils.GetEnv("DB_PASSWORD", "postgres")
	dbName := utils.GetEnv("DB_NAME", "postgres")
	dbPort := utils.GetEnv("DB_PORT", "5432")

	return Config{
		DatabaseDSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			dbHost, dbUser, dbPassword, dbName, dbPort),
	}
}
