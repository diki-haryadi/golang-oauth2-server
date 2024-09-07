package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SessionNoSqlSchema struct {
	Id        primitive.ObjectID `bson:"_id"`
	Code      string             `bson:"code"`
	HostId    primitive.ObjectID `bson:"host_id"`
	CreatedAt time.Time          `bson:"created_at"`
	ExpireAt  time.Time          `bson:"expire_at"`
}
