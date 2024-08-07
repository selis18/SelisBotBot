package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/joho/godotenv"
	"github.com/selis18/agents"
)

func main() {
	agents.GetAgents()
	//load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env")
	}

	token := os.Getenv("TOKEN")
	b, err := bot.New(token)
	if err != nil {
		log.Fatalf("Error creating bot: %v", err)
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "/randa", bot.MatchTypeContains, randAgentHandler)

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
	chatId := os.Getenv("CHATID")

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    chatId,
		Text:      "Самый лучший на " + randAgent.Name,
		ParseMode: models.ParseModeMarkdown,
	})
}
