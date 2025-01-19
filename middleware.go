package main

import "net/http"

func limitSizeMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, 1<<30)
		if err := r.ParseForm(); err != nil {
			respondWithError(w, http.StatusBadRequest, "File exceeds the limit", err)
		}
		handler(w, r)
	}
}
