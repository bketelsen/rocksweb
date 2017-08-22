package models

import (
	"encoding/json"

	"github.com/markbates/validate"
	"github.com/markbates/validate/validators"
)

// String is not required by pop and may be deleted
func (a Author) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// Authors is not required by pop and may be deleted
type Authors []Author

// String is not required by pop and may be deleted
func (a Authors) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// Validate gets run every time you call a "pop.Validate" method.
// This method is not required and may be deleted.
func (a *Author) Validate() (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: a.Name, Name: "Name"},
	), nil
}

// ValidateSave gets run every time you call "pop.ValidateSave" method.
// This method is not required and may be deleted.
func (a *Author) ValidateSave() (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateUpdate" method.
// This method is not required and may be deleted.
func (a *Author) ValidateUpdate() (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
