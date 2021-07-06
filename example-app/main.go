package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello! %s\n", request.Host)
	})

	http.ListenAndServe(":8080", nil)
}
