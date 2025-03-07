package huggingface

import (
	"bytes"
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io"
	"main.go/internal/config"
	"main.go/pkg/logger"
	"net/http"
	"time"
)

func GenerateImageWithRetries(bot *tgbotapi.BotAPI, chatID int64, statusMsgID int, description string, cfg *config.Config) ([]byte, error) {
	var lastError error

	for i := 0; i < cfg.MaxRetries; i++ {
		editMsg := tgbotapi.NewEditMessageText(chatID, statusMsgID, fmt.Sprintf("🔄 Генерация %d/%d...", i+1, cfg.MaxRetries))
		bot.Send(editMsg)

		imageBytes, err := GenerateImage(description, cfg)
		if err == nil {
			return imageBytes, nil
		}

		lastError = err
		logger.Errorf("Попытка %d: Ошибка при генерации изображения: %v", i+1, err)

		delay := time.Duration(float64(cfg.InitialDelay) * float64(i+1))
		if delay > cfg.MaxDelay {
			delay = cfg.MaxDelay
		}

		time.Sleep(delay)
	}

	return nil, fmt.Errorf("не удалось сгенерировать изображение после %d попыток: %v", cfg.MaxRetries, lastError)
}

func GenerateImage(description string, cfg *config.Config) ([]byte, error) {
	url := fmt.Sprintf("https://api-inference.huggingface.co/models/%s", cfg.ModelID)

	payload := map[string]string{
		"inputs": description,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+cfg.HuggingFaceToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	logger.Debugf("Ответ от Hugging Face: %s", string(body))

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ошибка при генерации изображения: %s", resp.Status)
	}

	return body, nil
}
