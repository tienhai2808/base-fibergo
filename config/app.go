package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	App   AppConfig
	Mongo MongoDBConfig
	Redis RedisConfig
}

type AppConfig struct {
	Port      string
	JWTSecret string
	ApiPrefix string
}

type MongoDBConfig struct {
	URI      string
	Database string
}

type RedisConfig struct {
	Addr     string
	Password string
}

func LoadConfig() *Config {
	_ = godotenv.Load()

	return &Config{
		App: AppConfig{
			Port:      getEnv("PORT", "5003"),
			JWTSecret: getEnv("JWT_SECRET", "haha-123-huhu-222"),
		},
		Mongo: MongoDBConfig{
			URI:      getEnv("MONGO_URI", "mongodb://localhost:27017"),
			Database: getEnv("MONGO_DB", "appdb"),
		},
		Redis: RedisConfig{
			Addr:     getEnv("REDIS_ADDR", "localhost:6379"),
			Password: getEnv("REDIS_PASS", ""),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
