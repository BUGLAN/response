package response

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var quick *GinQuickResp

func context(w *httptest.ResponseRecorder) *gin.Context {
	ctx, _ := gin.CreateTestContext(w)
	return ctx
}

func getResp(w *httptest.ResponseRecorder) (Response, int, error) {
	var resp Response
	b, _ := ioutil.ReadAll(w.Body)
	err := json.Unmarshal(b, &resp)
	return resp, w.Code, err
}

func TestQuickResp_RespOk(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := context(w)
	quick.RespOk(ctx)
	resp, code, err := getResp(w)
	assert.Nil(t, err)
	assert.Equal(t, code, http.StatusOK)
	assert.Equal(t, resp, Response{ErrMsg: "ok"})
}

func TestQuickResp_RespBadRequest(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := context(w)
	quick.RespNotFound(ctx)
	resp, code, err := getResp(w)
	assert.Nil(t, err)
	assert.Equal(t, code, http.StatusNotFound)
	assert.Equal(t, resp, Response{ErrMsg: "not found", ErrCode: http.StatusNotFound})
}

func TestQuickResp_RespData(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := context(w)
	quick.RespData(ctx, map[string]string{"name": "lan"})
	resp, code, err := getResp(w)
	assert.Nil(t, err)
	assert.Equal(t, code, http.StatusOK)
	assert.Equal(t, resp, Response{ErrMsg: "ok", Data: map[string]interface{}{"name": "lan"}})
}

func TestQuickResp_Resp(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := context(w)
	quick.Resp(ctx, http.StatusCreated, Response{ErrMsg: "not ok", ErrCode: -2, Data: map[string]string{"name": "lan"}})
	resp, code, err := getResp(w)
	assert.Nil(t, err)
	assert.Equal(t, code, http.StatusCreated)
	assert.Equal(t, resp, Response{ErrMsg: "not ok", ErrCode: -2, Data: map[string]interface{}{"name": "lan"}})
}

func TestQuickResp_RespFail(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := context(w)
	quick.RespFail(ctx)
	resp, code, err := getResp(w)
	assert.Nil(t, err)
	assert.Equal(t, code, http.StatusOK)
	assert.Equal(t, resp, Response{ErrCode: -1, ErrMsg: "not ok"})
}
