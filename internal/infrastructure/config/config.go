package config

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	App         AppConfig
	Database    DatabaseConfig
	CORS        CORSConfig
	JWT         JWTConfig
	GoogleDrive GoogleDriveConfig
	Storage     StorageConfig
	Xendit      XenditConfig
}

type AppConfig struct {
	Name  string
	Env   string
	Port  string
	Debug bool
}

type DatabaseConfig struct {
	Host            string
	Port            string
	User            string
	Password        string
	Name            string
	SSLMode         string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

type CORSConfig struct {
	AllowedOrigins []string
}

type JWTConfig struct {
	Secret     string
	Expiration time.Duration
}

type GoogleDriveConfig struct {
	CredentialsPath string
	FolderID        string
}

type StorageConfig struct {
	Type string
}

type XenditConfig struct {
	APIKey        string
	CallbackToken string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	return &Config{
		App: AppConfig{
			Name:  getEnv("APP_NAME", "ecommerce-backend"),
			Env:   getEnv("APP_ENV", "development"),
			Port:  getEnv("APP_PORT", "8080"),
			Debug: getEnvAsBool("APP_DEBUG", true),
		},
		Database: DatabaseConfig{
			Host:            getEnv("DB_HOST", "localhost"),
			Port:            getEnv("DB_PORT", "5432"),
			User:            getEnv("DB_USER", "postgres"),
			Password:        getEnv("DB_PASSWORD", "postgres"),
			Name:            getEnv("DB_NAME", "ecommerce_db"),
			SSLMode:         getEnv("DB_SSL_MODE", "disable"),
			MaxOpenConns:    getEnvAsInt("DB_MAX_OPEN_CONNS", 25),
			MaxIdleConns:    getEnvAsInt("DB_MAX_IDLE_CONNS", 5),
			ConnMaxLifetime: getEnvAsDuration("DB_CONN_MAX_LIFETIME", 5*time.Minute),
		},
		CORS: CORSConfig{
			AllowedOrigins: getEnvAsSlice("CORS_ALLOWED_ORIGINS", []string{"*"}),
		},
		JWT: JWTConfig{
			Secret:     getEnv("JWT_SECRET", ""),
			Expiration: getEnvAsDuration("JWT_EXPIRATION", 24*time.Hour),
		},
		GoogleDrive: GoogleDriveConfig{
			CredentialsPath: getEnv("GOOGLE_DRIVE_CREDENTIALS_PATH", "./credentials.json"),
			FolderID:        getEnv("GOOGLE_DRIVE_FOLDER_ID", ""),
		},
		Storage: StorageConfig{
			Type: getEnv("STORAGE_TYPE", "local"),
		},
		Xendit: XenditConfig{
			APIKey:        getEnv("XENDIT_API_KEY", ""),
			CallbackToken: getEnv("XENDIT_CALLBACK_TOKEN", ""),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultVal
}

func getEnvAsInt(key string, defaultVal int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}

func getEnvAsBool(key string, defaultVal bool) bool {
	valueStr := getEnv(key, "")
	if value, err := strconv.ParseBool(valueStr); err == nil {
		return value
	}
	return defaultVal
}

func getEnvAsDuration(key string, defaultVal time.Duration) time.Duration {
	valueStr := getEnv(key, "")
	if value, err := time.ParseDuration(valueStr); err == nil {
		return value
	}
	return defaultVal
}

func getEnvAsSlice(key string, defaultVal []string) []string {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return defaultVal
	}
	return strings.Split(valueStr, ",")
}
