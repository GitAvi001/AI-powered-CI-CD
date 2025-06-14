package model

import (
	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/knn"
	"github.com/GitAvi001/AI-powered-CI-CD/internal/data"
)

// Model represents the AI model.
type Model struct {
	Classifier *knn.KNNClassifier
}

// NewModel creates and trains a new k-NN model.
func NewModel() (*Model, error) {
	dataset, err := data.LoadIrisDataset()
	if err != nil {
		return nil, err
	}
	cls := data.TrainKNN(dataset)
	return &Model{Classifier: cls}, nil
}

// Predict makes a prediction using the trained model.
func (m *Model) Predict(instance *base.DenseInstances) (string, error) {
	pred, err := m.Classifier.Predict(instance)
	if err != nil {
		return "", err
	}
	return pred.String(), nil
}