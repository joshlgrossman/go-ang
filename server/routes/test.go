package routes

import (
	"net/http"
)

// Test endpoint
func Test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
