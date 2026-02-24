package main

import (
	"fmt"
	"net/http"
)

func AuthMidleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token != "secret_token" {
			http.Error(w, "Unauthorized!", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

func dashBoardHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the super secret Admin Dashboard!")

}
func main() {
	http.HandleFunc("GET /admin", AuthMidleWare(dashBoardHandler))
	http.ListenAndServe(":8080", nil)
}
