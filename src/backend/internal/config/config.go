package config

import "os"

type (
	Config struct {
		PostgresConf *PostgresConfig
		RedisConf    *RedisConfig
		ServerConfig *ServerConfig
		PaymentConfig *PaymentApiConfig
	}

	PostgresConfig struct {
		Host     string
		Port     string
		Password string
		User     string
		DBName   string
		SSLMode  string
	}

	RedisConfig struct {
		Host     string
		Port     string
		Password string
	}

	ServerConfig struct {
		Host string
		Port string
	}

	PaymentApiConfig struct {
		ApiKey string
		ShopID string
	}
)

func paymentApiConfFromEnv() *PaymentApiConfig {
	return &PaymentApiConfig{
		ApiKey: os.Getenv("API_KEY"),
		ShopID: os.Getenv("SHOP_ID"),
	}
}

func serverConfFromEnv() *ServerConfig {
	return &ServerConfig{
		Host: os.Getenv("GOLANG_HOST"),
		Port: os.Getenv("GOLANG_PORT"),
	}
}

func redisConfFromEnv() *RedisConfig {
	return &RedisConfig{
		Host:     os.Getenv("REDIS_HOST"),
		Port:     os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
	}
}

func postgresConfFromEnv() *PostgresConfig {
	return &PostgresConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		User:     os.Getenv("POSTGRES_USER"),
		DBName:   os.Getenv("POSTGRES_DB"),
		SSLMode:  os.Getenv("POSTGRES_SSLMODE"),
	}
}

func NewConfFromEnv() *Config {
	return &Config{
		PostgresConf: postgresConfFromEnv(),
		RedisConf:    redisConfFromEnv(),
		ServerConfig: serverConfFromEnv(),
		PaymentConfig: paymentApiConfFromEnv(),
	}
}
