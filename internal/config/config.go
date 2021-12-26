package config

import (
	"github.com/Netflix/go-env"
)

type Config struct {
	DB struct {
		Username string `env:"DB_USERNAME, required=true"`
		Password string `env:"DB_PASSWORD, required=true"`
		Port     int    `env:"DB_PORT, default=5432"`
		Address  string `env:"DB_ADDRESS, default=localhost"`
		DBName   string `env:"DB_NAME, required=true"`
		SSLMode  string `env:"DB_SSLMODE, default=require"`
	}
	OAuth struct {
		ClientID     string `env:"OAUTH_CLIENT_ID, required=true"`
		ClientSecret string `env:"OAUTH_CLIENT_SECRET, required=true"`
		TenantID     string `env:"OAUTH_TENANT_ID, required=true"`
	}
	JWTSecret    string `env:"JWT_SECRET, required=true"`
	MigrationDir string `env:"MIGRATION_DIR"`
	Debug        bool   `env:"DEBUG, default=false"`
}

func Get() (*Config, error) {
	var cfg Config
	_, err := env.UnmarshalFromEnviron(&cfg)
	return &cfg, err
}
