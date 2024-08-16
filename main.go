package main

import (
	"context"
	"fmt"
	"os"

	"github.com/go-telegram/bot"
	"github.com/joho/godotenv"
	"github.com/selis18/apis"

	"github.com/selis18/errs"
	"github.com/selis18/fun"
)

type Users struct {
	Users []User
}

type User struct {
	Id    string `json:"id"`
	Login string `json:"login"`
	Num   string `json:"num"`
}

func main() {
	// //load .env
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	token := os.Getenv("TOKEN")
	b, err := bot.New(token)
	errs.CheckErr("Bot blocked", err)

	var agentResponse *apis.AgentResponse
	var sprayResponse *apis.SprayResponse
	var collectionResponse *apis.CollectionResponse
	b.RegisterHandler(bot.HandlerTypeMessageText, "/randa", bot.MatchTypeContains, agentResponse.Handler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/randt", bot.MatchTypeContains, agentResponse.TeamHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/rands", bot.MatchTypeContains, sprayResponse.Handler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/randc", bot.MatchTypeContains, collectionResponse.Handler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/yn", bot.MatchTypeContains, fun.HandlerRandomAnswer)

	fmt.Println("Bot started")
	b.Start(context.Background())

	// // //user := update.Message.Chat.ID
	// paramss := &bot.ForwardMessageParams{
	// 	ChatID:     495361324,
	// 	MessageID:  347,
	// 	FromChatID: "495361324",
	// }

	// id, err := b.ForwardMessage(context.Background(), paramss)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(id.Text) //1723443037

	// var user User
	// var str = []byte(`{"id": "1", "login": "selis18", "num": "1"}`)

	// err = json.Unmarshal(str, &user)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(user.Id)

}
