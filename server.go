package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/DanielTitkov/gowebbp/app"
	"github.com/DanielTitkov/gowebbp/config"
)

func sampleHandler(w http.ResponseWriter, r *http.Request, mode string) {
	title, err := app.ProduceTitle(mode)
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
	runOpts, err := config.LoadRunOptions("app")
	confPath := runOpts.ConfigPath // get from env
	port := runOpts.Port           // get from env

	conf, err := config.LoadYamlConfig(confPath)
	if err != nil {
		log.Fatalf("Config is not loaded: %v", err)
	}

	log.Printf("Server is listening at port %s", port)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { sampleHandler(w, r, conf.Mode) })
	http.HandleFunc("/funky", funkyHandler)
	http.HandleFunc("/foods", jsonHandler)
	http.ListenAndServe(":"+port, nil)
}
