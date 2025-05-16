package config

import (
	"flag"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DatabaseDSN string
	JWTSecret   string
}

func Load() *Config {
	_ = godotenv.Load()

	var dsnFlag string
	flag.StringVar(&dsnFlag, "d", "", "PostgreSQL DSN")
	flag.Parse()

	dsn := os.Getenv("DATABASE_URI")

	log.Printf("Connecting to DB at: %s", dsn)

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
		DatabaseDSN: dsn,
		JWTSecret:   jwtSecret,
	}
}
