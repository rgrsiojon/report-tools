package models

import (
	"github.com/adlio/trello"
)

type TrelloAPI struct{}

//@ Get all cardon board from trello api
func (TrelloAPI) GetCardsOnTrelloAPI(key, token, id string) ([]*trello.Card, error) {
	client := trello.NewClient(key, token)
	board, err := client.GetBoard(id, trello.Defaults())
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
func (TrelloAPI) GetListbByIdOnTrelloAPI(appKey string, token string, ID string) (*trello.List, error) {
	client := trello.NewClient(appKey, token)
	list, err := client.GetList(ID, trello.Defaults())
	if err != nil {
		return nil, err
	}
	return list, nil
}
