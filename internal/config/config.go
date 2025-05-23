package config

import (
	"flag"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	RunAddress     string
	AccrualAddress string
	DatabaseDSN    string
	JWTSecret      string
}

func Load() *Config {
	_ = godotenv.Load()

	var (
		addressFlag string
		accrualFlag string
		dsnFlag     string
	)

	flag.StringVar(&addressFlag, "a", "", "Service run address")
	flag.StringVar(&accrualFlag, "r", "", "Accrual service run address")
	flag.StringVar(&dsnFlag, "d", "", "PostgreSQL DSN")
	flag.Parse()

	address := os.Getenv("RUN_ADDRESS")
	if address == "" {
		address = addressFlag
	}

	if address == "" {
		address = ":8080"
	}

	accrualAddress := os.Getenv("ACCRUAL_SYSTEM_ADDRESS")
	if accrualAddress == "" {
		accrualAddress = accrualFlag
	}

	if accrualAddress == "" {
		accrualAddress = ":8081"
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
