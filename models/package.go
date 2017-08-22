package models

import (
	"encoding/json"

	"github.com/markbates/validate"
	"github.com/markbates/validate/validators"
)

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
func (p *Package) Validate() (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: p.ImportPath, Name: "ImportPath"},
		&validators.StringIsPresent{Field: p.SourceURL, Name: "SourceUrl"},
		&validators.StringIsPresent{Field: p.License, Name: "License"},
		&validators.StringIsPresent{Field: p.Description, Name: "Description"},
	), nil
}

// ValidateSave gets run every time you call "pop.ValidateSave" method.
// This method is not required and may be deleted.
func (p *Package) ValidateSave() (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateUpdate" method.
// This method is not required and may be deleted.
func (p *Package) ValidateUpdate() (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
