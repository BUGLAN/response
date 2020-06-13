package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GinResponse interface {
	RespOk(ctx *gin.Context, errmsg ...string)
	RespBadRequest(ctx *gin.Context, errmsg ...string)
	RespNotFound(ctx *gin.Context, errmsg ...string)
	RespData(ctx *gin.Context, data interface{}, errmsg ...string)
	Resp(ctx *gin.Context, code int, resp Response)
	RespFail(ctx *gin.Context, errmsg ...string)
}

type GinQuickResp struct{}

// RespOk response mean's ok
func (r *GinQuickResp) RespOk(ctx *gin.Context, errmsg ...string) {
	s := "ok"
	if len(errmsg) > 0 {
		s = errmsg[0]
	}
	r.Resp(ctx, http.StatusOK, Response{ErrMsg: s})
}

// RespFail business fail code
func (r *GinQuickResp) RespFail(ctx *gin.Context, errmsg ...string) {
	s := "not ok"
	if len(errmsg) > 0 {
		s = errmsg[0]
	}
	r.Resp(ctx, http.StatusOK, Response{ErrMsg: s, ErrCode: -1})
}

// RespBadRequest means the client param error
func (r *GinQuickResp) RespBadRequest(ctx *gin.Context, errmsg ...string) {
	s := "bad request"
	if len(errmsg) > 0 {
		s = errmsg[0]
	}
	r.Resp(ctx, http.StatusBadRequest, Response{ErrMsg: s})
}

// RespNotFound means resource not found
func (r *GinQuickResp) RespNotFound(ctx *gin.Context, errmsg ...string) {
	s := "not found"
	if len(errmsg) > 0 {
		s = errmsg[0]
	}
	r.Resp(ctx, http.StatusNotFound, Response{ErrMsg: s, ErrCode: http.StatusNotFound})
}

// RespData data response
func (r *GinQuickResp) RespData(ctx *gin.Context, data interface{}, errmsg ...string) {
	s := "ok"
	if len(errmsg) > 0 {
		s = errmsg[0]
	}
	r.Resp(ctx, http.StatusOK, Response{ErrMsg: s, Data: data})
}

// Resp consume resp body
func (r *GinQuickResp) Resp(ctx *gin.Context, code int, resp Response) {
	ctx.JSON(code, resp)
}
