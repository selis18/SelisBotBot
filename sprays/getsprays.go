package sprays

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
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

func GetSprays() (SprayResponse, error) {
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

	var sprayResponse SprayResponse
	err = json.Unmarshal(body, &sprayResponse)
	if err != nil {
		return SprayResponse{}, err
	}

	// for _, agent := range agentResponse.Data {
	// 	fmt.Println("Full Name:", agent.Name)
	// 	fmt.Println("Uuid:", agent.Uuid)
	// }

	//randAgent := GetRandomAgent(agentResponse)
	//fmt.Println("Random Agent:", randAgent.Name)
	return sprayResponse, nil
}

func GetRandomSpray(agents SprayResponse) Spray {
	return agents.Data[rand.Intn(len(agents.Data))]
}
