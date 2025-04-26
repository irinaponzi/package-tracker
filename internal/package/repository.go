package _package

import (
	"database/sql"
)

// IPackageRepository is an interface that represents the package repository
type IPackageRepository interface {
	FindAll() ([]Package, error)
}

// PackageRepository is a struct that represents a package repository
type PackageRepository struct {
	db *sql.DB
}

// NewPackageRepository creates a new PackageRepository instance
func NewPackageRepository(db *sql.DB) *PackageRepository {
	return &PackageRepository{
		db: db,
	}
}

// FindAll retrieves all packages from the database
func (r *PackageRepository) FindAll() ([]Package, error) {
	query := `
		SELECT tracking, status, size, weight_kg, destination
    	FROM packages
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var packages []Package
	for rows.Next() {
		var p Package

		err := rows.Scan(&p.TrackingCode, &p.Status, &p.Size, &p.WeightKg, &p.Destination)
		if err != nil {
			return nil, err
		}
		packages = append(packages, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return packages, nil
}
