package middleware

import (
	"fmt"
	"net/http"
)

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("This is where I would handle authentication for the endpoint")
		next(w, r)
	}
}
