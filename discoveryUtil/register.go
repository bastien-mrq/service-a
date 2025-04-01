package discoveryutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/charmbracelet/log"
)

type RegisterPayload struct {
	Name string `json:"name"`
}
type RegisterResponse struct {
	Message string `json:"message"`
	Port    int    `json:"port"`
	Uuid    string `json:"uuid"`
}

func Register() RegisterResponse {
	payload := RegisterPayload{Name: "service-a"}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Error:%v", err)
	}

	resp, err := http.Post(Address+"/register", "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Fatalf("Error:%v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error:%v", err)
	}
	fmt.Printf("message : %v \n", string(body))

	var response RegisterResponse
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatalf("Error:%v", err)
	}
	return response
}
