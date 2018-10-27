package mongo

import (
	"os"
	"time"

	mgo "gopkg.in/mgo.v2"
)

// GLOBAL CONFIGURATION VARS

var MONGO_CONTAINER string
var MONGO_DATABASE_AUTH string
var MONGO_DATABASE string
var MONGO_PASSWORD string
var MONGO_USERNAME string

func init() {
	MONGO_CONTAINER = os.Getenv("MONGO_CONTAINER")
	MONGO_DATABASE_AUTH = os.Getenv("MONGO_DATABASE_AUTH")
	MONGO_DATABASE = os.Getenv("MONGO_DATABASE")
	MONGO_PASSWORD = os.Getenv("MONGO_PASSWORD")
	MONGO_USERNAME = os.Getenv("MONGO_USER")
}

// Get the session pointer

var session *mgo.Session

func getSession() (*mgo.Session, error) {
	if session == nil {
		MONGO_CONNECTION_PARAMS := &mgo.DialInfo{
			Addrs:    []string{MONGO_CONTAINER},
			Database: MONGO_DATABASE,
			Password: MONGO_PASSWORD,
			Source:   MONGO_DATABASE_AUTH,
			Timeout:  60 * time.Second,
			Username: MONGO_USERNAME,
		}

		newSession, err := mgo.DialWithInfo(MONGO_CONNECTION_PARAMS)
		if err != nil {
			return nil, err
		}

		newSession.SetMode(mgo.Monotonic, true)
		session = newSession
	}

	return session, nil
}

// Get the collection pointer

var collection *mgo.Collection
var currentCollectionName string

func GetCollection(collectionName string) (*mgo.Collection, error) {
	// if currentCollectionName and collectionName are different
	if collectionName != currentCollectionName {
		// redefine collection name
		currentCollectionName = collectionName

		session, err := getSession()
		if err != nil {
			return nil, err
		}

		// get collection pointer
		collection = session.DB(MONGO_DATABASE).C(currentCollectionName)
	}

	return collection, nil
}
