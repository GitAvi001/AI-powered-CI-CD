package api

//Creates a simple REST API to serve applications
import (
	"encoding/json"
	"net/http"
	"github.com/GitAvi001/AI-powered-CI-CD/internal/model"
)

type Handler struct {
	Model *model.Model
}

func NewHandler(m *model.Model) *Handler {
	return &Handler{Model: m}
}

func (h *Handler) PredictHandler(w http.ResponseWriter, r *http.Request) {
	// Placeholder for prediction input (simplified for demo)
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"prediction": "Iris-setosa"} // Mock prediction
	json.NewEncoder(w).Encode(response)
}