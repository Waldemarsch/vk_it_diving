package handler

import (
	"github.com/Waldemarsch/vk_it_diving/gateway/internal/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) InitRoutes() *gin.Engine {
	api := gin.New()

	api.POST("/", h.PostCallback)

	return api
}

func (h *Handler) PostCallback(c *gin.Context) {
	event := new(models.Event)
	err := c.BindJSON(&event)
	if err != nil {
		log.Println("error in Handler: ", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if event.Type == "confirmation" && event.GroupId == 225066783 {
		c.String(http.StatusAccepted, "1f9d7f10")
		return
	}
	status := h.s.Routing.ServeRequest(event)
	c.String(http.StatusOK, status)
}
