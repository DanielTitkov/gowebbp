package main

import (
	"fmt"
	"log"
	"net/http"
)

func sampleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Trabajadores al poder!")
}

func funkyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Papa’s Got a Brand New Bag</h1>")
}

func main() {
	const port string = "8000"
	log.Printf("Server is listening at port %s", port)
	http.HandleFunc("/", sampleHandler)
	http.HandleFunc("/funky", funkyHandler)
	http.ListenAndServe(":"+port, nil)
}
