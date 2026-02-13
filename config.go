package filestore

import (
	"errors"
)

type Config struct {
	Endpoint        string
	AccessKey       string
	SecretKey       string
	UseSSL          bool
	Bucket          string
	EnableVersioning bool
	EnableSSE       bool
}

func LoadFromEnv() (*Config, error) {
	cfg := &Config{
    Endpoint:  "localhost:9000",
    AccessKey: "admin",
    SecretKey: "StrongSuperSecurePassword123!",
    Bucket:    "images", // همان bucket که ساختی
}

	if cfg.Endpoint == "" || cfg.AccessKey == "" || cfg.SecretKey == "" || cfg.Bucket == "" {
		return nil, errors.New("missing required minio environment variables")
	}

	return cfg, nil
}
