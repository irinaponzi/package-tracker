package batch

import (
	"github.com/google/uuid"
	"time"
)

type BatchStatus string

const (
	BatchStatusPending BatchStatus = "pending"
	BatchStatusDone    BatchStatus = "done"
	BatchStatusFail    BatchStatus = "fail"
)

type Batch struct {
	ID       uuid.UUID `json:"id"`
	Quantity int       `json:"quantity"`
	// todo: transport company will be a struct in the next iteration
	TransportCompany *string     `json:"transport_company"`
	Status           BatchStatus `json:"status"`
	ErrorDetails     *string     `json:"error_details,omitempty"`
	CreatedAt        time.Time   `json:"created_at"`
}
