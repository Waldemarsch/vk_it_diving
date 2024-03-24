package handler

import "github.com/Waldemarsch/vk_it_diving/gateway/internal/service"

type Handler struct {
	s *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		s: s,
	}
}
