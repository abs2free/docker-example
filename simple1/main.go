package main

import (
	"fmt"

	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world! from %s\n", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("start http serve localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
