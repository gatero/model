package main

import "os"

// ENVIRONMENT VARS
// ARTICLE
var (
	CATALOG_CONTAINER string = os.Getenv("CATALOG_CONTAINER")
	CATALOG_PORT      string = os.Getenv("CATALOG_PORT")
	CATALOG_SERVICE   string = CATALOG_CONTAINER + ":" + CATALOG_PORT
)
