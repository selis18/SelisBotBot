package fun

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

var variants = []string{
	"Сомневаюсь в этом",
	"Скорее всего, да",
	"Не исключено",
	"Да, но с оговорками",
	"Нет, но есть исключения",
	"Попробуйте поверить в это",
	"Да, и это замечательно",
	"Нет, это не в наших планах",
	"Нет, и это требует дальнейшего изучения",
	"Да, безусловно!",
}

func YesOrNo() string {
	return variants[rand.Intn(len(variants))]
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
