package _package

import "github.com/irinaponzi/package-tracker/pkg/web"

// IPackageService is an interface that represents the package service
type IPackageService interface {
	FindAll() ([]PackageDto, error)
}

// PackageService is a struct that represents a package service
type PackageService struct {
	rp IPackageRepository
}

// NewPackageService creates a new PackageService instance
func NewPackageService(rp IPackageRepository) *PackageService {
	return &PackageService{
		rp: rp,
	}
}

// FindAll retrieves a DTO of all packages
func (s *PackageService) FindAll() ([]PackageDto, error) {
	packages, err := s.rp.FindAll()
	if err != nil {
		return nil, web.ErrInternal
	}
	// Check if there are no packages
	if len(packages) == 0 {
		return nil, web.ErrNotFound
	}
	
	// Map warehouses to DTO
	var response []PackageDto
	for _, p := range packages {
		response = append(response, *MapPackageToDto(&p))
	}

	return response, nil
}
