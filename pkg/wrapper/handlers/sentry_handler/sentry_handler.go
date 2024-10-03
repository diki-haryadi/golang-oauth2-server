package wrapperSentryhandler

import (
	"context"
	"golang-oauth2-server/config"

	"github.com/getsentry/sentry-go"

	sentryUtils "golang-oauth2-server/pkg/sentry/sentry_utils"
	"golang-oauth2-server/pkg/wrapper"
)

var SentryHandler = func(f wrapper.HandlerFunc) wrapper.HandlerFunc {
	return func(ctx context.Context, args ...interface{}) (interface{}, error) {
		hub := sentry.GetHubFromContext(ctx)
		if hub == nil {
			hub = sentry.CurrentHub().Clone()
			ctx = sentry.SetHubOnContext(ctx, hub)
		}
		hub.Scope().SetExtra("args", args)
		hub.Scope().SetTag("application", config.BaseConfig.App.AppName)
		hub.Scope().SetTag("AppEnv", config.BaseConfig.App.AppEnv)

		opts := &sentryUtils.Options{
			Repanic: true,
		}
		defer sentryUtils.RecoverWithSentry(hub, ctx, opts)

		return f(ctx, args)
	}
}
