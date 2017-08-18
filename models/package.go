package models

import (
	"encoding/json"
	"time"

	"github.com/markbates/pop"
	"github.com/markbates/pop/slices"
	"github.com/markbates/validate"
	"github.com/markbates/validate/validators"
	"github.com/satori/go.uuid"
)

type Package struct {
	ID          uuid.UUID     `json:"id" db:"id"`
	CreatedAt   time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at" db:"updated_at"`
	AuthorIDs   slices.String `json:"author_ids" db:"author_ids"`
	ImportPath  string        `json:"import_path" db:"import_path"`
	SourceURL   string        `json:"source_url" db:"source_url"`
	License     string        `json:"license" db:"license"`
	Keywords    slices.String `json:"keywords" db:"keywords"`
	Description string        `json:"description" db:"description"`
	VersionIDs  slices.String `json:"version_ids" db:"version_ids"`
	Authors     Authors       `json:"authors" db:"-"`
	Versions    Versions      `json:"versions" db:"-"`
}

// String is not required by pop and may be deleted
func (p Package) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Packages is not required by pop and may be deleted
type Packages []Package

// String is not required by pop and may be deleted
func (p Packages) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Validate gets run every time you call a "pop.Validate" method.
// This method is not required and may be deleted.
func (p *Package) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: p.ImportPath, Name: "ImportPath"},
		&validators.StringIsPresent{Field: p.SourceURL, Name: "SourceUrl"},
		&validators.StringIsPresent{Field: p.License, Name: "License"},
		&validators.StringIsPresent{Field: p.Description, Name: "Description"},
	), nil
}

// ValidateSave gets run every time you call "pop.ValidateSave" method.
// This method is not required and may be deleted.
func (p *Package) ValidateSave(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateUpdate" method.
// This method is not required and may be deleted.
func (p *Package) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
