package config

import (
	"flag"
	"os"
)

type PGConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type Config struct {
	PG   *PGConfig
	Port string
}

var (
	pgHostFlag     = flag.String("pg-host", "", "Database host")
	pgPortFlag     = flag.String("pg-port", "", "Database port")
	pgUserFlag     = flag.String("pg-user", "", "Database user")
	pgPasswordFlag = flag.String("pg-password", "", "Database password")
	pgNameFlag     = flag.String("pg-name", "", "Database name")
)

var (
	portFlag = flag.String("port", "8080", "Application port")
)

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func Load() *Config {
	flag.Parse()

	pghost := *pgHostFlag
	if *pgHostFlag == "" {
		pghost = getEnv("PG_HOST", "db")
	}
	pgport := *pgPortFlag
	if *pgPortFlag == "" {
		pgport = getEnv("PG_PORT", "5432")
	}
	pguser := *pgUserFlag
	if *pgUserFlag == "" {
		pguser = getEnv("PG_USER", "postgres")
	}
	pgpassword := *pgPasswordFlag
	if *pgPasswordFlag == "" {
		pgpassword = getEnv("PG_PASSWORD", "password")
	}
	pgname := *pgNameFlag
	if *pgNameFlag == "" {
		pgname = getEnv("PG_NAME", "db")
	}

	dbConfig := PGConfig{
		Host:     pghost,
		Port:     pgport,
		User:     pguser,
		Password: pgpassword,
		Name:     pgname,
	}

	port := *portFlag
	if *portFlag == "" {
		port = getEnv("PG_NAME", "db")
	}

	return &Config{
		PG:   &dbConfig,
		Port: port,
	}
}
