package apis

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/selis18/errs"
)

type CollectionResponse struct {
	Status int          `json:"status"`
	Data   []Collection `json:"data"`
}

type Collection struct {
	Uuid string `json:"uuid"`
	Name string `json:"displayName"`
	Icon string `json:"displayIcon"`
}

func (cl *CollectionResponse) GetAllEntity() (CollectionResponse, error) {
	url := "https://valorant-api.com/v1/bundles"

	resp, err := http.Get(url)
	if err != nil {
		return CollectionResponse{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return CollectionResponse{}, err
	}

	var collectionResponse CollectionResponse
	err = json.Unmarshal(body, &collectionResponse)
	if err != nil {
		return CollectionResponse{}, err
	}

	return collectionResponse, nil
}

func (cl *CollectionResponse) GetRandomEntity(collections CollectionResponse) Collection {
	return collections.Data[rand.Intn(len(collections.Data))]
}

func (cl *CollectionResponse) Handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	arrCollection, err := cl.GetAllEntity()
	errs.CheckErr("Can't get all collections", err)

	randSpray := cl.GetRandomEntity(arrCollection)

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
	errs.CheckErr("Can't send photo", err)
}
