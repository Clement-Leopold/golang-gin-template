package common

// standard response struct
type Response struct {
	Code    ErrorCode
	Message ErrorMessage
	Data    interface{}
}
