package common

// standard response struct
type Response struct {
	Code    ErrorCode
	Message ErrorMessage
	Data    interface{}
}

func SucResponse(data interface{}) *Response {
	return &Response{
		Code:    Success,
		Message: "OK",
		Data:    data,
	}
}

func ParamResponse() *Response {
	return &Response{
		Code:    ParamCode,
		Message: Param,
		Data:    nil,
	}
}

func SystemResponse() *Response {
	return &Response{
		Code:    DatabaseCode,
		Message: Database,
		Data:    nil,
	}
}
