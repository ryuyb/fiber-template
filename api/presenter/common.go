package presenter

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    *any   `json:"data"`
}

func (r ErrorResponse) Error() string {
	return r.Message
}
