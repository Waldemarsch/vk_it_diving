package api

import (
	"bytes"
	"fmt"
	"github.com/Waldemarsch/vk_it_diving/gateway/internal/models"
	"io"
	"net/http"
)

func (a *API) RequestMessagesAPI(event *models.Event) string {
	requestBody := event.MarshalToJSON()
	responseBody := bytes.NewBuffer(requestBody)
	response, err := http.Post("localhost:8003", "application/json", responseBody)
	if err != nil {
		fmt.Println("error in API: ", err)
		return ""
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error in API: ", err)
		return ""
	}
	return string(body)
}
