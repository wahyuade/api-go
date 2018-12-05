package controllers

import (
	"net/http"
	p "github.com/wahyuade/api-go/repositories"
	h "github.com/wahyuade/api-go/helper"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	response := h.Response{}
	response.Build(true,p.ApiWelcome(),nil)
	response.Send(w)
}