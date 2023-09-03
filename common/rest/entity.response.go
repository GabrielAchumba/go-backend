package rest

type Response struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	StatusCode int         `json:"statusCode"`
}

func (ap Response) GetSuccess(statusCode int, data interface{}) *Response {
	return &Response{
		StatusCode: statusCode,
		Data:       data,
		Message:    "Success",
		Success:    true,
	}
}

func (ap Response) GetError(statusCode int, msg string) *Response {
	return &Response{
		StatusCode: statusCode,
		Message:    msg,
		Success:    false,
	}
}

func (ap Response) NotAuthorized() *Response {
	return &Response{
		StatusCode: 400,
		Message:    "Not Authorized",
		Success:    false,
	}
}

func GetSuccess(statusCode int, data interface{}) *Response {
	return &Response{
		StatusCode: statusCode,
		Data:       data,
		Message:    "Success",
		Success:    true,
	}
}

func GetError(statusCode int, msg string) *Response {
	return &Response{
		StatusCode: statusCode,
		Message:    msg,
	}
}
