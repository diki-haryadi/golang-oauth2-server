package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Module struct {
	mongoManager *mongo.Database
}

type Opts struct {
	MongoManager *mongo.Database
}

func New(o Opts) *Module {
	return &Module{
		mongoManager: o.MongoManager,
	}
}
