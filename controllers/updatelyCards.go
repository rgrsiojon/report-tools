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
var TrelloAPI = new(models.TrelloAPI)
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
	for {
		cards, err := TrelloAPI.GetCardsOnTrelloAPI(key, token, id)
		if err != nil {
			chanCard <- nil
		} else {
			chanCard <- cards
		}
	}
}

//@ Handel data on chanel card
func HandelData(chanCard chan []*trello.Card, key, token string) {
	for {
		cards := <-chanCard
		if cards != nil {
			for _, value := range cards {
				myCard := models.MyCard{
					ID:               value.ID,
					Name:             value.Name,
					IdList:           value.IDList,
					DateLastActivity: value.DateLastActivity,
					Due:              value.Due,
				}
				list, err := TrelloAPI.GetListbByIdOnTrelloAPI(key, token, value.IDList)
				if err != nil {

				}
				myCard.ListName = list.Name
				myCard.TimeRealForDone = UtilString.GetRealTimeOfDone(value.Name)
				myCard.TimeGuessForDone = UtilString.GetTimeGuessForDone(value.Name)
				myCard.HistoryChangeDueDate = UtilTime.AppendTime(myCard.HistoryChangeDueDate, value.Due)

				result, err := store.FindOne(myCard.ID)
				if err != nil {
					store.InsertData(myCard, func(err error) {
						if err != nil {
							fmt.Println("Can't insert")
						}
						fmt.Println("Inserted !")
					})
				} else {
					newCard := utilCard.CompareTwoCards(result, myCard)
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
