package main

import (
	"net/http"

	"github.com/code-shivy/go/HTTP-Server/api"
)
func main(){
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}
