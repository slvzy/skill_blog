package response

type ApiResponse struct {
	Status int         `json:"status"`
	Msg    interface{} `json:"msg"`
	Error  interface{} `json:"error"`
	Data   interface{} `json:"data"`
}

// 响应结果
func ApiResult(status int, msg string, objects interface{}) (apijson *ApiResponse) {
	apijson = &ApiResponse{Status: status, Data: objects, Msg: msg}
	return
}

// 带错误码的
func ApiResultAndError(status int,error string, msg string, objects interface{}) (apijson *ApiResponse) {
	apijson = &ApiResponse{Status: status,Error:error, Data: objects, Msg: msg}
	return
}

// 成功响应
func ApiResultSuccess(msg string, objects interface{}) (apijson *ApiResponse) {
	apijson = &ApiResponse{Status: 1, Data: objects, Msg: msg}
	return
}

// 失败响应
func ApiResultFail(msg string, objects interface{}) (apijson *ApiResponse) {
	apijson = &ApiResponse{Status: 0, Data: objects, Msg: msg}
	return
}
