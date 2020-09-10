package models

type SimpleMessage struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

type SimpleMessageRequest struct {
	Message string `json:"message"`
}

type SimpleMessageResponse struct {
	ID      string `json:"id"`
	Message string `json:"message"`
	Error   error  `json:"error"`
}
