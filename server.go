package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DanielTitkov/gowebbp/app"
	"github.com/DanielTitkov/gowebbp/config"
	lg "github.com/DanielTitkov/gowebbp/logger"

	"go.uber.org/zap"
)

func requestLogger(handler http.Handler, logger *zap.SugaredLogger) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			lg.LogHTTPVisit(logger, r)
			handler.ServeHTTP(w, r)
		})
}

func sampleHandler(w http.ResponseWriter, _ *http.Request, logger *zap.SugaredLogger, mode string) {
	title, err := app.ProduceTitle(mode)
	if err == nil {
		fmt.Fprintf(w, title)
	} else {
		fmt.Fprintf(w, "Some error occurred")
		logger.Error(err)
	}
}

func funkyHandler(w http.ResponseWriter, _ *http.Request, logger *zap.SugaredLogger) {
	fmt.Fprintf(w, "<h1>Papa’s Got a Brand New Bag</h1>")
}

func jsonHandler(w http.ResponseWriter, _ *http.Request, logger *zap.SugaredLogger) {
	foods := app.GetFoods()
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(foods)
	if err != nil {
		logger.Errorf("JSON encoding failed: %v", err)
	}
}

func main() {
	// build logger
	logger, err := lg.CreateLogger("logger.json")
	if err != nil {
		panic(err)
	}

	// load env var
	runOpts, err := config.LoadRunOptions("app")
	if err != nil {
		logger.Fatalf("Env vars not loaded: %v", err)
	}

	// load config from file
	conf, err := config.LoadYamlConfig(runOpts.ConfigPath)
	if err != nil {
		logger.Fatalf("Config is not loaded: %v", err)
	}

	logger.Infof("Server is listening at port %s", runOpts.Port)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { sampleHandler(w, r, logger, conf.Mode) })
	http.HandleFunc("/funky", func(w http.ResponseWriter, r *http.Request) { funkyHandler(w, r, logger) })
	http.HandleFunc("/foods", func(w http.ResponseWriter, r *http.Request) { jsonHandler(w, r, logger) })

	mux := http.DefaultServeMux
	err = http.ListenAndServe(":"+runOpts.Port, requestLogger(mux, logger))
	if err != nil {
		logger.Fatalf("Server start failed: %v", err)
	}
}
