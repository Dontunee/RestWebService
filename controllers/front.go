package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegisterControllers() {
	controller := newUserController()
	http.Handle("/users", *controller)
	http.Handle("/users/", *controller)
}

func encodeResponseAsJson(data interface{}, writer io.Writer) {
	encoder := json.NewEncoder(writer)
	encoder.Encode(data)
}
