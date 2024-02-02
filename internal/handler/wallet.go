package handler

import (
	"errors"
	"github.com/Warh40k/infotecs_task/internal/app"
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
	wallet, err := h.services.GetWallet(walletId)
	var notFound *app.NotFoundError
	if err != nil {
		if errors.As(err, &notFound) {
			//c.AbortWithStatusJSON(http.StatusNotFound, statusResponse{err.Error()})
			c.AbortWithStatus(http.StatusNotFound)

		} else {
			//c.AbortWithStatusJSON(http.StatusBadRequest, statusResponse{err.Error()})
			c.AbortWithStatus(http.StatusBadRequest)
		}
		return
	}

	c.JSON(http.StatusOK, wallet)
}

func (h *Handler) getWalletHistory(c *gin.Context) {
	walletId := c.Param("walletId")

	trs, err := h.services.GetWalletHistory(walletId)
	if err != nil {
		//c.AbortWithStatusJSON(http.StatusNotFound, statusResponse{err.Error()})
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, trs)
}

func (h *Handler) sendMoney(c *gin.Context) {
	var transaction domain.Transaction
	transaction.From = c.Param("walletId")

	if err := c.BindJSON(&transaction); err != nil {
		//c.AbortWithStatusJSON(http.StatusBadRequest, statusResponse{"error: incorrect transaction format"})
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if transaction.From == transaction.To {
		//c.AbortWithStatusJSON(http.StatusBadRequest, statusResponse{"error: sender's and receivers's ids are the same"})
		c.AbortWithStatus(http.StatusOK)
		return
	}

	err := h.services.SendMoney(transaction)

	var notFound app.NotFoundError
	var badRequest app.BadRequestError
	if err != nil {
		if errors.As(err, &notFound) {
			//c.AbortWithStatusJSON(http.StatusNotFound, statusResponse{err.Error()})
			c.AbortWithStatus(http.StatusNotFound)
		} else if errors.As(err, &badRequest) {
			//c.AbortWithStatusJSON(http.StatusBadRequest, statusResponse{err.Error()})
			c.AbortWithStatus(http.StatusNotFound)
		}
		return
	}

	c.Status(http.StatusOK)
}
