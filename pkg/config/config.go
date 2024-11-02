package config

import (
	"os"
	"strconv"
)

var (
	OS_PREFIX   = "IGGY"
	SYSTEM_USER = "SYSTEM"
)

type Configuration struct {
	AdminUser       string
	AdminPassword   string
	BindingPort     string
	BindingIP       string
	DBUser          string
	DBHost          string
	DBPort          string
	DBPassword      string
	DBDatabase      string
	DBDialect       string
	Logger          string
	TokenExpiration int
	TokenKeyLength  int
	PasswordCost    int
}

func Load() (Configuration, error) {
	tokenExpiration, err := strconv.Atoi(getEnv("TOKEN_EXPIRATION", "2"))
	if err != nil {
		return Configuration{}, nil
	}

	return Configuration{
		AdminUser:       getEnv("ADMIN_User", "admin"),
		AdminPassword:   getEnv("ADMIN_PASSWORD", ""),
		BindingIP:       getEnv("BINDING_IP", "0.0.0.0"),
		BindingPort:     getEnv("BINDING_PORT", "7800"),
		DBUser:          getEnv("DB_USER", ""),
		DBPort:          getEnv("DB_PORT", ""),
		DBPassword:      getEnv("DB_PASSWORD", ""),
		DBHost:          getEnv("DB_HOST", ""),
		DBDatabase:      getEnv("DB_DATABASE", ""),
		DBDialect:       getEnv("DB_DIALECT", "postgres"),
		Logger:          getEnv("LOGGER", "zap"),
		TokenExpiration: tokenExpiration,
		PasswordCost:    11,
		TokenKeyLength:  32,
	}, nil
}

func getEnv(key string, defaultValue string) string {
	fullKey := OS_PREFIX + "_" + key
	val := os.Getenv(fullKey)
	if val == "" {
		if defaultValue != "" {
			return defaultValue
		}
	}
	return val
}
