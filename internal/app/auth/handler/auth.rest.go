package handler

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"golang-standards-project-layout/internal/app/auth/model"
	tm "golang-standards-project-layout/internal/app/token/model"
	um "golang-standards-project-layout/internal/app/user/model"
	"golang-standards-project-layout/internal/util"
	errlib "golang-standards-project-layout/pkg/err"
	"golang-standards-project-layout/pkg/response"
)

// Spotify Auth Callback
// @Summary Spotify Auth Callback
// @Description Spotify Auth Callback
// @Tags Auth
// @Accept json
// @Produce json
// @Param state query string true "state from callback"
// @Param user_id path string true "User ID"
// @Success 200 {object} response.JSONResponse{data=model.LinkageCallbackBodyRes}
// @Failure 400 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /auth/callback [GET]
func (m *RestModule) HandleLinkageCallback(fc *fiber.Ctx) error {
	//ctx, span := tracer.StartSpan(fc.Context(), "auth.rest.HandleLinkageCallback", nil)
	//defer span.End()

	var (
		//beginTs = time.Now()
		qParam  = &model.LinkageCallbackQParams{}
		res     response.JSONResponse
		err     error
		udata   *um.UserNoSqlSchema
		tdata   *tm.TokenNoSqlSchema
		resData *model.LinkageCallbackBodyRes
	)

	err = fc.QueryParser(qParam)

	if err != nil {
		e := errlib.GetError(err)
		return util.ReturnErrorToFiberResponse(fc, e)
	}

	udata, tdata, err = m.authUsecase.ProcessLinkageCallback(context.Background(), qParam.ToLinkageCallback())

	if err != nil {
		e := errlib.GetError(err)
		return util.ReturnErrorToFiberResponse(fc, e)
	}

	resData = model.BuildLinkageCallbackBodyRes(udata, tdata)

	res = *(response.NewJSONResponse().APIStatusCreated().SetData(resData))

	return fc.JSON(res)
}
