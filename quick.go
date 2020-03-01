package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type GinResponse interface {
	RespOk(ctx *gin.Context, errmsg ...string)
	RespBadRequest(ctx *gin.Context, errmsg ...string)
	RespNotFound(ctx *gin.Context, errmsg ...string)
	RespData(ctx *gin.Context, data interface{}, errmsg ...string)
	Resp(ctx *gin.Context, code int, resp Response)
}

type QuickResp struct{}

//RespOk response mean's ok
func (r *QuickResp) RespOk(ctx *gin.Context, errmsg ...string) {
	s := "ok"
	if len(errmsg) > 0 {
		s = errmsg[0]
	}
	r.Resp(ctx, http.StatusOK, Response{ErrMsg: s})
}

//RespBadRequest means the client param error
func (r *QuickResp) RespBadRequest(ctx *gin.Context, errmsg ...string) {
	s := "bad request"
	if len(errmsg) > 0 {
		s = errmsg[0]
	}
	r.Resp(ctx, http.StatusBadRequest, Response{ErrMsg: s})
}

//RespNotFound means resource not found
func (r *QuickResp) RespNotFound(ctx *gin.Context, errmsg ...string) {
	s := "not found"
	if len(errmsg) > 0 {
		s = errmsg[0]
	}
	r.Resp(ctx, http.StatusNotFound, Response{ErrMsg: s, ErrCode: http.StatusNotFound})
}

//RespData data response
func (r *QuickResp) RespData(ctx *gin.Context, data interface{}, errmsg ...string) {
	s := "ok"
	if len(errmsg) > 0 {
		s = errmsg[0]
	}
	r.Resp(ctx, http.StatusOK, Response{ErrMsg: s, Data: data})
}

//Resp consume resp body
func (r *QuickResp) Resp(ctx *gin.Context, code int, resp Response) {
	ctx.JSON(code, resp)
}
