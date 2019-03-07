package models

import (
	"../utils"
)

var utilTime = new(utils.UtilTime)

//@ compare two cards
func (Board) CompareTwoCards(cardOnDb Card, cardOnTrello Card) Card {
	if utilTime.CompareTwoTime(cardOnDb.DateLastActivity, cardOnTrello.DateLastActivity) == false {
		cardOnDb.DateLastActivity = cardOnTrello.DateLastActivity
	}
	if utilTime.CompareTwoTime(cardOnDb.Due, cardOnTrello.Due) == false {
		cardOnDb.ChangeDueDate = true
		cardOnDb.Due = cardOnTrello.Due
		cardOnDb.HistoryChangeDueDate = utilTime.AppendTime(cardOnDb.HistoryChangeDueDate, cardOnTrello.Due)
	}
	return cardOnDb
}
