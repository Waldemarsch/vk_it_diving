package infrastructure

import (
	"github.com/Waldemarsch/vk_it_diving/gateway/internal/infrastructure/api"
	"github.com/Waldemarsch/vk_it_diving/gateway/internal/models"
)

type API interface {
	RequestGroupAPI(event *models.Event) string
	RequestMessagesAPI(event *models.Event) string
}

type Infrastructure struct {
	API
}

func NewInfrastructure() *Infrastructure {
	return &Infrastructure{
		API: api.NewAPI(),
	}
}
