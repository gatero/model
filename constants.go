package main

import "os"

// ENVIRONMENT VARS
// ARTICLE
var (
	ARTICLE_CONTAINER string = os.Getenv("ARTICLE_CONTAINER")
	ARTICLE_PORT      string = os.Getenv("ARTICLE_PORT")
	ARTICLE_SERVICE   string = ARTICLE_CONTAINER + ":" + ARTICLE_PORT
)
