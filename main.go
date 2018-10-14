package main

import (
	"log"
	"net"
	"os"
	"time"

	"app/catalog"
	pb "app/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

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

// CHANNELS
var currentSession *mgo.Session

func createSession(sessionChannel chan *mgo.Session, errorChannel chan error) {
	if currentSession == nil {
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
			errorChannel <- err
		}

		newSession.SetMode(mgo.Monotonic, true)
		currentSession = newSession
	}

	sessionChannel <- currentSession
}

func main() {
	sessionChannel := make(chan *mgo.Session)
	errorChannel := make(chan error)

	go createSession(sessionChannel, errorChannel)

	session := <-sessionChannel
	log.Printf("SESSION: %+v", session)
	lis, err := net.Listen("tcp", ":"+CATALOG_PORT)
	if err != nil {
		log.Fatalf("ERROR: Failed listening %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterCatalogServer(server, &catalog.RPC{})

	reflection.Register(server)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("ERROR: Failed to serve %v", err)
	}
}
