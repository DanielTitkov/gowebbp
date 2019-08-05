package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/DanielTitkov/gowebbp/app"
)

func sampleHandler(w http.ResponseWriter, r *http.Request) {
	title, err := app.ProduceTitle("Foo")
	if err == nil {
		fmt.Fprintf(w, title)
	} else {
		fmt.Fprintf(w, "Some error occured")
		log.Print(err)
	}
}

func funkyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Papa’s Got a Brand New Bag</h1>")
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	foods := app.GetFoods()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(foods)
}

func main() {
	const port string = "8000"
	log.Printf("Server is listening at port %s", port)
	http.HandleFunc("/", sampleHandler)
	http.HandleFunc("/funky", funkyHandler)
	http.HandleFunc("/foods", jsonHandler)
	http.ListenAndServe(":"+port, nil)
}
