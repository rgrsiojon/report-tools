package controllers

import (
	"fmt"
	"sync"

	"../config"
	"../models"

	"../store"
	"../utils"
	"github.com/adlio/trello"
)

var wg sync.WaitGroup
var UtilString = new(utils.UtilString)
var UtilTime = new(utils.UtilTime)

type RoutineCard struct{}

//@ Call two go routine
func (RoutineCard) UpdateDataOnDB() {
	Config := config.ReadConfig()
	wg.Add(1)
	chanCard := make(chan []*trello.Card, 3)
	go WriteData(chanCard, Config.App.Keyapp, Config.App.Token, Config.App.Idboard)
	go HandelData(chanCard, Config.App.Keyapp, Config.App.Token)
	wg.Wait()
}

// @ Write data on chanel card
func WriteData(chanCard chan []*trello.Card, key, token, id string) {
	var TrelloAPI = models.TrelloAPI{
		Key:   key,
		Token: token,
	}
	for {
		cards, err := TrelloAPI.GetCardsOnTrelloAPI(id)
		if err != nil {
			chanCard <- nil
		} else {
			chanCard <- cards
		}
	}
}

//@ Handel data on chanel card
func HandelData(chanCard chan []*trello.Card, key, token string) {
	var TrelloAPI = models.TrelloAPI{
		Key:   key,
		Token: token,
	}
	for {
		cards := <-chanCard
		if cards != nil {
			for _, value := range cards {
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

				result, err := store.FindOne(card.ID)
				if err != nil {
					store.InsertData(card, func(err error) {
						if err != nil {
							fmt.Println("Can't insert")
						}
						fmt.Println("Inserted !")
					})
				} else {
					newCard := board.CompareTwoCards(result, card)
					err := store.UpdateCard(newCard.ID, newCard)
					if err != nil {
						fmt.Println("Can't Update")
					} else {
						fmt.Println("Updated !")
					}
				}
			}
		} else {
			fmt.Println("false")
		}
	}
}
