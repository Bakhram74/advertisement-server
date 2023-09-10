package handler

import (
	"github.com/Bakhram74/advertisement-server.git/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)

	}

	return router
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
