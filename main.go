package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/selis18/agents"
	"github.com/selis18/sprays"
)

func main() {
	//load .env
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("Error loading .env")
	// }

	token := os.Getenv("TOKEN")
	b, err := bot.New(token)
	if err != nil {
		log.Fatalf("Error creating bot: %v", err)
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "/randa", bot.MatchTypeContains, randAgentHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/rands", bot.MatchTypeContains, randSprayHandler)

	fmt.Println("Bot started")
	b.Start(context.Background())
}

// Обработчик команд бота
func randAgentHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	arrAgents, err := agents.GetAgents()
	if err != nil {
		fmt.Println(err)
	}

	randAgent := agents.GetRandomAgent(arrAgents)

	photo := randAgent.Icon

	// Параметры для отправки сообщения с изображением и текстом
	params := &bot.SendPhotoParams{
		ChatID:    update.Message.Chat.ID,
		Photo:     &models.InputFileString{Data: photo},
		Caption:   "Самый лучший на " + randAgent.Name,
		ParseMode: models.ParseModeMarkdown,
	}

	// Отправка сообщения с изображением
	_, err = b.SendPhoto(ctx, params)
	if err != nil {
		log.Fatalf("Error sending photo: %v", err)
	}
}

func randSprayHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	arrSprays, err := sprays.GetSprays()
	if err != nil {
		fmt.Println(err)
	}

	randSpray := sprays.GetRandomSpray(arrSprays)

	photo := randSpray.Icon

	// Параметры для отправки сообщения с изображением и текстом
	params := &bot.SendPhotoParams{
		ChatID:    update.Message.Chat.ID,
		Photo:     &models.InputFileString{Data: photo},
		Caption:   randSpray.Name,
		ParseMode: models.ParseModeMarkdown,
	}

	// Отправка сообщения с изображением
	_, err = b.SendPhoto(ctx, params)
	if err != nil {
		log.Fatalf("Error sending photo: %v", err)
	}
}
