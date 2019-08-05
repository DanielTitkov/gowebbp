package logger

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"go.uber.org/zap"
)

// CreateLogger builds logger from json config
func CreateLogger(path string) (*zap.SugaredLogger, error) {
	var emptyLogger zap.SugaredLogger

	rawJSONConfig, err := ioutil.ReadFile(path)
	if err != nil {
		return &emptyLogger, err
	}

	config := zap.Config{}
	if err := json.Unmarshal(rawJSONConfig, &config); err != nil {
		return &emptyLogger, err
	}
	logger, err := config.Build()
	if err != nil {
		return &emptyLogger, err
	}

	return logger.Sugar(), nil
}

// LogHTTPVisit logs basic request data
func LogHTTPVisit(l *zap.SugaredLogger, r *http.Request) {
	l.Infof("%v%v requested from %v", r.Host, r.URL, r.RemoteAddr)
}
