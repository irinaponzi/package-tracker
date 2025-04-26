package batch

// IBatchService is an interface that represents the batch service
type IBatchService interface {
}

// BatchService is a struct that represents a batch service
type BatchService struct {
	rp IBatchRepository
}

// NewBatchService creates a new BatchService instance
func NewBatchService(rp IBatchRepository) *BatchService {
	return &BatchService{
		rp: rp,
	}
}
