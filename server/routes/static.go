package routes

import (
	"net/http"
)

// Static page handler
func Static(path string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, path+r.URL.Path[1:])
	}
}
