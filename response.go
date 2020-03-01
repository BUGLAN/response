package response

// Response gin framework response
type Response struct {
	ErrMsg  string      `json:"errmsg"`
	ErrCode int         `json:"errcode"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}
