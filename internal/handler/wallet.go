package handler

import (
	"errors"
	"github.com/Warh40k/infotecs_task/internal/domain"
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
	/*	if walletId == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}*/
	wallet, err := h.services.GetWallet(walletId)
	var notFound *domain.NotFoundError
	if err != nil {
		if errors.As(err, &notFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, err)
			return
		} else {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
	}

	c.JSON(http.StatusOK, wallet)
}

func (h *Handler) getHistory(c *gin.Context) {

}

func (h *Handler) sendMoney(c *gin.Context) {
	var transaction domain.Transaction
	walletId := c.Param("walletId")
	transaction.From = walletId

	if err := c.BindJSON(&transaction); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if transaction.From == transaction.To {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := h.services.SendMoney(transaction)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.Status(http.StatusOK)
}
