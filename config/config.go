package config

import (
	"fmt"
	"log"
	"os"
)

type Config struct {
	DatabaseDSN string
}

func LoadConfig() Config {
	dbHost := getEnv("DB_HOST", "localhost")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "postgres")
	dbName := getEnv("DB_NAME", "postgres")
	dbPort := getEnv("DB_PORT", "5432")

	//print db config
	log.Println(dbHost)

	return Config{
		DatabaseDSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			dbHost, dbUser, dbPassword, dbName, dbPort),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
