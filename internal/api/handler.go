package api

import (
	"encoding/json"
	"net/http"
	"github.com/sjwhitworth/golearn/base"
	"github.com/GitAvi001/AI-powered-CI-CD/internal/model"
)

type Handler struct {
	Model *model.Model
}

func NewHandler(m *model.Model) *Handler {
	return &Handler{Model: m}
}

func (h *Handler) PredictHandler(w http.ResponseWriter, r *http.Request) {
	// Create a new DenseInstances (placeholder, replace with real preprocessing)
	instance := base.NewDenseInstances()
	// TODO: Add logic to populate instance with input data (e.g., from request body)

	w.Header().Set("Content-Type", "application/json")
	prediction, err := h.Model.Predict(instance)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := map[string]string{"prediction": prediction}
	json.NewEncoder(w).Encode(response)
}