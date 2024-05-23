package entity

type ResponseWebWithData struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ResponseWebWithoutData struct {
	Error   bool   `json:"error`
	Message string `json:"message"`
}
