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

// ParamErrorResponse: helper for param-error response.
func ParamErrorResponse() *Response {
	return &Response{
		Code:    ParamCode,
		Message: Param,
		Data:    nil,
	}
}

// ParamResponse: helper for system-error response.
func SystemErrorResponse() *Response {
	return &Response{
		Code:    DatabaseCode,
		Message: Database,
		Data:    nil,
	}
}
