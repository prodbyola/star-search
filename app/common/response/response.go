package response

// Error codes: https://en.wikipedia.org/wiki/List_of_HTTP_status_codes
type Response struct {
	Error   bool        `json:"error"`
	Code    uint        `json:"code"`
	Message interface{} `json:"message"`
}

func New() Response {
	return Response{
		Code:    417,
		Message: "Error is not defined",
		Error:   true,
	}
}

func (r *Response) SetData(data interface{}) {
	r.Code = 200
	r.Message = data
	r.Error = false
}

func (r *Response) SetError(code uint, err error) {
	r.Code = code
	r.Message = err.Error()
	r.Error = true
}
