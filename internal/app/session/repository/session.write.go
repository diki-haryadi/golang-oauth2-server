package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang-standards-project-layout/internal/app/session/model"
)

func (x *Module) StoreSession(ctx context.Context, session *model.SessionNoSqlSchema) error {
	//ctx, span := tracer.StartSpan(ctx, "repo.session.StoreSession", nil)
	//defer span.End()

	if session.Id == primitive.NilObjectID {
		session.Id = primitive.NewObjectIDFromTimestamp(time.Now())
	}

	var (
		e error
		c *mongo.Collection
	)

	c = x.mongoManager.Collection(model.SessionCollectionName)

	_, e = c.InsertOne(ctx, session)

	if e != nil {
		return e
	}

	return nil
}

func (x *Module) UpdateSession(ctx context.Context, id primitive.ObjectID, session *model.SessionNoSqlSchema) error {
	//ctx, span := tracer.StartSpan(ctx, "repo.session.UpdateSession", nil)
	//defer span.End()

	var (
		e error
		c *mongo.Collection
	)

	c = x.mongoManager.Collection(model.SessionCollectionName)

	_, e = c.UpdateByID(ctx, id, bson.M{"$set": session})

	if e != nil {
		return e
	}

	return nil
}
