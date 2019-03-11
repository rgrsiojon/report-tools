package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rgrsiojon/report-tools/service"
	"github.com/rgrsiojon/report-tools/utils"
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
