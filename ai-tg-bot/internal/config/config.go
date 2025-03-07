package config

import (
	"os"
	"time"
)

type Config struct {
	TelegramToken    string
	HuggingFaceToken string
	ModelID          string
	MaxRetries       int
	InitialDelay     time.Duration
	MaxDelay         time.Duration
}

func Load() (*Config, error) {
	return &Config{
		TelegramToken:    os.Getenv("TELEGRAM_TOKEN"),
		HuggingFaceToken: os.Getenv("HUGGINGFACE_TOKEN"),
		ModelID:          "stabilityai/stable-diffusion-2-1",
		MaxRetries:       5,
		InitialDelay:     2 * time.Second,
		MaxDelay:         30 * time.Second,
	}, nil
}
