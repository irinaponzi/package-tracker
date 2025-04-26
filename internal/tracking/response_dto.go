package tracking

const (
	msgSuccess      = "success"
	msgFail         = "fail"
	msgCardsCreated = "%d cards has been created successfully"
)

type TrackingResponse struct {
	Output string `json:"output"`
	Detail string `json:"detail"`
}

func NewTrackingResponse(output, detail string) TrackingResponse {
	return TrackingResponse{
		Output: output,
		Detail: detail,
	}
}
