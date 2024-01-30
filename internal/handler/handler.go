package handler

import (
	"github.com/Warh40k/infotecs_task/internal/service"
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
	wallet := router.Group("/api/v1/wallet")
	{
		wallet.POST("/", h.createWallet)
		wallet.POST("/:walletId/send", h.sendMoney)
		wallet.GET("/:walletId", h.getWallet)
		wallet.GET("/:walletId/history", h.getHistory)
	}

	return router
}
