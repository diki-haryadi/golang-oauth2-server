package chatbot

import (
	"context"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
	sm "golang-standards-project-layout/internal/app/session/model"
	tm "golang-standards-project-layout/internal/app/token/model"
	um "golang-standards-project-layout/internal/app/user/model"
	gam "golang-standards-project-layout/internal/model/auth"
	cbm "golang-standards-project-layout/internal/model/chatbot"
)

func (x *ChatbotMiddlewareModule) ParseAndValidateSenderData(ctx context.Context, ci cbm.ChatInfo) (res *gam.CommandMetadata, err error) {
	//ctx, span := tracer.StartSpan(ctx, "chatbot.mw.ParseAndValidateSenderData", nil)
	//defer span.End()

	var (
		u *um.UserNoSqlSchema
		s *sm.SessionNoSqlSchema
		t *tm.TokenNoSqlSchema
	)

	u, err = x.userRepository.FindUserByChatbotUserId(ctx, ci.SenderId, ci.Channel)

	if err != nil {
		log.Err(err).Msgf(err.Error())
		x.chatbotManager.SendErrorMessage(ctx, ci.ChatId, err.Error())
		return nil, err
	}

	if u.ActiveSessionId != primitive.NilObjectID {
		s, err = x.sessionRepository.FindSessionById(ctx, u.ActiveSessionId)
		if err != nil {
			log.Err(err).Msgf(err.Error())
			x.chatbotManager.SendErrorMessage(ctx, ci.ChatId, err.Error())
			return nil, err
		}

		t, err = x.tokenRepository.FindAndValidateTokenByUserId(ctx, s.HostId)

		if err != nil {
			log.Err(err).Msgf(err.Error())
			x.chatbotManager.SendErrorMessage(ctx, ci.ChatId, err.Error())
			return nil, err
		}
	}

	res = &gam.CommandMetadata{
		User:      *u,
		Session:   *s,
		HostToken: *t,
	}

	return res, nil
}
