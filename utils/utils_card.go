package utils

import (
	"strings"
	"time"

	"../models"
)

type UtilCard struct{}

var utilTime = new(UtilTime)

func (UtilCard) ParerCardsWithNameList(cards []models.MyCard, result []string) []interface{} {
	var data []interface{}
	for _, v := range result {
		data = append(data, filter(cards, func(item models.MyCard) bool {
			return strings.ToLower(v) == strings.ToLower(item.ListName)
		}))
	}
	return data
}

func (UtilCard) FilterCardWithTime(mc []models.MyCard, count int) []models.MyCard {
	now := time.Now()
	return filter(mc, func(i models.MyCard) bool {
		return func() int {
			return int(i.DateLastActivity.Sub(now).Hours() / 24)
		}()+count > 0
	})
}

func (UtilCard) FilterChangeDue(mc []models.MyCard) []models.MyCard {
	return filter(mc, func(i models.MyCard) bool {
		return i.ChangeDueDate
	})
}

func (UtilCard) CompareTwoCards(cardOnDb models.MyCard, cardOnTrello models.MyCard) models.MyCard {
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
func filter(vs []models.MyCard, f func(models.MyCard) bool) (vsf []models.MyCard) {
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return
}
