package models

import (
	"strings"
	"time"

	"../utils"
)

var utilTime = new(utils.UtilTime)

func (Board) ParerCardsWithNameList(cards []Card, result []string) []interface{} {
	var data []interface{}
	for _, v := range result {
		data = append(data, filter(cards, func(item Card) bool {
			return strings.ToLower(v) == strings.ToLower(item.ListName)
		}))
	}
	return data
}

func (Board) FilterCardWithTime(mc []Card, count int) []Card {
	now := time.Now()
	return filter(mc, func(i Card) bool {
		return func() int {
			return int(i.DateLastActivity.Sub(now).Hours() / 24)
		}()+count > 0
	})
}

func (Board) FilterChangeDue(mc []Card) []Card {
	return filter(mc, func(i Card) bool {
		return i.ChangeDueDate
	})
}

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

//@ Filter []modules.MyCard
func filter(vs []Card, f func(Card) bool) (vsf []Card) {
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return
}
