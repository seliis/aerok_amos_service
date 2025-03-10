package responses

type Response struct {
	IsOK    bool    `json:"is_ok"`
	Message *string `json:"message"`
	Data    any     `json:"data"`
}

func NewSuccessResponse(data any) *Response {
	return &Response{
		IsOK:    true,
		Message: nil,
		Data:    data,
	}
}

func NewErrorResponse(err error) *Response {
	message := err.Error()

	return &Response{
		IsOK:    false,
		Message: &message,
		Data:    nil,
	}
}
