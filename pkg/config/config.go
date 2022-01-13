package config

import (
	"log"
	"os"
	"strconv"
)

func ProvideConfig() Config {
	return Config{
		BasePath: requireEnv("BASE_PATH"),
		UserService: service{
			Host:     requireEnv("USER_SERVICE_HOST"),
			BasePath: requireEnv("USER_SERVICE_BASE_PATH"),
		},
		Authentication: Authentication{
			Jwks: Jwks{
				Host:                   requireEnv("JWKS_HOST"),
				Index:                  requireEnvAsInt("JWKS_INDEX"),
				MinimumRefreshInterval: requireEnvAsInt("JWKS_MINIMUM_REFRESH_INTERVAL"),
			},
		},
		Dhis2Database: dhis2Database{
			Username: requireEnv("DHIS2_DATABASE_USERNAME"),
			Password: requireEnv("DHIS2_DATABASE_PASSWORD"),
			Database: requireEnv("DHIS2_DATABASE_DATABASE"),
		},
	}
}

type Config struct {
	BasePath       string
	UserService    service
	Authentication Authentication
	Bucket         string
	Dhis2Database  dhis2Database
}

type service struct {
	Host     string
	BasePath string
	Username string
	Password string
}

type dhis2Database struct {
	Username string
	Password string
	Database string
}

type Authentication struct {
	Jwks Jwks
}

type Jwks struct {
	Host                   string
	Index                  int
	MinimumRefreshInterval int
}

func requireEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Can't find environment varialbe: %s\n", key)
	}
	return value
}

func requireEnvAsInt(key string) int {
	valueStr := requireEnv(key)
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Fatalf("Can't parse value as integer: %s", err.Error())
	}
	return value
}
