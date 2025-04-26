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

// GetAll is a temporary endpoint to retrieve all packages only for testing purposes
// This approach is not optimal for large-scale systems as it performs a costly query
// In next iterations this will be replaced by a key-value store to allow querying by ID
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
	}
}
