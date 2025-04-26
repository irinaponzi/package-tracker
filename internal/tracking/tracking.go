package tracking

import (
	"encoding/json"
	"fmt"
	"github.com/irinaponzi/package-tracker/pkg/web"
	"net/http"
)

// TrackingHandler is a struct that represents a tracking code handler
type TrackingHandler struct {
	sv ITrackingService
}

// NewTrackingHandler creates a new TrackingHandler instance
func NewTrackingHandler(sv ITrackingService) *TrackingHandler {
	return &TrackingHandler{
		sv: sv,
	}
}

func (h *TrackingHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body TrackingRequestDTO

		// decode JSON body
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			web.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		// todo: validator which validate the struct

		// process
		err := h.sv.Create(body.Amount, body.Country)

		// error response // todo: improve error handling
		if err != nil {
			web.Error(w, http.StatusInternalServerError, err.Error())
			return
		}

		// success response
		message := fmt.Sprintf(msgCardsCreated, body.Amount)
		web.JSON(w, http.StatusCreated, NewTrackingResponse(msgSuccess, message))
		return
	}
}
