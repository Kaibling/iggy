package config

import (
	"os"
	"strconv"
)

const (
	OS_PREFIX   = "IGGY"
	SYSTEM_USER = "SYSTEM"
)

type Configuration struct {
	App    AppConfig
	DB     DBConfig
	Broker BrokerConfig
}
type AppConfig struct {
	AdminUser       string
	AdminPassword   string
	BindingPort     string
	BindingIP       string
	Logger          string
	TokenExpiration int
	TokenKeyLength  int
	PasswordCost    int
}

type DBConfig struct {
	DBUser     string
	DBHost     string
	DBPort     string
	DBPassword string
	DBDatabase string
	DBDialect  string
}
type BrokerConfig struct {
	Channel string
	Broker  string
}

func Load() (Configuration, error) {
	tokenExpiration, err := strconv.Atoi(getEnv("TOKEN_EXPIRATION", "2"))
	if err != nil {
		return Configuration{}, nil
	}

	return Configuration{
		App: AppConfig{
			Logger:          getEnv("LOGGER", "zap"),
			TokenExpiration: tokenExpiration,
			AdminUser:       getEnv("ADMIN_User", "admin"),
			AdminPassword:   getEnv("ADMIN_PASSWORD", ""),
			BindingIP:       getEnv("BINDING_IP", "0.0.0.0"),
			BindingPort:     getEnv("BINDING_PORT", "7800"),
			PasswordCost:    11,
			TokenKeyLength:  32,
		},
		DB: DBConfig{
			DBUser:     getEnv("DB_USER", ""),
			DBPort:     getEnv("DB_PORT", ""),
			DBPassword: getEnv("DB_PASSWORD", ""),
			DBHost:     getEnv("DB_HOST", ""),
			DBDatabase: getEnv("DB_DATABASE", ""),
			DBDialect:  getEnv("DB_DIALECT", "postgres"),
		},
		Broker: BrokerConfig{
			Channel: getEnv("BROKER_CHANNEL", "iggy"),
			Broker:  getEnv("BROKER_BROKER", "loopback"),
		},
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
