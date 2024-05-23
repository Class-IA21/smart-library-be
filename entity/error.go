package entity

type ErrorResponse struct {
	Error   bool   `json:"error"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ErrorResponseWithErrors struct {
	Error   bool     `json:"error"`
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Errors  []string `json:"errors,omitempty"`
}
