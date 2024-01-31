package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createWallet(c *gin.Context) {
	wallet, err := h.services.CreateWallet()
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, wallet)
}

func (h *Handler) getWallet(c *gin.Context) {
	walletId := c.Param("walletId")
	if walletId == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	wallet, err := h.services.GetWallet(walletId)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, wallet)
}

func (h *Handler) getHistory(c *gin.Context) {

}

func (h *Handler) sendMoney(c *gin.Context) {

}
