package client

import "github.com/drasko/edgex-export/mongo"

var repo *mongo.Repository

// InitMongoRepository - Init Mongo DB
func InitMongoRepository(r *mongo.Repository) {
	repo = r
	return
}
