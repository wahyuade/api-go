package controllers

import (
	"fmt"
	"net/http"
	p "github.com/wahyuade/api-go/repositories"
	h "github.com/wahyuade/api-go/helper"
)

func PostLogin(w http.ResponseWriter, r *http.Request) {
	response := h.Response{}
	fmt.Println(r.FormValue("user"))
	response.Build(true,p.PostLogin(),nil)
	response.Send(w)
}