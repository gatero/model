package datastore

import (
	c "app/src/constants"
	"time"

	mgo "gopkg.in/mgo.v2"
)

// Get the session pointer

type Mongo struct {
	Session        *mgo.Session
	Collection     *mgo.Collection
	CollectionName string
}

func (mongo *Mongo) SetSession() error {
	MONGO_CONNECTION_PARAMS := &mgo.DialInfo{
		Addrs:    []string{c.MONGO_CONTAINER},
		Database: c.MONGO_DATABASE,
		Password: c.MONGO_PASSWORD,
		Source:   c.MONGO_DATABASE_AUTH,
		Timeout:  time.Minute,
		Username: c.MONGO_USERNAME,
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
func (mongo *Mongo) SetCollection(collectionName string) error {
	// if currentCollectionName and collectionName are different
	mongo.CollectionName = collectionName
	err := mongo.SetSession()
	if err != nil {
		return err
	}

	index := mgo.Index{
		Key:        []string{"providerId"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	// get collection pointer
	mongo.Collection = mongo.Session.DB(c.MONGO_DATABASE).C(mongo.CollectionName)
	mongo.Collection.EnsureIndex(index)

	return nil
}
