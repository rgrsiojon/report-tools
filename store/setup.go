package store

import (
	"github.com/rgrsiojon/report-tools/config"
	"gopkg.in/mgo.v2"
)

var Config = config.ReadConfig()

type DB struct {
	*mgo.Session
}

func NewDatabase() (*DB, error) {
	if Config.Database.Name == "mongodb" {
		db, err := mongoDB(Config.Database.Username, Config.Database.Password)
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
