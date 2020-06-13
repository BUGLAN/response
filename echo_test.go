package response

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var echoResp *EchoQuickResp

func echoContext(w http.ResponseWriter) echo.Context {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	return e.NewContext(req, w)
}

func TestEchoQuickResp_RespOk(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := echoContext(w)
	assert.NotNil(t, ctx)
	err := echoResp.RespOk(ctx)
	assert.Nil(t, err)
	resp, code, err := getResp(w)
	assert.Nil(t, err)
	assert.Equal(t, 200, code)
	assert.Equal(t, resp, Response{ErrMsg: "ok"})
}

func TestEchoQuickResp_RespBadRequest(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := echoContext(w)
	err := echoResp.RespNotFound(ctx)
	assert.Nil(t, err)
	resp, code, err := getResp(w)
	assert.Nil(t, err)
	assert.Equal(t, code, http.StatusNotFound)
	assert.Equal(t, resp, Response{ErrMsg: "not found", ErrCode: http.StatusNotFound})
}

func TestEchoQuickResp_RespData(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := echoContext(w)
	err := echoResp.RespData(ctx, map[string]string{"name": "lan"})
	assert.Nil(t, err)
	resp, code, err := getResp(w)
	assert.Nil(t, err)
	assert.Equal(t, code, http.StatusOK)
	assert.Equal(t, resp, Response{ErrMsg: "ok", Data: map[string]interface{}{"name": "lan"}})
}

func TestEchoQuickResp_Resp(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := echoContext(w)
	err := echoResp.Resp(ctx, http.StatusCreated, Response{ErrMsg: "not ok", ErrCode: -2, Data: map[string]string{"name": "lan"}})
	assert.Nil(t, err)
	resp, code, err := getResp(w)
	assert.Nil(t, err)
	assert.Equal(t, code, http.StatusCreated)
	assert.Equal(t, resp, Response{ErrMsg: "not ok", ErrCode: -2, Data: map[string]interface{}{"name": "lan"}})
}

func TestEchoQuickResp_RespFail(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := echoContext(w)
	err := echoResp.RespFail(ctx)
	assert.Nil(t, err)
	resp, code, err := getResp(w)
	assert.Nil(t, err)
	assert.Equal(t, code, http.StatusOK)
	assert.Equal(t, resp, Response{ErrCode: -1, ErrMsg: "not ok"})
}
