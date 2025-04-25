package api

import "net/http"

func goodbye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Goodbye!\n"))
}