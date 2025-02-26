package responses

type Response struct {
	IsOK    bool    `json:"is_ok"`
	Message *string `json:"message"`
	Result  any     `json:"result"`
}

func NewSuccessResponse(result any) *Response {
	return &Response{
		IsOK:    true,
		Message: nil,
		Result:  result,
	}
}

func NewErrorResponse(err error) *Response {
	message := err.Error()

	return &Response{
		IsOK:    false,
		Message: &message,
		Result:  nil,
	}
}
