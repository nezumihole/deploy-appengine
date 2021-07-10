package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprintf(writer, "Hello! %s\n", request.Host)
		if err != nil {
			return
		}
	})

	http.ListenAndServe(":8080", nil)
}
