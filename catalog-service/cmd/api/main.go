package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Print("Couldnt run the server")
	}
}
