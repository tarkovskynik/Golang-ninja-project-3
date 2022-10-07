package v1

import (
	"github.com/tarkovskynik/Golang-ninja-project-3/internal/services"

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

func (h *Handlers) InitV1(api *gin.RouterGroup) *gin.RouterGroup {
	// TODO: add handlers for auth
	//authAPI := api.Group("/auth")
	//{
	//	authAPI.POST("/signup", h.SignUp)
	//}

	return api
}
