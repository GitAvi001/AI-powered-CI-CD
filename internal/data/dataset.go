package data

import (
	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/knn"
	"log"
)

// LoadIrisDataset loads the Iris dataset for training.
func LoadIrisDataset() (*base.DenseInstances, error) {
	data, err := base.ParseCSVToInstances("https://archive.ics.uci.edu/ml/machine-learning-databases/iris/iris.data", true)
	if err != nil {
		log.Printf("Failed to load dataset: %v", err)
		return nil, err
	}
	return data, nil
}

// TrainKNN trains a k-NN classifier on the provided dataset.
func TrainKNN(data *base.DenseInstances) *knn.KNNClassifier {
	cls := knn.NewKnnClassifier("euclidean", "linear", 3)
	cls.Fit(data)
	return cls
}