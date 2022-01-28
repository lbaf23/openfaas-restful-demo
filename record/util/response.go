package util

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, code int, res interface{}) {
	w.WriteHeader(code)
	data, _ := json.Marshal(res)
	w.Write(data)
}
