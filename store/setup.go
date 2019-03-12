package store

import (
	"gopkg.in/mgo.v2"
)

var DatabaseName = "mongodb"

type DB struct {
	*mgo.Session
}

func NewDatabase() (*DB, error) {
	if DatabaseName == "mongodb" {
		db, err := mongoDB("roger", "roger123")
		if err != nil {
			return &DB{nil}, nil
		}
		return &DB{db}, nil
	}
	return &DB{nil}, nil
}

func mongoDB(username, password string) (*mgo.Session, error) {
	session, err := mgo.Dial("mongodb://" + username + ":" + password + "@ds213255.mlab.com:13255/report-tools")
	if err != nil {
		return nil, err
	}
	session.SetMode(mgo.Monotonic, true)
	return session, nil
}
