package response

// Response web response struct
type Response struct {
	ErrMsg  string      `json:"errmsg"`
	ErrCode int         `json:"errcode"`
	Data    interface{} `json:"data"`
}
