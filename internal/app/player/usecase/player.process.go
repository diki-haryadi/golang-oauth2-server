package usecase

import (
	"context"

	cbm "golang-standards-project-layout/internal/model/chatbot"
)

func (x *Module) ProcessQueueTrack(ctx context.Context, ci cbm.ChatInfo) (err error) {
	//ctx, span := tracer.StartSpan(ctx, "player.uc.ProcessQueueTrack", nil)
	//defer span.End()

	var (
		keyword = ci.Message
	)

	x.spotifyPlayerApiCall.SearchTracks(ctx, "", keyword)

	return err
}
