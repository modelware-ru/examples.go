package main

import (
	"fmt"
	"log"
	"net/http"
)

func main()  {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "backend server:")
		fmt.Fprintf(writer, "%v", request)
	})
	log.Fatal(http.ListenAndServe(":10001", nil))
}
