package mongo

import (
	"os"
	"time"

	mgo "gopkg.in/mgo.v2"
)

var instance *mgo.Session

var MONGO_HOST string
var MONGO_DATABASE string
var MONGO_USER string
var MONGO_PASSWORD string
var MONGO_DATABASE_AUTH string

var MONGO_ARTICLE_COLLECTION string

func init() {
	MONGO_DATABASE = os.Getenv("MONGO_DATABASE")
	MONGO_HOST = os.Getenv("DB_HOST")
	MONGO_USER = os.Getenv("MONGO_USER")
	MONGO_PASSWORD = os.Getenv("MONGO_PASSWORD")
	MONGO_DATABASE_AUTH = os.Getenv("MONGO_DATABASE_AUTH")
}

// MONGO DB
func setupConnection() (*mgo.Session, error) {
	MONGO_CONNECTION_PARAMS := &mgo.DialInfo{
		Addrs:    []string{MONGO_HOST},
		Timeout:  60 * time.Second,
		Database: MONGO_DATABASE,
		Username: MONGO_USER,
		Password: MONGO_PASSWORD,
		Source:   MONGO_DATABASE_AUTH,
	}

	return mgo.DialWithInfo(MONGO_CONNECTION_PARAMS)
}

func Session() (*mgo.Session, error) {
	if instance == nil {
		instance, err = setupConnection()
	}

	return instance, instanceError
}

func getCollection(collection string) (*mgo.Collection, error) {
	session, err := Session()
	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Monotonic, true)

	return session.DB(MONGO_DATABASE).C(collection), nil
}
