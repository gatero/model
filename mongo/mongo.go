package mongo

import (
	"os"
	"time"

	mgo "gopkg.in/mgo.v2"
)

// GLOBAL CONFIGURATION VARS

var MONGO_AUTH string
var MONGO_CONTAINER string
var MONGO_DATABASE string
var MONGO_PASSWORD string
var MONGO_USERNAME string

func init() {
	MONGO_AUTH = os.Getenv("MONGO_AUTH")
	MONGO_CONTAINER = os.Getenv("MONGO_CONTAINER")
	MONGO_DATABASE = os.Getenv("MONGO_DATABASE")
	MONGO_PASSWORD = os.Getenv("MONGO_INITDB_PASSWORD")
	MONGO_USERNAME = os.Getenv("MONGO_INITDB_USERNAME")
}

// Get the session pointer

var session *mgo.Session

func getSession() (*mgo.Session, error) {
	if session == nil {
		MONGO_CONNECTION_PARAMS := &mgo.DialInfo{
			Addrs:    []string{MONGO_CONTAINER},
			Database: MONGO_DATABASE,
			Password: MONGO_PASSWORD,
			Source:   MONGO_AUTH,
			Timeout:  60 * time.Second,
			Username: MONGO_USERNAME,
		}

		session, err = mgo.DialWithInfo(MONGO_CONNECTION_PARAMS)
		if err != nil {
			return nil, err
		}
	}

	return session, nil
}

// Get the collection pointer

var collection string
var currentCollectionName string

func getCollection(collectionName string) (*mgo.Collection, error) {
	if collection == nil {
		session, err := getSession()
		if err != nil {
			return nil, err
		}

		session.SetMode(mgo.Monotonic, true)

		if collectionName != currentCollectionName {
			currentCollectionName = collectionName
		}

		collection, err = session.DB(MONGO_DATABASE).C(currentCollectionName)
		if err != nil {
			return nil, err
		}
	}

	return collection, nil
}
