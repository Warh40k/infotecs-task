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
	/*	if walletId == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}*/
	wallet, err := h.services.GetWallet(walletId)
	var notFound *app.NotFoundError
	if err != nil {
		if errors.As(err, &notFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, statusResponse{err.Error()})
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, statusResponse{err.Error()})
		}
		return
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
		c.AbortWithStatusJSON(http.StatusBadRequest, &app.BadRequestError{
			Message: "Некорректный формат перевода",
		})
		return
	}

	if transaction.From == transaction.To {
		c.AbortWithStatusJSON(http.StatusBadRequest, &app.BadRequestError{
			Message: "Совпадение id отправителя и получателя",
		})
		return
	}

	err := h.services.SendMoney(transaction)

	var notFound app.NotFoundError
	var badRequest app.BadRequestError
	if err != nil {
		if errors.As(err, &notFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, statusResponse{err.Error()})
		} else if errors.As(err, &badRequest) {
			c.AbortWithStatusJSON(http.StatusBadRequest, statusResponse{err.Error()})
		}
		return
	}

	c.Status(http.StatusOK)
}
