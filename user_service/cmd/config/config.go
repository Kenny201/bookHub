package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI     string `env:"MONGO_URI"`
	KafkaBroker  string `env:"KAFKA_BROKER"`
	DatabaseName string `env:"DATABASE_NAME"`
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("could not load .env file: %v", err)
	}

	return &Config{
		MongoURI:     getEnv("MONGO_URI", "mongodb://userdb:27017"),
		KafkaBroker:  getEnv("KAFKA_BROKER", "kafka:9092"),
		DatabaseName: getEnv("DATABASE_NAME", "users"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
