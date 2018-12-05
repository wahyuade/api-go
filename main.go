package main

import (
	"net/http"	
	r "github.com/wahyuade/api-go/router"
	"fmt"
)

func main() {
	fmt.Print("Server is running")
	http.ListenAndServe(":8000", r.Initialize())
}