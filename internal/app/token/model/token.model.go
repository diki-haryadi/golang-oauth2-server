package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	sam "golang-standards-project-layout/internal/integration/spotify_api/model"
)

type TokenNoSqlSchema struct {
	Id           primitive.ObjectID `bson:"_id"`
	UserId       primitive.ObjectID `bson:"user_id"`
	AccessToken  string             `bson:"access_token"`
	RefreshToken string             `bson:"refresh_token"`
	ExpiresAt    time.Time          `bson:"expires_in"`
	TokenType    string             `bson:"token_type"`
	Scope        string             `bson:"scope"`
	Issuer       string             `bson:"issuer"`
}

func (x *TokenNoSqlSchema) BuildFromSpotifyAuthorizeBodyRes(spr *sam.SpotifyAuthorizeBodyRes, userId primitive.ObjectID) *TokenNoSqlSchema {
	if spr.RefreshToken == "" {
		spr.RefreshToken = x.RefreshToken
	}

	return &TokenNoSqlSchema{
		Id:           primitive.NewObjectIDFromTimestamp(time.Now()),
		UserId:       userId,
		AccessToken:  spr.AccessToken,
		RefreshToken: spr.RefreshToken,
		ExpiresAt:    time.Now().Add(time.Duration(spr.ExpiresIn) * time.Second),
		TokenType:    spr.TokenType,
		Scope:        spr.Scope,
		Issuer:       "spotify",
	}
}
