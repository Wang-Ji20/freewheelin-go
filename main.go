package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello, world\n")
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
