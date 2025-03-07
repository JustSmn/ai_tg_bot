package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"main.go/internal/config"
	"main.go/internal/huggingface"
	"main.go/pkg/logger"
)

const (
	helpText = `Привет🐙, я могу сгенерировать картинку по твоему тексту!🍄🍄🍄
Просто отправь текстовое описание, и я сгенерирую изображение.🔥
Есть одно но - пиши на английском языке🍔🦅🗽💵🥀
(Мой AI немного вредный, поэтому может генерировать изображение не с первого раза😭🤧😿🤒🤕 Но не грусти, попробуй еще раз отправить мне запрос, я постараюсь справиться с ним🥺😚😋😇🥰)`
	infoText = `Этот бот использует Hugging Face API для генерации изображений по текстовому описанию.
Используемая модель: stabilityai/stable-diffusion-2-1.
Контакт для связи: @Just_Semen228
JustSMN 2025 v.1.0.0`
)

func Start(cfg *config.Config) error {
	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		return fmt.Errorf("ошибка инициализации бота: %v", err)
	}

	bot.Debug = true
	logger.Infof("Авторизован как %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		switch update.Message.Command() {
		case "start", "help":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, helpText)
			bot.Send(msg)
			continue
		case "info":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, infoText)
			bot.Send(msg)
			continue
		}

		description := update.Message.Text

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "🔄 Генерация изображения началась. Пожалуйста, подождите...")
		statusMsg, _ := bot.Send(msg)

		// Генерация изображения с помощью Hugging Face API
		imageBytes, err := huggingface.GenerateImageWithRetries(bot, update.Message.Chat.ID, statusMsg.MessageID, description, cfg)
		if err != nil {
			logger.Errorf("Ошибка при генерации изображения: %v", err)

			deleteMsg := tgbotapi.NewDeleteMessage(update.Message.Chat.ID, statusMsg.MessageID)
			bot.Send(deleteMsg)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Произошла ошибка при генерации изображения😞😞😞 Попробуйте ещё раз 🥺 🥺 🥺")
			bot.Send(msg)
			continue
		}

		deleteMsg := tgbotapi.NewDeleteMessage(update.Message.Chat.ID, statusMsg.MessageID)
		bot.Send(deleteMsg)

		file := tgbotapi.FileBytes{Name: "image.png", Bytes: imageBytes}
		photo := tgbotapi.NewPhoto(update.Message.Chat.ID, file)
		_, err = bot.Send(photo)
		if err != nil {
			logger.Errorf("Ошибка при отправке изображения: %v", err)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Не удалось отправить изображение. Попробуйте ещё раз.")
			bot.Send(msg)
		}
	}

	return nil
}
