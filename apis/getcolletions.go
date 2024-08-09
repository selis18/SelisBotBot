package apis

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
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
	arrSprays, err := cl.GetAllEntity()
	if err != nil {
		fmt.Println(err)
	}

	randSpray := cl.GetRandomEntity(arrSprays)

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
