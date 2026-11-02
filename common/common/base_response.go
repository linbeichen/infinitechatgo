package common

type BaseResponse[T any] struct {
	Code    int    `json:"code"`
	Data    T      `json:"data"`
	Message string `json:"message"`
}

func NewBaseResponse[T any](code int, data T, message string) *BaseResponse[T] {
	return &BaseResponse[T]{Code: code, Data: data, Message: message}
}

func NewBaseResponseWithData[T any](code int, data T) *BaseResponse[T] {
	return NewBaseResponse(code, data, "")
}

func NewBaseResponseFromError(errorCode ErrorCode) *BaseResponse[any] {
	return NewBaseResponse[any](errorCode.Code, nil, errorCode.Message)
}

func NewBaseResponseFromMessage(errorCode ErrorCode, message string) *BaseResponse[any] {
	return NewBaseResponse[any](errorCode.Code, nil, message)
}
