package tracking

type TrackingRequestDTO struct {
	Amount  int    `json:"amount" validate:"min=1,max=1000000"`
	Country string `json:"country" validate:"required"`
}
