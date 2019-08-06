package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DanielTitkov/gowebbp/app"
	"go.uber.org/zap"
)

// SampleHandler writes phrase base on mode
func SampleHandler(w http.ResponseWriter, _ *http.Request, logger *zap.SugaredLogger, mode string) {
	title, err := app.ProduceTitle(mode)
	if err == nil {
		fmt.Fprintf(w, title)
	} else {
		fmt.Fprintf(w, "Some error occurred")
		logger.Error(err)
	}
}

// FunkyHandler writes stupid header
func FunkyHandler(w http.ResponseWriter, _ *http.Request, logger *zap.SugaredLogger) {
	fmt.Fprintf(w, "<h1>Papa’s Got a Brand New Bag</h1>")
}

// JSONHandler writes json
func JSONHandler(w http.ResponseWriter, _ *http.Request, logger *zap.SugaredLogger) {
	foods := app.GetFoods()
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(foods)
	if err != nil {
		logger.Errorf("JSON encoding failed: %v", err)
	}
}
