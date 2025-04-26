package _package

import (
	"errors"
	"github.com/irinaponzi/package-tracker/pkg/web"
	"net/http"
)

// PackageHandler is a struct that represents a package handler
type PackageHandler struct {
	sv IPackageService
}

// NewPackageHandler creates a new PackageHandler instance
func NewPackageHandler(sv IPackageService) *PackageHandler {
	return &PackageHandler{
		sv: sv,
	}
}

// GetAll is a method that returns a handler function to get all packages
func (h *PackageHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		packages, err := h.sv.FindAll()

		// Response if an error occurred
		if err != nil {
			switch {
			case errors.Is(err, web.ErrNotFound):
				web.Error(w, http.StatusNotFound, err.Error())
				return
			default:
				web.Error(w, http.StatusInternalServerError, err.Error())
				return
			}
		}

		// Response OK
		web.JSON(w, http.StatusOK, map[string]any{
			"output": "success",
			"detail": packages,
		})
	}
}

// Create is a method that returns a handler function to create a new package
func (h *PackageHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/*var body CardRequestDTO

		// decode JSON body
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			return web.EncodeJSON(w, NewCardResponse(msgFail, msgInvalidJSON), http.StatusBadRequest)
		}

		// validate struct
		if validationErrors := ValidateCardRequest(body); validationErrors != nil {
			return web.EncodeJSON(w, validationErrors, http.StatusBadRequest)
		}

		// process
		err := h.sv.Create(body.Amount, body.Bin, body.Range, body.FundingType)

		// error response
		if err != nil {
			var stockErr *card.BINAndRangeStockError
			switch {
			case errors.As(err, &stockErr):
				return web.EncodeJSON(w, NewCardResponse(msgFail, err.Error()), http.StatusConflict)
			default:
				return web.EncodeJSON(w, NewCardResponse(msgFail, err.Error()), http.StatusInternalServerError)
			}
		}

		// success response
		message := fmt.Sprintf(msgCardsCreated, body.Amount)
		return web.EncodeJSON(w, NewCardResponse(msgSuccess, message), http.StatusCreated)*/
	}
}
