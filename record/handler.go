package function

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Data    string `json:"data"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		Get(w, r)
	} else if r.Method == http.MethodPost {
		Post(w, r)
	} else if r.Method == http.MethodPut {
		Put(w, r)
	} else if r.Method == http.MethodDelete {
		Delete(w, r)
	} else {
		w.WriteHeader(http.StatusBadGateway)
	}
}

func Get(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Data:    "get",
		Message: "ok",
		Code:    200,
	}

	res, err := json.Marshal(response)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Put(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Data:    "put",
		Message: "ok",
		Code:    200,
	}

	res, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Post(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Data:    "post",
		Message: "ok",
		Code:    200,
	}

	res, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Data:    "delete",
		Message: "ok",
		Code:    200,
	}

	res, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
