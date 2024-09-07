package repository

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang-standards-project-layout/internal/app/user/model"
	"golang-standards-project-layout/internal/constants"
)

func (x *Module) StoreUser(ctx context.Context, user *model.UserNoSqlSchema) error {
	//ctx, span := tracer.StartSpan(ctx, "repo.user.GetUserInfo", nil)
	//defer span.End()

	if user.Id == primitive.NilObjectID {
		user.Id = primitive.NewObjectIDFromTimestamp(time.Now())
	}

	var (
		e error
		c *mongo.Collection
	)

	c = x.mongoManager.Collection(model.UserCollectionName)

	_, e = c.InsertOne(ctx, user)

	if e != nil {
		return e
	}

	return nil
}

func (x *Module) UpdateUser(ctx context.Context, id primitive.ObjectID, user *model.UserNoSqlSchema) error {
	//ctx, span := tracer.StartSpan(ctx, "repo.user.GetUserInfo", nil)
	//defer span.End()

	var (
		e error
		c *mongo.Collection
	)

	c = x.mongoManager.Collection(model.UserCollectionName)

	_, e = c.UpdateByID(ctx, id, bson.M{"$set": user})

	if e != nil {
		return e
	}

	return nil
}

func (x *Module) FindUserById(ctx context.Context, id primitive.ObjectID) (*model.UserNoSqlSchema, error) {
	//ctx, span := tracer.StartSpan(ctx, "repo.user.FindUserById", nil)
	//defer span.End()

	var (
		e    error
		c    *mongo.Collection
		f    = bson.D{{"_id", id}}
		res  *mongo.SingleResult
		data = &model.UserNoSqlSchema{}
	)

	c = x.mongoManager.Collection(model.UserCollectionName)

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

func (x *Module) FindUserByChatbotUserId(ctx context.Context, chatbotUserId string, chatbotChannel constants.ChatbotChannelEnum) (*model.UserNoSqlSchema, error) {
	//ctx, span := tracer.StartSpan(ctx, "repo.user.FindUserByChatbotUserId", nil)
	//defer span.End()

	var (
		e    error
		c    *mongo.Collection
		f    = bson.D{{"chatbot_user_id", chatbotUserId}, {"chatbot_channel", chatbotChannel}}
		res  *mongo.SingleResult
		data = &model.UserNoSqlSchema{}
	)

	c = x.mongoManager.Collection(model.UserCollectionName)

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
