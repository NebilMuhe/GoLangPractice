package main

import "net/http"

func handler(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, 200, struct{}{})
}
