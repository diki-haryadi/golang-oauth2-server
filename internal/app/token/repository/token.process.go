package repository

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang-standards-project-layout/internal/app/token/model"
	sam "golang-standards-project-layout/internal/integration/spotify_api/model"
)

func (x *Module) StoreToken(ctx context.Context, token *model.TokenNoSqlSchema) error {
	//ctx, span := tracer.StartSpan(ctx, "api_call.spotify.GetUserInfo", nil)
	//defer span.End()

	var (
		e error
		c *mongo.Collection
	)

	c = x.mongoManager.Collection(model.TokenCollectionName)

	_, e = c.InsertOne(ctx, token)

	if e != nil {
		return e
	}

	return nil
}

func (x *Module) updateToken(ctx context.Context, id primitive.ObjectID, token *model.TokenNoSqlSchema) error {
	//ctx, span := tracer.StartSpan(ctx, "api_call.spotify.GetUserInfo", nil)
	//defer span.End()

	var (
		e error
		c *mongo.Collection
	)

	c = x.mongoManager.Collection(model.TokenCollectionName)

	token.Id = id
	_, e = c.UpdateByID(ctx, id, bson.M{"$set": token})

	if e != nil {
		return e
	}

	return nil
}

func (x *Module) FindAndValidateTokenByUserId(ctx context.Context, id primitive.ObjectID) (*model.TokenNoSqlSchema, error) {
	//ctx, span := tracer.StartSpan(ctx, "repo.token.FindAndValidateTokenByUserId", nil)
	//defer span.End()

	var (
		e     error
		c     *mongo.Collection
		f     = bson.D{{"user_id", id}}
		res   *mongo.SingleResult
		data  = &model.TokenNoSqlSchema{}
		stres *sam.SpotifyAuthorizeBodyRes
	)

	c = x.mongoManager.Collection(model.TokenCollectionName)

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

	if time.Now().After(data.ExpiresAt) {
		stres, e = x.spotifyAuthApiCall.RefreshToken(ctx, data.RefreshToken)
		if e != nil {
			log.Err(e).Msg(e.Error())
			return nil, e
		}

		udata := data.BuildFromSpotifyAuthorizeBodyRes(stres, data.UserId)

		e = x.updateToken(ctx, data.Id, udata)

		if e != nil {
			log.Err(e).Msg(e.Error())
			return nil, e
		}
	}

	return data, nil
}
