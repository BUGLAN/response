package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// EchoResponse Echo Response interface
type EchoResponse interface {
	RespOk(ctx echo.Context, errmsg ...string) error
	RespBadRequest(ctx echo.Context, errmsg ...string) error
	RespNotFound(ctx echo.Context, errmsg ...string) error
	RespData(ctx echo.Context, data interface{}, errmsg ...string) error
	RespFail(ctx echo.Context, errmsg ...string) error
	Resp(ctx echo.Context, code int, resp Response) error
}

// EchoQuickResp echo quick response
type EchoQuickResp struct {
}

// RespOk response mean's ok
func (e *EchoQuickResp) RespOk(ctx echo.Context, errmsg ...string) error {
	msg := "ok"
	if len(errmsg) > 0 {
		msg = errmsg[0]
	}
	return e.Resp(ctx, http.StatusOK, Response{ErrCode: 0, ErrMsg: msg})
}

// RespBadRequest means the client param error
func (e *EchoQuickResp) RespBadRequest(ctx echo.Context, errmsg ...string) error {
	msg := "bad request"
	if len(errmsg) > 0 {
		msg = errmsg[0]
	}
	return e.Resp(ctx, http.StatusBadRequest, Response{ErrCode: http.StatusBadRequest, ErrMsg: msg})
}

// RespNotFound means resource not found
func (e *EchoQuickResp) RespNotFound(ctx echo.Context, errmsg ...string) error {
	msg := "not found"
	if len(errmsg) > 0 {
		msg = errmsg[0]
	}
	return e.Resp(ctx, http.StatusNotFound, Response{ErrCode: http.StatusNotFound, ErrMsg: msg})
}

// RespData data response
func (e *EchoQuickResp) RespData(ctx echo.Context, data interface{}, errmsg ...string) error {
	msg := "ok"
	if len(errmsg) > 0 {
		msg = errmsg[0]
	}
	return e.Resp(ctx, http.StatusOK, Response{Data: data, ErrMsg: msg})
}

// RespFail business fail code
func (e *EchoQuickResp) RespFail(ctx echo.Context, errmsg ...string) error {
	msg, code := "not ok", -1
	if len(errmsg) > 0 {
		msg = errmsg[0]
	}
	return e.Resp(ctx, http.StatusOK, Response{ErrCode: code, ErrMsg: msg})
}

// Resp consume resp body
func (e *EchoQuickResp) Resp(ctx echo.Context, code int, resp Response) error {
	return ctx.JSON(code, resp)
}
