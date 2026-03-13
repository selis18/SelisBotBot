package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	telegramToken := os.Getenv("TG_TOKEN")

	amveraClient := NewAmveraClient()

	b, err := bot.New(telegramToken)
	if err != nil {
		log.Fatal("Bot creation failed:", err)
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "", bot.MatchTypePrefix, func(ctx context.Context, b *bot.Bot, update *models.Update) {
		// Игнорируем команды
		if strings.HasPrefix(update.Message.Text, "/") {
			return
		}

		text := update.Message.Text
		if !strings.Contains(strings.ToLower(text), "чел") {
			return
		}

		// Показываем "печатает..."
		b.SendChatAction(ctx, &bot.SendChatActionParams{
			ChatID: update.Message.Chat.ID,
			Action: models.ChatActionTyping,
		})

		reply, err := amveraClient.Ask(update.Message.Text)
		if err != nil {
			log.Printf("Amvera API error: %v", err)
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "Ошибка при обращении к модели. Попробуйте позже.",
				ReplyParameters: &models.ReplyParameters{
					MessageID: update.Message.ID,
				},
			})
			return
		}

		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   reply,
			ReplyParameters: &models.ReplyParameters{
				MessageID: update.Message.ID,
			},
		})
	})

	fmt.Println("Bot started")
	b.Start(context.Background())
}
