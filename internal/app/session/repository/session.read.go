package repository

import (
	"context"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang-standards-project-layout/internal/app/session/model"
)

func (x *Module) FindSessionById(ctx context.Context, id primitive.ObjectID) (*model.SessionNoSqlSchema, error) {
	//ctx, span := tracer.StartSpan(ctx, "repo.session.FindSessionById", nil)
	//defer span.End()

	var (
		e    error
		c    *mongo.Collection
		f    = bson.D{{"_id", id}}
		res  *mongo.SingleResult
		data = &model.SessionNoSqlSchema{}
	)

	c = x.mongoManager.Collection(model.SessionCollectionName)

	res = c.FindOne(ctx, f)

	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, res.Err()
	}

	e = res.Decode(data)
	if e != nil {
		log.Err(e).Msg(e.Error())
		return nil, e
	}

	return data, nil
}

func (x *Module) FindSessionByCode(ctx context.Context, code int64) (*model.SessionNoSqlSchema, error) {
	//ctx, span := tracer.StartSpan(ctx, "repo.session.FindSessionByCode", nil)
	//defer span.End()

	var (
		e    error
		c    *mongo.Collection
		f    = bson.D{{"code", code}}
		res  *mongo.SingleResult
		data = &model.SessionNoSqlSchema{}
	)

	c = x.mongoManager.Collection(model.SessionCollectionName)

	res = c.FindOne(ctx, f)

	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, res.Err()
	}

	e = res.Decode(data)
	if e != nil {
		log.Err(e).Msg(e.Error())
		return nil, e
	}

	return data, nil
}
