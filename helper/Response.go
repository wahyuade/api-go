package helper

import (
	"net/http"
	"encoding/json"
)

type Response struct {
	Status bool `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data,omitempty"`
}

func (r *Response) SetData(d interface{}) {
	r.Data = d
}

func (r *Response) SetMessage(m string) {
	r.Message = m
}

func (r *Response) SetStatus(status bool) {
	r.Status = status
}

func (r *Response) Build(status bool, m string, d interface{}) {
	r.Status = status
	r.Message = m
	r.Data = d
}

func (r *Response) Send(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(r)
}