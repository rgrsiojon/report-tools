package models

import (
	"time"
)

type (
	MyCard struct {
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
)
