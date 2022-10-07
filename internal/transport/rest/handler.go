package rest

import (
	"github.com/tarkovskynik/Golang-ninja-project-3/internal/services"
	v1 "github.com/tarkovskynik/Golang-ninja-project-3/internal/transport/rest/v1"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	service *services.Service
}

func NewHandler(service *services.Service) *Handlers {
	return &Handlers{
		service: service,
	}
}

func (h *Handlers) Init() *gin.Engine {
	router := gin.Default()

	apiV1 := router.Group("/api/v1")
	{
		handlersV1 := v1.NewHandler(h.service)
		handlersV1.InitV1(apiV1)
	}

	return router
}
