package routing

import (
	"github.com/Waldemarsch/vk_it_diving/gateway/internal/infrastructure"
	"github.com/Waldemarsch/vk_it_diving/gateway/internal/models"
)

type Routing struct {
	inf *infrastructure.Infrastructure
}

func NewRouting(inf *infrastructure.Infrastructure) *Routing {
	return &Routing{
		inf: inf,
	}
}

func (s *Routing) ServeRequest(event *models.Event) string {
	var status string
	if event.Type[:5] == "group" {
		status = s.inf.API.RequestGroupAPI(event)
	} else if event.Type[:7] == "message" {
		status = s.inf.API.RequestMessagesAPI(event)
	}
	return status
}
