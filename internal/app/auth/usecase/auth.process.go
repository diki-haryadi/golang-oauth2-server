package usecase

import (
	"context"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang-standards-project-layout/internal/app/auth/model"
	tm "golang-standards-project-layout/internal/app/token/model"
	um "golang-standards-project-layout/internal/app/user/model"
	sam "golang-standards-project-layout/internal/integration/spotify_api/model"
	cbm "golang-standards-project-layout/internal/model/chatbot"
	"golang-standards-project-layout/internal/util"
	"golang.org/x/sync/errgroup"
)

// TODO: implement transaction
func (x *Module) ProcessLinkageCallback(ctx context.Context, data *model.LinkageCallback) (*um.UserNoSqlSchema, *tm.TokenNoSqlSchema, error) {
	//ctx, span := tracer.StartSpan(ctx, "auth.uc.ProcessLinkageCallback", nil)
	//defer span.End()

	var (
		err   error
		ares  *sam.SpotifyAuthorizeBodyRes
		ures  *sam.SpotifyGetUserProfileBodyRes
		user  *um.UserNoSqlSchema
		uid   primitive.ObjectID
		tdata *tm.TokenNoSqlSchema
		g     *errgroup.Group
	)

	// PRE: VALIDATE DATA
	g, _ = errgroup.WithContext(ctx)

	g.Go(func() error {
		return data.ValidateLinkageCallback()
	})

	g.Go(func() error {
		uid, err = primitive.ObjectIDFromHex(data.State)
		return err
	})

	if err = g.Wait(); err != nil {
		return nil, nil, err
	}

	// STEP 1: FIND USER IN THE DB
	user, err = x.userRepository.FindUserById(ctx, uid)

	if err != nil {
		return nil, nil, err
	}

	if user == nil {
		return nil, nil, util.NewDataNotFoundErr("User")
	}

	// STEP 2: GENERATE TOKEN TO SPOTIFY
	ares, err = x.spotifyAuthApiCall.ObtainToken(ctx, data.ToSpotifyAuthorizeData())

	if err != nil {
		return nil, nil, err
	}

	tdata = tdata.BuildFromSpotifyAuthorizeBodyRes(ares, user.Id)

	// STEP 3: STORE GENERATED TOKEN TO DB
	err = x.tokenRepository.StoreToken(ctx, tdata)

	if err != nil {
		log.Err(err).Msg(err.Error())
		return nil, nil, err
	}

	// STEP 4: GET USER INFO FROM TOKEN
	ures, err = x.spotifyAuthApiCall.GetUserInfo(ctx, ares.AccessToken)

	if err != nil {
		return nil, nil, err
	}

	// STEP 5: IF FIRST LINKAGE TO SPOTIFY, UPDATE SPOTIFY DATA TO DB
	if user.SpotifyData == um.NilUserSpotifyDataNoSqlSchema {
		user.AssignSpotifyData(ures)

		err = x.userRepository.UpdateUser(ctx, uid, user)
	}

	// STEP 6: SEND NOTIFICATION TO USER
	x.chatbotManager.SendLinkageSuccessMessage(ctx, user.ChatId)

	tdata = tdata.BuildFromSpotifyAuthorizeBodyRes(ares, user.Id)

	return user, tdata, nil
}

func (x *Module) ProcessHostAuthentication(ctx context.Context, ci cbm.ChatInfo) error {
	//ctx, span := tracer.StartSpan(ctx, "auth.uc.ProcessLinkageCallback", nil)
	//defer span.End()

	var (
		err   error
		udata *um.UserNoSqlSchema
		tdata *tm.TokenNoSqlSchema
	)

	udata, err = x.userRepository.FindUserByChatbotUserId(ctx, ci.SenderId, ci.Channel)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return err
	}

	tdata, err = x.tokenRepository.FindAndValidateTokenByUserId(ctx, udata.Id)

	if err != nil {
		log.Err(err).Msg(err.Error())
		return err
	}

	if tdata == nil {
		linkage_url := x.spotifyAuthApiCall.GenerateAuthorizeLink(ctx, udata.Id.Hex())
		err = x.chatbotManager.SendInitializeLinkageMessage(ctx, ci.ChatId, linkage_url)

		if err != nil {
			log.Err(err).Msg(err.Error())
		}
		return err
	}

	if udata == nil {
		udata = udata.BuildFromChatInfo(ci)
		err = x.userRepository.StoreUser(ctx, udata)
		if err != nil {
			log.Err(err).Msg(err.Error())
			return err
		}
	}

	if udata.ActiveSessionId == primitive.NilObjectID {
		err = x.chatbotManager.SendNoActiveSessionMessage(ctx, ci.ChatId)

		if err != nil {
			log.Err(err).Msg(err.Error())
			return err
		}
	} else {
		err = x.chatbotManager.SendHostActionsMessage(ctx, ci.ChatId)

		if err != nil {
			log.Err(err).Msg(err.Error())
			return err
		}
	}

	return err
}
