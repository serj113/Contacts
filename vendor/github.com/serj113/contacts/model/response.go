package model

type Response struct {
	Status  int    `json: status`
	Message string `json: message`
	Data    []Contact
}
