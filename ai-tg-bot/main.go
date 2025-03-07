package main

import (
	"log"
	"main.go/internal/bot"
	"main.go/internal/config"
)

func main() {
	// Инициализация логгера
	//logger.Init()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	if err := bot.Start(cfg); err != nil {
		log.Fatalf("Ошибка запуска бота: %v", err)
	}
}
