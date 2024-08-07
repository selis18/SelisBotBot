package agents

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
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

func GetAgents() (AgentResponse, error) {
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

	// for _, agent := range agentResponse.Data {
	// 	fmt.Println("Full Name:", agent.Name)
	// 	fmt.Println("Uuid:", agent.Uuid)
	// }

	//randAgent := GetRandomAgent(agentResponse)
	//fmt.Println("Random Agent:", randAgent.Name)
	return agentResponse, nil
}

func GetRandomAgent(agents AgentResponse) Agent {
	return agents.Data[rand.Intn(len(agents.Data))]
}
