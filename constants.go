package main

import "os"

// ENVIRONMENT VARS
// ARTICLE
var (
	MODEL_CONTAINER string = os.Getenv("MODEL_CONTAINER")
	MODEL_PORT      string = os.Getenv("MODEL_PORT")
	MODEL_SERVICE   string = MODEL_CONTAINER + ":" + MODEL_PORT
)
