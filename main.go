package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-telegram/bot"
	"github.com/joho/godotenv"
	"github.com/selis18/apis"
	"github.com/selis18/fun"
)

func main() {
	//load .env
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	token := os.Getenv("TOKEN")
	b, err := bot.New(token)
	if err != nil {
		log.Fatalf("Error creating bot: %v", err)
	}

	var agentResponse *apis.AgentResponse
	var sprayResponse *apis.SprayResponse
	var collectionResponse *apis.CollectionResponse
	b.RegisterHandler(bot.HandlerTypeMessageText, "/randa", bot.MatchTypeContains, agentResponse.Handler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/randt", bot.MatchTypeContains, agentResponse.TeamHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/rands", bot.MatchTypeContains, sprayResponse.Handler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/randc", bot.MatchTypeContains, collectionResponse.Handler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/yn", bot.MatchTypeContains, fun.HandlerYesOrNo)

	fmt.Println("Bot started")
	b.Start(context.Background())

}
