package fun

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func YesOrNo() string {
	rand.Seed(time.Now().UnixNano())
	var str string
	if rand.Intn(2) == 0 {
		str = "Да"
	} else {
		str = "Нет"
	}
	return str
}

func HandlerYesOrNo(ctx context.Context, b *bot.Bot, update *models.Update) {
	params := &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      "Мой ответ: " + YesOrNo(),
		ParseMode: models.ParseModeMarkdown,
	}

	_, err := b.SendMessage(ctx, params)
	if err != nil {
		fmt.Println(err)
	}
}
