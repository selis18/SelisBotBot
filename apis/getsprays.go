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

type SprayResponse struct {
	Status int     `json:"status"`
	Data   []Spray `json:"data"`
}

type Spray struct {
	Uuid string `json:"uuid"`
	Name string `json:"displayName"`
	Icon string `json:"fullTransparentIcon"`
}

func (sr *SprayResponse) GetAllEntity() (SprayResponse, error) {
	url := "https://valorant-api.com/v1/sprays"

	resp, err := http.Get(url)
	if err != nil {
		return SprayResponse{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return SprayResponse{}, err
	}

	err = json.Unmarshal(body, &sr)
	if err != nil {
		return SprayResponse{}, err
	}

	// for _, agent := range agentResponse.Data {
	// 	fmt.Println("Full Name:", agent.Name)
	// 	fmt.Println("Uuid:", agent.Uuid)
	// }

	//randAgent := GetRandomAgent(agentResponse)
	//fmt.Println("Random Agent:", randAgent.Name)
	return *sr, nil
}

func (sr *SprayResponse) GetRandomEntity(sprays SprayResponse) Spray {
	return sprays.Data[rand.Intn(len(sprays.Data))]
}

func (sr *SprayResponse) Handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	arrSprays, err := sr.GetAllEntity()
	errs.CheckErr("Can't get all sprays", err)

	randSpray := sr.GetRandomEntity(arrSprays)

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
