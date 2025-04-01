package discoveryutil

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/charmbracelet/log"
)

type UnregisterPayload struct {
	Uuid string `json:"uuid"`
}

func UnregisterService(uuid string) {
	log.Info("Unregister in process\n")

	payload := UnregisterPayload{Uuid: uuid}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Error:%v", err)
	}
	resp, err := http.Post(Address+"/unregister", "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Errorf("Error:%v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		log.Info("Unregister finished")
	} else {
		log.Fatalf("Server responded with %v, %v", resp.StatusCode, resp.Body.Close().Error())
	}
}
