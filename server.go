package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting server on 8080...")

	log.Fatal(
		http.ListenAndServe(":8080",
			http.FileServer(
				http.Dir("./"),
			),
		),
	)
}
