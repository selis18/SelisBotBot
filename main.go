package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-telegram/bot"
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

	fmt.Println("Bot started")
	arrAgents, err := agents.GetAgents()
	if err != nil {
		fmt.Println(err)
	}

	randAgent := agents.GetRandomAgent(arrAgents)

	chatId := os.Getenv("CHATID")
	message := "Ты пупсик на " + randAgent.Name

	params := &bot.SendMessageParams{
		ChatID: chatId,
		Text:   message,
	}

	_, err = b.SendMessage(context.Background(), params)
	if err != nil {
		log.Fatalf("Error sending message: %v", err)
	}

}
