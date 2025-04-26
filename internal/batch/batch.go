package batch

import (
	"net/http"
)

// BatchHandler is a struct that represents a batch handler
type BatchHandler struct {
	sv IBatchService
}

// NewBatchHandler is a function that returns a new instance of BatchHandler
func NewBatchHandler(sv IBatchService) *BatchHandler {
	return &BatchHandler{
		sv: sv,
	}
}

func (h *BatchHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
	}
}
