package config

import (
	"flag"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	RunAddress  string
	DatabaseDSN string
	JWTSecret   string
}

func Load() *Config {
	_ = godotenv.Load()

	var (
		addressFlag string
		dsnFlag     string
	)

	flag.StringVar(&addressFlag, "a", "", "service run address, e.g., :8080 or 127.0.0.1:9000")
	flag.StringVar(&dsnFlag, "d", "", "PostgreSQL DSN")
	flag.Parse()

	address := os.Getenv("RUN_ADDRESS")
	if address == "" {
		address = addressFlag
	}

	if address == "" {
		address = ":8080"
	}

	dsn := os.Getenv("DATABASE_URI")
	if dsn == "" {
		dsn = dsnFlag
	}
	if dsn == "" {
		log.Fatal("DATABASE_URI is not set")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET is not set")
	}

	return &Config{
		RunAddress:  address,
		DatabaseDSN: dsn,
		JWTSecret:   jwtSecret,
	}
}
