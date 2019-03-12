package controllers

import (
	"fmt"
	"sync"

	"github.com/rgrsiojon/report-tools/models"

	"github.com/adlio/trello"
	"github.com/rgrsiojon/report-tools/delivery"
	"github.com/rgrsiojon/report-tools/store"
	"github.com/rgrsiojon/report-tools/utils"
)

var keyapp = "91fc328d734261ff45775ebe00fdb13c"
var token = "0859930bb7d2177e2fd3a1a8c4bed14f253ba2f93b044e4b5e9e59241b9f61d3"
var idboard = "iCBtQXmr"

var wg sync.WaitGroup
var UtilTime = new(utils.UtilTime)
var board = new(models.Board)
var deliveryBoard = new(delivery.Board)

type RoutineCard struct{}

//@ Call two go routine
func (RoutineCard) UpdateDataOnDB() {
	wg.Add(1)
	chanCard := make(chan []*trello.Card, 3)
	go WriteData(chanCard, keyapp, token, idboard)
	go HandelData(chanCard, keyapp, token)
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
							// fmt.Println("Can't insert")
						}
						// fmt.Println("Inserted !")
					})
				} else {
					newCard := board.CompareTwoCards(cardOnDB, value)
					err := store.UpdateCard(newCard.ID, newCard)
					if err != nil {
						// handel
						// fmt.Println("Can't Update")
					} else {
						// handel
						// fmt.Println("Updated !")
					}
				}
			}
		} else {
			fmt.Println("false")
		}
	}
}
