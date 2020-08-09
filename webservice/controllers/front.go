package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegisterController(){
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
	http.Handle("/insert", *uc)
	http.Handle("/", *uc)
}

func encodeResponseAsJSON(data interface{}, w io.Writer){
	enc := json.NewEncoder(w)
	enc.Encode(data)
}