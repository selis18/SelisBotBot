package fun

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	_ "github.com/lib/pq"
	"github.com/selis18/db"
	"github.com/selis18/errs"
	"golang.org/x/exp/rand"
)

func getRandomAnswer() string {
	db, err := db.GetDB()
	errs.CheckErr("Can't get database", err)

	rows, err := db.Query("SELECT random_string FROM randomstrings")
	errs.CheckErr("Can't get strokes from database", err)
	defer db.Close()

	var randomStrings []string
	for rows.Next() {
		var randomString string
		err = rows.Scan(&randomString)
		errs.CheckErr("failed to scan random string: %v", err)

		randomStrings = append(randomStrings, randomString)
	}
 
	// Рандомная строка
	randIndex := rand.Intn(len(randomStrings))
	randomString := randomStrings[randIndex]

	return randomString
}

func HandlerRandomAnswer(ctx context.Context, b *bot.Bot, update *models.Update) {
	params := &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      getRandomAnswer(),
		ParseMode: models.ParseModeMarkdown,
	}

	_, err := b.SendMessage(ctx, params)
	errs.CheckErr("Can't send message", err)
}
