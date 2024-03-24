package service

import (
	"github.com/Waldemarsch/vk_it_diving/gateway/internal/infrastructure"
	"github.com/Waldemarsch/vk_it_diving/gateway/internal/models"
	"github.com/Waldemarsch/vk_it_diving/gateway/internal/service/routing"
)

type Routing interface {
	ServeRequest(event *models.Event) string
}

type Service struct {
	Routing
}

func NewService(inf *infrastructure.Infrastructure) *Service {
	return &Service{
		Routing: routing.NewRouting(inf),
	}
}
