package controllers

import (
	"../service"
	"../utils"
	"github.com/gin-gonic/gin"
)

type Card struct{}

var serverBoard = new(service.Board)
var utilUrl = new(utils.UtilUrl)

func (card Card) AllCardReview(c *gin.Context) {
	q := c.Request.URL.Query()
	c.JSON(serverBoard.GetAllCardsOnReview(q))
}

func (card Card) AllCardChangeDueDate(c *gin.Context) {
	q := c.Request.URL.Query()
	c.JSON(serverBoard.GetAllCardOnChangeDue(q))
}
