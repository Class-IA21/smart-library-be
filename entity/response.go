package entity

type ResponseWebWithData struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Code    int    `json:"code"`
	Data    any    `json:"data"`
}

type ResponseWebWithoutData struct {
	Error   bool   `json:"error`
	Code    int    `json:"code"`
	Message string `json:"message"`
}
