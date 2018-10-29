package db

import (
	"os"
	"time"

	mgo "gopkg.in/mgo.v2"
)

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

type Mongo struct {
	Session        *mgo.Session
	Collection     *mgo.Collection
	CollectionName string
}

var session *mgo.Session

func (mongo *Mongo) SetSession() error {
	MONGO_CONNECTION_PARAMS := &mgo.DialInfo{
		Addrs:    []string{MONGO_CONTAINER},
		Database: MONGO_DATABASE,
		Password: MONGO_PASSWORD,
		Source:   MONGO_DATABASE_AUTH,
		Timeout:  time.Minute,
		Username: MONGO_USERNAME,
	}

	newSession, err := mgo.DialWithInfo(MONGO_CONNECTION_PARAMS)
	if err != nil {
		return err
	}

	newSession.SetMode(mgo.Monotonic, true)

	mongo.Session = newSession

	return nil
}

// Get the collection pointer
func (mongo *Mongo) SetCollection() error {
	// if currentCollectionName and collectionName are different

	err := mongo.SetSession()
	if err != nil {
		return err
	}

	// get collection pointer
	mongo.Collection = mongo.Session.DB(MONGO_DATABASE).C(mongo.CollectionName)

	return nil
}
