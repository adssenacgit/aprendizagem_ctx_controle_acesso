package config

import (
	"strconv"
)

type Config struct {
	Port       int
	Timeout    int
	DBHost     string
	DBPort     string
	DBUser     string
	DBPw       string
	DBTimeZone string
	DBDatabase string
}

func GetConfig() Config {
	return Config{
		Port:       parseEnvToInt("PORT", "8080"),
		Timeout:    parseEnvToInt("TIMEOUT", "30"),
		DBHost:     GetEnv("DB_HOST", "127.0.0.1"),
		DBPort:     GetEnv("DB_PORT", "3306"),
		DBUser:     GetEnv("DB_USER", "postgres"),
		DBPw:       GetEnv("DB_PASSWORD", "postgres"),
		DBDatabase: GetEnv("DB_DATABASE", "users"),
	}
}

func parseEnvToInt(envName, defaultValue string) int {
	num, err := strconv.Atoi(GetEnv(envName, defaultValue))

	if err != nil {
		return 0
	}

	return num
}
