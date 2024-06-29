package config

import (
	"edusync/pkg/logger"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

var (
	instance *Config
	once     sync.Once
)

type Config struct {
	Host        string
	Port        int
	Environment string
	Debug       bool

	DBPostgreDriver string
	DBPostgreDsn    string
	DBPostgreURL    string

	JWTSecret  string
	JWTExpired int

	HashKey string

	MinioEndpoint    string
	MinioAccessKeyId string
	MinioSecretKey   string
	MinioBucketName  string
	MinioLocation    string
	MinioUseSSL      bool
}

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{
			Host:        cast.ToString(getOrReturnDefault("HOST", "localhost")),
			Port:        cast.ToInt(getOrReturnDefault("PORT", "8080")),
			Environment: cast.ToString(getOrReturnDefault("ENVIRONMENT", "")),
			Debug:       cast.ToBool(getOrReturnDefault("DEBUG", "")),

			DBPostgreDriver: cast.ToString(getOrReturnDefault("DB_POSTGRE_DRIVER", "")),
			DBPostgreDsn:    cast.ToString(getOrReturnDefault("DB_POSTGRE_DSN", "")),
			DBPostgreURL:    cast.ToString(getOrReturnDefault("DB_POSTGRE_URL", "")),

			JWTSecret:  cast.ToString(getOrReturnDefault("JWT_SECRET", "")),
			JWTExpired: cast.ToInt(getOrReturnDefault("JWT_EXPIRED", "")),

			HashKey: cast.ToString(getOrReturnDefault("HASH_KEY", "")),

			MinioEndpoint:    cast.ToString(getOrReturnDefault("MINIO_ENDPOINT", "")),
			MinioAccessKeyId: cast.ToString(getOrReturnDefault("MINIO_ACCESS_KEY_ID", "")),
			MinioSecretKey:   cast.ToString(getOrReturnDefault("MINIO_SECRET_KEY", "")),
			MinioBucketName:  cast.ToString(getOrReturnDefault("MINIO_BUCKET_NAME", "")),
			MinioLocation:    cast.ToString(getOrReturnDefault("MINIO_LOCATION", "")),
			MinioUseSSL:      cast.ToBool(getOrReturnDefault("MINIO_USE_SLL", false)),
		}
	})

	return instance
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	err := godotenv.Load("internal/config/.env")
	if err != nil {
		logger.GetLogger().Error(err)
	}
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
