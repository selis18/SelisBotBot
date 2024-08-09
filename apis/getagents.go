package apis

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type AgentResponse struct {
	Status int     `json:"status"`
	Data   []Agent `json:"data"`
}

type Agent struct {
	Uuid     string `json:"uuid"`
	Name     string `json:"displayName"`
	Icon     string `json:"displayIcon"`
	Portrait string `json:"fullPortrait"`
}

func (ar *AgentResponse) GetAllEntity() (AgentResponse, error) {
	url := "https://valorant-api.com/v1/agents"

	resp, err := http.Get(url)
	if err != nil {
		return AgentResponse{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return AgentResponse{}, err
	}

	var agentResponse AgentResponse
	err = json.Unmarshal(body, &agentResponse)
	if err != nil {
		return AgentResponse{}, err
	}
	return agentResponse, nil
}

func (ar *AgentResponse) GetRandomEntity(agents AgentResponse) Agent {
	return agents.Data[rand.Intn(len(agents.Data))]
}

func (ar *AgentResponse) GetTeamAgents(agents AgentResponse) string {
	str := ""
	count := 5
	for i := 0; i < count; i++ {
		agent := ar.GetRandomEntity(agents)
		if !strings.Contains(str, agent.Name) {
			str += agent.Name + ", "
		} else {
			count++
		}
	}
	return str
}

func (ar *AgentResponse) Handler(ctx context.Context, b *bot.Bot, update *models.Update) {

	arrAgents, err := ar.GetAllEntity()
	if err != nil {
		fmt.Println(err)
	}

	randAgent := ar.GetRandomEntity(arrAgents)
	str := ""
	for i := 0; i < 5; i++ {
		randAgent = ar.GetRandomEntity(arrAgents)
		str += randAgent.Name + ", "
	}
	fmt.Println(str)
	photo := randAgent.Icon

	// Параметры для отправки сообщения с изображением и текстом
	params := &bot.SendPhotoParams{
		ChatID:    update.Message.Chat.ID,
		Photo:     &models.InputFileString{Data: photo},
		Caption:   "Самый лучший на " + randAgent.Name,
		ParseMode: models.ParseModeMarkdown,
	}

	// Отправка сообщения с изображением
	_, err = b.SendPhoto(ctx, params)
	if err != nil {
		log.Fatalf("Error sending photo: %v", err)
	}
}

func (ar *AgentResponse) TeamHandler(ctx context.Context, b *bot.Bot, update *models.Update) {

	arrAgents, err := ar.GetAllEntity()
	if err != nil {
		fmt.Println(err)
	}

	params := &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      "Рандомная тима: " + ar.GetTeamAgents(arrAgents),
		ParseMode: models.ParseModeMarkdown,
	}

	_, err = b.SendMessage(ctx, params)
	if err != nil {
		log.Fatalf("Error sending photo: %v", err)
	}
}
