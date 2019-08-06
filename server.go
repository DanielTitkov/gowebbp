package main

import (
	"net/http"

	"github.com/DanielTitkov/gowebbp/config"
	hd "github.com/DanielTitkov/gowebbp/handlers"
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

func main() {
	// build logger
	logger, err := lg.CreateLogger("./configuration_files/logger.json")
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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { hd.SampleHandler(w, r, logger, conf.Mode) })
	http.HandleFunc("/funky", func(w http.ResponseWriter, r *http.Request) { hd.FunkyHandler(w, r, logger) })
	http.HandleFunc("/foods", func(w http.ResponseWriter, r *http.Request) { hd.JSONHandler(w, r, logger) })

	mux := http.DefaultServeMux
	err = http.ListenAndServe(":"+runOpts.Port, requestLogger(mux, logger))
	if err != nil {
		logger.Fatalf("Server start failed: %v", err)
	}
}
