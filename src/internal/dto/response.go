package dto

type Response struct {
	IsSuccess bool    `json:"is_success"`
	Message   *string `json:"message"`
	Data      any     `json:"data"`
}

func NewSuccessResponse(data any) *Response {
	return &Response{
		IsSuccess: true,
		Message:   nil,
		Data:      data,
	}
}

func NewErrorResponse(err error) *Response {
	message := err.Error()

	return &Response{
		IsSuccess: false,
		Message:   &message,
		Data:      nil,
	}
}
