package controllers

import (
	"../models"
	"../store"
	"../utils"
	"github.com/gin-gonic/gin"
)

type Card struct{}

var board = models.Board{}
var utilUrl = new(utils.UtilUrl)

func (card Card) AllCardReview(c *gin.Context) {
	cards, err := store.GetAllCard()
	if err != nil {
		c.JSON(200, gin.H{"Err": "Not Found"})
	}
	listNames, time := utilUrl.HandelResQuery(c)

	filterTime := board.FilterCardWithTime(cards, time)
	if listNames == nil {
		c.JSON(200, filterTime)
	}
	result := board.ParerCardsWithNameList(filterTime, listNames)
	c.JSON(200, result)
}

func (card Card) AllCardChangeDueDate(c *gin.Context) {
	cards, err := store.GetAllCard()
	if err != nil {
		c.JSON(200, gin.H{"Err": "Not Found"})
	}
	_, time := utilUrl.HandelResQuery(c)
	filterTime := board.FilterCardWithTime(cards, time)

	result := board.FilterChangeDue(filterTime)
	c.JSON(200, result)
}
