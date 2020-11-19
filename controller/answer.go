package controller

import (
    "net/http"
    "encoding/json"
)

func Answer(w http.ResponseWriter, HTTPCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(HTTPCode)

	js, err := json.Marshal(body)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	w.Write(js)
}