package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	fs := http.FileServer(http.Dir("./"))
	fmt.Println("Starting server on 8080...")

	http.ListenAndServe(":8080", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if strings.HasSuffix(r.URL.Path, ".wasm") {
				w.Header().Set("content-type", "application/wasm")
			}

			fs.ServeHTTP(w, r)
		}))
}
