package controllers

import (
	"fmt"
	"sync"

	"../config"
	"../models"

	"../delivery"
	"../store"
	"../utils"
	"github.com/adlio/trello"
)

var wg sync.WaitGroup
var UtilTime = new(utils.UtilTime)
var board = new(models.Board)
var deliveryBoard = new(delivery.Board)

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
	for {
		cards := <-chanCard
		if cards != nil {
			result := deliveryBoard.ConvertCard(cards)
			for _, value := range result {
				cardOnDB, err := store.FindOne(value.ID)
				if err != nil {
					store.InsertData(value, func(err error) {
						if err != nil {
							fmt.Println("Can't insert")
						}
						fmt.Println("Inserted !")
					})
				} else {
					newCard := board.CompareTwoCards(cardOnDB, value)
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
