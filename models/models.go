package models

import (
	"time"
)

type (
	Card struct {
		ID                   string
		Name                 string
		ListName             string
		IdList               string
		TimeGuessForDone     int
		TimeRealForDone      int
		DateLastActivity     *time.Time
		Due                  *time.Time
		ChangeDueDate        bool
		HistoryChangeDueDate []*time.Time
	}
	Board struct {
		Id string
	}
	TrelloAPI struct {
		Key   string
		Token string
	}
)
