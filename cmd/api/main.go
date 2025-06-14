package main

import (
	"log"
	"net/http"
	"github.com/GitAvi001/AI-powered-CI-CD/internal/api"
	"github.com/GitAvi001/AI-powered-CI-CD/internal/model"
)

func main() {
	//Model initializing
	m, err := model.NewModel()
	if err != nil {
		log.Fatalf("Failed to initialize model: %v", err)
	}

	handler := api.NewHandler(m)
	http.HandleFunc("/predict", handler.PredictHandler)

	//HTTP server starting
	log.Println("Server starting on :9091...")
	if err := http.ListenAndServe(":9091", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}