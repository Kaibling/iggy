package config

import (
	"os"
	"strconv"
	"strings"
)

const (
	osPrefix     = "IGGY"
	SystemUser   = "SYSTEM"
	passwordCost = 11
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
	ImportRepoURL   string
	ImportLocalPath string
	ImportOnStartup bool
	// ExportRepoURL   string
	ExportLocalPath string
	GitToken        string
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
	Channel          string
	BrokerName       string
	ConnectionString string
	Username         string
	Password         string
}

func Load() (Configuration, error) {
	tokenExpiration, err := strconv.Atoi(getEnv("TOKEN_EXPIRATION", "2"))
	if err != nil {
		return Configuration{}, err
	}

	tokenLength, err := strconv.Atoi(getEnv("TOKEN_LENGTH", "32"))
	if err != nil {
		return Configuration{}, err
	}

	return Configuration{
		App: AppConfig{
			Logger:          getEnv("LOGGER", "zap"),
			TokenExpiration: tokenExpiration,
			AdminUser:       getEnv("ADMIN_User", "admin"),
			AdminPassword:   getEnv("ADMIN_PASSWORD", ""),
			BindingIP:       getEnv("BINDING_IP", "0.0.0.0"),
			BindingPort:     getEnv("BINDING_PORT", "7800"),
			PasswordCost:    passwordCost,
			TokenKeyLength:  tokenLength,
			ImportRepoURL:   getEnv("IMPORT_REPO_URL", ""),
			ImportLocalPath: getEnv("IMPORT_LOCAL_PATH", "/tmp/workflows"),
			ImportOnStartup: toBool(getEnv("IMPORT_ON_STARTUP", "false")),
			ExportLocalPath: getEnv("EXPORT_LOCAL_PATH", "/tmp/workflows"),
			GitToken:        getEnv("GIT_TOKEN", ""),
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
			Channel:          getEnv("BROKER_CHANNEL", "iggy"),
			BrokerName:       getEnv("BROKER_NAME", "nats"),
			ConnectionString: getEnv("BROKER_CONNECTION_STRING", "nats://0.0.0.0:4222"),
			Username:         getEnv("BROKER_USERNAME", ""),
			Password:         getEnv("BROKER_PASSWORD", ""),
		},
	}, nil
}

func getEnv(key string, defaultValue string) string {
	fullKey := osPrefix + "_" + key

	val := os.Getenv(fullKey)
	if val == "" {
		if defaultValue != "" {
			return defaultValue
		}
	}

	return val
}

func toBool(s string) bool {
	return strings.ToLower(s) == "true"
}
