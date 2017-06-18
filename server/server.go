package main

import (
	"net/http"

	"github.com/joshlgrossman/go-ang/server/routes"
)

func main() {

	http.HandleFunc("/", routes.Static("client/build/"))
	http.HandleFunc("/ws/test", routes.Test)
	http.ListenAndServe(":8080", nil)

}
