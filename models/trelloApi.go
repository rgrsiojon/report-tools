package models

import (
	"github.com/adlio/trello"
)

//@ Get all cardon board from trello api
func (trll TrelloAPI) GetCardsOnTrelloAPI(ID string) ([]*trello.Card, error) {
	client := trello.NewClient(trll.Key, trll.Token)
	board, err := client.GetBoard(ID, trello.Defaults())
	if err != nil {
		return nil, err
	}
	cards, err := board.GetCards(trello.Defaults())
	if err != nil {
		return nil, err
	}
	return cards, nil
}

//@ Get data list with id list, key app and token
func (trll TrelloAPI) GetListbByIdOnTrelloAPI(ID string) (*trello.List, error) {
	client := trello.NewClient(trll.Key, trll.Token)
	list, err := client.GetList(ID, trello.Defaults())
	if err != nil {
		return nil, err
	}
	return list, nil
}
