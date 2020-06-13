package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type EchoResponse interface {
	RespOk(ctx echo.Context, errmsg ...string) error
	RespBadRequest(ctx echo.Context, errmsg ...string) error
	RespNotFound(ctx echo.Context, errmsg ...string) error
	RespData(ctx echo.Context, data interface{}, errmsg ...string) error
	RespFail(ctx echo.Context, errmsg ...string) error
	Resp(ctx echo.Context, code int, resp Response) error
}

type EchoQuickResp struct {
}

func (e *EchoQuickResp) RespOk(ctx echo.Context, errmsg ...string) error {
	msg := "ok"
	if len(errmsg) > 0 {
		msg = errmsg[0]
	}
	return e.Resp(ctx, http.StatusOK, Response{ErrCode: 0, ErrMsg: msg})
}

func (e *EchoQuickResp) RespBadRequest(ctx echo.Context, errmsg ...string) error {
	msg := "bad request"
	if len(errmsg) > 0 {
		msg = errmsg[0]
	}
	return e.Resp(ctx, http.StatusBadRequest, Response{ErrCode: http.StatusBadRequest, ErrMsg: msg})
}

func (e *EchoQuickResp) RespNotFound(ctx echo.Context, errmsg ...string) error {
	msg := "not found"
	if len(errmsg) > 0 {
		msg = errmsg[0]
	}
	return e.Resp(ctx, http.StatusNotFound, Response{ErrCode: http.StatusNotFound, ErrMsg: msg})
}

func (e *EchoQuickResp) RespData(ctx echo.Context, data interface{}, errmsg ...string) error {
	msg := "ok"
	if len(errmsg) > 0 {
		msg = errmsg[0]
	}
	return e.Resp(ctx, http.StatusOK, Response{Data: data, ErrMsg: msg})
}

func (e *EchoQuickResp) RespFail(ctx echo.Context, errmsg ...string) error {
	msg, code := "not ok", -1
	if len(errmsg) > 0 {
		msg = errmsg[0]
	}
	return e.Resp(ctx, http.StatusOK, Response{ErrCode: code, ErrMsg: msg})
}

func (e *EchoQuickResp) Resp(ctx echo.Context, code int, resp Response) error {
	return ctx.JSON(code, resp)
}
