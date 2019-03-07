package delivery

import (
	"net/url"
	"strings"

	"time"

	"../config"
	"../models"
	"../utils"
	"github.com/adlio/trello"
)

type Board struct{}

var Config = config.ReadConfig()
var utilUrl = new(utils.UtilUrl)
var UtilString = new(utils.UtilString)
var UtilTime = new(utils.UtilTime)
var TrelloAPI = models.TrelloAPI{
	Config.App.Keyapp,
	Config.App.Token,
}

func (Board) FilterCardsOnReview(cards []models.Card, q url.Values) (result [][]models.Card) {
	listNames, count := utilUrl.HandelResQuery(q)

	now := time.Now()
	filterTime := filter(cards, func(i models.Card) bool {
		return func() int {
			return int(i.DateLastActivity.Sub(now).Hours() / 24)
		}()+count > 0
	})

	for _, v := range listNames {
		result = append(result, filter(filterTime, func(item models.Card) bool {
			return strings.ToLower(v) == strings.ToLower(item.ListName)
		}))
	}
	return
}

func (Board) FilterCardsOnChangeDue(cards []models.Card, q url.Values) []models.Card {
	_, count := utilUrl.HandelResQuery(q)
	now := time.Now()
	return filter(filter(cards, func(i models.Card) bool {
		return func() int {
			return int(i.DateLastActivity.Sub(now).Hours() / 24)
		}()+count > 0
	}), func(i models.Card) bool {
		return i.ChangeDueDate
	})
}

func (Board) ConvertCard(cardsAPI []*trello.Card) []models.Card {
	var cards []models.Card
	for _, value := range cardsAPI {
		card := models.Card{
			ID:               value.ID,
			Name:             value.Name,
			IdList:           value.IDList,
			DateLastActivity: value.DateLastActivity,
			Due:              value.Due,
		}
		list, err := TrelloAPI.GetListbByIdOnTrelloAPI(value.IDList)
		if err != nil {

		}
		card.ListName = list.Name
		card.TimeRealForDone = UtilString.GetRealTimeOfDone(value.Name)
		card.TimeGuessForDone = UtilString.GetTimeGuessForDone(value.Name)
		card.HistoryChangeDueDate = UtilTime.AppendTime(card.HistoryChangeDueDate, value.Due)
		cards = append(cards, card)
	}
	return cards
}

//@ Filter []modules.MyCard
func filter(vs []models.Card, f func(models.Card) bool) (vsf []models.Card) {
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return
}
