package catalog

import "app/db"

type RPC struct {
	Mongo db.Mongo
}

const (
	STATUS_DELETED string = "STATUS_DELETED"
)
