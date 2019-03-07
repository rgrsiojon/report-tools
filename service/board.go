package service

import (
	"net/url"

	"../delivery"
	"../store"
)

type Board struct{}

var deliveryBoard = new(delivery.Board)

func (Board) GetAllCardsOnReview(q url.Values) (int, interface{}) {
	cards, err := store.GetAllCard()
	if err != nil {
		return 404, nil
	}
	return 200, deliveryBoard.FilterCardsOnReview(cards, q)
}

func (Board) GetAllCardOnChangeDue(q url.Values) (int, interface{}) {
	cards, err := store.GetAllCard()
	if err != nil {
		return 404, nil
	}
	return 200, deliveryBoard.FilterCardsOnChangeDue(cards, q)
}
